// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForMarshallsList struct {
	fnVariantToBase64 gdc.MethodBindPtr
	fnBase64ToVariant gdc.MethodBindPtr
	fnRawToBase64     gdc.MethodBindPtr
	fnBase64ToRaw     gdc.MethodBindPtr
	fnUtf8ToBase64    gdc.MethodBindPtr
	fnBase64ToUtf8    gdc.MethodBindPtr
}

var ptrsForMarshalls ptrsForMarshallsList

func initMarshallsPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Marshalls")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("variant_to_base64")
		defer methodName.Destroy()
		ptrsForMarshalls.fnVariantToBase64 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3876248563))
	}
	{
		methodName := StringNameFromStr("base64_to_variant")
		defer methodName.Destroy()
		ptrsForMarshalls.fnBase64ToVariant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 218087648))
	}
	{
		methodName := StringNameFromStr("raw_to_base64")
		defer methodName.Destroy()
		ptrsForMarshalls.fnRawToBase64 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3999417757))
	}
	{
		methodName := StringNameFromStr("base64_to_raw")
		defer methodName.Destroy()
		ptrsForMarshalls.fnBase64ToRaw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659035735))
	}
	{
		methodName := StringNameFromStr("utf8_to_base64")
		defer methodName.Destroy()
		ptrsForMarshalls.fnUtf8ToBase64 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1703090593))
	}
	{
		methodName := StringNameFromStr("base64_to_utf8")
		defer methodName.Destroy()
		ptrsForMarshalls.fnBase64ToUtf8 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1703090593))
	}
}

type Marshalls struct {
	Object
}

func (me *Marshalls) BaseClass() string {
	return "Marshalls"
}

func NewMarshalls() *Marshalls {
	str := StringNameFromStr("Marshalls") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Marshalls{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Marshalls) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Marshalls) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Marshalls) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Marshalls) VariantToBase64(variant Variant, full_objects bool) String {
	cargs := []gdc.ConstTypePtr{variant.AsCTypePtr(), gdc.ConstTypePtr(&full_objects)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&full_objects)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMarshalls.fnVariantToBase64), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Marshalls) Base64ToVariant(base64_str String, allow_objects bool) Variant {
	cargs := []gdc.ConstTypePtr{base64_str.AsCTypePtr(), gdc.ConstTypePtr(&allow_objects)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&allow_objects)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMarshalls.fnBase64ToVariant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Marshalls) RawToBase64(array PackedByteArray) String {
	cargs := []gdc.ConstTypePtr{array.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMarshalls.fnRawToBase64), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Marshalls) Base64ToRaw(base64_str String) PackedByteArray {
	cargs := []gdc.ConstTypePtr{base64_str.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMarshalls.fnBase64ToRaw), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Marshalls) Utf8ToBase64(utf8_str String) String {
	cargs := []gdc.ConstTypePtr{utf8_str.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMarshalls.fnUtf8ToBase64), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Marshalls) Base64ToUtf8(base64_str String) String {
	cargs := []gdc.ConstTypePtr{base64_str.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMarshalls.fnBase64ToUtf8), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
