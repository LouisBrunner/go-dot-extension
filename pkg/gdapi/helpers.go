package gdapi

import (
	"C"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)
import (
	"fmt"
	"unsafe"
)

// BClass

func BClassFromPtr(typ gdc.VariantType, ptr gdc.ConstTypePtr) (interface{}, error) {
	ctr, found := bclassPtrConstructor[typ]
	if !found {
		return nil, fmt.Errorf("no constructor found for type %v", typ)
	}
	return ctr(ptr), nil
}

func dataFromPtr(data []byte, ptr gdc.ConstTypePtr) {
	slice := unsafe.Slice((*byte)(ptr), len(data))
	copy(data, slice)
}

// Object

type objectLike interface {
	SetBaseObject(ptr gdc.ObjectPtr)
}

func createObject(obj objectLike, ptr gdc.ConstTypePtr) any {
	obj.SetBaseObject(gdc.ObjectPtr(ptr))
	return obj
}

func ObjectFromPtr(ptr gdc.ObjectPtr) Object {
	return Object{
		obj: ptr,
	}
}

func DObjectFromPtr(className string, ptr gdc.ObjectPtr) (interface{}, error) {
	return nil, nil // TODO: HERE!
}

// String

func StringFromStr(str string) *String {
	me := newString()
	giface.StringNewWithUtf8CharsAndLen(gdc.UninitializedStringPtr(me.asUninitialized()), str, gdc.Int(len(str)))
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
