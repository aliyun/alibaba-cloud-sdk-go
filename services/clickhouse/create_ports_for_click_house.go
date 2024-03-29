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

// CreatePortsForClickHouse invokes the clickhouse.CreatePortsForClickHouse API synchronously
func (client *Client) CreatePortsForClickHouse(request *CreatePortsForClickHouseRequest) (response *CreatePortsForClickHouseResponse, err error) {
	response = CreateCreatePortsForClickHouseResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePortsForClickHouseWithChan invokes the clickhouse.CreatePortsForClickHouse API asynchronously
func (client *Client) CreatePortsForClickHouseWithChan(request *CreatePortsForClickHouseRequest) (<-chan *CreatePortsForClickHouseResponse, <-chan error) {
	responseChan := make(chan *CreatePortsForClickHouseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePortsForClickHouse(request)
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

// CreatePortsForClickHouseWithCallback invokes the clickhouse.CreatePortsForClickHouse API asynchronously
func (client *Client) CreatePortsForClickHouseWithCallback(request *CreatePortsForClickHouseRequest, callback func(response *CreatePortsForClickHouseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePortsForClickHouseResponse
		var err error
		defer close(result)
		response, err = client.CreatePortsForClickHouse(request)
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

// CreatePortsForClickHouseRequest is the request struct for api CreatePortsForClickHouse
type CreatePortsForClickHouseRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PortType             string           `position:"Query" name:"PortType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreatePortsForClickHouseResponse is the response struct for api CreatePortsForClickHouse
type CreatePortsForClickHouseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreatePortsForClickHouseRequest creates a request to invoke CreatePortsForClickHouse API
func CreateCreatePortsForClickHouseRequest() (request *CreatePortsForClickHouseRequest) {
	request = &CreatePortsForClickHouseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "CreatePortsForClickHouse", "service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatePortsForClickHouseResponse creates a response to parse from CreatePortsForClickHouse response
func CreateCreatePortsForClickHouseResponse() (response *CreatePortsForClickHouseResponse) {
	response = &CreatePortsForClickHouseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
