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

type ptrsForRDTextureFormatList struct {
	fnSetFormat             gdc.MethodBindPtr
	fnGetFormat             gdc.MethodBindPtr
	fnSetWidth              gdc.MethodBindPtr
	fnGetWidth              gdc.MethodBindPtr
	fnSetHeight             gdc.MethodBindPtr
	fnGetHeight             gdc.MethodBindPtr
	fnSetDepth              gdc.MethodBindPtr
	fnGetDepth              gdc.MethodBindPtr
	fnSetArrayLayers        gdc.MethodBindPtr
	fnGetArrayLayers        gdc.MethodBindPtr
	fnSetMipmaps            gdc.MethodBindPtr
	fnGetMipmaps            gdc.MethodBindPtr
	fnSetTextureType        gdc.MethodBindPtr
	fnGetTextureType        gdc.MethodBindPtr
	fnSetSamples            gdc.MethodBindPtr
	fnGetSamples            gdc.MethodBindPtr
	fnSetUsageBits          gdc.MethodBindPtr
	fnGetUsageBits          gdc.MethodBindPtr
	fnAddShareableFormat    gdc.MethodBindPtr
	fnRemoveShareableFormat gdc.MethodBindPtr
}

var ptrsForRDTextureFormat ptrsForRDTextureFormatList

func initRDTextureFormatPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDTextureFormat")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_format")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnSetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 565531219))
	}
	{
		methodName := StringNameFromStr("get_format")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2235804183))
	}
	{
		methodName := StringNameFromStr("set_width")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_width")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_height")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_height")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_depth")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnSetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_depth")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnGetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_array_layers")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnSetArrayLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_array_layers")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnGetArrayLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_mipmaps")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnSetMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_mipmaps")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnGetMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_texture_type")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnSetTextureType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 652343381))
	}
	{
		methodName := StringNameFromStr("get_texture_type")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnGetTextureType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4036357416))
	}
	{
		methodName := StringNameFromStr("set_samples")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnSetSamples = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3774171498))
	}
	{
		methodName := StringNameFromStr("get_samples")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnGetSamples = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 407791724))
	}
	{
		methodName := StringNameFromStr("set_usage_bits")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnSetUsageBits = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 245642367))
	}
	{
		methodName := StringNameFromStr("get_usage_bits")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnGetUsageBits = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1313398998))
	}
	{
		methodName := StringNameFromStr("add_shareable_format")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnAddShareableFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 565531219))
	}
	{
		methodName := StringNameFromStr("remove_shareable_format")
		defer methodName.Destroy()
		ptrsForRDTextureFormat.fnRemoveShareableFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 565531219))
	}

}

type RDTextureFormat struct {
	RefCounted
}

func (me *RDTextureFormat) BaseClass() string {
	return "RDTextureFormat"
}

func NewRDTextureFormat() *RDTextureFormat {
	str := StringNameFromStr("RDTextureFormat") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDTextureFormat{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDTextureFormat) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDTextureFormat) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDTextureFormat) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDTextureFormat) SetFormat(p_member RenderingDeviceDataFormat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnSetFormat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDTextureFormat) GetFormat() RenderingDeviceDataFormat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceDataFormat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDTextureFormat) SetWidth(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDTextureFormat) GetWidth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDTextureFormat) SetHeight(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDTextureFormat) GetHeight() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDTextureFormat) SetDepth(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnSetDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDTextureFormat) GetDepth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnGetDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDTextureFormat) SetArrayLayers(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnSetArrayLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDTextureFormat) GetArrayLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnGetArrayLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDTextureFormat) SetMipmaps(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnSetMipmaps), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDTextureFormat) GetMipmaps() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnGetMipmaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDTextureFormat) SetTextureType(p_member RenderingDeviceTextureType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnSetTextureType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDTextureFormat) GetTextureType() RenderingDeviceTextureType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceTextureType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnGetTextureType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDTextureFormat) SetSamples(p_member RenderingDeviceTextureSamples) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnSetSamples), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDTextureFormat) GetSamples() RenderingDeviceTextureSamples {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceTextureSamples

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnGetSamples), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDTextureFormat) SetUsageBits(p_member RenderingDeviceTextureUsageBits) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnSetUsageBits), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDTextureFormat) GetUsageBits() RenderingDeviceTextureUsageBits {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceTextureUsageBits

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnGetUsageBits), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDTextureFormat) AddShareableFormat(format RenderingDeviceDataFormat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnAddShareableFormat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDTextureFormat) RemoveShareableFormat(format RenderingDeviceDataFormat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureFormat.fnRemoveShareableFormat), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
