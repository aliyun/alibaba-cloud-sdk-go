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

// DeleteDataStream invokes the elasticsearch.DeleteDataStream API synchronously
func (client *Client) DeleteDataStream(request *DeleteDataStreamRequest) (response *DeleteDataStreamResponse, err error) {
	response = CreateDeleteDataStreamResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDataStreamWithChan invokes the elasticsearch.DeleteDataStream API asynchronously
func (client *Client) DeleteDataStreamWithChan(request *DeleteDataStreamRequest) (<-chan *DeleteDataStreamResponse, <-chan error) {
	responseChan := make(chan *DeleteDataStreamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDataStream(request)
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

// DeleteDataStreamWithCallback invokes the elasticsearch.DeleteDataStream API asynchronously
func (client *Client) DeleteDataStreamWithCallback(request *DeleteDataStreamRequest, callback func(response *DeleteDataStreamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDataStreamResponse
		var err error
		defer close(result)
		response, err = client.DeleteDataStream(request)
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

// DeleteDataStreamRequest is the request struct for api DeleteDataStream
type DeleteDataStreamRequest struct {
	*requests.RoaRequest
	DataStream  string `position:"Path" name:"DataStream"`
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

// DeleteDataStreamResponse is the response struct for api DeleteDataStream
type DeleteDataStreamResponse struct {
	*responses.BaseResponse
	Result    bool   `json:"Result" xml:"Result"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDataStreamRequest creates a request to invoke DeleteDataStream API
func CreateDeleteDataStreamRequest() (request *DeleteDataStreamRequest) {
	request = &DeleteDataStreamRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DeleteDataStream", "/openapi/instances/[InstanceId]/data-streams/[DataStream]", "elasticsearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteDataStreamResponse creates a response to parse from DeleteDataStream response
func CreateDeleteDataStreamResponse() (response *DeleteDataStreamResponse) {
	response = &DeleteDataStreamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
