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

package auth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	"net/url"
	"sort"
	"strings"
)

func signRpcRequest(request *requests.RpcRequest, signer Signer, regionId string) {
	completeRpcSignParams(request, signer, regionId)
	stringToSign := buildRpcStringToSign(request)
	signature := signer.Sign(stringToSign, "&")
	request.QueryParams["Signature"] = signature
}

func completeRpcSignParams(request *requests.RpcRequest, signer Signer, regionId string) {
	queryParams := request.QueryParams
	queryParams["Version"] = request.GetVersion()
	queryParams["Action"] = request.GetActionName()
	queryParams["Format"] = request.GetAcceptFormat()
	queryParams["Timestamp"] = utils.GetTimeInFormatISO8601()
	queryParams["SignatureMethod"] = signer.GetName()
	queryParams["SignatureType"] = signer.GetType()
	queryParams["SignatureVersion"] = signer.GetVersion()
	queryParams["SignatureNonce"] = utils.GetUUIDV4()
	queryParams["AccessKeyId"] = signer.GetAccessKeyId()

	if _, contains := queryParams["RegionId"]; !contains {
		queryParams["RegionId"] = regionId
	}
	if extraParam := signer.GetExtraParam(); extraParam != nil {
		for key, value := range extraParam {
			queryParams[key] = value
		}
	}

	request.Headers["Content-Type"] = requests.Form
	formString := utils.GetUrlFormedMap(request.FormParams)
	request.Content = []byte(formString)
}

func buildRpcStringToSign(request *requests.RpcRequest) (stringToSign string) {
	signParams := make(map[string]string)
	for key, value := range request.QueryParams {
		signParams[key] = value
	}
	for key, value := range request.FormParams {
		signParams[key] = value
	}

	// sort params by key
	var paramKeySlice []string
	for key := range signParams {
		paramKeySlice = append(paramKeySlice, key)
	}
	sort.Strings(paramKeySlice)
	stringToSign = utils.GetUrlFormedMap(signParams)
	stringToSign = strings.Replace(stringToSign, "+", "%20", -1)
	stringToSign = strings.Replace(stringToSign, "*", "%2A", -1)
	stringToSign = strings.Replace(stringToSign, "%7E", "~", -1)
	stringToSign = url.QueryEscape(stringToSign)
	stringToSign = request.Method + "&%2F&" + stringToSign
	return
}
