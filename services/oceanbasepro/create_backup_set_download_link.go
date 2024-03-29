package oceanbasepro

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

// CreateBackupSetDownloadLink invokes the oceanbasepro.CreateBackupSetDownloadLink API synchronously
func (client *Client) CreateBackupSetDownloadLink(request *CreateBackupSetDownloadLinkRequest) (response *CreateBackupSetDownloadLinkResponse, err error) {
	response = CreateCreateBackupSetDownloadLinkResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBackupSetDownloadLinkWithChan invokes the oceanbasepro.CreateBackupSetDownloadLink API asynchronously
func (client *Client) CreateBackupSetDownloadLinkWithChan(request *CreateBackupSetDownloadLinkRequest) (<-chan *CreateBackupSetDownloadLinkResponse, <-chan error) {
	responseChan := make(chan *CreateBackupSetDownloadLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBackupSetDownloadLink(request)
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

// CreateBackupSetDownloadLinkWithCallback invokes the oceanbasepro.CreateBackupSetDownloadLink API asynchronously
func (client *Client) CreateBackupSetDownloadLinkWithCallback(request *CreateBackupSetDownloadLinkRequest, callback func(response *CreateBackupSetDownloadLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBackupSetDownloadLinkResponse
		var err error
		defer close(result)
		response, err = client.CreateBackupSetDownloadLink(request)
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

// CreateBackupSetDownloadLinkRequest is the request struct for api CreateBackupSetDownloadLink
type CreateBackupSetDownloadLinkRequest struct {
	*requests.RpcRequest
	BackupSetId string `position:"Body" name:"BackupSetId"`
	InstanceId  string `position:"Body" name:"InstanceId"`
}

// CreateBackupSetDownloadLinkResponse is the response struct for api CreateBackupSetDownloadLink
type CreateBackupSetDownloadLinkResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DownloadTaskId int64  `json:"DownloadTaskId" xml:"DownloadTaskId"`
}

// CreateCreateBackupSetDownloadLinkRequest creates a request to invoke CreateBackupSetDownloadLink API
func CreateCreateBackupSetDownloadLinkRequest() (request *CreateBackupSetDownloadLinkRequest) {
	request = &CreateBackupSetDownloadLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "CreateBackupSetDownloadLink", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBackupSetDownloadLinkResponse creates a response to parse from CreateBackupSetDownloadLink response
func CreateCreateBackupSetDownloadLinkResponse() (response *CreateBackupSetDownloadLinkResponse) {
	response = &CreateBackupSetDownloadLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
