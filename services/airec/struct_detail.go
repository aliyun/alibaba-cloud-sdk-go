package airec

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

// Detail is a nested struct in airec response
type Detail struct {
	Author       string `json:"author" xml:"author"`
	Title        string `json:"title" xml:"title"`
	Duration     string `json:"duration" xml:"duration"`
	ItemId       string `json:"itemId" xml:"itemId"`
	PubTime      string `json:"pubTime" xml:"pubTime"`
	Channel      string `json:"channel" xml:"channel"`
	ItemType     string `json:"itemType" xml:"itemType"`
	BrandId      string `json:"brandId" xml:"brandId"`
	CategoryPath string `json:"categoryPath" xml:"categoryPath"`
	Status       string `json:"status" xml:"status"`
	ShopId       string `json:"shopId" xml:"shopId"`
	ExpireTime   string `json:"expireTime" xml:"expireTime"`
}
