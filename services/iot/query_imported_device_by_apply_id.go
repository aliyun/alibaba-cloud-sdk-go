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

// QueryImportedDeviceByApplyId invokes the iot.QueryImportedDeviceByApplyId API synchronously
func (client *Client) QueryImportedDeviceByApplyId(request *QueryImportedDeviceByApplyIdRequest) (response *QueryImportedDeviceByApplyIdResponse, err error) {
	response = CreateQueryImportedDeviceByApplyIdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryImportedDeviceByApplyIdWithChan invokes the iot.QueryImportedDeviceByApplyId API asynchronously
func (client *Client) QueryImportedDeviceByApplyIdWithChan(request *QueryImportedDeviceByApplyIdRequest) (<-chan *QueryImportedDeviceByApplyIdResponse, <-chan error) {
	responseChan := make(chan *QueryImportedDeviceByApplyIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryImportedDeviceByApplyId(request)
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

// QueryImportedDeviceByApplyIdWithCallback invokes the iot.QueryImportedDeviceByApplyId API asynchronously
func (client *Client) QueryImportedDeviceByApplyIdWithCallback(request *QueryImportedDeviceByApplyIdRequest, callback func(response *QueryImportedDeviceByApplyIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryImportedDeviceByApplyIdResponse
		var err error
		defer close(result)
		response, err = client.QueryImportedDeviceByApplyId(request)
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

// QueryImportedDeviceByApplyIdRequest is the request struct for api QueryImportedDeviceByApplyId
type QueryImportedDeviceByApplyIdRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	ApplyId       requests.Integer `position:"Query" name:"ApplyId"`
	PageNo        requests.Integer `position:"Query" name:"PageNo"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryImportedDeviceByApplyIdResponse is the response struct for api QueryImportedDeviceByApplyId
type QueryImportedDeviceByApplyIdResponse struct {
	*responses.BaseResponse
	RequestId    string                                   `json:"RequestId" xml:"RequestId"`
	Success      bool                                     `json:"Success" xml:"Success"`
	Code         string                                   `json:"Code" xml:"Code"`
	ErrorMessage string                                   `json:"ErrorMessage" xml:"ErrorMessage"`
	PageNo       int                                      `json:"PageNo" xml:"PageNo"`
	PageSize     int                                      `json:"PageSize" xml:"PageSize"`
	TotalPage    int                                      `json:"TotalPage" xml:"TotalPage"`
	ProductKey   string                                   `json:"ProductKey" xml:"ProductKey"`
	DeviceList   DeviceListInQueryImportedDeviceByApplyId `json:"DeviceList" xml:"DeviceList"`
}

// CreateQueryImportedDeviceByApplyIdRequest creates a request to invoke QueryImportedDeviceByApplyId API
func CreateQueryImportedDeviceByApplyIdRequest() (request *QueryImportedDeviceByApplyIdRequest) {
	request = &QueryImportedDeviceByApplyIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryImportedDeviceByApplyId", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryImportedDeviceByApplyIdResponse creates a response to parse from QueryImportedDeviceByApplyId response
func CreateQueryImportedDeviceByApplyIdResponse() (response *QueryImportedDeviceByApplyIdResponse) {
	response = &QueryImportedDeviceByApplyIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
