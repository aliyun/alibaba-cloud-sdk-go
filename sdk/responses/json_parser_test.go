package responses

import (
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	from := []byte(`{}`)
	to := &struct{}{}
	// support auto json type trans

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{}`, string(str))
}

func TestUnmarshal_int(t *testing.T) {
	to := &struct {
		INT int
	}{}
	from := []byte(`{"INT":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, 100, to.INT)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT":100}`, string(str))

	from = []byte(`{"INT":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, 100, to.INT)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT":100}`, string(str))

	// string to int
	from = []byte(`{"INT":"100"}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, 100, to.INT)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT":100}`, string(str))

	from = []byte(`{"INT":""}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, 0, to.INT)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT":0}`, string(str))

	// bool to int
	from = []byte(`{"INT":true}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, 1, to.INT)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT":1}`, string(str))

	from = []byte(`{"INT":false}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, 0, to.INT)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT":0}`, string(str))

	// nil to int
	from = []byte(`{"INT":null}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, 0, to.INT)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT":0}`, string(str))

	// fuzzy decode int
	from = []byte(`{"INT":100000000000000000000000000.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "INT: fuzzy decode int: exceed range, error found in")

	from = []byte(`{"INT":{}}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "INT: readUint64: unexpected character: \xff")
}

func TestUnmarshal_uint(t *testing.T) {
	to := &struct {
		UINT uint
	}{}
	from := []byte(`{"UINT":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, uint(100), to.UINT)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"UINT":100}`, string(str))

	from = []byte(`{"UINT":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, uint(100), to.UINT)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"UINT":100}`, string(str))

	// fuzzy decode uint
	from = []byte(`{"UINT":100000000000000000000000000.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "UINT: fuzzy decode uint: exceed range, error found in #10")
}

func TestUnmarshal_int8(t *testing.T) {
	to := &struct {
		INT8 int8
	}{}
	from := []byte(`{"INT8":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, int8(100), to.INT8)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT8":100}`, string(str))

	from = []byte(`{"INT8":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, int8(100), to.INT8)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT8":100}`, string(str))

	// fuzzy decode uint
	from = []byte(`{"INT8":100000000000000000000000000.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "INT8: fuzzy decode int8: exceed range")
}

func TestUnmarshal_uint8(t *testing.T) {
	to := &struct {
		UINT8 uint8
	}{}
	from := []byte(`{"UINT8":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, uint8(100), to.UINT8)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"UINT8":100}`, string(str))

	from = []byte(`{"UINT8":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, uint8(100), to.UINT8)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"UINT8":100}`, string(str))

	// fuzzy decode uint
	from = []byte(`{"UINT8":100000000000000000000000000.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "UINT8: fuzzy decode uint8: exceed range")
}

func TestUnmarshal_int16(t *testing.T) {
	to := &struct {
		INT16 int16
	}{}
	from := []byte(`{"INT16":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, int16(100), to.INT16)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT16":100}`, string(str))

	from = []byte(`{"INT16":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, int16(100), to.INT16)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT16":100}`, string(str))

	// fuzzy decode uint
	from = []byte(`{"INT16":100000000000000000000000000.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "INT16: fuzzy decode int16: exceed range")
}

func TestUnmarshal_uint16(t *testing.T) {
	to := &struct {
		UINT16 uint16
	}{}
	from := []byte(`{"UINT16":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, uint16(100), to.UINT16)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"UINT16":100}`, string(str))

	from = []byte(`{"UINT16":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, uint16(100), to.UINT16)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"UINT16":100}`, string(str))

	// fuzzy decode uint
	from = []byte(`{"UINT16":100000000000000000000000000.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "UINT16: fuzzy decode uint16: exceed range", err.Error())
}

func TestUnmarshal_int32(t *testing.T) {
	to := &struct {
		INT32 int32
	}{}
	from := []byte(`{"INT32":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, int32(100), to.INT32)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT32":100}`, string(str))

	from = []byte(`{"INT32":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, int32(100), to.INT32)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT32":100}`, string(str))

	// fuzzy decode uint
	from = []byte(`{"INT32":100000000000000000000000000.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "INT32: fuzzy decode int32: exceed range")
}

func TestUnmarshal_uint32(t *testing.T) {
	to := &struct {
		UINT32 uint32
	}{}
	from := []byte(`{"UINT32":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, uint32(100), to.UINT32)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"UINT32":100}`, string(str))

	from = []byte(`{"UINT32":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, uint32(100), to.UINT32)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"UINT32":100}`, string(str))

	// fuzzy decode uint
	from = []byte(`{"UINT32":100000000000000000000000000.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "UINT32: fuzzy decode uint32: exceed range")
}

