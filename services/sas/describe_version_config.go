package sas

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

// DescribeVersionConfig invokes the sas.DescribeVersionConfig API synchronously
func (client *Client) DescribeVersionConfig(request *DescribeVersionConfigRequest) (response *DescribeVersionConfigResponse, err error) {
	response = CreateDescribeVersionConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVersionConfigWithChan invokes the sas.DescribeVersionConfig API asynchronously
func (client *Client) DescribeVersionConfigWithChan(request *DescribeVersionConfigRequest) (<-chan *DescribeVersionConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeVersionConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVersionConfig(request)
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

// DescribeVersionConfigWithCallback invokes the sas.DescribeVersionConfig API asynchronously
func (client *Client) DescribeVersionConfigWithCallback(request *DescribeVersionConfigRequest, callback func(response *DescribeVersionConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVersionConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeVersionConfig(request)
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

// DescribeVersionConfigRequest is the request struct for api DescribeVersionConfig
type DescribeVersionConfigRequest struct {
	*requests.RpcRequest
	SourceIp                   string `position:"Query" name:"SourceIp"`
	ResourceDirectoryAccountId string `position:"Query" name:"ResourceDirectoryAccountId"`
}

// DescribeVersionConfigResponse is the response struct for api DescribeVersionConfig
type DescribeVersionConfigResponse struct {
	*responses.BaseResponse
	MVAuthCount             int    `json:"MVAuthCount" xml:"MVAuthCount"`
	SasLog                  int    `json:"SasLog" xml:"SasLog"`
	LogCapacity             int    `json:"LogCapacity" xml:"LogCapacity"`
	SasScreen               int    `json:"SasScreen" xml:"SasScreen"`
	HoneypotCapacity        int64  `json:"HoneypotCapacity" xml:"HoneypotCapacity"`
	CreateTime              int64  `json:"CreateTime" xml:"CreateTime"`
	MVUnusedAuthCount       int    `json:"MVUnusedAuthCount" xml:"MVUnusedAuthCount"`
	WebLock                 int    `json:"WebLock" xml:"WebLock"`
	AppWhiteListAuthCount   int64  `json:"AppWhiteListAuthCount" xml:"AppWhiteListAuthCount"`
	RequestId               string `json:"RequestId" xml:"RequestId"`
	LogTime                 int    `json:"LogTime" xml:"LogTime"`
	Flag                    int    `json:"Flag" xml:"Flag"`
	LastInstanceReleaseTime int64  `json:"LastInstanceReleaseTime" xml:"LastInstanceReleaseTime"`
	LastTrailEndTime        int64  `json:"LastTrailEndTime" xml:"LastTrailEndTime"`
	Version                 int    `json:"Version" xml:"Version"`
	WebLockAuthCount        int64  `json:"WebLockAuthCount" xml:"WebLockAuthCount"`
	ReleaseTime             int64  `json:"ReleaseTime" xml:"ReleaseTime"`
	HighestVersion          int    `json:"HighestVersion" xml:"HighestVersion"`
	AssetLevel              int    `json:"AssetLevel" xml:"AssetLevel"`
	AvdsFlag                int    `json:"AvdsFlag" xml:"AvdsFlag"`
	IsPaidUser              bool   `json:"IsPaidUser" xml:"IsPaidUser"`
	IsOverBalance           bool   `json:"IsOverBalance" xml:"IsOverBalance"`
	InstanceId              string `json:"InstanceId" xml:"InstanceId"`
	SlsCapacity             int64  `json:"SlsCapacity" xml:"SlsCapacity"`
	VmCores                 int    `json:"VmCores" xml:"VmCores"`
	AllowPartialBuy         int    `json:"AllowPartialBuy" xml:"AllowPartialBuy"`
	AppWhiteList            int    `json:"AppWhiteList" xml:"AppWhiteList"`
	IsSasOpening            bool   `json:"IsSasOpening" xml:"IsSasOpening"`
	GmtCreate               int64  `json:"GmtCreate" xml:"GmtCreate"`
	ImageScanCapacity       int64  `json:"ImageScanCapacity" xml:"ImageScanCapacity"`
	IsTrialVersion          int    `json:"IsTrialVersion" xml:"IsTrialVersion"`
	UserDefinedAlarms       int    `json:"UserDefinedAlarms" xml:"UserDefinedAlarms"`
	OpenTime                int64  `json:"OpenTime" xml:"OpenTime"`
	IsNewContainerVersion   bool   `json:"IsNewContainerVersion" xml:"IsNewContainerVersion"`
}

// CreateDescribeVersionConfigRequest creates a request to invoke DescribeVersionConfig API
func CreateDescribeVersionConfigRequest() (request *DescribeVersionConfigRequest) {
	request = &DescribeVersionConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeVersionConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeVersionConfigResponse creates a response to parse from DescribeVersionConfig response
func CreateDescribeVersionConfigResponse() (response *DescribeVersionConfigResponse) {
	response = &DescribeVersionConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
