package vpc

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":    "vpc.aliyuncs.com",
			"cn-qingdao":     "vpc.aliyuncs.com",
			"cn-beijing":     "vpc.aliyuncs.com",
			"cn-shanghai":    "vpc.aliyuncs.com",
			"cn-hongkong":    "vpc.aliyuncs.com",
			"ap-southeast-1": "vpc.aliyuncs.com",
			"us-east-1":      "vpc.aliyuncs.com",
			"us-west-1":      "vpc.aliyuncs.com",
			"cn-hangzhou":    "vpc.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