func TestUnmarshal_int64(t *testing.T) {
	to := &struct {
		INT64 int64
	}{}
	from := []byte(`{"INT64":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, int64(100), to.INT64)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT64":100}`, string(str))

	from = []byte(`{"INT64":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, int64(100), to.INT64)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"INT64":100}`, string(str))

	// fuzzy decode uint
	from = []byte(`{"INT64":100000000000000000000000000.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "INT64: fuzzy decode int64: exceed range")
}

func TestUnmarshal_uint64(t *testing.T) {
	to := &struct {
		UINT64 uint64
	}{}
	from := []byte(`{"UINT64":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, uint64(100), to.UINT64)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"UINT64":100}`, string(str))

	from = []byte(`{"UINT64":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, uint64(100), to.UINT64)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"UINT64":100}`, string(str))

	// fuzzy decode uint
	from = []byte(`{"UINT64":100000000000000000000000000.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "UINT64: fuzzy decode uint64: exceed range")
}

func TestUnmarshal_string(t *testing.T) {
	to := &struct {
		STRING string
	}{}
	// string to string
	from := []byte(`{"STRING":""}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, "", to.STRING)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"STRING":""}`, string(str))

	// number to string
	from = []byte(`{"STRING":100}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, "100", to.STRING)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"STRING":"100"}`, string(str))

	// bool to string
	from = []byte(`{"STRING":true}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, "true", to.STRING)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"STRING":"true"}`, string(str))

	// nil to string
	from = []byte(`{"STRING":null}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, "", to.STRING)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"STRING":""}`, string(str))

	// other to string
	from = []byte(`{"STRING":{}}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "STRING: fuzzyStringDecoder: not number or string or bool")
}

func TestUnmarshal_bool(t *testing.T) {
	to := &struct {
		BOOL bool
	}{}
	// bool to bool
	from := []byte(`{"BOOL":true}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, true, to.BOOL)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"BOOL":true}`, string(str))

	// number to bool
	from = []byte(`{"BOOL":100}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, true, to.BOOL)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"BOOL":true}`, string(str))

	from = []byte(`{"BOOL":0}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, false, to.BOOL)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"BOOL":false}`, string(str))

	// invalid number literal
	from = []byte(`{"BOOL": 1000000000000000000000000000000000000000}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "BOOL: fuzzyBoolDecoder: get value from json.number failed")

	// bool to string
	from = []byte(`{"BOOL":"true"}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, true, to.BOOL)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"BOOL":true}`, string(str))

	from = []byte(`{"BOOL":"false"}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, false, to.BOOL)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"BOOL":false}`, string(str))

	from = []byte(`{"BOOL":"other"}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "BOOL: fuzzyBoolDecoder: unsupported bool value: other")

	// nil to bool
	from = []byte(`{"BOOL":null}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, false, to.BOOL)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"BOOL":false}`, string(str))

	// other to string
	from = []byte(`{"BOOL":{}}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "BOOL: fuzzyBoolDecoder: not number or string or nil")
}

func TestUnmarshal_array(t *testing.T) {
	to := &struct {
		Array []string
	}{}
	// bool to bool
	from := []byte(`{"Array":[]}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(to.Array))
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"Array":[]}`, string(str))
}

func TestUnmarshal_float32(t *testing.T) {
	to := &struct {
		FLOAT32 float32
	}{}
	from := []byte(`{"FLOAT32":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float32(100), to.FLOAT32)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT32":100}`, string(str))

	from = []byte(`{"FLOAT32":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float32(100.1), to.FLOAT32)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT32":100.1}`, string(str))

	// string to float32
	from = []byte(`{"FLOAT32":"100.1"}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float32(100.1), to.FLOAT32)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT32":100.1}`, string(str))

	from = []byte(`{"FLOAT32":""}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float32(0), to.FLOAT32)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT32":0}`, string(str))

	// error branch
	from = []byte(`{"FLOAT32":"."}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "FLOAT32: readFloat32: leading dot is invalid")

	// bool to float32
	from = []byte(`{"FLOAT32":true}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float32(1), to.FLOAT32)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT32":1}`, string(str))

	from = []byte(`{"FLOAT32":false}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float32(0), to.FLOAT32)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT32":0}`, string(str))

	// nil to float32
	from = []byte(`{"FLOAT32":null}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float32(0), to.FLOAT32)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT32":0}`, string(str))

	// others to float32
	from = []byte(`{"FLOAT32":{}}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "FLOAT32: nullableFuzzyFloat32Decoder: not number or string")
}

func TestUnmarshal_float64(t *testing.T) {
	to := &struct {
		FLOAT64 float64
	}{}
	from := []byte(`{"FLOAT64":100}`)

	err := jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float64(100), to.FLOAT64)
	str, err := jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT64":100}`, string(str))

	from = []byte(`{"FLOAT64":100.1}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float64(100.1), to.FLOAT64)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT64":100.1}`, string(str))

	// string to float64
	from = []byte(`{"FLOAT64":"100.1"}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float64(100.1), to.FLOAT64)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT64":100.1}`, string(str))

	from = []byte(`{"FLOAT64":""}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float64(0), to.FLOAT64)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT64":0}`, string(str))

	// error branch
	from = []byte(`{"FLOAT64":"."}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "FLOAT64: readFloat64: leading dot is invalid")

	// bool to float64
	from = []byte(`{"FLOAT64":true}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float64(1), to.FLOAT64)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT64":1}`, string(str))

	from = []byte(`{"FLOAT64":false}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float64(0), to.FLOAT64)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT64":0}`, string(str))

	// nil to float64
	from = []byte(`{"FLOAT64":null}`)

	err = jsonParser.Unmarshal(from, to)
	assert.Nil(t, err)
	assert.Equal(t, float64(0), to.FLOAT64)
	str, err = jsonParser.Marshal(to)
	assert.Nil(t, err)
	assert.Equal(t, `{"FLOAT64":0}`, string(str))

	// others to float64
	from = []byte(`{"FLOAT64":{}}`)

	err = jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "FLOAT64: nullableFuzzyFloat64Decoder: not number or string")
}

