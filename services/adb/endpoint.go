package adb

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "adb.aliyuncs.com",
			"cn-beijing":            "adb.aliyuncs.com",
			"ap-south-1":            "adb.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "adb.aliyuncs.com",
			"me-east-1":             "adb.ap-northeast-1.aliyuncs.com",
			"cn-north-2-gov-1":      "adb.aliyuncs.com",
			"cn-qingdao":            "adb.aliyuncs.com",
			"cn-shanghai":           "adb.aliyuncs.com",
			"cn-shanghai-finance-1": "adb.aliyuncs.com",
			"cn-hongkong":           "adb.aliyuncs.com",
			"ap-southeast-1":        "adb.aliyuncs.com",
			"ap-southeast-2":        "adb.ap-northeast-1.aliyuncs.com",
			"cn-huhehaote":          "adb.aliyuncs.com",
			"us-east-1":             "adb.aliyuncs.com",
			"ap-southeast-5":        "adb.ap-northeast-1.aliyuncs.com",
			"us-west-1":             "adb.aliyuncs.com",
			"cn-hangzhou":           "adb.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
