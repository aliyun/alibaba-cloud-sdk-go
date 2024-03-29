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

// MountNFS invokes the ehpc.MountNFS API synchronously
func (client *Client) MountNFS(request *MountNFSRequest) (response *MountNFSResponse, err error) {
	response = CreateMountNFSResponse()
	err = client.DoAction(request, response)
	return
}

// MountNFSWithChan invokes the ehpc.MountNFS API asynchronously
func (client *Client) MountNFSWithChan(request *MountNFSRequest) (<-chan *MountNFSResponse, <-chan error) {
	responseChan := make(chan *MountNFSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MountNFS(request)
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

// MountNFSWithCallback invokes the ehpc.MountNFS API asynchronously
func (client *Client) MountNFSWithCallback(request *MountNFSRequest, callback func(response *MountNFSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MountNFSResponse
		var err error
		defer close(result)
		response, err = client.MountNFS(request)
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

// MountNFSRequest is the request struct for api MountNFS
type MountNFSRequest struct {
	*requests.RpcRequest
	MountDir     string `position:"Query" name:"MountDir"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	RemoteDir    string `position:"Query" name:"RemoteDir"`
	NfsDir       string `position:"Query" name:"NfsDir"`
	ProtocolType string `position:"Query" name:"ProtocolType"`
}

// MountNFSResponse is the response struct for api MountNFS
type MountNFSResponse struct {
	*responses.BaseResponse
	InvokeId  string `json:"InvokeId" xml:"InvokeId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMountNFSRequest creates a request to invoke MountNFS API
func CreateMountNFSRequest() (request *MountNFSRequest) {
	request = &MountNFSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "MountNFS", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateMountNFSResponse creates a response to parse from MountNFS response
func CreateMountNFSResponse() (response *MountNFSResponse) {
	response = &MountNFSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
