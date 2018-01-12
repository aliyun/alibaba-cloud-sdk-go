package responses

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuzzyFieldUnmarshal(t *testing.T) {
	from, err := getJsonBytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("From:")
	fmt.Println(string(from))
	to := &To{}
	// support auto json type trans
	initJsonParserOnce()
	err = jsonParser.Unmarshal(from, to)
	if err != nil {
		panic(err)
	}

	fmt.Println("To:")
	bytesTo, _ := json.MarshalIndent(to, "", "    ")
	fmt.Println(string(bytesTo))

	assert.Equal(t, "demo string", to.StrToStr)
	assert.Equal(t, 123, to.StrToInt)
	assert.Equal(t, int64(2147483648), to.StrToInt64)
	assert.Equal(t, 1.23, to.StrToFloat)
	assert.True(t, to.UpperFullStrToBool)
	assert.True(t, to.UpperCamelStrToBool)
	assert.True(t, to.LowerStrToBool)

	assert.Equal(t, "123", to.IntToStr)
	assert.Equal(t, 123, to.IntToInt)

	assert.Equal(t, "2147483648", to.Int64ToStr)
	assert.Equal(t, int64(2147483648), to.Int64ToInt64)

	assert.Equal(t, "true", to.BoolToStr)
	assert.Equal(t, true, to.BoolToBool)

	assert.Equal(t, 0, to.EmptyStrToInt)
	assert.Equal(t, int64(0), to.EmptyStrToInt64)
	assert.Equal(t, float64(0), to.EmptyStrToFloat)
	assert.Equal(t, false, to.EmptyStrToBool)
	assert.Equal(t, "", to.EmptyStrToStr)
}

//func TestFuzzyTypeUnmarshal(t *testing.T){
//	arrayJson := "[{\"instance_type\":\"ecs.n4.large\",\"vpc_id\":\"\",\"vswitch_id\":\"\",\"vswitch_cidr\":\"\",\"image_id\":\"registry-internal.cn-beijing.aliyuncs.com/acs/agent\",\"data_disk_size\":0,\"data_disk_category\":\"cloud_ssd\",\"security_group_id\":\"sg-2ze57kc2cf36f9mrsrjx\",\"tags\":\"\",\"zone_id\":\"cn-beijing-a\",\"-\":\"PayByTraffic\",\"name\":\"Hello\",\"cluster_id\":\"ca737c2c04143464eaf439e245ceb1bf4\",\"size\":3,\"region_id\":\"cn-beijing\",\"network_mode\":\"classic\",\"subnet_cidr\":\"172.18.1.1/24\",\"state\":\"running\",\"master_url\":\"https://master4g7.cs-cn-beijing.aliyun.com:20019\",\"agent_version\":\"0.9-cdb96d4\",\"external_loadbalancer_id\":\"lb-2zegrgbsmjvxx1r1v26pn\",\"internal_master_url\":\"https://master4g7.cs-cn-beijing.aliyun-inc.com:20019\",\"port\":20019,\"node_status\":\"{\\\"health\\\":0,\\\"unhealth\\\":3}\",\"cluster_healthy\":\"unhealth\",\"docker_version\":\"17.06.1-ce\",\"cluster_type\":\"aliyun\",\"swarm_mode\":true,\"init_version\":\"\",\"current_version\":\"\",\"meta_data\":\"\",\"upgrade_components\":null,\"capabilities\":{\"acslogging\":true,\"acsrouting\":true,\"blue-green_confirm\":true,\"blue-green_deployment\":true,\"cluster_event\":true,\"compose_v2\":true,\"config_map\":true,\"cron\":true,\"default_update_method\":true,\"drain\":true,\"logging_sls\":true,\"monitoring\":true,\"node_scaling\":true,\"porject_event_deletion\":true,\"porject_trigger\":true,\"rebalance\":true,\"reset_node\":true,\"routing_custom_root_domain\":true,\"routing_slb\":true,\"scalingtrigger\":true,\"slb-blue-green_deployment\":true,\"update_force_reschedule\":true,\"volume_ossfs\":true,\"volume_yunpan\":true},\"need_update_agent\":false,\"created\":\"2017-09-12T07:26:58Z\",\"updated\":\"2017-12-26T11:36:41Z\",\"outputs\":null,\"parameters\":null}]"
//	commonResponse := NewCommonResponse()
//	// support auto json type trans
//	initJsonParserOnce()
//	err := jsonParser.Unmarshal([]byte(arrayJson), commonResponse)
//	if err != nil {
//		panic(err)
//	}
//}

func getJsonBytes() ([]byte, error) {
	from := &From{
		StrToStr:            "demo string",
		StrToInt:            "123",
		StrToInt64:          "2147483648",
		StrToFloat:          "1.23",
		UpperFullStrToBool:  "TRUE",
		UpperCamelStrToBool: "True",
		LowerStrToBool:      "true",

		IntToStr: 123,
		IntToInt: 123,

		Int64ToStr:   int64(2147483648),
		Int64ToInt64: int64(2147483648),

		FloatToStr:   float64(1.23),
		FloatToFloat: float64(1.23),

		BoolToStr:  true,
		BoolToBool: true,

		EmptyStrToInt:   "",
		EmptyStrToInt64: "",
		EmptyStrToFloat: "",
		EmptyStrToBool:  "",
		EmptyStrToStr:   "",
	}

	return json.MarshalIndent(from, "", "    ")
}

type From struct {
	StrToStr            string
	StrToInt            string
	StrToInt64          string
	StrToFloat          string
	UpperFullStrToBool  string
	UpperCamelStrToBool string
	LowerStrToBool      string

	IntToStr int
	IntToInt int

	Int64ToStr   int64
	Int64ToInt64 int64

	FloatToStr   float64
	FloatToFloat float64

	BoolToStr  bool
	BoolToBool bool

	EmptyStrToInt   string
	EmptyStrToInt64 string
	EmptyStrToFloat string
	EmptyStrToBool  string
	EmptyStrToStr   string
}

type To struct {
	StrToStr            string
	StrToInt            int
	StrToInt64          int64
	StrToFloat          float64
	UpperFullStrToBool  bool
	UpperCamelStrToBool bool
	LowerStrToBool      bool

	IntToStr string
	IntToInt int

	Int64ToStr   string
	Int64ToInt64 int64

	FloatToStr   string
	FloatToFloat float64

	BoolToStr  string
	BoolToBool bool

	EmptyStrToInt   int
	EmptyStrToInt64 int64
	EmptyStrToFloat float64
	EmptyStrToBool  bool
	EmptyStrToStr   string

	NilToInt   int
	NilToInt64 int64
	NilToFloat float64
	NilToBool  bool
	NilToStr   string
}
