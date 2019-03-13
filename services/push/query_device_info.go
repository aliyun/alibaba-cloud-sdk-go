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

package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryDeviceInfo invokes the push.QueryDeviceInfo API synchronously
// api document: https://help.aliyun.com/api/push/querydeviceinfo.html
func (client *Client) QueryDeviceInfo(request *QueryDeviceInfoRequest) (response *QueryDeviceInfoResponse, err error) {
	response = CreateQueryDeviceInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceInfoWithChan invokes the push.QueryDeviceInfo API asynchronously
// api document: https://help.aliyun.com/api/push/querydeviceinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceInfoWithChan(request *QueryDeviceInfoRequest) (<-chan *QueryDeviceInfoResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceInfo(request)
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

// QueryDeviceInfoWithCallback invokes the push.QueryDeviceInfo API asynchronously
// api document: https://help.aliyun.com/api/push/querydeviceinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceInfoWithCallback(request *QueryDeviceInfoRequest, callback func(response *QueryDeviceInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceInfo(request)
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

// QueryDeviceInfoRequest is the request struct for api QueryDeviceInfo
type QueryDeviceInfoRequest struct {
	*requests.RpcRequest
	AccessKeyId string           `position:"Query" name:"AccessKeyId"`
	AppKey      requests.Integer `position:"Query" name:"AppKey"`
	DeviceId    string           `position:"Query" name:"DeviceId"`
}

// QueryDeviceInfoResponse is the response struct for api QueryDeviceInfo
type QueryDeviceInfoResponse struct {
	*responses.BaseResponse
	RequestId  string                     `json:"RequestId" xml:"RequestId"`
	DeviceInfo QueryDeviceInfoDeviceInfo0 `json:"DeviceInfo" xml:"DeviceInfo"`
}

type QueryDeviceInfoDeviceInfo0 struct {
	DeviceId       string `json:"DeviceId" xml:"DeviceId"`
	DeviceType     string `json:"DeviceType" xml:"DeviceType"`
	Account        string `json:"Account" xml:"Account"`
	DeviceToken    string `json:"DeviceToken" xml:"DeviceToken"`
	Tags           string `json:"Tags" xml:"Tags"`
	Alias          string `json:"Alias" xml:"Alias"`
	LastOnlineTime string `json:"LastOnlineTime" xml:"LastOnlineTime"`
	Online         bool   `json:"Online" xml:"Online"`
	PhoneNumber    string `json:"PhoneNumber" xml:"PhoneNumber"`
	PushEnabled    bool   `json:"PushEnabled" xml:"PushEnabled"`
}

// CreateQueryDeviceInfoRequest creates a request to invoke QueryDeviceInfo API
func CreateQueryDeviceInfoRequest() (request *QueryDeviceInfoRequest) {
	request = &QueryDeviceInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryDeviceInfo", "push", "openAPI")
	return
}

// CreateQueryDeviceInfoResponse creates a response to parse from QueryDeviceInfo response
func CreateQueryDeviceInfoResponse() (response *QueryDeviceInfoResponse) {
	response = &QueryDeviceInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
