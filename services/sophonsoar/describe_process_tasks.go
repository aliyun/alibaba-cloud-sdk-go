package sophonsoar

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

// DescribeProcessTasks invokes the sophonsoar.DescribeProcessTasks API synchronously
func (client *Client) DescribeProcessTasks(request *DescribeProcessTasksRequest) (response *DescribeProcessTasksResponse, err error) {
	response = CreateDescribeProcessTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProcessTasksWithChan invokes the sophonsoar.DescribeProcessTasks API asynchronously
func (client *Client) DescribeProcessTasksWithChan(request *DescribeProcessTasksRequest) (<-chan *DescribeProcessTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeProcessTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProcessTasks(request)
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

// DescribeProcessTasksWithCallback invokes the sophonsoar.DescribeProcessTasks API asynchronously
func (client *Client) DescribeProcessTasksWithCallback(request *DescribeProcessTasksRequest, callback func(response *DescribeProcessTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProcessTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeProcessTasks(request)
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

// DescribeProcessTasksRequest is the request struct for api DescribeProcessTasks
type DescribeProcessTasksRequest struct {
	*requests.RpcRequest
	EntityKey           string           `position:"Query" name:"EntityKey"`
	RoleFor             string           `position:"Query" name:"RoleFor"`
	EntityName          string           `position:"Query" name:"EntityName"`
	TaskName            string           `position:"Query" name:"TaskName"`
	YunCode             string           `position:"Query" name:"YunCode"`
	Source              string           `position:"Query" name:"Source"`
	PageNumber          requests.Integer `position:"Query" name:"PageNumber"`
	TaskStatus          string           `position:"Query" name:"TaskStatus"`
	ProcessRemoveEnd    requests.Integer `position:"Query" name:"ProcessRemoveEnd"`
	ParamContent        string           `position:"Query" name:"ParamContent"`
	Scope               string           `position:"Query" name:"Scope"`
	PageSize            requests.Integer `position:"Query" name:"PageSize"`
	TriggerSource       string           `position:"Query" name:"TriggerSource"`
	ProcessRemoveStart  requests.Integer `position:"Query" name:"ProcessRemoveStart"`
	RoleType            string           `position:"Query" name:"RoleType"`
	Lang                string           `position:"Query" name:"Lang"`
	TaskId              string           `position:"Query" name:"TaskId"`
	OrderField          string           `position:"Query" name:"OrderField"`
	Direction           string           `position:"Query" name:"Direction"`
	SceneCode           string           `position:"Query" name:"SceneCode"`
	ProcessActionStart  requests.Integer `position:"Query" name:"ProcessActionStart"`
	ProcessActionEnd    requests.Integer `position:"Query" name:"ProcessActionEnd"`
	ProcessStrategyUuid string           `position:"Query" name:"ProcessStrategyUuid"`
	EntityType          string           `position:"Query" name:"EntityType"`
	EntityUuid          string           `position:"Query" name:"EntityUuid"`
	EventUuid           string           `position:"Query" name:"EventUuid"`
}

// DescribeProcessTasksResponse is the response struct for api DescribeProcessTasks
type DescribeProcessTasksResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Page         Page   `json:"Page" xml:"Page"`
	ProcessTasks []Data `json:"ProcessTasks" xml:"ProcessTasks"`
}

// CreateDescribeProcessTasksRequest creates a request to invoke DescribeProcessTasks API
func CreateDescribeProcessTasksRequest() (request *DescribeProcessTasksRequest) {
	request = &DescribeProcessTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "DescribeProcessTasks", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeProcessTasksResponse creates a response to parse from DescribeProcessTasks response
func CreateDescribeProcessTasksResponse() (response *DescribeProcessTasksResponse) {
	response = &DescribeProcessTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
