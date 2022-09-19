package threedvision

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

// Data is a nested struct in threedvision response
type Data struct {
	DepthVisUrl     string `json:"DepthVisUrl" xml:"DepthVisUrl"`
	JobId           string `json:"JobId" xml:"JobId"`
	PointCloudURL   string `json:"PointCloudURL" xml:"PointCloudURL"`
	Result          string `json:"Result" xml:"Result"`
	DisparityMapURL string `json:"DisparityMapURL" xml:"DisparityMapURL"`
	ErrorMessage    string `json:"ErrorMessage" xml:"ErrorMessage"`
	DisparityVisURL string `json:"DisparityVisURL" xml:"DisparityVisURL"`
	DepthUrl        string `json:"DepthUrl" xml:"DepthUrl"`
	DepthURL        string `json:"DepthURL" xml:"DepthURL"`
	DepthMapUrl     string `json:"DepthMapUrl" xml:"DepthMapUrl"`
	Status          string `json:"Status" xml:"Status"`
	MeshURL         string `json:"MeshURL" xml:"MeshURL"`
	ErrorCode       string `json:"ErrorCode" xml:"ErrorCode"`
	DepthToColorUrl string `json:"DepthToColorUrl" xml:"DepthToColorUrl"`
}
