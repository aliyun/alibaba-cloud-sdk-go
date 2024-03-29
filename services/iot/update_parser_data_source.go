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

// UpdateParserDataSource invokes the iot.UpdateParserDataSource API synchronously
func (client *Client) UpdateParserDataSource(request *UpdateParserDataSourceRequest) (response *UpdateParserDataSourceResponse, err error) {
	response = CreateUpdateParserDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateParserDataSourceWithChan invokes the iot.UpdateParserDataSource API asynchronously
func (client *Client) UpdateParserDataSourceWithChan(request *UpdateParserDataSourceRequest) (<-chan *UpdateParserDataSourceResponse, <-chan error) {
	responseChan := make(chan *UpdateParserDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateParserDataSource(request)
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

// UpdateParserDataSourceWithCallback invokes the iot.UpdateParserDataSource API asynchronously
func (client *Client) UpdateParserDataSourceWithCallback(request *UpdateParserDataSourceRequest, callback func(response *UpdateParserDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateParserDataSourceResponse
		var err error
		defer close(result)
		response, err = client.UpdateParserDataSource(request)
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

// UpdateParserDataSourceRequest is the request struct for api UpdateParserDataSource
type UpdateParserDataSourceRequest struct {
	*requests.RpcRequest
	Description   string           `position:"Query" name:"Description"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	DataSourceId  requests.Integer `position:"Query" name:"DataSourceId"`
	Name          string           `position:"Query" name:"Name"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// UpdateParserDataSourceResponse is the response struct for api UpdateParserDataSource
type UpdateParserDataSourceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateParserDataSourceRequest creates a request to invoke UpdateParserDataSource API
func CreateUpdateParserDataSourceRequest() (request *UpdateParserDataSourceRequest) {
	request = &UpdateParserDataSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateParserDataSource", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateParserDataSourceResponse creates a response to parse from UpdateParserDataSource response
func CreateUpdateParserDataSourceResponse() (response *UpdateParserDataSourceResponse) {
	response = &UpdateParserDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
