package sas

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "central"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-et2-b01": "tds.cn-shanghai-et2-b01.aliyuncs.com",
			"ap-southeast-3":      "tds.ap-southeast-3.aliyuncs.com",
			"cn-hangzhou":         "tds.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
