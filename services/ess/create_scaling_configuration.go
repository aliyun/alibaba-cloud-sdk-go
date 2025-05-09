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

// CreateScalingConfiguration invokes the ess.CreateScalingConfiguration API synchronously
func (client *Client) CreateScalingConfiguration(request *CreateScalingConfigurationRequest) (response *CreateScalingConfigurationResponse, err error) {
	response = CreateCreateScalingConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateScalingConfigurationWithChan invokes the ess.CreateScalingConfiguration API asynchronously
func (client *Client) CreateScalingConfigurationWithChan(request *CreateScalingConfigurationRequest) (<-chan *CreateScalingConfigurationResponse, <-chan error) {
	responseChan := make(chan *CreateScalingConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateScalingConfiguration(request)
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

// CreateScalingConfigurationWithCallback invokes the ess.CreateScalingConfiguration API asynchronously
func (client *Client) CreateScalingConfigurationWithCallback(request *CreateScalingConfigurationRequest, callback func(response *CreateScalingConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateScalingConfigurationResponse
		var err error
		defer close(result)
		response, err = client.CreateScalingConfiguration(request)
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

// CreateScalingConfigurationRequest is the request struct for api CreateScalingConfiguration
type CreateScalingConfigurationRequest struct {
	*requests.RpcRequest
	HpcClusterId                    string                                            `position:"Query" name:"HpcClusterId"`
	SecurityOptions                 CreateScalingConfigurationSecurityOptions         `position:"Query" name:"SecurityOptions"  type:"Struct"`
	SecurityEnhancementStrategy     string                                            `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                     string                                            `position:"Query" name:"KeyPairName"`
	SpotPriceLimit                  *[]CreateScalingConfigurationSpotPriceLimit       `position:"Query" name:"SpotPriceLimit"  type:"Repeated"`
	DeletionProtection              requests.Boolean                                  `position:"Query" name:"DeletionProtection"`
	ResourceGroupId                 string                                            `position:"Query" name:"ResourceGroupId"`
	PrivatePoolOptionsMatchCriteria string                                            `position:"Query" name:"PrivatePoolOptions.MatchCriteria"`
	HostName                        string                                            `position:"Query" name:"HostName"`
	Password                        string                                            `position:"Query" name:"Password"`
	InstanceDescription             string                                            `position:"Query" name:"InstanceDescription"`
	StorageSetPartitionNumber       requests.Integer                                  `position:"Query" name:"StorageSetPartitionNumber"`
	CustomPriorities                *[]CreateScalingConfigurationCustomPriorities     `position:"Query" name:"CustomPriorities"  type:"Repeated"`
	SystemDiskAutoSnapshotPolicyId  string                                            `position:"Query" name:"SystemDisk.AutoSnapshotPolicyId"`
	PrivatePoolOptionsId            string                                            `position:"Query" name:"PrivatePoolOptions.Id"`
	ImageOptionsLoginAsNonRoot      requests.Boolean                                  `position:"Query" name:"ImageOptions.LoginAsNonRoot"`
	Ipv6AddressCount                requests.Integer                                  `position:"Query" name:"Ipv6AddressCount"`
	Cpu                             requests.Integer                                  `position:"Query" name:"Cpu"`
	SystemDiskCategories            *[]string                                         `position:"Query" name:"SystemDiskCategories"  type:"Repeated"`
	OwnerId                         requests.Integer                                  `position:"Query" name:"OwnerId"`
	ScalingConfigurationName        string                                            `position:"Query" name:"ScalingConfigurationName"`
	Tags                            string                                            `position:"Query" name:"Tags"`
	SpotStrategy                    string                                            `position:"Query" name:"SpotStrategy"`
	CpuOptions                      CreateScalingConfigurationCpuOptions              `position:"Query" name:"CpuOptions"  type:"Struct"`
	SystemDiskBurstingEnabled       requests.Boolean                                  `position:"Query" name:"SystemDisk.BurstingEnabled"`
	InstanceName                    string                                            `position:"Query" name:"InstanceName"`
	InternetChargeType              string                                            `position:"Query" name:"InternetChargeType"`
	ZoneId                          string                                            `position:"Query" name:"ZoneId"`
	InternetMaxBandwidthIn          requests.Integer                                  `position:"Query" name:"InternetMaxBandwidthIn"`
	InstancePatternInfo             *[]CreateScalingConfigurationInstancePatternInfo  `position:"Query" name:"InstancePatternInfo"  type:"Repeated"`
	Affinity                        string                                            `position:"Query" name:"Affinity"`
	NetworkInterfaces               *[]CreateScalingConfigurationNetworkInterfaces    `position:"Query" name:"NetworkInterfaces"  type:"Repeated"`
	ImageId                         string                                            `position:"Query" name:"ImageId"`
	Memory                          requests.Integer                                  `position:"Query" name:"Memory"`
	ClientToken                     string                                            `position:"Query" name:"ClientToken"`
	SpotInterruptionBehavior        string                                            `position:"Query" name:"SpotInterruptionBehavior"`
	ScalingGroupId                  string                                            `position:"Query" name:"ScalingGroupId"`
	IoOptimized                     string                                            `position:"Query" name:"IoOptimized"`
	InstanceTypes                   *[]string                                         `position:"Query" name:"InstanceTypes"  type:"Repeated"`
	SecurityGroupId                 string                                            `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut         requests.Integer                                  `position:"Query" name:"InternetMaxBandwidthOut"`
	SystemDiskKMSKeyId              string                                            `position:"Query" name:"SystemDisk.KMSKeyId"`
	SystemDiskCategory              string                                            `position:"Query" name:"SystemDisk.Category"`
	SystemDiskPerformanceLevel      string                                            `position:"Query" name:"SystemDisk.PerformanceLevel"`
	UserData                        string                                            `position:"Query" name:"UserData"`
	PasswordInherit                 requests.Boolean                                  `position:"Query" name:"PasswordInherit"`
	ResourcePoolOptions             CreateScalingConfigurationResourcePoolOptions     `position:"Query" name:"ResourcePoolOptions"  type:"Struct"`
	ImageName                       string                                            `position:"Query" name:"ImageName"`
	HttpEndpoint                    string                                            `position:"Query" name:"HttpEndpoint"`
	DedicatedHostClusterId          string                                            `position:"Query" name:"DedicatedHostClusterId"`
	InstanceType                    string                                            `position:"Query" name:"InstanceType"`
	SchedulerOptions                map[string]interface{}                            `position:"Query" name:"SchedulerOptions"`
	DeploymentSetId                 string                                            `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerAccount            string                                            `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                    string                                            `position:"Query" name:"OwnerAccount"`
	Tenancy                         string                                            `position:"Query" name:"Tenancy"`
	SystemDiskDiskName              string                                            `position:"Query" name:"SystemDisk.DiskName"`
	RamRoleName                     string                                            `position:"Query" name:"RamRoleName"`
	SystemDiskEncryptAlgorithm      string                                            `position:"Query" name:"SystemDisk.EncryptAlgorithm"`
	DedicatedHostId                 string                                            `position:"Query" name:"DedicatedHostId"`
	CreditSpecification             string                                            `position:"Query" name:"CreditSpecification"`
	SpotDuration                    requests.Integer                                  `position:"Query" name:"SpotDuration"`
	SecurityGroupIds                *[]string                                         `position:"Query" name:"SecurityGroupIds"  type:"Repeated"`
	DataDisk                        *[]CreateScalingConfigurationDataDisk             `position:"Query" name:"DataDisk"  type:"Repeated"`
	InstanceTypeOverride            *[]CreateScalingConfigurationInstanceTypeOverride `position:"Query" name:"InstanceTypeOverride"  type:"Repeated"`
	SystemDiskProvisionedIops       requests.Integer                                  `position:"Query" name:"SystemDisk.ProvisionedIops"`
	LoadBalancerWeight              requests.Integer                                  `position:"Query" name:"LoadBalancerWeight"`
	StorageSetId                    string                                            `position:"Query" name:"StorageSetId"`
	SystemDiskSize                  requests.Integer                                  `position:"Query" name:"SystemDisk.Size"`
	ImageFamily                     string                                            `position:"Query" name:"ImageFamily"`
	HttpTokens                      string                                            `position:"Query" name:"HttpTokens"`
	SystemDiskDescription           string                                            `position:"Query" name:"SystemDisk.Description"`
	SystemDiskEncrypted             requests.Boolean                                  `position:"Query" name:"SystemDisk.Encrypted"`
}

// CreateScalingConfigurationSecurityOptions is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationSecurityOptions struct {
	ConfidentialComputingMode string `name:"ConfidentialComputingMode"`
}

// CreateScalingConfigurationSpotPriceLimit is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationSpotPriceLimit struct {
	InstanceType string `name:"InstanceType"`
	PriceLimit   string `name:"PriceLimit"`
}

// CreateScalingConfigurationCustomPriorities is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationCustomPriorities struct {
	VswitchId    string `name:"VswitchId"`
	InstanceType string `name:"InstanceType"`
}

// CreateScalingConfigurationCpuOptions is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationCpuOptions struct {
	Accelerators *[]string `name:"Accelerators" type:"Repeated"`
}

// CreateScalingConfigurationInstancePatternInfo is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationInstancePatternInfo struct {
	InstanceCategories                 *[]string `name:"InstanceCategories" type:"Repeated"`
	Memory                             string    `name:"Memory"`
	MaximumCpuCoreCount                string    `name:"MaximumCpuCoreCount"`
	MaxPrice                           string    `name:"MaxPrice"`
	MinimumGpuAmount                   string    `name:"MinimumGpuAmount"`
	MaximumGpuAmount                   string    `name:"MaximumGpuAmount"`
	MaximumMemorySize                  string    `name:"MaximumMemorySize"`
	MinimumInitialCredit               string    `name:"MinimumInitialCredit"`
	ExcludedInstanceType               *[]string `name:"ExcludedInstanceType" type:"Repeated"`
	MinimumEniIpv6AddressQuantity      string    `name:"MinimumEniIpv6AddressQuantity"`
	MinimumEniPrivateIpAddressQuantity string    `name:"MinimumEniPrivateIpAddressQuantity"`
	BurstablePerformance               string    `name:"BurstablePerformance"`
	PhysicalProcessorModels            *[]string `name:"PhysicalProcessorModels" type:"Repeated"`
	MinimumCpuCoreCount                string    `name:"MinimumCpuCoreCount"`
	GpuSpecs                           *[]string `name:"GpuSpecs" type:"Repeated"`
	MinimumEniQuantity                 string    `name:"MinimumEniQuantity"`
	MinimumMemorySize                  string    `name:"MinimumMemorySize"`
	Cores                              string    `name:"Cores"`
	InstanceFamilyLevel                string    `name:"InstanceFamilyLevel"`
	InstanceTypeFamilies               *[]string `name:"InstanceTypeFamilies" type:"Repeated"`
	MinimumBaselineCredit              string    `name:"MinimumBaselineCredit"`
	CpuArchitectures                   *[]string `name:"CpuArchitectures" type:"Repeated"`
	Architecture                       *[]string `name:"Architecture" type:"Repeated"`
}

