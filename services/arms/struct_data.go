package arms

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

// Data is a nested struct in arms response
type Data struct {
	EnableTag             bool                     `json:"EnableTag" xml:"EnableTag"`
	IsControllerInstalled bool                     `json:"isControllerInstalled" xml:"isControllerInstalled"`
	Info                  string                   `json:"Info" xml:"Info"`
	Version               string                   `json:"Version" xml:"Version"`
	PageSize              int                      `json:"PageSize" xml:"PageSize"`
	City                  string                   `json:"City" xml:"City"`
	District              string                   `json:"District" xml:"District"`
	CityCode              int64                    `json:"CityCode" xml:"CityCode"`
	NetServiceName        string                   `json:"NetServiceName" xml:"NetServiceName"`
	Success               bool                     `json:"Success" xml:"Success"`
	Page                  int                      `json:"Page" xml:"Page"`
	Total                 string                   `json:"Total" xml:"Total"`
	ClientType            int64                    `json:"ClientType" xml:"ClientType"`
	TaskId                int64                    `json:"TaskId" xml:"TaskId"`
	Size                  string                   `json:"Size" xml:"Size"`
	NetServiceId          int64                    `json:"NetServiceId" xml:"NetServiceId"`
	Msg                   string                   `json:"Msg" xml:"Msg"`
	UploadTime            string                   `json:"UploadTime" xml:"UploadTime"`
	Fid                   string                   `json:"Fid" xml:"Fid"`
	FileName              string                   `json:"FileName" xml:"FileName"`
	Busy                  int64                    `json:"Busy" xml:"Busy"`
	Items                 []map[string]interface{} `json:"Items" xml:"Items"`
	Products              []ProductsItem           `json:"Products" xml:"Products"`
}
