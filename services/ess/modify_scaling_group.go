package ess

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

// ModifyScalingGroup invokes the ess.ModifyScalingGroup API synchronously
func (client *Client) ModifyScalingGroup(request *ModifyScalingGroupRequest) (response *ModifyScalingGroupResponse, err error) {
	response = CreateModifyScalingGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyScalingGroupWithChan invokes the ess.ModifyScalingGroup API asynchronously
func (client *Client) ModifyScalingGroupWithChan(request *ModifyScalingGroupRequest) (<-chan *ModifyScalingGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyScalingGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScalingGroup(request)
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

// ModifyScalingGroupWithCallback invokes the ess.ModifyScalingGroup API asynchronously
func (client *Client) ModifyScalingGroupWithCallback(request *ModifyScalingGroupRequest, callback func(response *ModifyScalingGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScalingGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyScalingGroup(request)
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

// ModifyScalingGroupRequest is the request struct for api ModifyScalingGroup
type ModifyScalingGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                     requests.Integer                            `position:"Query" name:"ResourceOwnerId"`
	AzBalance                           requests.Boolean                            `position:"Query" name:"AzBalance"`
	VSwitchIds                          *[]string                                   `position:"Query" name:"VSwitchIds"  type:"Repeated"`
	MaxInstanceLifetime                 requests.Integer                            `position:"Query" name:"MaxInstanceLifetime"`
	ActiveScalingConfigurationId        string                                      `position:"Query" name:"ActiveScalingConfigurationId"`
	SpotInstanceRemedy                  requests.Boolean                            `position:"Query" name:"SpotInstanceRemedy"`
	ScaleOutAmountCheck                 requests.Boolean                            `position:"Query" name:"ScaleOutAmountCheck"`
	CustomPolicyARN                     string                                      `position:"Query" name:"CustomPolicyARN"`
	StopInstanceTimeout                 requests.Integer                            `position:"Query" name:"StopInstanceTimeout"`
	DefaultCooldown                     requests.Integer                            `position:"Query" name:"DefaultCooldown"`
	HealthCheckTypes                    *[]string                                   `position:"Query" name:"HealthCheckTypes"  type:"Repeated"`
	MultiAZPolicy                       string                                      `position:"Query" name:"MultiAZPolicy"`
	LaunchTemplateId                    string                                      `position:"Query" name:"LaunchTemplateId"`
	DesiredCapacity                     requests.Integer                            `position:"Query" name:"DesiredCapacity"`
	LaunchTemplateOverride              *[]ModifyScalingGroupLaunchTemplateOverride `position:"Query" name:"LaunchTemplateOverride"  type:"Repeated"`
	CompensateWithOnDemand              requests.Boolean                            `position:"Query" name:"CompensateWithOnDemand"`
	CapacityOptions                     ModifyScalingGroupCapacityOptions           `position:"Query" name:"CapacityOptions"  type:"Struct"`
	MinSize                             requests.Integer                            `position:"Query" name:"MinSize"`
	OwnerId                             requests.Integer                            `position:"Query" name:"OwnerId"`
	MaxSize                             requests.Integer                            `position:"Query" name:"MaxSize"`
	ScalingGroupId                      string                                      `position:"Query" name:"ScalingGroupId"`
	OnDemandBaseCapacity                requests.Integer                            `position:"Query" name:"OnDemandBaseCapacity"`
	OnDemandPercentageAboveBaseCapacity requests.Integer                            `position:"Query" name:"OnDemandPercentageAboveBaseCapacity"`
	SpotAllocationStrategy              string                                      `position:"Query" name:"SpotAllocationStrategy"`
	DisableDesiredCapacity              requests.Boolean                            `position:"Query" name:"DisableDesiredCapacity"`
	RemovalPolicy1                      string                                      `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2                      string                                      `position:"Query" name:"RemovalPolicy.2"`
	RemovalPolicy3                      string                                      `position:"Query" name:"RemovalPolicy.3"`
	HealthCheckType                     string                                      `position:"Query" name:"HealthCheckType"`
	ResourceOwnerAccount                string                                      `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName                    string                                      `position:"Query" name:"ScalingGroupName"`
	OwnerAccount                        string                                      `position:"Query" name:"OwnerAccount"`
	SpotInstancePools                   requests.Integer                            `position:"Query" name:"SpotInstancePools"`
	GroupDeletionProtection             requests.Boolean                            `position:"Query" name:"GroupDeletionProtection"`
	LaunchTemplateVersion               string                                      `position:"Query" name:"LaunchTemplateVersion"`
	ScalingPolicy                       string                                      `position:"Query" name:"ScalingPolicy"`
	AllocationStrategy                  string                                      `position:"Query" name:"AllocationStrategy"`
}

// ModifyScalingGroupLaunchTemplateOverride is a repeated param struct in ModifyScalingGroupRequest
type ModifyScalingGroupLaunchTemplateOverride struct {
	WeightedCapacity string `name:"WeightedCapacity"`
	InstanceType     string `name:"InstanceType"`
	SpotPriceLimit   string `name:"SpotPriceLimit"`
}

// ModifyScalingGroupCapacityOptions is a repeated param struct in ModifyScalingGroupRequest
type ModifyScalingGroupCapacityOptions struct {
	CompensateWithOnDemand              string `name:"CompensateWithOnDemand"`
	OnDemandBaseCapacity                string `name:"OnDemandBaseCapacity"`
	SpotAutoReplaceOnDemand             string `name:"SpotAutoReplaceOnDemand"`
	OnDemandPercentageAboveBaseCapacity string `name:"OnDemandPercentageAboveBaseCapacity"`
}

// ModifyScalingGroupResponse is the response struct for api ModifyScalingGroup
type ModifyScalingGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyScalingGroupRequest creates a request to invoke ModifyScalingGroup API
func CreateModifyScalingGroupRequest() (request *ModifyScalingGroupRequest) {
	request = &ModifyScalingGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ModifyScalingGroup", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyScalingGroupResponse creates a response to parse from ModifyScalingGroup response
func CreateModifyScalingGroupResponse() (response *ModifyScalingGroupResponse) {
	response = &ModifyScalingGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
