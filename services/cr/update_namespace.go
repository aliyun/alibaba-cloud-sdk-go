package cr

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

// UpdateNamespace invokes the cr.UpdateNamespace API synchronously
// api document: https://help.aliyun.com/api/cr/updatenamespace.html
func (client *Client) UpdateNamespace(request *UpdateNamespaceRequest) (response *UpdateNamespaceResponse, err error) {
	response = CreateUpdateNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateNamespaceWithChan invokes the cr.UpdateNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/updatenamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateNamespaceWithChan(request *UpdateNamespaceRequest) (<-chan *UpdateNamespaceResponse, <-chan error) {
	responseChan := make(chan *UpdateNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateNamespace(request)
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

// UpdateNamespaceWithCallback invokes the cr.UpdateNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/updatenamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateNamespaceWithCallback(request *UpdateNamespaceRequest, callback func(response *UpdateNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateNamespaceResponse
		var err error
		defer close(result)
		response, err = client.UpdateNamespace(request)
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

// UpdateNamespaceRequest is the request struct for api UpdateNamespace
type UpdateNamespaceRequest struct {
	*requests.RoaRequest
	Namespace string `position:"Path" name:"Namespace"`
}

// UpdateNamespaceResponse is the response struct for api UpdateNamespace
type UpdateNamespaceResponse struct {
	*responses.BaseResponse
}

// CreateUpdateNamespaceRequest creates a request to invoke UpdateNamespace API
func CreateUpdateNamespaceRequest() (request *UpdateNamespaceRequest) {
	request = &UpdateNamespaceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "UpdateNamespace", "/namespace/[Namespace]", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateNamespaceResponse creates a response to parse from UpdateNamespace response
func CreateUpdateNamespaceResponse() (response *UpdateNamespaceResponse) {
	response = &UpdateNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
