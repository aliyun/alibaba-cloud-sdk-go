package dysmsapi

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

// ConversionData invokes the dysmsapi.ConversionData API synchronously
func (client *Client) ConversionData(request *ConversionDataRequest) (response *ConversionDataResponse, err error) {
	response = CreateConversionDataResponse()
	err = client.DoAction(request, response)
	return
}

// ConversionDataWithChan invokes the dysmsapi.ConversionData API asynchronously
func (client *Client) ConversionDataWithChan(request *ConversionDataRequest) (<-chan *ConversionDataResponse, <-chan error) {
	responseChan := make(chan *ConversionDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConversionData(request)
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

// ConversionDataWithCallback invokes the dysmsapi.ConversionData API asynchronously
func (client *Client) ConversionDataWithCallback(request *ConversionDataRequest, callback func(response *ConversionDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConversionDataResponse
		var err error
		defer close(result)
		response, err = client.ConversionData(request)
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

// ConversionDataRequest is the request struct for api ConversionData
type ConversionDataRequest struct {
	*requests.RpcRequest
	ReportTime     requests.Integer `position:"Body" name:"ReportTime"`
	ConversionRate string           `position:"Body" name:"ConversionRate"`
}

// ConversionDataResponse is the response struct for api ConversionData
type ConversionDataResponse struct {
	*responses.BaseResponse
	ResponseCode        string `json:"ResponseCode" xml:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription" xml:"ResponseDescription"`
	RequestId           string `json:"RequestId" xml:"RequestId"`
}

// CreateConversionDataRequest creates a request to invoke ConversionData API
func CreateConversionDataRequest() (request *ConversionDataRequest) {
	request = &ConversionDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2018-05-01", "ConversionData", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConversionDataResponse creates a response to parse from ConversionData response
func CreateConversionDataResponse() (response *ConversionDataResponse) {
	response = &ConversionDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
