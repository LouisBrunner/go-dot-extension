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

func  (me *RenderSceneBuffersRD) HasTexture(context StringName, name StringName, ) bool {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderSceneBuffersRD) CreateTexture(context StringName, name StringName, data_format RenderingDeviceDataFormat, usage_bits int64, texture_samples RenderingDeviceTextureSamples, size Vector2i, layers int64, mipmaps int64, unique bool, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3559915770) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&data_format) , gdc.ConstTypePtr(&usage_bits) , gdc.ConstTypePtr(&texture_samples) , size.AsCTypePtr(), gdc.ConstTypePtr(&layers) , gdc.ConstTypePtr(&mipmaps) , gdc.ConstTypePtr(&unique) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&data_format)
  pinner.Pin(&usage_bits)
  pinner.Pin(&texture_samples)
  pinner.Pin(&layers)
  pinner.Pin(&mipmaps)
  pinner.Pin(&unique)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) CreateTextureFromFormat(context StringName, name StringName, format RDTextureFormat, view RDTextureView, unique bool, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_texture_from_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3344669382) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), format.AsCTypePtr(), view.AsCTypePtr(), gdc.ConstTypePtr(&unique) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&unique)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) CreateTextureView(context StringName, name StringName, view_name StringName, view RDTextureView, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_texture_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 283055834) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), view_name.AsCTypePtr(), view.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetTexture(context StringName, name StringName, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 750006389) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetTextureFormat(context StringName, name StringName, ) RDTextureFormat {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 371461758) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRDTextureFormat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetTextureSlice(context StringName, name StringName, layer int64, mipmap int64, layers int64, mipmaps int64, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_slice")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 588440706) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&mipmap) , gdc.ConstTypePtr(&layers) , gdc.ConstTypePtr(&mipmaps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&layer)
  pinner.Pin(&mipmap)
  pinner.Pin(&layers)
  pinner.Pin(&mipmaps)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetTextureSliceView(context StringName, name StringName, layer int64, mipmap int64, layers int64, mipmaps int64, view RDTextureView, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_slice_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 682451778) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&mipmap) , gdc.ConstTypePtr(&layers) , gdc.ConstTypePtr(&mipmaps) , view.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&layer)
  pinner.Pin(&mipmap)
  pinner.Pin(&layers)
  pinner.Pin(&mipmaps)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetTextureSliceSize(context StringName, name StringName, mipmap int64, ) Vector2i {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_slice_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2617625368) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&mipmap) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()
  pinner.Pin(&mipmap)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) ClearContext(context StringName, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_context")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{context.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersRD) GetColorTexture() RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetColorLayer(layer int64, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 937000113) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetDepthTexture() RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetDepthLayer(layer int64, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 937000113) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetVelocityTexture() RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetVelocityLayer(layer int64, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 937000113) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetRenderTarget() RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_render_target")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetViewCount() int64 {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_view_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderSceneBuffersRD) GetInternalSize() Vector2i {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_internal_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersRD) GetUseTaa() bool {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_taa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
