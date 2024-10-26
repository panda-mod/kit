package value

type Value interface {
	Data() any
	IsNil() bool
	Bool() bool
	Equal(other Value) bool
	Bytes() []byte
	Slice() []Value
	Int(def ...int) int
	Int8(def ...int8) int8
	Int16(def ...int16) int16
	Int32(def ...int32) int32
	Int64(def ...int64) int64
	Uint(def ...uint) uint
	Uint8(def ...uint8) uint8
	Uint16(def ...uint16) uint16
	Uint32(def ...uint32) uint32
	Uint64(def ...uint64) uint64
	Float32(def ...float32) float32
	Float64(def ...float64) float64
	String(def ...string) string
}
