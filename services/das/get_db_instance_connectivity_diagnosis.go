package das

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

// GetDBInstanceConnectivityDiagnosis invokes the das.GetDBInstanceConnectivityDiagnosis API synchronously
func (client *Client) GetDBInstanceConnectivityDiagnosis(request *GetDBInstanceConnectivityDiagnosisRequest) (response *GetDBInstanceConnectivityDiagnosisResponse, err error) {
	response = CreateGetDBInstanceConnectivityDiagnosisResponse()
	err = client.DoAction(request, response)
	return
}

// GetDBInstanceConnectivityDiagnosisWithChan invokes the das.GetDBInstanceConnectivityDiagnosis API asynchronously
func (client *Client) GetDBInstanceConnectivityDiagnosisWithChan(request *GetDBInstanceConnectivityDiagnosisRequest) (<-chan *GetDBInstanceConnectivityDiagnosisResponse, <-chan error) {
	responseChan := make(chan *GetDBInstanceConnectivityDiagnosisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDBInstanceConnectivityDiagnosis(request)
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

// GetDBInstanceConnectivityDiagnosisWithCallback invokes the das.GetDBInstanceConnectivityDiagnosis API asynchronously
func (client *Client) GetDBInstanceConnectivityDiagnosisWithCallback(request *GetDBInstanceConnectivityDiagnosisRequest, callback func(response *GetDBInstanceConnectivityDiagnosisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDBInstanceConnectivityDiagnosisResponse
		var err error
		defer close(result)
		response, err = client.GetDBInstanceConnectivityDiagnosis(request)
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

// GetDBInstanceConnectivityDiagnosisRequest is the request struct for api GetDBInstanceConnectivityDiagnosis
type GetDBInstanceConnectivityDiagnosisRequest struct {
	*requests.RpcRequest
	SrcIp      string `position:"Query" name:"SrcIp"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetDBInstanceConnectivityDiagnosisResponse is the response struct for api GetDBInstanceConnectivityDiagnosis
type GetDBInstanceConnectivityDiagnosisResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetDBInstanceConnectivityDiagnosisRequest creates a request to invoke GetDBInstanceConnectivityDiagnosis API
func CreateGetDBInstanceConnectivityDiagnosisRequest() (request *GetDBInstanceConnectivityDiagnosisRequest) {
	request = &GetDBInstanceConnectivityDiagnosisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetDBInstanceConnectivityDiagnosis", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDBInstanceConnectivityDiagnosisResponse creates a response to parse from GetDBInstanceConnectivityDiagnosis response
func CreateGetDBInstanceConnectivityDiagnosisResponse() (response *GetDBInstanceConnectivityDiagnosisResponse) {
	response = &GetDBInstanceConnectivityDiagnosisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