func TestUnmarshalWithArray(t *testing.T) {
	from := []byte(`[]`)
	to := &struct{}{}
	// TODO: Must support Array
	// support auto json type trans
	err := jsonParser.Unmarshal(from, to)
	assert.NotNil(t, err)
}

func TestNewBetterFuzzyExtension(t *testing.T) {
	betterFuzzyExtension := newBetterFuzzyExtension()
	assert.NotNil(t, betterFuzzyExtension)

	decoder := betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.String)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyStringDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Bool)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &fuzzyBoolDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Float32)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyFloat32Decoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Float64)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyFloat64Decoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Int)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyIntegerDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Uint)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyIntegerDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Int8)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyIntegerDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Uint8)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyIntegerDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Int16)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyIntegerDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Uint16)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyIntegerDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Int32)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyIntegerDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Uint32)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyIntegerDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Int64)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyIntegerDecoder{}, decoder)

	decoder = betterFuzzyExtension[reflect2.DefaultTypeOfKind(reflect.Uint64)]
	assert.NotNil(t, decoder)
	assert.IsType(t, &nullableFuzzyIntegerDecoder{}, decoder)
}

func TestUnmarshalWithDefaultDecoders(t *testing.T) {
	// should not be valid with default decoders
	from := []byte(`{"STRING":true}`)
	toString := &struct {
		STRING string
	}{}

	err := jsoniter.Unmarshal(from, toString)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"BOOL":""}`)
	toBool := &struct {
		BOOL bool
	}{}

	err = jsoniter.Unmarshal(from, toBool)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"FLOAT32":""}`)
	toFloat32 := &struct {
		FLOAT32 float32
	}{}

	err = jsoniter.Unmarshal(from, toFloat32)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"FLOAT64":""}`)
	toFloat64 := &struct {
		FLOAT64 float64
	}{}

	err = jsoniter.Unmarshal(from, toFloat64)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"INT":""}`)
	toInt := &struct {
		INT int
	}{}

	err = jsoniter.Unmarshal(from, toInt)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"UINT":""}`)
	toUint := &struct {
		UINT uint
	}{}

	err = jsoniter.Unmarshal(from, toUint)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"INT8":""}`)
	toInt8 := &struct {
		INT8 int8
	}{}

	err = jsoniter.Unmarshal(from, toInt8)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"UINT8":""}`)
	toUint8 := &struct {
		UINT8 uint8
	}{}

	err = jsoniter.Unmarshal(from, toUint8)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"INT16":""}`)
	toInt16 := &struct {
		INT16 int16
	}{}

	err = jsoniter.Unmarshal(from, toInt16)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"UINT16":""}`)
	toUint16 := &struct {
		UINT16 uint16
	}{}

	err = jsoniter.Unmarshal(from, toUint16)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"INT32":""}`)
	toInt32 := &struct {
		INT32 int32
	}{}

	err = jsoniter.Unmarshal(from, toInt32)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"UINT32":""}`)
	toUint32 := &struct {
		UINT32 uint32
	}{}

	err = jsoniter.Unmarshal(from, toUint32)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"INT64":""}`)
	toInt64 := &struct {
		INT64 int64
	}{}

	err = jsoniter.Unmarshal(from, toInt64)
	assert.NotNil(t, err)

	// should not be valid with default decoders
	from = []byte(`{"UINT64":""}`)
	toUint64 := &struct {
		UINT64 uint64
	}{}

	err = jsoniter.Unmarshal(from, toUint64)
	assert.NotNil(t, err)
}
