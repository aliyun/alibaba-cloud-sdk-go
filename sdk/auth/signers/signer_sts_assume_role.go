/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package signers

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/jmespath/go-jmespath"
	"net/http"
	"strconv"
	"time"
)

const (
	defaultDurationSeconds = 3600
)

type SignerStsAssumeRole struct {
	*credentialUpdater
	roleSessionName   string
	sessionCredential *sessionCredential
	credential        *credentials.StsAssumeRoleCredential
}

type sessionCredential struct {
	accessKeyId     string
	accessKeySecret string
	securityToken   string
}

func NewSignerStsAssumeRole(credential *credentials.StsAssumeRoleCredential, commonApi func(request *requests.CommonRequest) (response *responses.CommonResponse, err error)) (signer *SignerStsAssumeRole, err error) {
	signer = &SignerStsAssumeRole{
		credentialUpdater: &credentialUpdater{
			credentialExpiration: credential.RoleSessionExpiration,
			buildRequestMethod:   signer.buildCommonRequest,
			responseCallBack:     signer.refreshCredential,
			commonApi:            commonApi,
		},
		credential: credential,
	}
	if len(credential.RoleSessionName) > 0 {
		signer.roleSessionName = credential.RoleSessionName
	} else {
		signer.roleSessionName = "aliyun-go-sdk-" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
	}
	if credential.RoleSessionExpiration > 0 {
		if credential.RoleSessionExpiration > 900 && credential.RoleSessionExpiration < 3600 {
			signer.credentialExpiration = credential.RoleSessionExpiration
		} else {
			err = errors.NewClientError(errors.InvalidParamCode, "Assume Role session duration should be in the range of 15min - 1Hr", nil)
		}
	} else {
		signer.credentialExpiration = defaultDurationSeconds
	}
	return
}

func (*SignerStsAssumeRole) GetName() string {
	return "HMAC-SHA1"
}

func (*SignerStsAssumeRole) GetType() string {
	return ""
}

func (*SignerStsAssumeRole) GetVersion() string {
	return "1.0"
}

func (signer *SignerStsAssumeRole) GetAccessKeyId() string {
	return signer.credential.AccessKeyId
}

func (signer *SignerStsAssumeRole) GetExtraParam() map[string]string {
	return map[string]string{"SecurityToken": signer.credential.RoleArn}
}

func (signer *SignerStsAssumeRole) Sign(stringToSign, secretSuffix string) string {
	if signer.sessionCredential == nil || signer.needUpdateCredential() {
		signer.updateCredential()
	}
	secret := signer.credential.AccessKeySecret + secretSuffix
	return ShaHmac1(stringToSign, secret)
}

func (signer *SignerStsAssumeRole) buildCommonRequest() (request *requests.CommonRequest, err error) {
	request = requests.NewCommonRequest()
	request.Product = "Sts"
	request.Version = "2015-04-01"
	request.ApiName = "AssumeRole"
	request.QueryParams["RoleArn"] = signer.credential.RoleArn
	request.QueryParams["RoleSessionName"] = signer.credential.RoleSessionName
	request.QueryParams["DurationSeconds"] = strconv.Itoa(signer.credentialExpiration)
	return
}

func (signer *SignerStsAssumeRole) refreshCredential(response *responses.CommonResponse) (err error) {
	if response.GetHttpStatus() != http.StatusOK {
		message := "refresh session token failed, message = " + response.GetHttpContentString()
		err = errors.NewServerError(response.GetHttpStatus(), response.GetOriginHttpResponse().Status, message)
		return
	}
	var data interface{}
	err = json.Unmarshal(response.GetHttpContentBytes(), data)
	if err != nil {
		return
	}
	accessKeyId, err := jmespath.Search("AssumeRoleResponse.Credentials.AccessKeyId", data)
	if err != nil {
		return
	}
	accessKeySecret, err := jmespath.Search("AssumeRoleResponse.Credentials.AccessKeySecret", data)
	if err != nil {
		return
	}
	securityToken, err := jmespath.Search("AssumeRoleResponse.Credentials.SecurityToken", data)
	if err != nil {
		return
	}
	signer.sessionCredential = &sessionCredential{
		accessKeyId:     accessKeyId.(string),
		accessKeySecret: accessKeySecret.(string),
		securityToken:   securityToken.(string),
	}
	return
}

func (signer *SignerStsAssumeRole) Shutdown() {

}