// CreateScalingConfigurationNetworkInterfaces is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationNetworkInterfaces struct {
	Ipv6AddressCount            string    `name:"Ipv6AddressCount"`
	InstanceType                string    `name:"InstanceType"`
	SecurityGroupIds            *[]string `name:"SecurityGroupIds" type:"Repeated"`
	NetworkInterfaceTrafficMode string    `name:"NetworkInterfaceTrafficMode"`
}

// CreateScalingConfigurationResourcePoolOptions is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationResourcePoolOptions struct {
	Strategy       string    `name:"Strategy"`
	PrivatePoolIds *[]string `name:"PrivatePoolIds" type:"Repeated"`
}

// CreateScalingConfigurationDataDisk is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationDataDisk struct {
	SnapshotId           string    `name:"SnapshotId"`
	PerformanceLevel     string    `name:"PerformanceLevel"`
	AutoSnapshotPolicyId string    `name:"AutoSnapshotPolicyId"`
	Description          string    `name:"Description"`
	BurstingEnabled      string    `name:"BurstingEnabled"`
	DiskName             string    `name:"DiskName"`
	ProvisionedIops      string    `name:"ProvisionedIops"`
	Encrypted            string    `name:"Encrypted"`
	Size                 string    `name:"Size"`
	Categories           *[]string `name:"Categories" type:"Repeated"`
	Category             string    `name:"Category"`
	KMSKeyId             string    `name:"KMSKeyId"`
	Device               string    `name:"Device"`
	DeleteWithInstance   string    `name:"DeleteWithInstance"`
}

// CreateScalingConfigurationInstanceTypeOverride is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationInstanceTypeOverride struct {
	WeightedCapacity string `name:"WeightedCapacity"`
	InstanceType     string `name:"InstanceType"`
}

// CreateScalingConfigurationResponse is the response struct for api CreateScalingConfiguration
type CreateScalingConfigurationResponse struct {
	*responses.BaseResponse
	ScalingConfigurationId string `json:"ScalingConfigurationId" xml:"ScalingConfigurationId"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateScalingConfigurationRequest creates a request to invoke CreateScalingConfiguration API
func CreateCreateScalingConfigurationRequest() (request *CreateScalingConfigurationRequest) {
	request = &CreateScalingConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingConfiguration", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateScalingConfigurationResponse creates a response to parse from CreateScalingConfiguration response
func CreateCreateScalingConfigurationResponse() (response *CreateScalingConfigurationResponse) {
	response = &CreateScalingConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
