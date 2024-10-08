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

// DataSourceBlood invokes the quickbi_public.DataSourceBlood API synchronously
func (client *Client) DataSourceBlood(request *DataSourceBloodRequest) (response *DataSourceBloodResponse, err error) {
	response = CreateDataSourceBloodResponse()
	err = client.DoAction(request, response)
	return
}

// DataSourceBloodWithChan invokes the quickbi_public.DataSourceBlood API asynchronously
func (client *Client) DataSourceBloodWithChan(request *DataSourceBloodRequest) (<-chan *DataSourceBloodResponse, <-chan error) {
	responseChan := make(chan *DataSourceBloodResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DataSourceBlood(request)
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

// DataSourceBloodWithCallback invokes the quickbi_public.DataSourceBlood API asynchronously
func (client *Client) DataSourceBloodWithCallback(request *DataSourceBloodRequest, callback func(response *DataSourceBloodResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DataSourceBloodResponse
		var err error
		defer close(result)
		response, err = client.DataSourceBlood(request)
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

// DataSourceBloodRequest is the request struct for api DataSourceBlood
type DataSourceBloodRequest struct {
	*requests.RpcRequest
	AccessPoint  string `position:"Query" name:"AccessPoint"`
	SignType     string `position:"Query" name:"SignType"`
	DataSourceId string `position:"Query" name:"DataSourceId"`
}

// DataSourceBloodResponse is the response struct for api DataSourceBlood
type DataSourceBloodResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Success   bool     `json:"Success" xml:"Success"`
	Result    []string `json:"Result" xml:"Result"`
}

// CreateDataSourceBloodRequest creates a request to invoke DataSourceBlood API
func CreateDataSourceBloodRequest() (request *DataSourceBloodRequest) {
	request = &DataSourceBloodRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "DataSourceBlood", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDataSourceBloodResponse creates a response to parse from DataSourceBlood response
func CreateDataSourceBloodResponse() (response *DataSourceBloodResponse) {
	response = &DataSourceBloodResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
