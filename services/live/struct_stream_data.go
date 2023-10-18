package live

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

// StreamData is a nested struct in live response
type StreamData struct {
	MediaProfile string  `json:"MediaProfile" xml:"MediaProfile"`
	RtsTraffic   int64   `json:"RtsTraffic" xml:"RtsTraffic"`
	RtmpTraffic  int64   `json:"RtmpTraffic" xml:"RtmpTraffic"`
	HlsCount     int64   `json:"HlsCount" xml:"HlsCount"`
	HlsTraffic   int64   `json:"HlsTraffic" xml:"HlsTraffic"`
	ReqBps       float64 `json:"ReqBps" xml:"ReqBps"`
	MediaType    string  `json:"MediaType" xml:"MediaType"`
	P2pTraffic   int64   `json:"P2pTraffic" xml:"P2pTraffic"`
	StreamName   string  `json:"StreamName" xml:"StreamName"`
	RtsCount     int64   `json:"RtsCount" xml:"RtsCount"`
	RtmpCount    int64   `json:"RtmpCount" xml:"RtmpCount"`
	FlvCount     int64   `json:"FlvCount" xml:"FlvCount"`
	TimeStamp    string  `json:"TimeStamp" xml:"TimeStamp"`
	RtsBps       float64 `json:"RtsBps" xml:"RtsBps"`
	FlvBps       float64 `json:"FlvBps" xml:"FlvBps"`
	Duration     int64   `json:"Duration" xml:"Duration"`
	FlvTraffic   int64   `json:"FlvTraffic" xml:"FlvTraffic"`
	RtmpBps      float64 `json:"RtmpBps" xml:"RtmpBps"`
	Bps          float64 `json:"Bps" xml:"Bps"`
	AppId        string  `json:"AppId" xml:"AppId"`
	P2pCount     int64   `json:"P2pCount" xml:"P2pCount"`
	ReqTraffic   int64   `json:"ReqTraffic" xml:"ReqTraffic"`
	HlsBps       float64 `json:"HlsBps" xml:"HlsBps"`
	P2pBps       float64 `json:"P2pBps" xml:"P2pBps"`
	Count        int64   `json:"Count" xml:"Count"`
	Traffic      int64   `json:"Traffic" xml:"Traffic"`
	AppName      string  `json:"AppName" xml:"AppName"`
}
