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

// CreateMySqlDataSource invokes the oceanbasepro.CreateMySqlDataSource API synchronously
func (client *Client) CreateMySqlDataSource(request *CreateMySqlDataSourceRequest) (response *CreateMySqlDataSourceResponse, err error) {
	response = CreateCreateMySqlDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMySqlDataSourceWithChan invokes the oceanbasepro.CreateMySqlDataSource API asynchronously
func (client *Client) CreateMySqlDataSourceWithChan(request *CreateMySqlDataSourceRequest) (<-chan *CreateMySqlDataSourceResponse, <-chan error) {
	responseChan := make(chan *CreateMySqlDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMySqlDataSource(request)
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

// CreateMySqlDataSourceWithCallback invokes the oceanbasepro.CreateMySqlDataSource API asynchronously
func (client *Client) CreateMySqlDataSourceWithCallback(request *CreateMySqlDataSourceRequest, callback func(response *CreateMySqlDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMySqlDataSourceResponse
		var err error
		defer close(result)
		response, err = client.CreateMySqlDataSource(request)
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

// CreateMySqlDataSourceRequest is the request struct for api CreateMySqlDataSource
type CreateMySqlDataSourceRequest struct {
	*requests.RpcRequest
	Schema       string           `position:"Body" name:"Schema"`
	Description  string           `position:"Body" name:"Description"`
	Type         string           `position:"Body" name:"Type"`
	UseSsl       requests.Boolean `position:"Body" name:"UseSsl"`
	Password     string           `position:"Body" name:"Password"`
	DgInstanceId string           `position:"Body" name:"DgInstanceId"`
	Ip           string           `position:"Body" name:"Ip"`
	InstanceId   string           `position:"Body" name:"InstanceId"`
	Port         requests.Integer `position:"Body" name:"Port"`
	VpcId        string           `position:"Body" name:"VpcId"`
	Name         string           `position:"Body" name:"Name"`
	UserName     string           `position:"Body" name:"UserName"`
}

// CreateMySqlDataSourceResponse is the response struct for api CreateMySqlDataSource
type CreateMySqlDataSourceResponse struct {
	*responses.BaseResponse
}

// CreateCreateMySqlDataSourceRequest creates a request to invoke CreateMySqlDataSource API
func CreateCreateMySqlDataSourceRequest() (request *CreateMySqlDataSourceRequest) {
	request = &CreateMySqlDataSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "CreateMySqlDataSource", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMySqlDataSourceResponse creates a response to parse from CreateMySqlDataSource response
func CreateCreateMySqlDataSourceResponse() (response *CreateMySqlDataSourceResponse) {
	response = &CreateMySqlDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
