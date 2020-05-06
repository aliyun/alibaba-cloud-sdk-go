package paistudio

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":    "pai.cn-shenzhen.aliyuncs.com",
			"cn-beijing":     "pai.cn-beijing.aliyuncs.com",
			"ap-south-1":     "pai.ap-south-1.aliyuncs.com",
			"me-east-1":      "pai.me-east-1.aliyuncs.com",
			"cn-qingdao":     "pai.cn-qingdao.aliyuncs.com",
			"cn-shanghai":    "pai.cn-shanghai.aliyuncs.com",
			"cn-hongkong":    "pai.cn-hongkong.aliyuncs.com",
			"ap-southeast-1": "pai.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2": "pai.ap-southeast-2.aliyuncs.com",
			"ap-southeast-3": "pai.ap-southeast-3.aliyuncs.com",
			"eu-central-1":   "pai.eu-central-1.aliyuncs.com",
			"ap-southeast-5": "pai.ap-southeast-5.aliyuncs.com",
			"cn-zhangjiakou": "pai.cn-zhangjiakou.aliyuncs.com",
			"us-west-1":      "pai.us-west-1.aliyuncs.com",
			"cn-hangzhou":    "pai-cn-hangzhou.data.aliyun.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
