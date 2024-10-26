package deepcopy

import (
	"reflect"
)

type builder[T any] struct {
	value reflect.Value
}

// build
func (b *builder[T]) build(target T) T {
	valueOf := reflect.ValueOf(target)
	return b.clone(valueOf).Interface().(T)
}

// create 创建一个对象
func (b *builder[T]) create(valueOf reflect.Value) reflect.Value {
	newValue := reflect.New(valueOf.Type())
	return newValue.Elem()
}

// clone 克隆一个对象
func (b *builder[T]) clone(target reflect.Value) reflect.Value {
	// 克隆一个空对象
	newValue := b.create(target)
	// 判断类型
	switch target.Kind() {
	case reflect.Pointer:
		return b.clonePointer(target, newValue)
	case reflect.Struct:
		return b.cloneStruct(target, newValue)
	case reflect.Interface:
		return b.cloneInterface(target, newValue)
	case reflect.Slice:
		return b.cloneSlice(target, newValue)
	case reflect.Map:
		return b.cloneMap(target, newValue)
	default:
		newValue.Set(target)
	}
	return newValue
}

// cloneSlice 克隆切片
func (b *builder[T]) cloneSlice(target reflect.Value, dstValue reflect.Value) reflect.Value {
	// 新建一个空的Slice
	dstValue.Set(reflect.MakeSlice(target.Type(), target.Len(), target.Cap()))
	// 遍历原Slice，克隆每个元素
	for i := 0; i < target.Len(); i++ {
		dstValue.Index(i).Set(b.clone(target.Index(i)))
	}
	return dstValue
}

// cloneMap 克隆映射
func (b *builder[T]) cloneMap(target reflect.Value, dstValue reflect.Value) reflect.Value {
	// 新建一个空的Map
	dstValue.Set(reflect.MakeMap(target.Type()))
	// 遍历原Map，克隆每个元素
	for _, key := range target.MapKeys() {
		dstValue.SetMapIndex(b.clone(key), b.clone(target.MapIndex(key)))
	}
	return dstValue
}

// clonePointer 克隆指针
func (b *builder[T]) clonePointer(target reflect.Value, dstValue reflect.Value) reflect.Value {
	// 新建一个空的指针类型
	dstValue.Set(reflect.New(target.Type().Elem()))
	// 克隆指针指向的对象
	dstValue.Elem().Set(b.clone(target.Elem()))
	return dstValue
}

// cloneStruct 克隆结构体
func (b *builder[T]) cloneStruct(target reflect.Value, dstValue reflect.Value) reflect.Value {
	// 遍历结构体字段，克隆每个字段
	for i := 0; i < target.NumField(); i++ {
		dstValue.Field(i).Set(b.clone(target.Field(i)))
	}
	return dstValue
}

// cloneInterface 克隆接口
func (b *builder[T]) cloneInterface(target reflect.Value, dstValue reflect.Value) reflect.Value {
	if target.IsNil() {
		return dstValue
	}
	dstValue.Set(b.clone(target.Elem()))
	return dstValue
}
