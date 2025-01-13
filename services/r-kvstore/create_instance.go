package r_kvstore

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

// CreateInstance invokes the r_kvstore.CreateInstance API synchronously
func (client *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
	response = CreateCreateInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateInstanceWithChan invokes the r_kvstore.CreateInstance API asynchronously
func (client *Client) CreateInstanceWithChan(request *CreateInstanceRequest) (<-chan *CreateInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateInstance(request)
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

// CreateInstanceWithCallback invokes the r_kvstore.CreateInstance API asynchronously
func (client *Client) CreateInstanceWithCallback(request *CreateInstanceRequest, callback func(response *CreateInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateInstance(request)
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

// CreateInstanceRequest is the request struct for api CreateInstance
type CreateInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer     `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string               `position:"Query" name:"ConnectionStringPrefix"`
	SecondaryZoneId        string               `position:"Query" name:"SecondaryZoneId"`
	SlaveReadOnlyCount     requests.Integer     `position:"Query" name:"SlaveReadOnlyCount"`
	CouponNo               string               `position:"Query" name:"CouponNo"`
	NetworkType            string               `position:"Query" name:"NetworkType"`
	EngineVersion          string               `position:"Query" name:"EngineVersion"`
	PhysicalInstanceId     string               `position:"Query" name:"PhysicalInstanceId"`
	ResourceGroupId        string               `position:"Query" name:"ResourceGroupId"`
	Password               string               `position:"Query" name:"Password"`
	SecurityToken          string               `position:"Query" name:"SecurityToken"`
	Tag                    *[]CreateInstanceTag `position:"Query" name:"Tag"  type:"Repeated"`
	GlobalSecurityGroupIds string               `position:"Query" name:"GlobalSecurityGroupIds"`
	BusinessInfo           string               `position:"Query" name:"BusinessInfo"`
	ShardCount             requests.Integer     `position:"Query" name:"ShardCount"`
	AutoRenewPeriod        string               `position:"Query" name:"AutoRenewPeriod"`
	Period                 string               `position:"Query" name:"Period"`
	DryRun                 requests.Boolean     `position:"Query" name:"DryRun"`
	BackupId               string               `position:"Query" name:"BackupId"`
	OwnerId                requests.Integer     `position:"Query" name:"OwnerId"`
	VSwitchId              string               `position:"Query" name:"VSwitchId"`
	PrivateIpAddress       string               `position:"Query" name:"PrivateIpAddress"`
	InstanceName           string               `position:"Query" name:"InstanceName"`
	AutoRenew              string               `position:"Query" name:"AutoRenew"`
	Port                   string               `position:"Query" name:"Port"`
	ZoneId                 string               `position:"Query" name:"ZoneId"`
	ReplicaCount           requests.Integer     `position:"Query" name:"ReplicaCount"`
	Appendonly             string               `position:"Query" name:"Appendonly"`
	NodeType               string               `position:"Query" name:"NodeType"`
	AutoUseCoupon          string               `position:"Query" name:"AutoUseCoupon"`
	InstanceClass          string               `position:"Query" name:"InstanceClass"`
	Capacity               requests.Integer     `position:"Query" name:"Capacity"`
	InstanceType           string               `position:"Query" name:"InstanceType"`
	DedicatedHostGroupId   string               `position:"Query" name:"DedicatedHostGroupId"`
	RestoreTime            string               `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount   string               `position:"Query" name:"ResourceOwnerAccount"`
	SrcDBInstanceId        string               `position:"Query" name:"SrcDBInstanceId"`
	OwnerAccount           string               `position:"Query" name:"OwnerAccount"`
	GlobalInstance         requests.Boolean     `position:"Query" name:"GlobalInstance"`
	RecoverConfigMode      string               `position:"Query" name:"RecoverConfigMode"`
	Token                  string               `position:"Query" name:"Token"`
	GlobalInstanceId       string               `position:"Query" name:"GlobalInstanceId"`
	ParamGroupId           string               `position:"Query" name:"ParamGroupId"`
	VpcId                  string               `position:"Query" name:"VpcId"`
	DeletePhysicalInstance requests.Boolean     `position:"Query" name:"DeletePhysicalInstance"`
	ReadOnlyCount          requests.Integer     `position:"Query" name:"ReadOnlyCount"`
	ChargeType             string               `position:"Query" name:"ChargeType"`
	Config                 string               `position:"Query" name:"Config"`
	SlaveReplicaCount      requests.Integer     `position:"Query" name:"SlaveReplicaCount"`
	ClusterBackupId        string               `position:"Query" name:"ClusterBackupId"`
}

// CreateInstanceTag is a repeated param struct in CreateInstanceRequest
type CreateInstanceTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateInstanceResponse is the response struct for api CreateInstance
type CreateInstanceResponse struct {
	*responses.BaseResponse
	VpcId            string `json:"VpcId" xml:"VpcId"`
	QPS              int64  `json:"QPS" xml:"QPS"`
	Capacity         int64  `json:"Capacity" xml:"Capacity"`
	ConnectionDomain string `json:"ConnectionDomain" xml:"ConnectionDomain"`
	ChargeType       string `json:"ChargeType" xml:"ChargeType"`
	NetworkType      string `json:"NetworkType" xml:"NetworkType"`
	InstanceId       string `json:"InstanceId" xml:"InstanceId"`
	Port             int    `json:"Port" xml:"Port"`
	Config           string `json:"Config" xml:"Config"`
	RegionId         string `json:"RegionId" xml:"RegionId"`
	EndTime          string `json:"EndTime" xml:"EndTime"`
	VSwitchId        string `json:"VSwitchId" xml:"VSwitchId"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	NodeType         string `json:"NodeType" xml:"NodeType"`
	Connections      int64  `json:"Connections" xml:"Connections"`
	Bandwidth        int64  `json:"Bandwidth" xml:"Bandwidth"`
	InstanceName     string `json:"InstanceName" xml:"InstanceName"`
	ZoneId           string `json:"ZoneId" xml:"ZoneId"`
	InstanceStatus   string `json:"InstanceStatus" xml:"InstanceStatus"`
	PrivateIpAddr    string `json:"PrivateIpAddr" xml:"PrivateIpAddr"`
	UserName         string `json:"UserName" xml:"UserName"`
	OrderId          int64  `json:"OrderId" xml:"OrderId"`
}

// CreateCreateInstanceRequest creates a request to invoke CreateInstance API
func CreateCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateInstance", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateInstanceResponse creates a response to parse from CreateInstance response
func CreateCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
