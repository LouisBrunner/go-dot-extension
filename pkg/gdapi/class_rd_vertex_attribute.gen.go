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

type ptrsForRDVertexAttributeList struct {
	fnSetLocation  gdc.MethodBindPtr
	fnGetLocation  gdc.MethodBindPtr
	fnSetOffset    gdc.MethodBindPtr
	fnGetOffset    gdc.MethodBindPtr
	fnSetFormat    gdc.MethodBindPtr
	fnGetFormat    gdc.MethodBindPtr
	fnSetStride    gdc.MethodBindPtr
	fnGetStride    gdc.MethodBindPtr
	fnSetFrequency gdc.MethodBindPtr
	fnGetFrequency gdc.MethodBindPtr
}

var ptrsForRDVertexAttribute ptrsForRDVertexAttributeList

func initRDVertexAttributePtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDVertexAttribute")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_location")
		defer methodName.Destroy()
		ptrsForRDVertexAttribute.fnSetLocation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_location")
		defer methodName.Destroy()
		ptrsForRDVertexAttribute.fnGetLocation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_offset")
		defer methodName.Destroy()
		ptrsForRDVertexAttribute.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_offset")
		defer methodName.Destroy()
		ptrsForRDVertexAttribute.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_format")
		defer methodName.Destroy()
		ptrsForRDVertexAttribute.fnSetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 565531219))
	}
	{
		methodName := StringNameFromStr("get_format")
		defer methodName.Destroy()
		ptrsForRDVertexAttribute.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2235804183))
	}
	{
		methodName := StringNameFromStr("set_stride")
		defer methodName.Destroy()
		ptrsForRDVertexAttribute.fnSetStride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_stride")
		defer methodName.Destroy()
		ptrsForRDVertexAttribute.fnGetStride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_frequency")
		defer methodName.Destroy()
		ptrsForRDVertexAttribute.fnSetFrequency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 522141836))
	}
	{
		methodName := StringNameFromStr("get_frequency")
		defer methodName.Destroy()
		ptrsForRDVertexAttribute.fnGetFrequency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4154106413))
	}

}

type RDVertexAttribute struct {
	RefCounted
}

func (me *RDVertexAttribute) BaseClass() string {
	return "RDVertexAttribute"
}

func NewRDVertexAttribute() *RDVertexAttribute {
	str := StringNameFromStr("RDVertexAttribute") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDVertexAttribute{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDVertexAttribute) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDVertexAttribute) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDVertexAttribute) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDVertexAttribute) SetLocation(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDVertexAttribute.fnSetLocation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDVertexAttribute) GetLocation() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDVertexAttribute.fnGetLocation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDVertexAttribute) SetOffset(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDVertexAttribute.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDVertexAttribute) GetOffset() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDVertexAttribute.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDVertexAttribute) SetFormat(p_member RenderingDeviceDataFormat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDVertexAttribute.fnSetFormat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDVertexAttribute) GetFormat() RenderingDeviceDataFormat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceDataFormat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDVertexAttribute.fnGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDVertexAttribute) SetStride(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDVertexAttribute.fnSetStride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDVertexAttribute) GetStride() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDVertexAttribute.fnGetStride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDVertexAttribute) SetFrequency(p_member RenderingDeviceVertexFrequency) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDVertexAttribute.fnSetFrequency), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDVertexAttribute) GetFrequency() RenderingDeviceVertexFrequency {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceVertexFrequency

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDVertexAttribute.fnGetFrequency), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
