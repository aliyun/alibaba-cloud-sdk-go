package sgw

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":    "sgw.cn-shanghai.aliyuncs.com",
			"cn-qingdao":     "sgw.cn-shanghai.aliyuncs.com",
			"cn-beijing":     "sgw.cn-shanghai.aliyuncs.com",
			"cn-hongkong":    "sgw.cn-shanghai.aliyuncs.com",
			"cn-huhehaote":   "sgw.cn-shanghai.aliyuncs.com",
			"us-east-1":      "sgw.us-west-1.aliyuncs.com",
			"cn-zhangjiakou": "sgw.cn-shanghai.aliyuncs.com",
			"cn-chengdu":     "sgw.cn-shanghai.aliyuncs.com",
			"cn-hangzhou":    "sgw.cn-shanghai.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
