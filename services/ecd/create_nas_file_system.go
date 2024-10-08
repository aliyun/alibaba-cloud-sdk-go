package ecd

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

// CreateNASFileSystem invokes the ecd.CreateNASFileSystem API synchronously
func (client *Client) CreateNASFileSystem(request *CreateNASFileSystemRequest) (response *CreateNASFileSystemResponse, err error) {
	response = CreateCreateNASFileSystemResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNASFileSystemWithChan invokes the ecd.CreateNASFileSystem API asynchronously
func (client *Client) CreateNASFileSystemWithChan(request *CreateNASFileSystemRequest) (<-chan *CreateNASFileSystemResponse, <-chan error) {
	responseChan := make(chan *CreateNASFileSystemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNASFileSystem(request)
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

// CreateNASFileSystemWithCallback invokes the ecd.CreateNASFileSystem API asynchronously
func (client *Client) CreateNASFileSystemWithCallback(request *CreateNASFileSystemRequest, callback func(response *CreateNASFileSystemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNASFileSystemResponse
		var err error
		defer close(result)
		response, err = client.CreateNASFileSystem(request)
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

// CreateNASFileSystemRequest is the request struct for api CreateNASFileSystem
type CreateNASFileSystemRequest struct {
	*requests.RpcRequest
	OfficeSiteId string `position:"Query" name:"OfficeSiteId"`
	Description  string `position:"Query" name:"Description"`
	StorageType  string `position:"Query" name:"StorageType"`
	EncryptType  string `position:"Query" name:"EncryptType"`
	Name         string `position:"Query" name:"Name"`
}

// CreateNASFileSystemResponse is the response struct for api CreateNASFileSystem
type CreateNASFileSystemResponse struct {
	*responses.BaseResponse
	FileSystemId      string `json:"FileSystemId" xml:"FileSystemId"`
	FileSystemName    string `json:"FileSystemName" xml:"FileSystemName"`
	MountTargetDomain string `json:"MountTargetDomain" xml:"MountTargetDomain"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
	OfficeSiteId      string `json:"OfficeSiteId" xml:"OfficeSiteId"`
}

// CreateCreateNASFileSystemRequest creates a request to invoke CreateNASFileSystem API
func CreateCreateNASFileSystemRequest() (request *CreateNASFileSystemRequest) {
	request = &CreateNASFileSystemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreateNASFileSystem", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNASFileSystemResponse creates a response to parse from CreateNASFileSystem response
func CreateCreateNASFileSystemResponse() (response *CreateNASFileSystemResponse) {
	response = &CreateNASFileSystemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
