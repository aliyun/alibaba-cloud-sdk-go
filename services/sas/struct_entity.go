package sas

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

// Entity is a nested struct in sas response
type Entity struct {
	Uuid         string `json:"Uuid" xml:"Uuid"`
	GroupId      int    `json:"GroupId" xml:"GroupId"`
	InternetIp   string `json:"InternetIp" xml:"InternetIp"`
	InstanceName string `json:"InstanceName" xml:"InstanceName"`
	Ip           string `json:"Ip" xml:"Ip"`
	Os           string `json:"Os" xml:"Os"`
	IntranetIp   string `json:"IntranetIp" xml:"IntranetIp"`
}
