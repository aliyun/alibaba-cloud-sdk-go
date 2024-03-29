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

// StartShiftLoadBalancerZones invokes the nlb.StartShiftLoadBalancerZones API synchronously
func (client *Client) StartShiftLoadBalancerZones(request *StartShiftLoadBalancerZonesRequest) (response *StartShiftLoadBalancerZonesResponse, err error) {
	response = CreateStartShiftLoadBalancerZonesResponse()
	err = client.DoAction(request, response)
	return
}

// StartShiftLoadBalancerZonesWithChan invokes the nlb.StartShiftLoadBalancerZones API asynchronously
func (client *Client) StartShiftLoadBalancerZonesWithChan(request *StartShiftLoadBalancerZonesRequest) (<-chan *StartShiftLoadBalancerZonesResponse, <-chan error) {
	responseChan := make(chan *StartShiftLoadBalancerZonesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartShiftLoadBalancerZones(request)
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

// StartShiftLoadBalancerZonesWithCallback invokes the nlb.StartShiftLoadBalancerZones API asynchronously
func (client *Client) StartShiftLoadBalancerZonesWithCallback(request *StartShiftLoadBalancerZonesRequest, callback func(response *StartShiftLoadBalancerZonesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartShiftLoadBalancerZonesResponse
		var err error
		defer close(result)
		response, err = client.StartShiftLoadBalancerZones(request)
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

// StartShiftLoadBalancerZonesRequest is the request struct for api StartShiftLoadBalancerZones
type StartShiftLoadBalancerZonesRequest struct {
	*requests.RpcRequest
	ClientToken    string                                     `position:"Body" name:"ClientToken"`
	DryRun         requests.Boolean                           `position:"Body" name:"DryRun"`
	ZoneMappings   *[]StartShiftLoadBalancerZonesZoneMappings `position:"Body" name:"ZoneMappings"  type:"Repeated"`
	LoadBalancerId string                                     `position:"Body" name:"LoadBalancerId"`
}

// StartShiftLoadBalancerZonesZoneMappings is a repeated param struct in StartShiftLoadBalancerZonesRequest
type StartShiftLoadBalancerZonesZoneMappings struct {
	VSwitchId string `name:"VSwitchId"`
	ZoneId    string `name:"ZoneId"`
}

// StartShiftLoadBalancerZonesResponse is the response struct for api StartShiftLoadBalancerZones
type StartShiftLoadBalancerZonesResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
}

// CreateStartShiftLoadBalancerZonesRequest creates a request to invoke StartShiftLoadBalancerZones API
func CreateStartShiftLoadBalancerZonesRequest() (request *StartShiftLoadBalancerZonesRequest) {
	request = &StartShiftLoadBalancerZonesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "StartShiftLoadBalancerZones", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartShiftLoadBalancerZonesResponse creates a response to parse from StartShiftLoadBalancerZones response
func CreateStartShiftLoadBalancerZonesResponse() (response *StartShiftLoadBalancerZonesResponse) {
	response = &StartShiftLoadBalancerZonesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
