package dysmsapi

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

// QueryMobilesCardSupport invokes the dysmsapi.QueryMobilesCardSupport API synchronously
func (client *Client) QueryMobilesCardSupport(request *QueryMobilesCardSupportRequest) (response *QueryMobilesCardSupportResponse, err error) {
	response = CreateQueryMobilesCardSupportResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMobilesCardSupportWithChan invokes the dysmsapi.QueryMobilesCardSupport API asynchronously
func (client *Client) QueryMobilesCardSupportWithChan(request *QueryMobilesCardSupportRequest) (<-chan *QueryMobilesCardSupportResponse, <-chan error) {
	responseChan := make(chan *QueryMobilesCardSupportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMobilesCardSupport(request)
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

// QueryMobilesCardSupportWithCallback invokes the dysmsapi.QueryMobilesCardSupport API asynchronously
func (client *Client) QueryMobilesCardSupportWithCallback(request *QueryMobilesCardSupportRequest, callback func(response *QueryMobilesCardSupportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMobilesCardSupportResponse
		var err error
		defer close(result)
		response, err = client.QueryMobilesCardSupport(request)
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

// QueryMobilesCardSupportRequest is the request struct for api QueryMobilesCardSupport
type QueryMobilesCardSupportRequest struct {
	*requests.RpcRequest
	ProductCode     string               `position:"Query" name:"ProductCode"`
	ResourceOwnerId requests.Integer     `position:"Query" name:"ResourceOwnerId"`
	Mobiles         *[]map[string]string `position:"Query" name:"Mobiles"  type:"Json"`
	TemplateCode    string               `position:"Query" name:"TemplateCode"`
}

// QueryMobilesCardSupportResponse is the response struct for api QueryMobilesCardSupport
type QueryMobilesCardSupportResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryMobilesCardSupportRequest creates a request to invoke QueryMobilesCardSupport API
func CreateQueryMobilesCardSupportRequest() (request *QueryMobilesCardSupportRequest) {
	request = &QueryMobilesCardSupportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "QueryMobilesCardSupport", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryMobilesCardSupportResponse creates a response to parse from QueryMobilesCardSupport response
func CreateQueryMobilesCardSupportResponse() (response *QueryMobilesCardSupportResponse) {
	response = &QueryMobilesCardSupportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
