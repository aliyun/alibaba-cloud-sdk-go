package actiontrail

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

// GetAccessKeyLastUsedInfo invokes the actiontrail.GetAccessKeyLastUsedInfo API synchronously
func (client *Client) GetAccessKeyLastUsedInfo(request *GetAccessKeyLastUsedInfoRequest) (response *GetAccessKeyLastUsedInfoResponse, err error) {
	response = CreateGetAccessKeyLastUsedInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetAccessKeyLastUsedInfoWithChan invokes the actiontrail.GetAccessKeyLastUsedInfo API asynchronously
func (client *Client) GetAccessKeyLastUsedInfoWithChan(request *GetAccessKeyLastUsedInfoRequest) (<-chan *GetAccessKeyLastUsedInfoResponse, <-chan error) {
	responseChan := make(chan *GetAccessKeyLastUsedInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccessKeyLastUsedInfo(request)
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

// GetAccessKeyLastUsedInfoWithCallback invokes the actiontrail.GetAccessKeyLastUsedInfo API asynchronously
func (client *Client) GetAccessKeyLastUsedInfoWithCallback(request *GetAccessKeyLastUsedInfoRequest, callback func(response *GetAccessKeyLastUsedInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAccessKeyLastUsedInfoResponse
		var err error
		defer close(result)
		response, err = client.GetAccessKeyLastUsedInfo(request)
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

// GetAccessKeyLastUsedInfoRequest is the request struct for api GetAccessKeyLastUsedInfo
type GetAccessKeyLastUsedInfoRequest struct {
	*requests.RpcRequest
	AccessKey string `position:"Query" name:"AccessKey"`
}

// GetAccessKeyLastUsedInfoResponse is the response struct for api GetAccessKeyLastUsedInfo
type GetAccessKeyLastUsedInfoResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	AccessKeyId   string `json:"AccessKeyId" xml:"AccessKeyId"`
	AccountId     string `json:"AccountId" xml:"AccountId"`
	OwnerId       string `json:"OwnerId" xml:"OwnerId"`
	UserName      string `json:"UserName" xml:"UserName"`
	AccountType   string `json:"AccountType" xml:"AccountType"`
	UsedTimestamp int64  `json:"UsedTimestamp" xml:"UsedTimestamp"`
	Detail        string `json:"Detail" xml:"Detail"`
	Source        string `json:"Source" xml:"Source"`
	ServiceName   string `json:"ServiceName" xml:"ServiceName"`
	ServiceNameCn string `json:"ServiceNameCn" xml:"ServiceNameCn"`
	ServiceNameEn string `json:"ServiceNameEn" xml:"ServiceNameEn"`
}

// CreateGetAccessKeyLastUsedInfoRequest creates a request to invoke GetAccessKeyLastUsedInfo API
func CreateGetAccessKeyLastUsedInfoRequest() (request *GetAccessKeyLastUsedInfoRequest) {
	request = &GetAccessKeyLastUsedInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Actiontrail", "2020-07-06", "GetAccessKeyLastUsedInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAccessKeyLastUsedInfoResponse creates a response to parse from GetAccessKeyLastUsedInfo response
func CreateGetAccessKeyLastUsedInfoResponse() (response *GetAccessKeyLastUsedInfoResponse) {
	response = &GetAccessKeyLastUsedInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
