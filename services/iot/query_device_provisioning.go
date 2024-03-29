package iot

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

// QueryDeviceProvisioning invokes the iot.QueryDeviceProvisioning API synchronously
func (client *Client) QueryDeviceProvisioning(request *QueryDeviceProvisioningRequest) (response *QueryDeviceProvisioningResponse, err error) {
	response = CreateQueryDeviceProvisioningResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceProvisioningWithChan invokes the iot.QueryDeviceProvisioning API asynchronously
func (client *Client) QueryDeviceProvisioningWithChan(request *QueryDeviceProvisioningRequest) (<-chan *QueryDeviceProvisioningResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceProvisioningResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceProvisioning(request)
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

// QueryDeviceProvisioningWithCallback invokes the iot.QueryDeviceProvisioning API asynchronously
func (client *Client) QueryDeviceProvisioningWithCallback(request *QueryDeviceProvisioningRequest, callback func(response *QueryDeviceProvisioningResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceProvisioningResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceProvisioning(request)
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

// QueryDeviceProvisioningRequest is the request struct for api QueryDeviceProvisioning
type QueryDeviceProvisioningRequest struct {
	*requests.RpcRequest
	ProductKey  string `position:"Body" name:"ProductKey"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
	DeviceName  string `position:"Query" name:"DeviceName"`
}

// QueryDeviceProvisioningResponse is the response struct for api QueryDeviceProvisioning
type QueryDeviceProvisioningResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	Success      bool                          `json:"Success" xml:"Success"`
	Code         string                        `json:"Code" xml:"Code"`
	ErrorMessage string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDeviceProvisioning `json:"Data" xml:"Data"`
}

// CreateQueryDeviceProvisioningRequest creates a request to invoke QueryDeviceProvisioning API
func CreateQueryDeviceProvisioningRequest() (request *QueryDeviceProvisioningRequest) {
	request = &QueryDeviceProvisioningRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceProvisioning", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceProvisioningResponse creates a response to parse from QueryDeviceProvisioning response
func CreateQueryDeviceProvisioningResponse() (response *QueryDeviceProvisioningResponse) {
	response = &QueryDeviceProvisioningResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
