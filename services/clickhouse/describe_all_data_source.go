package clickhouse

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

// DescribeAllDataSource invokes the clickhouse.DescribeAllDataSource API synchronously
func (client *Client) DescribeAllDataSource(request *DescribeAllDataSourceRequest) (response *DescribeAllDataSourceResponse, err error) {
	response = CreateDescribeAllDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAllDataSourceWithChan invokes the clickhouse.DescribeAllDataSource API asynchronously
func (client *Client) DescribeAllDataSourceWithChan(request *DescribeAllDataSourceRequest) (<-chan *DescribeAllDataSourceResponse, <-chan error) {
	responseChan := make(chan *DescribeAllDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAllDataSource(request)
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

// DescribeAllDataSourceWithCallback invokes the clickhouse.DescribeAllDataSource API asynchronously
func (client *Client) DescribeAllDataSourceWithCallback(request *DescribeAllDataSourceRequest, callback func(response *DescribeAllDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAllDataSourceResponse
		var err error
		defer close(result)
		response, err = client.DescribeAllDataSource(request)
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

// DescribeAllDataSourceRequest is the request struct for api DescribeAllDataSource
type DescribeAllDataSourceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TableName            string           `position:"Query" name:"TableName"`
	SchemaName           string           `position:"Query" name:"SchemaName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeAllDataSourceResponse is the response struct for api DescribeAllDataSource
type DescribeAllDataSourceResponse struct {
	*responses.BaseResponse
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	Tables    TablesInDescribeAllDataSource  `json:"Tables" xml:"Tables"`
	Columns   ColumnsInDescribeAllDataSource `json:"Columns" xml:"Columns"`
	Schemas   SchemasInDescribeAllDataSource `json:"Schemas" xml:"Schemas"`
}

// CreateDescribeAllDataSourceRequest creates a request to invoke DescribeAllDataSource API
func CreateDescribeAllDataSourceRequest() (request *DescribeAllDataSourceRequest) {
	request = &DescribeAllDataSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeAllDataSource", "service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAllDataSourceResponse creates a response to parse from DescribeAllDataSource response
func CreateDescribeAllDataSourceResponse() (response *DescribeAllDataSourceResponse) {
	response = &DescribeAllDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
