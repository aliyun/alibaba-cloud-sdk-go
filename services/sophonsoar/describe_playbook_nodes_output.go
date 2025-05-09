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

// DescribePlaybookNodesOutput invokes the sophonsoar.DescribePlaybookNodesOutput API synchronously
func (client *Client) DescribePlaybookNodesOutput(request *DescribePlaybookNodesOutputRequest) (response *DescribePlaybookNodesOutputResponse, err error) {
	response = CreateDescribePlaybookNodesOutputResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePlaybookNodesOutputWithChan invokes the sophonsoar.DescribePlaybookNodesOutput API asynchronously
func (client *Client) DescribePlaybookNodesOutputWithChan(request *DescribePlaybookNodesOutputRequest) (<-chan *DescribePlaybookNodesOutputResponse, <-chan error) {
	responseChan := make(chan *DescribePlaybookNodesOutputResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePlaybookNodesOutput(request)
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

// DescribePlaybookNodesOutputWithCallback invokes the sophonsoar.DescribePlaybookNodesOutput API asynchronously
func (client *Client) DescribePlaybookNodesOutputWithCallback(request *DescribePlaybookNodesOutputRequest, callback func(response *DescribePlaybookNodesOutputResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePlaybookNodesOutputResponse
		var err error
		defer close(result)
		response, err = client.DescribePlaybookNodesOutput(request)
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

// DescribePlaybookNodesOutputRequest is the request struct for api DescribePlaybookNodesOutput
type DescribePlaybookNodesOutputRequest struct {
	*requests.RpcRequest
	RoleFor      string `position:"Query" name:"RoleFor"`
	NodeName     string `position:"Query" name:"NodeName"`
	PlaybookUuid string `position:"Query" name:"PlaybookUuid"`
	RoleType     string `position:"Query" name:"RoleType"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribePlaybookNodesOutputResponse is the response struct for api DescribePlaybookNodesOutput
type DescribePlaybookNodesOutputResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	PlaybookNodesOutput PlaybookNodesOutput `json:"PlaybookNodesOutput" xml:"PlaybookNodesOutput"`
}

// CreateDescribePlaybookNodesOutputRequest creates a request to invoke DescribePlaybookNodesOutput API
func CreateDescribePlaybookNodesOutputRequest() (request *DescribePlaybookNodesOutputRequest) {
	request = &DescribePlaybookNodesOutputRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "DescribePlaybookNodesOutput", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribePlaybookNodesOutputResponse creates a response to parse from DescribePlaybookNodesOutput response
func CreateDescribePlaybookNodesOutputResponse() (response *DescribePlaybookNodesOutputResponse) {
	response = &DescribePlaybookNodesOutputResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
