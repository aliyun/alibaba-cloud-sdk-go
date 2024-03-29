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

// DescribeBackupSetDownloadLink invokes the oceanbasepro.DescribeBackupSetDownloadLink API synchronously
func (client *Client) DescribeBackupSetDownloadLink(request *DescribeBackupSetDownloadLinkRequest) (response *DescribeBackupSetDownloadLinkResponse, err error) {
	response = CreateDescribeBackupSetDownloadLinkResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupSetDownloadLinkWithChan invokes the oceanbasepro.DescribeBackupSetDownloadLink API asynchronously
func (client *Client) DescribeBackupSetDownloadLinkWithChan(request *DescribeBackupSetDownloadLinkRequest) (<-chan *DescribeBackupSetDownloadLinkResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupSetDownloadLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupSetDownloadLink(request)
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

// DescribeBackupSetDownloadLinkWithCallback invokes the oceanbasepro.DescribeBackupSetDownloadLink API asynchronously
func (client *Client) DescribeBackupSetDownloadLinkWithCallback(request *DescribeBackupSetDownloadLinkRequest, callback func(response *DescribeBackupSetDownloadLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupSetDownloadLinkResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupSetDownloadLink(request)
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

// DescribeBackupSetDownloadLinkRequest is the request struct for api DescribeBackupSetDownloadLink
type DescribeBackupSetDownloadLinkRequest struct {
	*requests.RpcRequest
	InstanceId     string `position:"Body" name:"InstanceId"`
	DownloadTaskId string `position:"Body" name:"DownloadTaskId"`
}

// DescribeBackupSetDownloadLinkResponse is the response struct for api DescribeBackupSetDownloadLink
type DescribeBackupSetDownloadLinkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeBackupSetDownloadLinkRequest creates a request to invoke DescribeBackupSetDownloadLink API
func CreateDescribeBackupSetDownloadLinkRequest() (request *DescribeBackupSetDownloadLinkRequest) {
	request = &DescribeBackupSetDownloadLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeBackupSetDownloadLink", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupSetDownloadLinkResponse creates a response to parse from DescribeBackupSetDownloadLink response
func CreateDescribeBackupSetDownloadLinkResponse() (response *DescribeBackupSetDownloadLinkResponse) {
	response = &DescribeBackupSetDownloadLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
