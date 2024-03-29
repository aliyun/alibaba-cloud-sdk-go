package linkvisual

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

// QueryVisionDeviceInfo invokes the linkvisual.QueryVisionDeviceInfo API synchronously
func (client *Client) QueryVisionDeviceInfo(request *QueryVisionDeviceInfoRequest) (response *QueryVisionDeviceInfoResponse, err error) {
	response = CreateQueryVisionDeviceInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryVisionDeviceInfoWithChan invokes the linkvisual.QueryVisionDeviceInfo API asynchronously
func (client *Client) QueryVisionDeviceInfoWithChan(request *QueryVisionDeviceInfoRequest) (<-chan *QueryVisionDeviceInfoResponse, <-chan error) {
	responseChan := make(chan *QueryVisionDeviceInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryVisionDeviceInfo(request)
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

// QueryVisionDeviceInfoWithCallback invokes the linkvisual.QueryVisionDeviceInfo API asynchronously
func (client *Client) QueryVisionDeviceInfoWithCallback(request *QueryVisionDeviceInfoRequest, callback func(response *QueryVisionDeviceInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryVisionDeviceInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryVisionDeviceInfo(request)
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

// QueryVisionDeviceInfoRequest is the request struct for api QueryVisionDeviceInfo
type QueryVisionDeviceInfoRequest struct {
	*requests.RpcRequest
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// QueryVisionDeviceInfoResponse is the response struct for api QueryVisionDeviceInfo
type QueryVisionDeviceInfoResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryVisionDeviceInfoRequest creates a request to invoke QueryVisionDeviceInfo API
func CreateQueryVisionDeviceInfoRequest() (request *QueryVisionDeviceInfoRequest) {
	request = &QueryVisionDeviceInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryVisionDeviceInfo", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryVisionDeviceInfoResponse creates a response to parse from QueryVisionDeviceInfo response
func CreateQueryVisionDeviceInfoResponse() (response *QueryVisionDeviceInfoResponse) {
	response = &QueryVisionDeviceInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
