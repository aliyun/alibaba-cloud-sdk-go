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

// QueryShareTaskDeviceList invokes the iot.QueryShareTaskDeviceList API synchronously
func (client *Client) QueryShareTaskDeviceList(request *QueryShareTaskDeviceListRequest) (response *QueryShareTaskDeviceListResponse, err error) {
	response = CreateQueryShareTaskDeviceListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryShareTaskDeviceListWithChan invokes the iot.QueryShareTaskDeviceList API asynchronously
func (client *Client) QueryShareTaskDeviceListWithChan(request *QueryShareTaskDeviceListRequest) (<-chan *QueryShareTaskDeviceListResponse, <-chan error) {
	responseChan := make(chan *QueryShareTaskDeviceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryShareTaskDeviceList(request)
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

// QueryShareTaskDeviceListWithCallback invokes the iot.QueryShareTaskDeviceList API asynchronously
func (client *Client) QueryShareTaskDeviceListWithCallback(request *QueryShareTaskDeviceListRequest, callback func(response *QueryShareTaskDeviceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryShareTaskDeviceListResponse
		var err error
		defer close(result)
		response, err = client.QueryShareTaskDeviceList(request)
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

// QueryShareTaskDeviceListRequest is the request struct for api QueryShareTaskDeviceList
type QueryShareTaskDeviceListRequest struct {
	*requests.RpcRequest
	PageId        requests.Integer `position:"Body" name:"PageId"`
	IotInstanceId string           `position:"Body" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	ShareTaskId   string           `position:"Body" name:"ShareTaskId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Body" name:"DeviceName"`
}

// QueryShareTaskDeviceListResponse is the response struct for api QueryShareTaskDeviceList
type QueryShareTaskDeviceListResponse struct {
	*responses.BaseResponse
	RequestId    string                         `json:"RequestId" xml:"RequestId"`
	Success      bool                           `json:"Success" xml:"Success"`
	Code         string                         `json:"Code" xml:"Code"`
	ErrorMessage string                         `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryShareTaskDeviceList `json:"Data" xml:"Data"`
}

// CreateQueryShareTaskDeviceListRequest creates a request to invoke QueryShareTaskDeviceList API
func CreateQueryShareTaskDeviceListRequest() (request *QueryShareTaskDeviceListRequest) {
	request = &QueryShareTaskDeviceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryShareTaskDeviceList", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryShareTaskDeviceListResponse creates a response to parse from QueryShareTaskDeviceList response
func CreateQueryShareTaskDeviceListResponse() (response *QueryShareTaskDeviceListResponse) {
	response = &QueryShareTaskDeviceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
