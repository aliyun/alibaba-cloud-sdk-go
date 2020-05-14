package schedulerx2

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen": "schedulerx.cn-shenzhen.aliyuncs.com",
			"cn-beijing":  "schedulerx.cn-beijing.aliyuncs.com",
			"cn-shanghai": "schedulerx.cn-shanghai.aliyuncs.com",
			"cn-hangzhou": "schedulerx.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
