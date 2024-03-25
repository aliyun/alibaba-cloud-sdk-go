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

// DescribeSampleSqlRawTexts invokes the oceanbasepro.DescribeSampleSqlRawTexts API synchronously
func (client *Client) DescribeSampleSqlRawTexts(request *DescribeSampleSqlRawTextsRequest) (response *DescribeSampleSqlRawTextsResponse, err error) {
	response = CreateDescribeSampleSqlRawTextsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSampleSqlRawTextsWithChan invokes the oceanbasepro.DescribeSampleSqlRawTexts API asynchronously
func (client *Client) DescribeSampleSqlRawTextsWithChan(request *DescribeSampleSqlRawTextsRequest) (<-chan *DescribeSampleSqlRawTextsResponse, <-chan error) {
	responseChan := make(chan *DescribeSampleSqlRawTextsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSampleSqlRawTexts(request)
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

// DescribeSampleSqlRawTextsWithCallback invokes the oceanbasepro.DescribeSampleSqlRawTexts API asynchronously
func (client *Client) DescribeSampleSqlRawTextsWithCallback(request *DescribeSampleSqlRawTextsRequest, callback func(response *DescribeSampleSqlRawTextsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSampleSqlRawTextsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSampleSqlRawTexts(request)
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

// DescribeSampleSqlRawTextsRequest is the request struct for api DescribeSampleSqlRawTexts
type DescribeSampleSqlRawTextsRequest struct {
	*requests.RpcRequest
	TraceId    string `position:"Body" name:"TraceId"`
	StartTime  string `position:"Body" name:"StartTime"`
	TenantId   string `position:"Body" name:"TenantId"`
	Limit      string `position:"Body" name:"Limit"`
	SqlId      string `position:"Body" name:"SqlId"`
	EndTime    string `position:"Body" name:"EndTime"`
	InstanceId string `position:"Body" name:"InstanceId"`
	DbName     string `position:"Body" name:"DbName"`
}

// DescribeSampleSqlRawTextsResponse is the response struct for api DescribeSampleSqlRawTexts
type DescribeSampleSqlRawTextsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeSampleSqlRawTextsRequest creates a request to invoke DescribeSampleSqlRawTexts API
func CreateDescribeSampleSqlRawTextsRequest() (request *DescribeSampleSqlRawTextsRequest) {
	request = &DescribeSampleSqlRawTextsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeSampleSqlRawTexts", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSampleSqlRawTextsResponse creates a response to parse from DescribeSampleSqlRawTexts response
func CreateDescribeSampleSqlRawTextsResponse() (response *DescribeSampleSqlRawTextsResponse) {
	response = &DescribeSampleSqlRawTextsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
