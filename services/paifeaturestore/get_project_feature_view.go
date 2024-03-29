package paifeaturestore

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

// GetProjectFeatureView invokes the paifeaturestore.GetProjectFeatureView API synchronously
func (client *Client) GetProjectFeatureView(request *GetProjectFeatureViewRequest) (response *GetProjectFeatureViewResponse, err error) {
	response = CreateGetProjectFeatureViewResponse()
	err = client.DoAction(request, response)
	return
}

// GetProjectFeatureViewWithChan invokes the paifeaturestore.GetProjectFeatureView API asynchronously
func (client *Client) GetProjectFeatureViewWithChan(request *GetProjectFeatureViewRequest) (<-chan *GetProjectFeatureViewResponse, <-chan error) {
	responseChan := make(chan *GetProjectFeatureViewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProjectFeatureView(request)
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

// GetProjectFeatureViewWithCallback invokes the paifeaturestore.GetProjectFeatureView API asynchronously
func (client *Client) GetProjectFeatureViewWithCallback(request *GetProjectFeatureViewRequest, callback func(response *GetProjectFeatureViewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProjectFeatureViewResponse
		var err error
		defer close(result)
		response, err = client.GetProjectFeatureView(request)
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

// GetProjectFeatureViewRequest is the request struct for api GetProjectFeatureView
type GetProjectFeatureViewRequest struct {
	*requests.RoaRequest
	InstanceId      string `position:"Path" name:"InstanceId"`
	FeatureViewName string `position:"Path" name:"FeatureViewName"`
	ProjectId       string `position:"Path" name:"ProjectId"`
}

// GetProjectFeatureViewResponse is the response struct for api GetProjectFeatureView
type GetProjectFeatureViewResponse struct {
	*responses.BaseResponse
	RequestId            string       `json:"RequestId" xml:"RequestId"`
	FeatureViewId        string       `json:"FeatureViewId" xml:"FeatureViewId"`
	ProjectId            string       `json:"ProjectId" xml:"ProjectId"`
	ProjectName          string       `json:"ProjectName" xml:"ProjectName"`
	FeatureEntityId      string       `json:"FeatureEntityId" xml:"FeatureEntityId"`
	FeatureEntityName    string       `json:"FeatureEntityName" xml:"FeatureEntityName"`
	JoinId               string       `json:"JoinId" xml:"JoinId"`
	Name                 string       `json:"Name" xml:"Name"`
	Owner                string       `json:"Owner" xml:"Owner"`
	Type                 string       `json:"Type" xml:"Type"`
	WriteMethod          string       `json:"WriteMethod" xml:"WriteMethod"`
	RegisterTable        string       `json:"RegisterTable" xml:"RegisterTable"`
	RegisterDatasourceId string       `json:"RegisterDatasourceId" xml:"RegisterDatasourceId"`
	SyncOnlineTable      bool         `json:"SyncOnlineTable" xml:"SyncOnlineTable"`
	TTL                  int          `json:"TTL" xml:"TTL"`
	Config               string       `json:"Config" xml:"Config"`
	GmtSyncTime          string       `json:"GmtSyncTime" xml:"GmtSyncTime"`
	LastSyncConfig       string       `json:"LastSyncConfig" xml:"LastSyncConfig"`
	Tags                 []string     `json:"Tags" xml:"Tags"`
	Fields               []FieldsItem `json:"Fields" xml:"Fields"`
}

// CreateGetProjectFeatureViewRequest creates a request to invoke GetProjectFeatureView API
func CreateGetProjectFeatureViewRequest() (request *GetProjectFeatureViewRequest) {
	request = &GetProjectFeatureViewRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "GetProjectFeatureView", "/api/v1/instances/[InstanceId]/projects/[ProjectId]/featureviews/[FeatureViewName]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetProjectFeatureViewResponse creates a response to parse from GetProjectFeatureView response
func CreateGetProjectFeatureViewResponse() (response *GetProjectFeatureViewResponse) {
	response = &GetProjectFeatureViewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
