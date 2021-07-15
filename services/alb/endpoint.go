package alb

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "central"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":    "alb.cn-shenzhen.aliyuncs.com",
			"cn-beijing":     "alb.cn-beijing.aliyuncs.com",
			"cn-wulanchabu":  "alb.cn-wulanchabu.aliyuncs.com",
			"ap-south-1":     "alb.ap-south-1.aliyuncs.com",
			"ap-northeast-1": "alb.ap-northeast-1.aliyuncs.com",
			"cn-chengdu":     "alb.cn-chengdu.aliyuncs.com",
			"cn-shanghai":    "alb.cn-shanghai.aliyuncs.com",
			"cn-hongkong":    "alb.cn-hongkong.aliyuncs.com",
			"ap-southeast-1": "alb.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2": "alb.ap-southeast-2.aliyuncs.com",
			"eu-central-1":   "alb.eu-central-1.aliyuncs.com",
			"ap-southeast-5": "alb.ap-southeast-5.aliyuncs.com",
			"us-east-1":      "alb.us-east-1.aliyuncs.com",
			"cn-zhangjiakou": "alb.cn-zhangjiakou.aliyuncs.com",
			"cn-hangzhou":    "alb.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
