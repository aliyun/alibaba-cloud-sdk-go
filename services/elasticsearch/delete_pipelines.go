package elasticsearch

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

// DeletePipelines invokes the elasticsearch.DeletePipelines API synchronously
func (client *Client) DeletePipelines(request *DeletePipelinesRequest) (response *DeletePipelinesResponse, err error) {
	response = CreateDeletePipelinesResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePipelinesWithChan invokes the elasticsearch.DeletePipelines API asynchronously
func (client *Client) DeletePipelinesWithChan(request *DeletePipelinesRequest) (<-chan *DeletePipelinesResponse, <-chan error) {
	responseChan := make(chan *DeletePipelinesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePipelines(request)
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

// DeletePipelinesWithCallback invokes the elasticsearch.DeletePipelines API asynchronously
func (client *Client) DeletePipelinesWithCallback(request *DeletePipelinesRequest, callback func(response *DeletePipelinesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePipelinesResponse
		var err error
		defer close(result)
		response, err = client.DeletePipelines(request)
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

// DeletePipelinesRequest is the request struct for api DeletePipelines
type DeletePipelinesRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	PipelineIds string `position:"Query" name:"pipelineIds"`
}

// DeletePipelinesResponse is the response struct for api DeletePipelines
type DeletePipelinesResponse struct {
	*responses.BaseResponse
	Result    bool   `json:"Result" xml:"Result"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeletePipelinesRequest creates a request to invoke DeletePipelines API
func CreateDeletePipelinesRequest() (request *DeletePipelinesRequest) {
	request = &DeletePipelinesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DeletePipelines", "/openapi/logstashes/[InstanceId]/pipelines", "elasticsearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeletePipelinesResponse creates a response to parse from DeletePipelines response
func CreateDeletePipelinesResponse() (response *DeletePipelinesResponse) {
	response = &DeletePipelinesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
