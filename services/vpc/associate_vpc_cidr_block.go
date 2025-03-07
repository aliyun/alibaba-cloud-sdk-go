package vpc

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

// AssociateVpcCidrBlock invokes the vpc.AssociateVpcCidrBlock API synchronously
func (client *Client) AssociateVpcCidrBlock(request *AssociateVpcCidrBlockRequest) (response *AssociateVpcCidrBlockResponse, err error) {
	response = CreateAssociateVpcCidrBlockResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateVpcCidrBlockWithChan invokes the vpc.AssociateVpcCidrBlock API asynchronously
func (client *Client) AssociateVpcCidrBlockWithChan(request *AssociateVpcCidrBlockRequest) (<-chan *AssociateVpcCidrBlockResponse, <-chan error) {
	responseChan := make(chan *AssociateVpcCidrBlockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateVpcCidrBlock(request)
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

// AssociateVpcCidrBlockWithCallback invokes the vpc.AssociateVpcCidrBlock API asynchronously
func (client *Client) AssociateVpcCidrBlockWithCallback(request *AssociateVpcCidrBlockRequest, callback func(response *AssociateVpcCidrBlockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateVpcCidrBlockResponse
		var err error
		defer close(result)
		response, err = client.AssociateVpcCidrBlock(request)
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

// AssociateVpcCidrBlockRequest is the request struct for api AssociateVpcCidrBlock
type AssociateVpcCidrBlockRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	IPv6CidrType         string           `position:"Query" name:"IPv6CidrType"`
	IpamPoolId           string           `position:"Query" name:"IpamPoolId"`
	Ipv6Isp              string           `position:"Query" name:"Ipv6Isp"`
	IpVersion            string           `position:"Query" name:"IpVersion"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	IPv6CidrBlock        string           `position:"Query" name:"IPv6CidrBlock"`
	SecondaryCidrMask    requests.Integer `position:"Query" name:"SecondaryCidrMask"`
	SecondaryCidrBlock   string           `position:"Query" name:"SecondaryCidrBlock"`
	VpcId                string           `position:"Query" name:"VpcId"`
}

// AssociateVpcCidrBlockResponse is the response struct for api AssociateVpcCidrBlock
type AssociateVpcCidrBlockResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateVpcCidrBlockRequest creates a request to invoke AssociateVpcCidrBlock API
func CreateAssociateVpcCidrBlockRequest() (request *AssociateVpcCidrBlockRequest) {
	request = &AssociateVpcCidrBlockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AssociateVpcCidrBlock", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateVpcCidrBlockResponse creates a response to parse from AssociateVpcCidrBlock response
func CreateAssociateVpcCidrBlockResponse() (response *AssociateVpcCidrBlockResponse) {
	response = &AssociateVpcCidrBlockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
