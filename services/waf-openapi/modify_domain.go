package waf_openapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyDomain invokes the waf_openapi.ModifyDomain API synchronously
func (client *Client) ModifyDomain(request *ModifyDomainRequest) (response *ModifyDomainResponse, err error) {
	response = CreateModifyDomainResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDomainWithChan invokes the waf_openapi.ModifyDomain API asynchronously
func (client *Client) ModifyDomainWithChan(request *ModifyDomainRequest) (<-chan *ModifyDomainResponse, <-chan error) {
	responseChan := make(chan *ModifyDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDomain(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyDomainWithCallback invokes the waf_openapi.ModifyDomain API asynchronously
func (client *Client) ModifyDomainWithCallback(request *ModifyDomainRequest, callback func(response *ModifyDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDomainResponse
		var err error
		defer close(result)
		response, err = client.ModifyDomain(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyDomainRequest is the request struct for api ModifyDomain
type ModifyDomainRequest struct {
	*requests.RpcRequest
	IpFollowStatus       requests.Integer `position:"Query" name:"IpFollowStatus"`
	Keepalive            requests.Boolean `position:"Query" name:"Keepalive"`
	SniHost              string           `position:"Query" name:"SniHost"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	HybridCloudType      requests.Integer `position:"Query" name:"HybridCloudType"`
	SourceIp             string           `position:"Query" name:"SourceIp"`
	HttpPort             string           `position:"Query" name:"HttpPort"`
	Http2Port            string           `position:"Query" name:"Http2Port"`
	WriteTime            requests.Integer `position:"Query" name:"WriteTime"`
	AccessHeaderMode     requests.Integer `position:"Query" name:"AccessHeaderMode"`
	AccessHeaders        string           `position:"Query" name:"AccessHeaders"`
	KeepaliveTimeout     requests.Integer `position:"Query" name:"KeepaliveTimeout"`
	ClusterType          requests.Integer `position:"Query" name:"ClusterType"`
	HttpsRedirect        requests.Integer `position:"Query" name:"HttpsRedirect"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Domain               string           `position:"Query" name:"Domain"`
	ReadTime             requests.Integer `position:"Query" name:"ReadTime"`
	HttpsPort            string           `position:"Query" name:"HttpsPort"`
	SniStatus            requests.Integer `position:"Query" name:"SniStatus"`
	Lang                 string           `position:"Query" name:"Lang"`
	Retry                requests.Boolean `position:"Query" name:"Retry"`
	KeepaliveRequests    requests.Integer `position:"Query" name:"KeepaliveRequests"`
	AccessType           string           `position:"Query" name:"AccessType"`
	BindingIpv6          requests.Integer `position:"Query" name:"BindingIpv6"`
	LogHeaders           string           `position:"Query" name:"LogHeaders"`
	ConnectionTime       requests.Integer `position:"Query" name:"ConnectionTime"`
	CloudNativeInstances string           `position:"Query" name:"CloudNativeInstances"`
	SourceIps            string           `position:"Query" name:"SourceIps"`
	IsAccessProduct      requests.Integer `position:"Query" name:"IsAccessProduct"`
	LoadBalancing        requests.Integer `position:"Query" name:"LoadBalancing"`
	HttpToUserIp         requests.Integer `position:"Query" name:"HttpToUserIp"`
}

// ModifyDomainResponse is the response struct for api ModifyDomain
type ModifyDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDomainRequest creates a request to invoke ModifyDomain API
func CreateModifyDomainRequest() (request *ModifyDomainRequest) {
	request = &ModifyDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "ModifyDomain", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDomainResponse creates a response to parse from ModifyDomain response
func CreateModifyDomainResponse() (response *ModifyDomainResponse) {
	response = &ModifyDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
