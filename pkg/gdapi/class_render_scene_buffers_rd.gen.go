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

type ptrsForRenderSceneBuffersRDList struct {
	fnHasTexture              gdc.MethodBindPtr
	fnCreateTexture           gdc.MethodBindPtr
	fnCreateTextureFromFormat gdc.MethodBindPtr
	fnCreateTextureView       gdc.MethodBindPtr
	fnGetTexture              gdc.MethodBindPtr
	fnGetTextureFormat        gdc.MethodBindPtr
	fnGetTextureSlice         gdc.MethodBindPtr
	fnGetTextureSliceView     gdc.MethodBindPtr
	fnGetTextureSliceSize     gdc.MethodBindPtr
	fnClearContext            gdc.MethodBindPtr
	fnGetColorTexture         gdc.MethodBindPtr
	fnGetColorLayer           gdc.MethodBindPtr
	fnGetDepthTexture         gdc.MethodBindPtr
	fnGetDepthLayer           gdc.MethodBindPtr
	fnGetVelocityTexture      gdc.MethodBindPtr
	fnGetVelocityLayer        gdc.MethodBindPtr
	fnGetRenderTarget         gdc.MethodBindPtr
	fnGetViewCount            gdc.MethodBindPtr
	fnGetInternalSize         gdc.MethodBindPtr
	fnGetUseTaa               gdc.MethodBindPtr
}

var ptrsForRenderSceneBuffersRD ptrsForRenderSceneBuffersRDList

func initRenderSceneBuffersRDPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RenderSceneBuffersRD")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("has_texture")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnHasTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("create_texture")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnCreateTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3559915770))
	}
	{
		methodName := StringNameFromStr("create_texture_from_format")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnCreateTextureFromFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3344669382))
	}
	{
		methodName := StringNameFromStr("create_texture_view")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnCreateTextureView = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 283055834))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 750006389))
	}
	{
		methodName := StringNameFromStr("get_texture_format")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetTextureFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 371461758))
	}
	{
		methodName := StringNameFromStr("get_texture_slice")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetTextureSlice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 588440706))
	}
	{
		methodName := StringNameFromStr("get_texture_slice_view")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetTextureSliceView = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 682451778))
	}
	{
		methodName := StringNameFromStr("get_texture_slice_size")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetTextureSliceSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2617625368))
	}
	{
		methodName := StringNameFromStr("clear_context")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnClearContext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_color_texture")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetColorTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("get_color_layer")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetColorLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 937000113))
	}
	{
		methodName := StringNameFromStr("get_depth_texture")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetDepthTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("get_depth_layer")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetDepthLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 937000113))
	}
	{
		methodName := StringNameFromStr("get_velocity_texture")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetVelocityTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("get_velocity_layer")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetVelocityLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 937000113))
	}
	{
		methodName := StringNameFromStr("get_render_target")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetRenderTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_view_count")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetViewCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_internal_size")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetInternalSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("get_use_taa")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffersRD.fnGetUseTaa = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type RenderSceneBuffersRD struct {
	RenderSceneBuffers
}

func (me *RenderSceneBuffersRD) BaseClass() string {
	return "RenderSceneBuffersRD"
}

func NewRenderSceneBuffersRD() *RenderSceneBuffersRD {
	str := StringNameFromStr("RenderSceneBuffersRD") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RenderSceneBuffersRD{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RenderSceneBuffersRD) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RenderSceneBuffersRD) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RenderSceneBuffersRD) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RenderSceneBuffersRD) HasTexture(context StringName, name StringName) bool {
	cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnHasTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderSceneBuffersRD) CreateTexture(context StringName, name StringName, data_format RenderingDeviceDataFormat, usage_bits int64, texture_samples RenderingDeviceTextureSamples, size Vector2i, layers int64, mipmaps int64, unique bool) RID {
	cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&data_format), gdc.ConstTypePtr(&usage_bits), gdc.ConstTypePtr(&texture_samples), size.AsCTypePtr(), gdc.ConstTypePtr(&layers), gdc.ConstTypePtr(&mipmaps), gdc.ConstTypePtr(&unique)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&data_format)
	pinner.Pin(&usage_bits)
	pinner.Pin(&texture_samples)
	pinner.Pin(&layers)
	pinner.Pin(&mipmaps)
	pinner.Pin(&unique)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnCreateTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) CreateTextureFromFormat(context StringName, name StringName, format RDTextureFormat, view RDTextureView, unique bool) RID {
	cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), format.AsCTypePtr(), view.AsCTypePtr(), gdc.ConstTypePtr(&unique)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&unique)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnCreateTextureFromFormat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) CreateTextureView(context StringName, name StringName, view_name StringName, view RDTextureView) RID {
	cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), view_name.AsCTypePtr(), view.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnCreateTextureView), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetTexture(context StringName, name StringName) RID {
	cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetTextureFormat(context StringName, name StringName) RDTextureFormat {
	cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRDTextureFormat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetTextureFormat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetTextureSlice(context StringName, name StringName, layer int64, mipmap int64, layers int64, mipmaps int64) RID {
	cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&mipmap), gdc.ConstTypePtr(&layers), gdc.ConstTypePtr(&mipmaps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&layer)
	pinner.Pin(&mipmap)
	pinner.Pin(&layers)
	pinner.Pin(&mipmaps)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetTextureSlice), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetTextureSliceView(context StringName, name StringName, layer int64, mipmap int64, layers int64, mipmaps int64, view RDTextureView) RID {
	cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&mipmap), gdc.ConstTypePtr(&layers), gdc.ConstTypePtr(&mipmaps), view.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&layer)
	pinner.Pin(&mipmap)
	pinner.Pin(&layers)
	pinner.Pin(&mipmaps)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetTextureSliceView), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetTextureSliceSize(context StringName, name StringName, mipmap int64) Vector2i {
	cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&mipmap)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&mipmap)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetTextureSliceSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) ClearContext(context StringName) {
	cargs := []gdc.ConstTypePtr{context.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnClearContext), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderSceneBuffersRD) GetColorTexture() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetColorTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetColorLayer(layer int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetColorLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetDepthTexture() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetDepthTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetDepthLayer(layer int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetDepthLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetVelocityTexture() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetVelocityTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetVelocityLayer(layer int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetVelocityLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetRenderTarget() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetRenderTarget), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetViewCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetViewCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderSceneBuffersRD) GetInternalSize() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetInternalSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneBuffersRD) GetUseTaa() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersRD.fnGetUseTaa), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
