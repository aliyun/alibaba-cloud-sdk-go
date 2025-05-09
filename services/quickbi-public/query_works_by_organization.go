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

// QueryWorksByOrganization invokes the quickbi_public.QueryWorksByOrganization API synchronously
func (client *Client) QueryWorksByOrganization(request *QueryWorksByOrganizationRequest) (response *QueryWorksByOrganizationResponse, err error) {
	response = CreateQueryWorksByOrganizationResponse()
	err = client.DoAction(request, response)
	return
}

// QueryWorksByOrganizationWithChan invokes the quickbi_public.QueryWorksByOrganization API asynchronously
func (client *Client) QueryWorksByOrganizationWithChan(request *QueryWorksByOrganizationRequest) (<-chan *QueryWorksByOrganizationResponse, <-chan error) {
	responseChan := make(chan *QueryWorksByOrganizationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryWorksByOrganization(request)
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

// QueryWorksByOrganizationWithCallback invokes the quickbi_public.QueryWorksByOrganization API asynchronously
func (client *Client) QueryWorksByOrganizationWithCallback(request *QueryWorksByOrganizationRequest, callback func(response *QueryWorksByOrganizationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryWorksByOrganizationResponse
		var err error
		defer close(result)
		response, err = client.QueryWorksByOrganization(request)
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

// QueryWorksByOrganizationRequest is the request struct for api QueryWorksByOrganization
type QueryWorksByOrganizationRequest struct {
	*requests.RpcRequest
	ThirdPartAuthFlag requests.Integer `position:"Query" name:"ThirdPartAuthFlag"`
	AccessPoint       string           `position:"Query" name:"AccessPoint"`
	PageNum           requests.Integer `position:"Query" name:"PageNum"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	WorksType         string           `position:"Query" name:"WorksType"`
	SignType          string           `position:"Query" name:"SignType"`
	Status            requests.Integer `position:"Query" name:"Status"`
}

// QueryWorksByOrganizationResponse is the response struct for api QueryWorksByOrganization
type QueryWorksByOrganizationResponse struct {
	*responses.BaseResponse
	Success   bool                             `json:"Success" xml:"Success"`
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	Result    ResultInQueryWorksByOrganization `json:"Result" xml:"Result"`
}

// CreateQueryWorksByOrganizationRequest creates a request to invoke QueryWorksByOrganization API
func CreateQueryWorksByOrganizationRequest() (request *QueryWorksByOrganizationRequest) {
	request = &QueryWorksByOrganizationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryWorksByOrganization", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryWorksByOrganizationResponse creates a response to parse from QueryWorksByOrganization response
func CreateQueryWorksByOrganizationResponse() (response *QueryWorksByOrganizationResponse) {
	response = &QueryWorksByOrganizationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
