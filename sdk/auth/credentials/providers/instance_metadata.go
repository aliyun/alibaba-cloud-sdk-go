package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
)

type ecsRAMRoleCredentials struct {
	SecurityToken   *string `json:"SecurityToken"`
	Expiration      *string `json:"Expiration"`
	AccessKeySecret *string `json:"AccessKeySecret"`
	AccessKeyId     *string `json:"AccessKeyId"`
	LastUpdated     *string `json:"LastUpdated"`
	Code            *string `json:"Code"`
}

var securityCredURL = "http://100.100.100.200/latest/meta-data/ram/security-credentials/"
var tokenURL = "http://100.100.100.200/latest/api/token"
var defaultMetadataTokenDuration = 21600 // 6 hours

type do func(req *http.Request) (*http.Response, error)

var hookDo = func(fn do) do {
	return fn
}

type newReuqest func(method, url string, body io.Reader) (*http.Request, error)

var hookNewRequest = func(fn newReuqest) newReuqest {
	return fn
}

func NewInstanceMetadataProvider() Provider {
	return &InstanceMetadataProvider{}
}

type InstanceMetadataProvider struct {
	RoleName      string
	DisableIMDSv1 bool
}

func (p *InstanceMetadataProvider) getMetadataToken() (metadataToken string, err error) {
	// PUT http://100.100.100.200/latest/api/token
	httpRequest, _err := hookNewRequest(http.NewRequest)("PUT", tokenURL, strings.NewReader(""))
	if _err != nil {
		if p.DisableIMDSv1 {
			err = fmt.Errorf("get metadata token failed: %s", _err.Error())
		}
		return
	}
	httpRequest.Header.Set("X-aliyun-ecs-metadata-token-ttl-seconds", strconv.Itoa(defaultMetadataTokenDuration))

	httpClient := &http.Client{}

	httpResponse, _err := hookDo(httpClient.Do)(httpRequest)
	if _err != nil {
		if p.DisableIMDSv1 {
			err = fmt.Errorf("get metadata token failed: %s", _err.Error())
		}
		return
	}

	defer httpResponse.Body.Close()

	responseBody, _err := ioutil.ReadAll(httpResponse.Body)
	if _err != nil {
		if p.DisableIMDSv1 {
			err = fmt.Errorf("get metadata token failed: %s", _err.Error())
		}
		return
	}

	if httpResponse.StatusCode != http.StatusOK {
		if p.DisableIMDSv1 {
			err = fmt.Errorf("received %d getting metadata token: %s", httpResponse.StatusCode, string(responseBody))
		}
		return
	}

	metadataToken = strings.TrimSpace(string(responseBody))
	return
}

func (p *InstanceMetadataProvider) GetRoleName() (roleName string, err error) {
	// Instances can have only one role name that never changes,
	// so attempt to populate it.
	// If this call is executed in an environment that doesn't support instance metadata,
	// it will time out after 30 seconds and return an err.
	httpRequest, err := hookNewRequest(http.NewRequest)("GET", securityCredURL, strings.NewReader(""))
	if err != nil {
		err = fmt.Errorf("get role name failed: %s", err.Error())
		return
	}

	metadataToken, err := p.getMetadataToken()
	if err != nil {
		return
	}
	if metadataToken != "" {
		httpRequest.Header.Set("X-aliyun-ecs-metadata-token", metadataToken)
	}

	httpClient := &http.Client{}

	httpResponse, err := hookDo(httpClient.Do)(httpRequest)
	if err != nil {
		err = fmt.Errorf("get role name failed: %s", err.Error())
		return
	}

	defer httpResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}

	roleName = strings.TrimSpace(string(responseBody))

	if httpResponse.StatusCode != http.StatusOK {
		err = fmt.Errorf("received %d getting role name: %s", httpResponse.StatusCode, roleName)
		return
	}

	if roleName == "" {
		err = fmt.Errorf("unable to retrieve role name, it may be unset")
		return
	}

	return
}

func (p *InstanceMetadataProvider) Retrieve() (auth.Credential, error) {
	if p.RoleName == "" {
		roleName, err := p.GetRoleName()
		if err != nil {
			return nil, err
		}
		p.RoleName = roleName
	}

	httpRequest, err := hookNewRequest(http.NewRequest)("GET", securityCredURL+p.RoleName, strings.NewReader(""))
	if err != nil {
		err = fmt.Errorf("refresh metadata token err: %s", err.Error())
		return nil, err
	}

	metadataToken, err := p.getMetadataToken()
	if err != nil {
		return nil, err
	}
	if metadataToken != "" {
		httpRequest.Header.Set("X-aliyun-ecs-metadata-token", metadataToken)
	}

	httpClient := &http.Client{}

	httpResponse, err := hookDo(httpClient.Do)(httpRequest)
	if err != nil {
		err = fmt.Errorf("refresh metadata token err: %s", err.Error())
		return nil, err
	}

	defer httpResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received %d getting security credentials for %s", httpResponse.StatusCode, p.RoleName)
	}

	var data ecsRAMRoleCredentials
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return nil, fmt.Errorf("refresh metadata err, json.Unmarshal fail: %s", err.Error())
	}

	if data.AccessKeyId == nil || data.AccessKeySecret == nil || data.SecurityToken == nil {
		return nil, fmt.Errorf("refresh metadata err, fail to get credentials, response: %s", string(responseBody))
	}

	if *data.Code != "Success" {
		return nil, fmt.Errorf("refresh metadata err, Code is not Success, response: %s", string(responseBody))
	}

	return credentials.NewStsTokenCredential(*data.AccessKeyId, *data.AccessKeySecret, *data.SecurityToken), nil
}
