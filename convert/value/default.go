package value

import (
	"reflect"
	"sync/atomic"

	"github.com/panda-mod/kit/convert"
	"github.com/panda-mod/kit/encoding"
	"github.com/panda-mod/kit/slice"
)

var (
	// Value 值接口
	_ Value = &value{}
	// defaultValue = New(nil)
	defaultValue = new(value)
)

// value 默认 Value 实现
type value struct {
	raw atomic.Value // 值
}

// Default 默认值对象
func Default() Value {
	return defaultValue
}

// New 创建一个值对象
func New(val any) Value {
	if val == nil {
		return defaultValue
	}
	return new(value).Set(val)
}

func (v *value) load() any {
	return v.raw.Load()
}

func (v *value) Data() any {
	return v.load()
}

func (v *value) Equal(other Value) bool {
	return reflect.DeepEqual(v.Data(), other.Data())
}

func (v *value) IsNil() bool {
	return v.load() == nil
}

func (v *value) Bool() bool {
	return convert.Bool(v.load())
}

func (v *value) Bytes() []byte {
	return convert.Bytes(v.load())
}

func (v *value) Int(def ...int) int {
	return convert.Int(v.load(), def...)
}

func (v *value) Int8(def ...int8) int8 {
	return convert.Int8(v.load(), def...)
}

func (v *value) Int16(def ...int16) int16 {
	return convert.Int16(v.load(), def...)
}

func (v *value) Int32(def ...int32) int32 {
	return convert.Int32(v.load(), def...)
}

func (v *value) Int64(def ...int64) int64 {
	return convert.Int64(v.load(), def...)
}

func (v *value) Uint(def ...uint) uint {
	return convert.Uint(v.load(), def...)
}

func (v *value) Uint8(def ...uint8) uint8 {
	return convert.Uint8(v.load(), def...)
}

func (v *value) Uint16(def ...uint16) uint16 {
	return convert.Uint16(v.load(), def...)
}

func (v *value) Uint32(def ...uint32) uint32 {
	return convert.Uint32(v.load(), def...)
}

func (v *value) Uint64(def ...uint64) uint64 {
	return convert.Uint64(v.load(), def...)
}

func (v *value) Float32(def ...float32) float32 {
	return convert.Float32(v.load(), def...)
}

func (v *value) Float64(def ...float64) float64 {
	return convert.Float64(v.load(), def...)
}

func (v *value) String(def ...string) string {
	return convert.String(v.load(), def...)
}

func (v *value) Set(val any) *value {
	v.raw.Store(val)
	return v
}

func (v *value) Slice() []Value {
	switch values := v.load().(type) {
	case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
		return []Value{New(values)}
	case []any:
		return slice.Map(values, func(val any) Value { return New(val) })
	default:
		return []Value{}
	}
}

func (v *value) Scan(pointer any, decode encoding.Decoder) error {
	switch p := pointer.(type) {
	case *int:
		*p = v.Int()
	case *uint:
		*p = v.Uint()
	case *float64:
		*p = v.Float64()
	case *bool:
		*p = v.Bool()
	case *string:
		*p = v.String()
	}
	return decode(v.Bytes(), pointer)
}
