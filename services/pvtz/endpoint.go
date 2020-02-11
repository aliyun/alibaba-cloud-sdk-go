package pvtz

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "central"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":    "pvtz.cn-shenzhen.aliyuncs.com",
			"cn-beijing":     "pvtz.cn-beijing.aliyuncs.com",
			"cn-shanghai":    "pvtz.cn-shanghai.aliyuncs.com",
			"cn-hongkong":    "pvtz.cn-hongkong.aliyuncs.com",
			"ap-southeast-1": "pvtz.ap-southeast-1.aliyuncs.com",
			"ap-southeast-3": "pvtz.ap-southeast-3.aliyuncs.com",
			"eu-central-1":   "pvtz.eu-central-1.aliyuncs.com",
			"cn-huhehaote":   "pvtz.cn-huhehaote.aliyuncs.com",
			"ap-southeast-5": "pvtz.ap-southeast-5.aliyuncs.com",
			"cn-zhangjiakou": "pvtz.cn-zhangjiakou.aliyuncs.com",
			"cn-chengdu":     "pvtz.cn-chengdu.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
