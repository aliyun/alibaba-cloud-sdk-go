package live

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

// LiveUpstreamQosData invokes the live.LiveUpstreamQosData API synchronously
func (client *Client) LiveUpstreamQosData(request *LiveUpstreamQosDataRequest) (response *LiveUpstreamQosDataResponse, err error) {
	response = CreateLiveUpstreamQosDataResponse()
	err = client.DoAction(request, response)
	return
}

// LiveUpstreamQosDataWithChan invokes the live.LiveUpstreamQosData API asynchronously
func (client *Client) LiveUpstreamQosDataWithChan(request *LiveUpstreamQosDataRequest) (<-chan *LiveUpstreamQosDataResponse, <-chan error) {
	responseChan := make(chan *LiveUpstreamQosDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LiveUpstreamQosData(request)
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

// LiveUpstreamQosDataWithCallback invokes the live.LiveUpstreamQosData API asynchronously
func (client *Client) LiveUpstreamQosDataWithCallback(request *LiveUpstreamQosDataRequest, callback func(response *LiveUpstreamQosDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LiveUpstreamQosDataResponse
		var err error
		defer close(result)
		response, err = client.LiveUpstreamQosData(request)
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

// LiveUpstreamQosDataRequest is the request struct for api LiveUpstreamQosData
type LiveUpstreamQosDataRequest struct {
	*requests.RpcRequest
	CdnDomains      string           `position:"Query" name:"CdnDomains"`
	StartTime       string           `position:"Query" name:"StartTime"`
	CdnProvinces    string           `position:"Query" name:"CdnProvinces"`
	KwaiSidcs       string           `position:"Query" name:"KwaiSidcs"`
	KwaiTsc         *[]string        `position:"Query" name:"KwaiTsc"  type:"Json"`
	UpstreamDomains string           `position:"Query" name:"UpstreamDomains"`
	EndTime         string           `position:"Query" name:"EndTime"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	CdnIsps         string           `position:"Query" name:"CdnIsps"`
	Region          string           `position:"Query" name:"Region"`
}

// LiveUpstreamQosDataResponse is the response struct for api LiveUpstreamQosData
type LiveUpstreamQosDataResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	StartTime string     `json:"StartTime" xml:"StartTime"`
	EndTime   string     `json:"EndTime" xml:"EndTime"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateLiveUpstreamQosDataRequest creates a request to invoke LiveUpstreamQosData API
func CreateLiveUpstreamQosDataRequest() (request *LiveUpstreamQosDataRequest) {
	request = &LiveUpstreamQosDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "LiveUpstreamQosData", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateLiveUpstreamQosDataResponse creates a response to parse from LiveUpstreamQosData response
func CreateLiveUpstreamQosDataResponse() (response *LiveUpstreamQosDataResponse) {
	response = &LiveUpstreamQosDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
