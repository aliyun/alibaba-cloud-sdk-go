package ddoscoo

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

// StatusCode is a nested struct in ddoscoo response
type StatusCode struct {
	Index     int   `json:"Index" xml:"Index"`
	Status502 int64 `json:"Status502" xml:"Status502"`
	Time      int64 `json:"Time" xml:"Time"`
	Status405 int64 `json:"Status405" xml:"Status405"`
	Status3XX int64 `json:"Status3XX" xml:"Status3XX"`
	Status503 int64 `json:"Status503" xml:"Status503"`
	Status4XX int64 `json:"Status4XX" xml:"Status4XX"`
	Status2XX int64 `json:"Status2XX" xml:"Status2XX"`
	Status5XX int64 `json:"Status5XX" xml:"Status5XX"`
	Status504 int64 `json:"Status504" xml:"Status504"`
	Status200 int64 `json:"Status200" xml:"Status200"`
	Status403 int64 `json:"Status403" xml:"Status403"`
	Status404 int64 `json:"Status404" xml:"Status404"`
	Status501 int64 `json:"Status501" xml:"Status501"`
}
