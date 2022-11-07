package dms_enterprise

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

// UpdateTaskFlowRelations invokes the dms_enterprise.UpdateTaskFlowRelations API synchronously
func (client *Client) UpdateTaskFlowRelations(request *UpdateTaskFlowRelationsRequest) (response *UpdateTaskFlowRelationsResponse, err error) {
	response = CreateUpdateTaskFlowRelationsResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTaskFlowRelationsWithChan invokes the dms_enterprise.UpdateTaskFlowRelations API asynchronously
func (client *Client) UpdateTaskFlowRelationsWithChan(request *UpdateTaskFlowRelationsRequest) (<-chan *UpdateTaskFlowRelationsResponse, <-chan error) {
	responseChan := make(chan *UpdateTaskFlowRelationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTaskFlowRelations(request)
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

// UpdateTaskFlowRelationsWithCallback invokes the dms_enterprise.UpdateTaskFlowRelations API asynchronously
func (client *Client) UpdateTaskFlowRelationsWithCallback(request *UpdateTaskFlowRelationsRequest, callback func(response *UpdateTaskFlowRelationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTaskFlowRelationsResponse
		var err error
		defer close(result)
		response, err = client.UpdateTaskFlowRelations(request)
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

// UpdateTaskFlowRelationsRequest is the request struct for api UpdateTaskFlowRelations
type UpdateTaskFlowRelationsRequest struct {
	*requests.RpcRequest
	Edges *[]UpdateTaskFlowRelationsEdges `position:"Query" name:"Edges"  type:"Json"`
	DagId requests.Integer                `position:"Query" name:"DagId"`
	Tid   requests.Integer                `position:"Query" name:"Tid"`
}

// UpdateTaskFlowRelationsEdges is a repeated param struct in UpdateTaskFlowRelationsRequest
type UpdateTaskFlowRelationsEdges struct {
	NodeEnd  string `name:"NodeEnd"`
	NodeFrom string `name:"NodeFrom"`
	EdgeType string `name:"EdgeType"`
	Id       string `name:"Id"`
}

// UpdateTaskFlowRelationsResponse is the response struct for api UpdateTaskFlowRelations
type UpdateTaskFlowRelationsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateUpdateTaskFlowRelationsRequest creates a request to invoke UpdateTaskFlowRelations API
func CreateUpdateTaskFlowRelationsRequest() (request *UpdateTaskFlowRelationsRequest) {
	request = &UpdateTaskFlowRelationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "UpdateTaskFlowRelations", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateTaskFlowRelationsResponse creates a response to parse from UpdateTaskFlowRelations response
func CreateUpdateTaskFlowRelationsResponse() (response *UpdateTaskFlowRelationsResponse) {
	response = &UpdateTaskFlowRelationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
