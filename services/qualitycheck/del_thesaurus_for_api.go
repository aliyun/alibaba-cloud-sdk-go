package qualitycheck

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

// DelThesaurusForApi invokes the qualitycheck.DelThesaurusForApi API synchronously
func (client *Client) DelThesaurusForApi(request *DelThesaurusForApiRequest) (response *DelThesaurusForApiResponse, err error) {
	response = CreateDelThesaurusForApiResponse()
	err = client.DoAction(request, response)
	return
}

// DelThesaurusForApiWithChan invokes the qualitycheck.DelThesaurusForApi API asynchronously
func (client *Client) DelThesaurusForApiWithChan(request *DelThesaurusForApiRequest) (<-chan *DelThesaurusForApiResponse, <-chan error) {
	responseChan := make(chan *DelThesaurusForApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DelThesaurusForApi(request)
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

// DelThesaurusForApiWithCallback invokes the qualitycheck.DelThesaurusForApi API asynchronously
func (client *Client) DelThesaurusForApiWithCallback(request *DelThesaurusForApiRequest, callback func(response *DelThesaurusForApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DelThesaurusForApiResponse
		var err error
		defer close(result)
		response, err = client.DelThesaurusForApi(request)
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

// DelThesaurusForApiRequest is the request struct for api DelThesaurusForApi
type DelThesaurusForApiRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId   requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// DelThesaurusForApiResponse is the response struct for api DelThesaurusForApi
type DelThesaurusForApiResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDelThesaurusForApiRequest creates a request to invoke DelThesaurusForApi API
func CreateDelThesaurusForApiRequest() (request *DelThesaurusForApiRequest) {
	request = &DelThesaurusForApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "DelThesaurusForApi", "", "")
	request.Method = requests.POST
	return
}

// CreateDelThesaurusForApiResponse creates a response to parse from DelThesaurusForApi response
func CreateDelThesaurusForApiResponse() (response *DelThesaurusForApiResponse) {
	response = &DelThesaurusForApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
