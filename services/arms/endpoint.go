package arms

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "arms.aliyuncs.com",
			"cn-beijing-gov-1":            "arms.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "arms.aliyuncs.com",
			"cn-shanghai-inner":           "arms.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "arms.aliyuncs.com",
			"cn-haidian-cm12-c01":         "arms.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "arms.aliyuncs.com",
			"cn-yushanfang":               "arms.aliyuncs.com",
			"cn-hongkong-finance-pop":     "arms.aliyuncs.com",
			"cn-qingdao-nebula":           "arms.aliyuncs.com",
			"cn-beijing-finance-pop":      "arms.aliyuncs.com",
			"cn-wuhan":                    "arms.aliyuncs.com",
			"cn-zhangbei":                 "arms.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "arms.aliyuncs.com",
			"rus-west-1-pop":              "arms.aliyuncs.com",
			"cn-shanghai-et15-b01":        "arms.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "arms.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "arms.aliyuncs.com",
			"eu-west-1-oxs":               "arms.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "arms.aliyuncs.com",
			"cn-beijing-finance-1":        "arms.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "arms.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "arms.aliyuncs.com",
			"me-east-1":                   "arms.aliyuncs.com",
			"cn-hangzhou-test-306":        "arms.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "arms.aliyuncs.com",
			"cn-shanghai-et2-b01":         "arms.aliyuncs.com",
			"cn-beijing-nu16-b01":         "arms.aliyuncs.com",
			"cn-edge-1":                   "arms.aliyuncs.com",
			"cn-fujian":                   "arms.aliyuncs.com",
			"ap-northeast-2-pop":          "arms.aliyuncs.com",
			"cn-shenzhen-inner":           "arms.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "arms.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
