package dysmsapi

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "central"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-beijing":     "dysmsapi-proxy.cn-beijing.aliyuncs.com",
			"cn-hongkong":    "dysmsapi-xman.cn-hongkong.aliyuncs.com",
			"ap-southeast-1": "dysmsapi.ap-southeast-1.aliyuncs.com",
			"eu-central-1":   "dysmsapi.eu-central-1.aliyuncs.com",
			"ap-southeast-5": "dysmsapi.ap-southeast-5.aliyuncs.com",
			"us-east-1":      "dysmsapi.us-east-1.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
