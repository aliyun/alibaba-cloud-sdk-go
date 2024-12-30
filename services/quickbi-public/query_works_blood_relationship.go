package quickbi_public

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

// QueryWorksBloodRelationship invokes the quickbi_public.QueryWorksBloodRelationship API synchronously
func (client *Client) QueryWorksBloodRelationship(request *QueryWorksBloodRelationshipRequest) (response *QueryWorksBloodRelationshipResponse, err error) {
	response = CreateQueryWorksBloodRelationshipResponse()
	err = client.DoAction(request, response)
	return
}

// QueryWorksBloodRelationshipWithChan invokes the quickbi_public.QueryWorksBloodRelationship API asynchronously
func (client *Client) QueryWorksBloodRelationshipWithChan(request *QueryWorksBloodRelationshipRequest) (<-chan *QueryWorksBloodRelationshipResponse, <-chan error) {
	responseChan := make(chan *QueryWorksBloodRelationshipResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryWorksBloodRelationship(request)
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

// QueryWorksBloodRelationshipWithCallback invokes the quickbi_public.QueryWorksBloodRelationship API asynchronously
func (client *Client) QueryWorksBloodRelationshipWithCallback(request *QueryWorksBloodRelationshipRequest, callback func(response *QueryWorksBloodRelationshipResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryWorksBloodRelationshipResponse
		var err error
		defer close(result)
		response, err = client.QueryWorksBloodRelationship(request)
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

// QueryWorksBloodRelationshipRequest is the request struct for api QueryWorksBloodRelationship
type QueryWorksBloodRelationshipRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	ApiLevel    string `position:"Query" name:"ApiLevel"`
	WorksId     string `position:"Query" name:"WorksId"`
}

// QueryWorksBloodRelationshipResponse is the response struct for api QueryWorksBloodRelationship
type QueryWorksBloodRelationshipResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    []Item `json:"Result" xml:"Result"`
}

// CreateQueryWorksBloodRelationshipRequest creates a request to invoke QueryWorksBloodRelationship API
func CreateQueryWorksBloodRelationshipRequest() (request *QueryWorksBloodRelationshipRequest) {
	request = &QueryWorksBloodRelationshipRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryWorksBloodRelationship", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryWorksBloodRelationshipResponse creates a response to parse from QueryWorksBloodRelationship response
func CreateQueryWorksBloodRelationshipResponse() (response *QueryWorksBloodRelationshipResponse) {
	response = &QueryWorksBloodRelationshipResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
