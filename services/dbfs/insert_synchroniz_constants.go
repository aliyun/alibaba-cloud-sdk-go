package dbfs

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

// InsertSynchronizConstants invokes the dbfs.InsertSynchronizConstants API synchronously
func (client *Client) InsertSynchronizConstants(request *InsertSynchronizConstantsRequest) (response *InsertSynchronizConstantsResponse, err error) {
	response = CreateInsertSynchronizConstantsResponse()
	err = client.DoAction(request, response)
	return
}

// InsertSynchronizConstantsWithChan invokes the dbfs.InsertSynchronizConstants API asynchronously
func (client *Client) InsertSynchronizConstantsWithChan(request *InsertSynchronizConstantsRequest) (<-chan *InsertSynchronizConstantsResponse, <-chan error) {
	responseChan := make(chan *InsertSynchronizConstantsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertSynchronizConstants(request)
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

// InsertSynchronizConstantsWithCallback invokes the dbfs.InsertSynchronizConstants API asynchronously
func (client *Client) InsertSynchronizConstantsWithCallback(request *InsertSynchronizConstantsRequest, callback func(response *InsertSynchronizConstantsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertSynchronizConstantsResponse
		var err error
		defer close(result)
		response, err = client.InsertSynchronizConstants(request)
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

// InsertSynchronizConstantsRequest is the request struct for api InsertSynchronizConstants
type InsertSynchronizConstantsRequest struct {
	*requests.RpcRequest
	ZoneData        string           `position:"Query" name:"ZoneData"`
	OsversionData   string           `position:"Query" name:"OsversionData"`
	EndpointData    string           `position:"Query" name:"EndpointData"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	MasterData      string           `position:"Query" name:"MasterData"`
	ProductCodeData string           `position:"Query" name:"ProductCodeData"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	AccessData      string           `position:"Query" name:"AccessData"`
	RegionData      string           `position:"Query" name:"RegionData"`
}

// InsertSynchronizConstantsResponse is the response struct for api InsertSynchronizConstants
type InsertSynchronizConstantsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Data       string `json:"Data" xml:"Data"`
	PageSize   int64  `json:"PageSize" xml:"PageSize"`
	TotalCount int64  `json:"TotalCount" xml:"TotalCount"`
	PageNumber int64  `json:"PageNumber" xml:"PageNumber"`
}

// CreateInsertSynchronizConstantsRequest creates a request to invoke InsertSynchronizConstants API
func CreateInsertSynchronizConstantsRequest() (request *InsertSynchronizConstantsRequest) {
	request = &InsertSynchronizConstantsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "InsertSynchronizConstants", "", "")
	request.Method = requests.POST
	return
}

// CreateInsertSynchronizConstantsResponse creates a response to parse from InsertSynchronizConstants response
func CreateInsertSynchronizConstantsResponse() (response *InsertSynchronizConstantsResponse) {
	response = &InsertSynchronizConstantsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
