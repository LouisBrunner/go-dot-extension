package gdapi

import (
	"strings"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BClass interface {
	Type() gdc.VariantType
	Destroy()
	AsTypePtr() gdc.TypePtr
}

type Variant struct {
	iface gdc.Interface
	inner BClass
	ptr   gdc.VariantPtr
}

// From pointers

func NewVariantWith(ptr gdc.VariantPtr) Variant {
	return Variant{
		iface: giface,
		ptr:   ptr,
	}
}

func NewVariantWithC(ptr gdc.ConstVariantPtr) Variant {
	return Variant{
		iface: giface,
		ptr:   gdc.VariantPtr(ptr),
	}
}

// From other types

func NewVariantFrom(inner BClass) Variant {
	ctr := giface.GetVariantFromTypeConstructor(inner.Type()) // FIXME: cache?
	ptr := (gdc.UninitializedVariantPtr)(cmalloc(classSizeVariant))
	giface.CallVariantFromTypeConstructorFunc(ctr, ptr, inner.AsTypePtr())
	return Variant{
		iface: giface,
		inner: inner,
		ptr:   gdc.VariantPtr(ptr),
	}
}

func NewVariantFromStr(str string) Variant {
	return NewVariantFrom(StringFromStr(str))
}

// Getting pointers

func (me Variant) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.ptr)
}

func (me Variant) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.ptr)
}

func (me Variant) AsPtr() gdc.VariantPtr {
	return me.ptr
}

func (me Variant) AsCPtr() gdc.ConstVariantPtr {
	return gdc.ConstVariantPtr(me.ptr)
}

// Methods

func (me Variant) Type() gdc.VariantType {
	return me.iface.VariantGetType(gdc.ConstVariantPtr(me.ptr))
}

func (me *Variant) Destroy() {
	if me.ptr == nil {
		giface.VariantDestroy(me.ptr)
		me.ptr = nil
	}
	if me.inner != nil {
		me.inner.Destroy()
	}
}

func (me Variant) String() string {
	const maxLength = 1024
	stringified := StringFromStr(strings.Repeat("a", maxLength))
	me.iface.VariantStringify(gdc.ConstVariantPtr(me.ptr), stringified.AsPtr())
	content := make([]byte, maxLength)
	len := me.iface.StringToUtf8Chars(stringified.AsCPtr(), unsafe.SliceData(content), gdc.Int(maxLength))
	return string(content[:len])
}
