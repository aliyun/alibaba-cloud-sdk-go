package ahas_openapi

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":    "ahas.cn-shenzhen.aliyuncs.com",
			"cn-beijing":     "ahas.cn-beijing.aliyuncs.com",
			"cn-shanghai":    "ahas.cn-shanghai.aliyuncs.com",
			"cn-hongkong":    "ahas.cn-hongkong.aliyuncs.com",
			"ap-southeast-1": "ahas.ap-southeast-1.aliyuncs.com",
			"eu-central-1":   "ahas.eu-central-1.aliyuncs.com",
			"cn-zhangjiakou": "ahas.cn-zhangjiakou.aliyuncs.com",
			"cn-hangzhou":    "ahas.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
