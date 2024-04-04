package gdapi

import (
	"C"

	"fmt"
	"strings"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)
import "runtime"

type BClass interface {
	Type() gdc.VariantType
	Destroy()
	AsTypePtr() gdc.TypePtr
	AsVariant() *Variant
}

type Variant struct {
	data   *[classSizeVariant]byte
	iface  gdc.Interface
	inner  BClass
	pinner runtime.Pinner
}

func newVariant() *Variant {
	me := &Variant{
		data:  new([classSizeVariant]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewVariant() *Variant {
	return newVariant()
}

// From pointers

func NewVariantWith(ptr gdc.VariantPtr) *Variant {
	return NewVariantWithC(gdc.ConstVariantPtr(ptr))
}

func NewVariantWithC(ptr gdc.ConstVariantPtr) *Variant {
	me := newVariant()
	me.iface.VariantNewCopy(me.asUninitialized(), ptr)
	return me
}

// Getting pointers

func (me *Variant) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.AsPtr())
}

func (me *Variant) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Variant) AsPtr() gdc.VariantPtr {
	return gdc.VariantPtr(unsafe.Pointer(me.data))
}

func (me *Variant) AsCPtr() gdc.ConstVariantPtr {
	return gdc.ConstVariantPtr(me.AsPtr())
}

func (me *Variant) asUninitialized() gdc.UninitializedVariantPtr {
	return gdc.UninitializedVariantPtr(me.AsPtr())
}

// Methods

func (me *Variant) Type() gdc.VariantType {
	return me.iface.VariantGetType(me.AsCPtr())
}

func (me *Variant) Destroy() {
	giface.VariantDestroy(me.AsPtr())
	if me.inner != nil {
		me.inner.Destroy()
	}
	me.pinner.Unpin()
}

func (me *Variant) Assign(value Variant) {
	me.Destroy()
	me.iface.VariantNewCopy(me.asUninitialized(), value.AsCPtr())
}

func (me *Variant) String() string {
	const maxLength = 1024
	stringified := StringFromStr(strings.Repeat("a", maxLength))
	me.iface.VariantStringify(me.AsCPtr(), stringified.AsPtr())
	content := make([]byte, maxLength)
	len := me.iface.StringToUtf8Chars(stringified.AsCPtr(), unsafe.SliceData(content), gdc.Int(maxLength))
	return string(content[:len])
}

func (me *Variant) AsObject() (interface{}, error) {
	if me.Type() != gdc.VariantTypeObject {
		return nil, fmt.Errorf("variant is not an object")
	}
	data := [classSizeObject]byte{}
	dataPtr := unsafe.Pointer(&data)
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, gdc.UninitializedTypePtr(dataPtr), me.AsPtr())
	objPtr := *(*gdc.ObjectPtr)(dataPtr)
	obj := ObjectFromPtr(objPtr)
	clazzStr := obj.GetClass()
	defer clazzStr.Destroy()
	return DObjectFromPtr(clazzStr.String(), objPtr)
}

func (me *Variant) AsBClass() (BClass, error) {
	if me.Type() == gdc.VariantTypeObject {
		return nil, fmt.Errorf("variant is an object, use AsObject instead")
	}
	return bclassFromVariant[me.Type()](me)
}
