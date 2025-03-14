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

// DescribeNodeParamTags invokes the sophonsoar.DescribeNodeParamTags API synchronously
func (client *Client) DescribeNodeParamTags(request *DescribeNodeParamTagsRequest) (response *DescribeNodeParamTagsResponse, err error) {
	response = CreateDescribeNodeParamTagsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNodeParamTagsWithChan invokes the sophonsoar.DescribeNodeParamTags API asynchronously
func (client *Client) DescribeNodeParamTagsWithChan(request *DescribeNodeParamTagsRequest) (<-chan *DescribeNodeParamTagsResponse, <-chan error) {
	responseChan := make(chan *DescribeNodeParamTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNodeParamTags(request)
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

// DescribeNodeParamTagsWithCallback invokes the sophonsoar.DescribeNodeParamTags API asynchronously
func (client *Client) DescribeNodeParamTagsWithCallback(request *DescribeNodeParamTagsRequest, callback func(response *DescribeNodeParamTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNodeParamTagsResponse
		var err error
		defer close(result)
		response, err = client.DescribeNodeParamTags(request)
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

// DescribeNodeParamTagsRequest is the request struct for api DescribeNodeParamTags
type DescribeNodeParamTagsRequest struct {
	*requests.RpcRequest
	RoleFor      string `position:"Query" name:"RoleFor"`
	NodeName     string `position:"Query" name:"NodeName"`
	PlaybookUuid string `position:"Query" name:"PlaybookUuid"`
	RoleType     string `position:"Query" name:"RoleType"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeNodeParamTagsResponse is the response struct for api DescribeNodeParamTags
type DescribeNodeParamTagsResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	ParamReferredPaths []Data `json:"ParamReferredPaths" xml:"ParamReferredPaths"`
}

// CreateDescribeNodeParamTagsRequest creates a request to invoke DescribeNodeParamTags API
func CreateDescribeNodeParamTagsRequest() (request *DescribeNodeParamTagsRequest) {
	request = &DescribeNodeParamTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "DescribeNodeParamTags", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeNodeParamTagsResponse creates a response to parse from DescribeNodeParamTags response
func CreateDescribeNodeParamTagsResponse() (response *DescribeNodeParamTagsResponse) {
	response = &DescribeNodeParamTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
