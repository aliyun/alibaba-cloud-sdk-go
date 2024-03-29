package oms

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

// DeleteMeasureData invokes the oms.DeleteMeasureData API synchronously
func (client *Client) DeleteMeasureData(request *DeleteMeasureDataRequest) (response *DeleteMeasureDataResponse, err error) {
	response = CreateDeleteMeasureDataResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMeasureDataWithChan invokes the oms.DeleteMeasureData API asynchronously
func (client *Client) DeleteMeasureDataWithChan(request *DeleteMeasureDataRequest) (<-chan *DeleteMeasureDataResponse, <-chan error) {
	responseChan := make(chan *DeleteMeasureDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMeasureData(request)
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

// DeleteMeasureDataWithCallback invokes the oms.DeleteMeasureData API asynchronously
func (client *Client) DeleteMeasureDataWithCallback(request *DeleteMeasureDataRequest, callback func(response *DeleteMeasureDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMeasureDataResponse
		var err error
		defer close(result)
		response, err = client.DeleteMeasureData(request)
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

// DeleteMeasureDataRequest is the request struct for api DeleteMeasureData
type DeleteMeasureDataRequest struct {
	*requests.RpcRequest
	Filter     string           `position:"Query" name:"Filter"`
	DomainCode string           `position:"Query" name:"DomainCode"`
	Data       string           `position:"Query" name:"Data"`
	DataType   string           `position:"Query" name:"DataType"`
	Compressed requests.Boolean `position:"Query" name:"Compressed"`
	ApiType    string           `position:"Query" name:"ApiType"`
}

// DeleteMeasureDataResponse is the response struct for api DeleteMeasureData
type DeleteMeasureDataResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainCode string `json:"DomainCode" xml:"DomainCode"`
	DataType   string `json:"DataType" xml:"DataType"`
	ApiType    string `json:"ApiType" xml:"ApiType"`
	Total      int64  `json:"Total" xml:"Total"`
}

// CreateDeleteMeasureDataRequest creates a request to invoke DeleteMeasureData API
func CreateDeleteMeasureDataRequest() (request *DeleteMeasureDataRequest) {
	request = &DeleteMeasureDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oms", "2016-06-15", "DeleteMeasureData", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMeasureDataResponse creates a response to parse from DeleteMeasureData response
func CreateDeleteMeasureDataResponse() (response *DeleteMeasureDataResponse) {
	response = &DeleteMeasureDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
