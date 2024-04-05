package gdapi

import (
	"C"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)
import (
	"fmt"
	"reflect"
	"unsafe"
)

// Generic

func ensurePtr[T any](ptr T) T {
	if reflect.ValueOf(ptr).IsNil() {
		panic("provided pointer is nil") // TODO: ugh, should be proper error handling
	}
	return ptr
}

// BClass

func BClassFromPtr(typ gdc.VariantType, ptr gdc.ConstTypePtr) (interface{}, error) {
	ctr, found := bclassPtrConstructor[typ]
	if !found {
		return nil, fmt.Errorf("no constructor found for type %v", typ)
	}
	return ctr(ptr), nil
}

func BClassForType(typ gdc.VariantType) (any, error) {
	ctr, found := bclassDefaultConstructor[typ]
	if !found {
		return nil, fmt.Errorf("no constructor found for type %v", typ)
	}
	return ctr(), nil
}

func dataFromPtr(data []byte, ptr gdc.ConstTypePtr) {
	slice := unsafe.Slice((*byte)(ptr), len(data))
	copy(data, slice)
}

// Object

type objectLike interface {
	SetBaseObject(ptr gdc.ObjectPtr)
}

func createObject(obj objectLike, ptr gdc.ObjectPtr) any {
	obj.SetBaseObject(gdc.ObjectPtr(ptr))
	return obj
}

func ObjectFromPtr(ptr gdc.ObjectPtr) *Object {
	return &Object{
		obj: ptr,
	}
}

func DObjectFromPtr(className string, ptr gdc.ObjectPtr) (interface{}, error) {
	ctr, found := objectPtrConstructor[className]
	if !found {
		return nil, fmt.Errorf("no constructor found for class %v", className)
	}
	return ctr(ptr), nil
}

func (me *Object) AsVariant() *Variant {
	va := newVariant()
	va.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForVariant.fromObjectFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

func (me *Object) AsPtr() gdc.ObjectPtr {
	return me.obj
}

// Array

func (me *Array) Get(i int64) Variant {
	varPtr := me.iface.ArrayOperatorIndexConst(me.AsCTypePtr(), gdc.Int(i))
	va := NewVariantWithC(gdc.ConstVariantPtr(varPtr))
	return *va
}

func (me *Array) Set(i int64, value Variant) {
	varPtr := me.iface.ArrayOperatorIndex(me.AsTypePtr(), gdc.Int(i))
	AssignVariant(varPtr, value)
}

func ConvertArrayToSlice[T any](array *Array) ([]T, error) {
	size := array.Size()
	slice := make([]T, size)
	for i := int64(0); i < size; i += 1 {
		item := array.Get(i)
		converted, err := NativeFromVariantT[T](item)
		if err != nil {
			return nil, fmt.Errorf("error converting item %d (%v): %w", i, item, err)
		}
		slice[i] = *converted
	}
	return slice, nil
}

func ConvertArrayToUnknownSlice(array *Array) ([]any, error) {
	size := array.Size()
	slice := make([]any, size)
	for i := int64(0); i < size; i += 1 {
		item := array.Get(i)
		converted, err := NativeFromVariant(item, nil)
		if err != nil {
			return nil, err
		}
		slice[i] = converted
	}
	return slice, nil
}

// Dictionary

func (me *Dictionary) LookupWithDefault(key string, defaultValue Variant) Variant {
	keyVar := StringFromStr(key).AsVariant()
	defer keyVar.Destroy()
	return me.Get(*keyVar, defaultValue)
}

func (me *Dictionary) Lookup(key string) Variant {
	keyVar := StringFromStr(key).AsVariant()
	defer keyVar.Destroy()
	return me.LookupWithVariant(*keyVar)
}

func (me *Dictionary) LookupWithVariant(key Variant) Variant {
	varPtr := me.iface.DictionaryOperatorIndexConst(me.AsCTypePtr(), key.AsCPtr())
	va := NewVariantWithC(gdc.ConstVariantPtr(varPtr))
	return *va
}

func (me *Dictionary) Set(key string, value Variant) {
	keyVar := StringFromStr(key).AsVariant()
	defer keyVar.Destroy()
	me.SetWithVariant(*keyVar, value)
}

func (me *Dictionary) SetWithVariant(key, value Variant) {
	varPtr := me.iface.DictionaryOperatorIndex(me.AsTypePtr(), key.AsCPtr())
	AssignVariant(varPtr, value)
}

// String

func StringFromStr(str string) *String {
	me := newString()
	me.iface.StringNewWithUtf8CharsAndLen(gdc.UninitializedStringPtr(me.asUninitialized()), str, gdc.Int(len(str)))
	return me
}

func StringFromStringPtr(ptr gdc.StringPtr) *String {
	return StringFromPtr(gdc.ConstTypePtr(ptr))
}

func (me *String) AsPtr() gdc.StringPtr {
	return gdc.StringPtr(me.AsTypePtr())
}

func (me *String) AsCPtr() gdc.ConstStringPtr {
	return gdc.ConstStringPtr(me.AsPtr())
}

func (me *String) String() string {
	size := me.iface.StringToUtf8Chars(me.AsCPtr(), nil, 0)
	buffer := make([]byte, size)
	me.iface.StringToUtf8Chars(me.AsCPtr(), unsafe.SliceData(buffer), size)
	return C.GoString((*C.char)(unsafe.Pointer(&buffer[0])))
}

// StringName

func StringNameFromStr(str string) *StringName {
	strPtr := StringFromStr(str)
	defer strPtr.Destroy()
	return NewStringNameFromString(*strPtr)
}

func StringNameFromStringPtr(ptr gdc.StringNamePtr) *StringName {
	return StringNameFromPtr(gdc.ConstTypePtr(ptr))
}

func StringNameFromCPtr(ptr gdc.ConstStringNamePtr) *StringName {
	return StringNameFromPtr(gdc.ConstTypePtr(ptr))
}

func (me *StringName) AsPtr() gdc.StringNamePtr {
	return gdc.StringNamePtr(me.AsTypePtr())
}

func (me *StringName) AsCPtr() gdc.ConstStringNamePtr {
	return gdc.ConstStringNamePtr(me.AsPtr())
}

// Error

func (me Error) Error() string {
	// FIXME: this is not working for some reason
	// varStr := Utilities.ErrorString(int64(me))
	// defer varStr.Destroy()
	// return fmt.Sprintf("godot error %d: %s", me, varStr.String())
	return fmt.Sprintf("godot error %d", me)
}
