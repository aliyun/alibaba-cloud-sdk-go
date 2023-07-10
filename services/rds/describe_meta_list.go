package rds

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

// DescribeMetaList invokes the rds.DescribeMetaList API synchronously
func (client *Client) DescribeMetaList(request *DescribeMetaListRequest) (response *DescribeMetaListResponse, err error) {
	response = CreateDescribeMetaListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMetaListWithChan invokes the rds.DescribeMetaList API asynchronously
func (client *Client) DescribeMetaListWithChan(request *DescribeMetaListRequest) (<-chan *DescribeMetaListResponse, <-chan error) {
	responseChan := make(chan *DescribeMetaListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMetaList(request)
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

// DescribeMetaListWithCallback invokes the rds.DescribeMetaList API asynchronously
func (client *Client) DescribeMetaListWithCallback(request *DescribeMetaListRequest, callback func(response *DescribeMetaListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMetaListResponse
		var err error
		defer close(result)
		response, err = client.DescribeMetaList(request)
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

// DescribeMetaListRequest is the request struct for api DescribeMetaList
type DescribeMetaListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Pattern              string           `position:"Query" name:"Pattern"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	PageIndex            requests.Integer `position:"Query" name:"PageIndex"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupSetID          requests.Integer `position:"Query" name:"BackupSetID"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	GetDbName            string           `position:"Query" name:"GetDbName"`
	RestoreType          string           `position:"Query" name:"RestoreType"`
}

// DescribeMetaListResponse is the response struct for api DescribeMetaList
type DescribeMetaListResponse struct {
	*responses.BaseResponse
	DBInstanceName   string                  `json:"DBInstanceName" xml:"DBInstanceName"`
	TotalPageCount   int                     `json:"TotalPageCount" xml:"TotalPageCount"`
	RequestId        string                  `json:"RequestId" xml:"RequestId"`
	PageRecordCount  int                     `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int                     `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                     `json:"PageNumber" xml:"PageNumber"`
	Items            ItemsInDescribeMetaList `json:"Items" xml:"Items"`
}

// CreateDescribeMetaListRequest creates a request to invoke DescribeMetaList API
func CreateDescribeMetaListRequest() (request *DescribeMetaListRequest) {
	request = &DescribeMetaListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeMetaList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeMetaListResponse creates a response to parse from DescribeMetaList response
func CreateDescribeMetaListResponse() (response *DescribeMetaListResponse) {
	response = &DescribeMetaListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
