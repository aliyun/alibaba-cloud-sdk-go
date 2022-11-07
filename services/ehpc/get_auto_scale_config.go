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

// GetAutoScaleConfig invokes the ehpc.GetAutoScaleConfig API synchronously
func (client *Client) GetAutoScaleConfig(request *GetAutoScaleConfigRequest) (response *GetAutoScaleConfigResponse, err error) {
	response = CreateGetAutoScaleConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetAutoScaleConfigWithChan invokes the ehpc.GetAutoScaleConfig API asynchronously
func (client *Client) GetAutoScaleConfigWithChan(request *GetAutoScaleConfigRequest) (<-chan *GetAutoScaleConfigResponse, <-chan error) {
	responseChan := make(chan *GetAutoScaleConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAutoScaleConfig(request)
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

// GetAutoScaleConfigWithCallback invokes the ehpc.GetAutoScaleConfig API asynchronously
func (client *Client) GetAutoScaleConfigWithCallback(request *GetAutoScaleConfigRequest, callback func(response *GetAutoScaleConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAutoScaleConfigResponse
		var err error
		defer close(result)
		response, err = client.GetAutoScaleConfig(request)
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

// GetAutoScaleConfigRequest is the request struct for api GetAutoScaleConfig
type GetAutoScaleConfigRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// GetAutoScaleConfigResponse is the response struct for api GetAutoScaleConfig
type GetAutoScaleConfigResponse struct {
	*responses.BaseResponse
	MaxNodesInCluster       int                        `json:"MaxNodesInCluster" xml:"MaxNodesInCluster"`
	GrowTimeoutInMinutes    int                        `json:"GrowTimeoutInMinutes" xml:"GrowTimeoutInMinutes"`
	SpotStrategy            string                     `json:"SpotStrategy" xml:"SpotStrategy"`
	EnableAutoShrink        bool                       `json:"EnableAutoShrink" xml:"EnableAutoShrink"`
	RequestId               string                     `json:"RequestId" xml:"RequestId"`
	EnableAutoGrow          bool                       `json:"EnableAutoGrow" xml:"EnableAutoGrow"`
	ClusterType             string                     `json:"ClusterType" xml:"ClusterType"`
	ExcludeNodes            string                     `json:"ExcludeNodes" xml:"ExcludeNodes"`
	ShrinkIntervalInMinutes int                        `json:"ShrinkIntervalInMinutes" xml:"ShrinkIntervalInMinutes"`
	GrowIntervalInMinutes   int                        `json:"GrowIntervalInMinutes" xml:"GrowIntervalInMinutes"`
	SpotPriceLimit          float64                    `json:"SpotPriceLimit" xml:"SpotPriceLimit"`
	ExtraNodesGrowRatio     int                        `json:"ExtraNodesGrowRatio" xml:"ExtraNodesGrowRatio"`
	ShrinkIdleTimes         int                        `json:"ShrinkIdleTimes" xml:"ShrinkIdleTimes"`
	ImageId                 string                     `json:"ImageId" xml:"ImageId"`
	GrowRatio               int                        `json:"GrowRatio" xml:"GrowRatio"`
	ClusterId               string                     `json:"ClusterId" xml:"ClusterId"`
	Uid                     string                     `json:"Uid" xml:"Uid"`
	Queues                  QueuesInGetAutoScaleConfig `json:"Queues" xml:"Queues"`
}

// CreateGetAutoScaleConfigRequest creates a request to invoke GetAutoScaleConfig API
func CreateGetAutoScaleConfigRequest() (request *GetAutoScaleConfigRequest) {
	request = &GetAutoScaleConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "GetAutoScaleConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateGetAutoScaleConfigResponse creates a response to parse from GetAutoScaleConfig response
func CreateGetAutoScaleConfigResponse() (response *GetAutoScaleConfigResponse) {
	response = &GetAutoScaleConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
