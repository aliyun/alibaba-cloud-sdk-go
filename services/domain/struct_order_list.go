package domain

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

// OrderList is a nested struct in domain response
type OrderList struct {
	OrderType  int    `json:"OrderType" xml:"OrderType"`
	BizId      string `json:"BizId" xml:"BizId"`
	UserId     string `json:"UserId" xml:"UserId"`
	Status     int    `json:"Status" xml:"Status"`
	Price      int64  `json:"Price" xml:"Price"`
	Domain     string `json:"Domain" xml:"Domain"`
	CreateTime int64  `json:"CreateTime" xml:"CreateTime"`
	UpdateTime int64  `json:"UpdateTime" xml:"UpdateTime"`
}
