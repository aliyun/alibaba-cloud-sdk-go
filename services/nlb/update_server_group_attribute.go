package nlb

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

// UpdateServerGroupAttribute invokes the nlb.UpdateServerGroupAttribute API synchronously
func (client *Client) UpdateServerGroupAttribute(request *UpdateServerGroupAttributeRequest) (response *UpdateServerGroupAttributeResponse, err error) {
	response = CreateUpdateServerGroupAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateServerGroupAttributeWithChan invokes the nlb.UpdateServerGroupAttribute API asynchronously
func (client *Client) UpdateServerGroupAttributeWithChan(request *UpdateServerGroupAttributeRequest) (<-chan *UpdateServerGroupAttributeResponse, <-chan error) {
	responseChan := make(chan *UpdateServerGroupAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateServerGroupAttribute(request)
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

// UpdateServerGroupAttributeWithCallback invokes the nlb.UpdateServerGroupAttribute API asynchronously
func (client *Client) UpdateServerGroupAttributeWithCallback(request *UpdateServerGroupAttributeRequest, callback func(response *UpdateServerGroupAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateServerGroupAttributeResponse
		var err error
		defer close(result)
		response, err = client.UpdateServerGroupAttribute(request)
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

// UpdateServerGroupAttributeRequest is the request struct for api UpdateServerGroupAttribute
type UpdateServerGroupAttributeRequest struct {
	*requests.RpcRequest
	QuicVersion             string                                      `position:"Body" name:"QuicVersion"`
	ServerGroupName         string                                      `position:"Body" name:"ServerGroupName"`
	ClientToken             string                                      `position:"Body" name:"ClientToken"`
	PreserveClientIpEnabled requests.Boolean                            `position:"Body" name:"PreserveClientIpEnabled"`
	HealthCheckConfig       UpdateServerGroupAttributeHealthCheckConfig `position:"Body" name:"HealthCheckConfig"  type:"Struct"`
	ServerGroupId           string                                      `position:"Body" name:"ServerGroupId"`
	Scheduler               string                                      `position:"Body" name:"Scheduler"`
	PersistenceEnabled      requests.Boolean                            `position:"Body" name:"PersistenceEnabled"`
	PersistenceTimeout      requests.Integer                            `position:"Body" name:"PersistenceTimeout"`
	DryRun                  requests.Boolean                            `position:"Body" name:"DryRun"`
	ConnectionDrainEnabled  requests.Boolean                            `position:"Body" name:"ConnectionDrainEnabled"`
	ConnectionDrainTimeout  requests.Integer                            `position:"Body" name:"ConnectionDrainTimeout"`
}

// UpdateServerGroupAttributeHealthCheckConfig is a repeated param struct in UpdateServerGroupAttributeRequest
type UpdateServerGroupAttributeHealthCheckConfig struct {
	HealthCheckEnabled        string    `name:"HealthCheckEnabled"`
	HealthCheckType           string    `name:"HealthCheckType"`
	HealthCheckConnectPort    string    `name:"HealthCheckConnectPort"`
	HealthyThreshold          string    `name:"HealthyThreshold"`
	UnhealthyThreshold        string    `name:"UnhealthyThreshold"`
	HealthCheckConnectTimeout string    `name:"HealthCheckConnectTimeout"`
	HealthCheckInterval       string    `name:"HealthCheckInterval"`
	HealthCheckDomain         string    `name:"HealthCheckDomain"`
	HealthCheckUrl            string    `name:"HealthCheckUrl"`
	HealthCheckHttpCode       *[]string `name:"HealthCheckHttpCode" type:"Repeated"`
	HttpCheckMethod           string    `name:"HttpCheckMethod"`
}

// UpdateServerGroupAttributeResponse is the response struct for api UpdateServerGroupAttribute
type UpdateServerGroupAttributeResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	ServerGroupId  string `json:"ServerGroupId" xml:"ServerGroupId"`
	JobId          string `json:"JobId" xml:"JobId"`
}

// CreateUpdateServerGroupAttributeRequest creates a request to invoke UpdateServerGroupAttribute API
func CreateUpdateServerGroupAttributeRequest() (request *UpdateServerGroupAttributeRequest) {
	request = &UpdateServerGroupAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "UpdateServerGroupAttribute", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateServerGroupAttributeResponse creates a response to parse from UpdateServerGroupAttribute response
func CreateUpdateServerGroupAttributeResponse() (response *UpdateServerGroupAttributeResponse) {
	response = &UpdateServerGroupAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
