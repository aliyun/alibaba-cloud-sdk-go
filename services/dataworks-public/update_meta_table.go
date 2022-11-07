package dataworks_public

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

// UpdateMetaTable invokes the dataworks_public.UpdateMetaTable API synchronously
func (client *Client) UpdateMetaTable(request *UpdateMetaTableRequest) (response *UpdateMetaTableResponse, err error) {
	response = CreateUpdateMetaTableResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMetaTableWithChan invokes the dataworks_public.UpdateMetaTable API asynchronously
func (client *Client) UpdateMetaTableWithChan(request *UpdateMetaTableRequest) (<-chan *UpdateMetaTableResponse, <-chan error) {
	responseChan := make(chan *UpdateMetaTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMetaTable(request)
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

// UpdateMetaTableWithCallback invokes the dataworks_public.UpdateMetaTable API asynchronously
func (client *Client) UpdateMetaTableWithCallback(request *UpdateMetaTableRequest, callback func(response *UpdateMetaTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMetaTableResponse
		var err error
		defer close(result)
		response, err = client.UpdateMetaTable(request)
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

// UpdateMetaTableRequest is the request struct for api UpdateMetaTable
type UpdateMetaTableRequest struct {
	*requests.RpcRequest
	Schema        string           `position:"Query" name:"Schema"`
	Visibility    requests.Integer `position:"Query" name:"Visibility"`
	Caption       string           `position:"Query" name:"Caption"`
	NewOwnerId    string           `position:"Query" name:"NewOwnerId"`
	TableGuid     string           `position:"Query" name:"TableGuid"`
	AddedLabels   string           `position:"Body" name:"AddedLabels"`
	RemovedLabels string           `position:"Body" name:"RemovedLabels"`
	EnvType       requests.Integer `position:"Query" name:"EnvType"`
	TableName     string           `position:"Query" name:"TableName"`
	ProjectId     requests.Integer `position:"Query" name:"ProjectId"`
	CategoryId    requests.Integer `position:"Query" name:"CategoryId"`
}

// UpdateMetaTableResponse is the response struct for api UpdateMetaTable
type UpdateMetaTableResponse struct {
	*responses.BaseResponse
	UpdateResult bool   `json:"UpdateResult" xml:"UpdateResult"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateMetaTableRequest creates a request to invoke UpdateMetaTable API
func CreateUpdateMetaTableRequest() (request *UpdateMetaTableRequest) {
	request = &UpdateMetaTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateMetaTable", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateMetaTableResponse creates a response to parse from UpdateMetaTable response
func CreateUpdateMetaTableResponse() (response *UpdateMetaTableResponse) {
	response = &UpdateMetaTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
