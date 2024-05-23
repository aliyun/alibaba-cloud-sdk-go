package ecs

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

// AttachDisk invokes the ecs.AttachDisk API synchronously
func (client *Client) AttachDisk(request *AttachDiskRequest) (response *AttachDiskResponse, err error) {
	response = CreateAttachDiskResponse()
	err = client.DoAction(request, response)
	return
}

// AttachDiskWithChan invokes the ecs.AttachDisk API asynchronously
func (client *Client) AttachDiskWithChan(request *AttachDiskRequest) (<-chan *AttachDiskResponse, <-chan error) {
	responseChan := make(chan *AttachDiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachDisk(request)
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

// AttachDiskWithCallback invokes the ecs.AttachDisk API asynchronously
func (client *Client) AttachDiskWithCallback(request *AttachDiskRequest, callback func(response *AttachDiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachDiskResponse
		var err error
		defer close(result)
		response, err = client.AttachDisk(request)
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

// AttachDiskRequest is the request struct for api AttachDisk
type AttachDiskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	KeyPairName          string           `position:"Query" name:"KeyPairName"`
	Bootable             requests.Boolean `position:"Query" name:"Bootable"`
	Password             string           `position:"Query" name:"Password"`
	DiskId               string           `position:"Query" name:"DiskId"`
	DeleteWithInstance   requests.Boolean `position:"Query" name:"DeleteWithInstance"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Force                requests.Boolean `position:"Query" name:"Force"`
	Device               string           `position:"Query" name:"Device"`
}

// AttachDiskResponse is the response struct for api AttachDisk
type AttachDiskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachDiskRequest creates a request to invoke AttachDisk API
func CreateAttachDiskRequest() (request *AttachDiskRequest) {
	request = &AttachDiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AttachDisk", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachDiskResponse creates a response to parse from AttachDisk response
func CreateAttachDiskResponse() (response *AttachDiskResponse) {
	response = &AttachDiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
