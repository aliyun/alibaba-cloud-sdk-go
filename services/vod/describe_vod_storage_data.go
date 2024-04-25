package vod

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

// DescribeVodStorageData invokes the vod.DescribeVodStorageData API synchronously
func (client *Client) DescribeVodStorageData(request *DescribeVodStorageDataRequest) (response *DescribeVodStorageDataResponse, err error) {
	response = CreateDescribeVodStorageDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodStorageDataWithChan invokes the vod.DescribeVodStorageData API asynchronously
func (client *Client) DescribeVodStorageDataWithChan(request *DescribeVodStorageDataRequest) (<-chan *DescribeVodStorageDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVodStorageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodStorageData(request)
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

// DescribeVodStorageDataWithCallback invokes the vod.DescribeVodStorageData API asynchronously
func (client *Client) DescribeVodStorageDataWithCallback(request *DescribeVodStorageDataRequest, callback func(response *DescribeVodStorageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodStorageDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodStorageData(request)
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

// DescribeVodStorageDataRequest is the request struct for api DescribeVodStorageData
type DescribeVodStorageDataRequest struct {
	*requests.RpcRequest
	StartTime   string           `position:"Query" name:"StartTime"`
	Storage     string           `position:"Query" name:"Storage"`
	StorageType string           `position:"Query" name:"StorageType"`
	EndTime     string           `position:"Query" name:"EndTime"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	AppId       string           `position:"Query" name:"AppId"`
	Interval    string           `position:"Query" name:"Interval"`
	Region      string           `position:"Query" name:"Region"`
}

// DescribeVodStorageDataResponse is the response struct for api DescribeVodStorageData
type DescribeVodStorageDataResponse struct {
	*responses.BaseResponse
	DataInterval string      `json:"DataInterval" xml:"DataInterval"`
	RequestId    string      `json:"RequestId" xml:"RequestId"`
	StorageData  StorageData `json:"StorageData" xml:"StorageData"`
}

// CreateDescribeVodStorageDataRequest creates a request to invoke DescribeVodStorageData API
func CreateDescribeVodStorageDataRequest() (request *DescribeVodStorageDataRequest) {
	request = &DescribeVodStorageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodStorageData", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVodStorageDataResponse creates a response to parse from DescribeVodStorageData response
func CreateDescribeVodStorageDataResponse() (response *DescribeVodStorageDataResponse) {
	response = &DescribeVodStorageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
