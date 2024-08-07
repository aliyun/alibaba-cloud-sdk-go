package elasticsearch

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

// ResultInDescribeLogstash is a nested struct in elasticsearch response
type ResultInDescribeLogstash struct {
	Config          map[string]interface{}     `json:"config" xml:"config"`
	PaymentType     string                     `json:"paymentType" xml:"paymentType"`
	ResourceGroupId string                     `json:"ResourceGroupId" xml:"ResourceGroupId"`
	NodeAmount      int                        `json:"nodeAmount" xml:"nodeAmount"`
	Description     string                     `json:"description" xml:"description"`
	CreatedAt       string                     `json:"createdAt" xml:"createdAt"`
	Status          string                     `json:"status" xml:"status"`
	VpcInstanceId   string                     `json:"vpcInstanceId" xml:"vpcInstanceId"`
	UpdatedAt       string                     `json:"updatedAt" xml:"updatedAt"`
	Version         string                     `json:"version" xml:"version"`
	InstanceId      string                     `json:"instanceId" xml:"instanceId"`
	ExtendConfigs   []map[string]interface{}   `json:"ExtendConfigs" xml:"ExtendConfigs"`
	NodeSpec        NodeSpecInDescribeLogstash `json:"nodeSpec" xml:"nodeSpec"`
	NetworkConfig   NetworkConfig              `json:"networkConfig" xml:"networkConfig"`
	EndpointList    []Endpoint                 `json:"endpointList" xml:"endpointList"`
	Tags            []TagsItem                 `json:"Tags" xml:"Tags"`
	ZoneInfos       []ZoneInfosItem            `json:"ZoneInfos" xml:"ZoneInfos"`
}
