package ehpc

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

// SetSchedulerInfo invokes the ehpc.SetSchedulerInfo API synchronously
func (client *Client) SetSchedulerInfo(request *SetSchedulerInfoRequest) (response *SetSchedulerInfoResponse, err error) {
	response = CreateSetSchedulerInfoResponse()
	err = client.DoAction(request, response)
	return
}

// SetSchedulerInfoWithChan invokes the ehpc.SetSchedulerInfo API asynchronously
func (client *Client) SetSchedulerInfoWithChan(request *SetSchedulerInfoRequest) (<-chan *SetSchedulerInfoResponse, <-chan error) {
	responseChan := make(chan *SetSchedulerInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetSchedulerInfo(request)
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

// SetSchedulerInfoWithCallback invokes the ehpc.SetSchedulerInfo API asynchronously
func (client *Client) SetSchedulerInfoWithCallback(request *SetSchedulerInfoRequest, callback func(response *SetSchedulerInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetSchedulerInfoResponse
		var err error
		defer close(result)
		response, err = client.SetSchedulerInfo(request)
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

// SetSchedulerInfoRequest is the request struct for api SetSchedulerInfo
type SetSchedulerInfoRequest struct {
	*requests.RpcRequest
	SlurmInfo *[]SetSchedulerInfoSlurmInfo `position:"Query" name:"SlurmInfo"  type:"Repeated"`
	ClusterId string                       `position:"Query" name:"ClusterId"`
	Scheduler *[]SetSchedulerInfoScheduler `position:"Query" name:"Scheduler"  type:"Repeated"`
	PbsInfo   *[]SetSchedulerInfoPbsInfo   `position:"Query" name:"PbsInfo"  type:"Repeated"`
}

// SetSchedulerInfoSlurmInfo is a repeated param struct in SetSchedulerInfoRequest
type SetSchedulerInfoSlurmInfo struct {
	SchedInterval    string `name:"SchedInterval"`
	BackfillInterval string `name:"BackfillInterval"`
}

// SetSchedulerInfoScheduler is a repeated param struct in SetSchedulerInfoRequest
type SetSchedulerInfoScheduler struct {
	SchedName string `name:"SchedName"`
}

// SetSchedulerInfoPbsInfo is a repeated param struct in SetSchedulerInfoRequest
type SetSchedulerInfoPbsInfo struct {
	SchedInterval      string                                  `name:"SchedInterval"`
	SchedMaxJobs       string                                  `name:"SchedMaxJobs"`
	AclLimit           *[]SetSchedulerInfoPbsInfoAclLimit      `name:"AclLimit" type:"Repeated"`
	ResourceLimit      *[]SetSchedulerInfoPbsInfoResourceLimit `name:"ResourceLimit" type:"Repeated"`
	SchedMaxQueuedJobs string                                  `name:"SchedMaxQueuedJobs"`
	JobHistoryDuration string                                  `name:"JobHistoryDuration"`
}

// SetSchedulerInfoPbsInfoAclLimit is a repeated param struct in SetSchedulerInfoRequest
type SetSchedulerInfoPbsInfoAclLimit struct {
	AclUsers string `name:"AclUsers"`
	Queue    string `name:"Queue"`
}

// SetSchedulerInfoPbsInfoResourceLimit is a repeated param struct in SetSchedulerInfoRequest
type SetSchedulerInfoPbsInfoResourceLimit struct {
	MaxJobs string `name:"MaxJobs"`
	Nodes   string `name:"Nodes"`
	Mem     string `name:"Mem"`
	Cpus    string `name:"Cpus"`
	User    string `name:"User"`
	Queue   string `name:"Queue"`
}

// SetSchedulerInfoResponse is the response struct for api SetSchedulerInfo
type SetSchedulerInfoResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetSchedulerInfoRequest creates a request to invoke SetSchedulerInfo API
func CreateSetSchedulerInfoRequest() (request *SetSchedulerInfoRequest) {
	request = &SetSchedulerInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "SetSchedulerInfo", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateSetSchedulerInfoResponse creates a response to parse from SetSchedulerInfo response
func CreateSetSchedulerInfoResponse() (response *SetSchedulerInfoResponse) {
	response = &SetSchedulerInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
