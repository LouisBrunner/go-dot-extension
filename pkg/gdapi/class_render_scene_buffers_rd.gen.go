// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RenderSceneBuffersRD struct {
  obj gdc.ObjectPtr
}

func (me *RenderSceneBuffersRD) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RenderSceneBuffersRD) BaseClass() string {
  return "RenderSceneBuffersRD"
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
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(context.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) CreateTexture(context StringName, name StringName, data_format RenderingDeviceDataFormat, usage_bits int, texture_samples RenderingDeviceTextureSamples, size Vector2i, layers int, mipmaps int, unique bool, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3559915770) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(context.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&data_format), gdc.ConstTypePtr(&usage_bits), gdc.ConstTypePtr(&texture_samples), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&layers), gdc.ConstTypePtr(&mipmaps), gdc.ConstTypePtr(&unique), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) CreateTextureFromFormat(context StringName, name StringName, format RDTextureFormat, view RDTextureView, unique bool, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_texture_from_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3344669382) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(context.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(format.AsCTypePtr()), gdc.ConstTypePtr(view.AsCTypePtr()), gdc.ConstTypePtr(&unique), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) CreateTextureView(context StringName, name StringName, view_name StringName, view RDTextureView, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_texture_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 283055834) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(context.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(view_name.AsCTypePtr()), gdc.ConstTypePtr(view.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetTexture(context StringName, name StringName, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 750006389) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(context.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetTextureFormat(context StringName, name StringName, ) RDTextureFormat {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 371461758) // FIXME: should cache?
  var ret RDTextureFormat
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(context.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetTextureSlice(context StringName, name StringName, layer int, mipmap int, layers int, mipmaps int, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_slice")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 588440706) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(context.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&mipmap), gdc.ConstTypePtr(&layers), gdc.ConstTypePtr(&mipmaps), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetTextureSliceView(context StringName, name StringName, layer int, mipmap int, layers int, mipmaps int, view RDTextureView, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_slice_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 682451778) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(context.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&mipmap), gdc.ConstTypePtr(&layers), gdc.ConstTypePtr(&mipmaps), gdc.ConstTypePtr(view.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetTextureSliceSize(context StringName, name StringName, mipmap int, ) Vector2i {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_slice_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2617625368) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(context.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&mipmap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) ClearContext(context StringName, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_context")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(context.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RenderSceneBuffersRD) GetColorTexture() RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetColorLayer(layer int, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 937000113) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetDepthTexture() RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetDepthLayer(layer int, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 937000113) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetVelocityTexture() RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetVelocityLayer(layer int, ) RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 937000113) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetRenderTarget() RID {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_render_target")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetViewCount() int {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_view_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetInternalSize() Vector2i {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_internal_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersRD) GetUseTaa() bool {
  classNameV := StringNameFromStr("RenderSceneBuffersRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_taa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
