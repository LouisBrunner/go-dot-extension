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

type ptrsForRDAttachmentFormatList struct {
	fnSetFormat     gdc.MethodBindPtr
	fnGetFormat     gdc.MethodBindPtr
	fnSetSamples    gdc.MethodBindPtr
	fnGetSamples    gdc.MethodBindPtr
	fnSetUsageFlags gdc.MethodBindPtr
	fnGetUsageFlags gdc.MethodBindPtr
}

var ptrsForRDAttachmentFormat ptrsForRDAttachmentFormatList

func initRDAttachmentFormatPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDAttachmentFormat")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_format")
		defer methodName.Destroy()
		ptrsForRDAttachmentFormat.fnSetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 565531219))
	}
	{
		methodName := StringNameFromStr("get_format")
		defer methodName.Destroy()
		ptrsForRDAttachmentFormat.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2235804183))
	}
	{
		methodName := StringNameFromStr("set_samples")
		defer methodName.Destroy()
		ptrsForRDAttachmentFormat.fnSetSamples = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3774171498))
	}
	{
		methodName := StringNameFromStr("get_samples")
		defer methodName.Destroy()
		ptrsForRDAttachmentFormat.fnGetSamples = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 407791724))
	}
	{
		methodName := StringNameFromStr("set_usage_flags")
		defer methodName.Destroy()
		ptrsForRDAttachmentFormat.fnSetUsageFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_usage_flags")
		defer methodName.Destroy()
		ptrsForRDAttachmentFormat.fnGetUsageFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type RDAttachmentFormat struct {
	RefCounted
}

func (me *RDAttachmentFormat) BaseClass() string {
	return "RDAttachmentFormat"
}

func NewRDAttachmentFormat() *RDAttachmentFormat {
	str := StringNameFromStr("RDAttachmentFormat") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDAttachmentFormat{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDAttachmentFormat) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDAttachmentFormat) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDAttachmentFormat) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDAttachmentFormat) SetFormat(p_member RenderingDeviceDataFormat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDAttachmentFormat.fnSetFormat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDAttachmentFormat) GetFormat() RenderingDeviceDataFormat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceDataFormat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDAttachmentFormat.fnGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDAttachmentFormat) SetSamples(p_member RenderingDeviceTextureSamples) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDAttachmentFormat.fnSetSamples), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDAttachmentFormat) GetSamples() RenderingDeviceTextureSamples {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceTextureSamples

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDAttachmentFormat.fnGetSamples), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDAttachmentFormat) SetUsageFlags(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDAttachmentFormat.fnSetUsageFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDAttachmentFormat) GetUsageFlags() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDAttachmentFormat.fnGetUsageFlags), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
