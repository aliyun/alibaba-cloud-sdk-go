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

// DescribeKibanaSettings invokes the elasticsearch.DescribeKibanaSettings API synchronously
func (client *Client) DescribeKibanaSettings(request *DescribeKibanaSettingsRequest) (response *DescribeKibanaSettingsResponse, err error) {
	response = CreateDescribeKibanaSettingsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeKibanaSettingsWithChan invokes the elasticsearch.DescribeKibanaSettings API asynchronously
func (client *Client) DescribeKibanaSettingsWithChan(request *DescribeKibanaSettingsRequest) (<-chan *DescribeKibanaSettingsResponse, <-chan error) {
	responseChan := make(chan *DescribeKibanaSettingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeKibanaSettings(request)
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

// DescribeKibanaSettingsWithCallback invokes the elasticsearch.DescribeKibanaSettings API asynchronously
func (client *Client) DescribeKibanaSettingsWithCallback(request *DescribeKibanaSettingsRequest, callback func(response *DescribeKibanaSettingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeKibanaSettingsResponse
		var err error
		defer close(result)
		response, err = client.DescribeKibanaSettings(request)
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

// DescribeKibanaSettingsRequest is the request struct for api DescribeKibanaSettings
type DescribeKibanaSettingsRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
}

// DescribeKibanaSettingsResponse is the response struct for api DescribeKibanaSettings
type DescribeKibanaSettingsResponse struct {
	*responses.BaseResponse
	Result    map[string]interface{} `json:"Result" xml:"Result"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeKibanaSettingsRequest creates a request to invoke DescribeKibanaSettings API
func CreateDescribeKibanaSettingsRequest() (request *DescribeKibanaSettingsRequest) {
	request = &DescribeKibanaSettingsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DescribeKibanaSettings", "/openapi/instances/[InstanceId]/kibana-settings", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeKibanaSettingsResponse creates a response to parse from DescribeKibanaSettings response
func CreateDescribeKibanaSettingsResponse() (response *DescribeKibanaSettingsResponse) {
	response = &DescribeKibanaSettingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
