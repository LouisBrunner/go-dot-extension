// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RenderSceneBuffersConfiguration struct {
  RefCounted
}

func (me *RenderSceneBuffersConfiguration) BaseClass() string {
  return "RenderSceneBuffersConfiguration"
}



// Enums

func (me *RenderSceneBuffersConfiguration) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RenderSceneBuffersConfiguration) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RenderSceneBuffersConfiguration) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RenderSceneBuffersConfiguration) GetRenderTarget() RID {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_render_target")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetRenderTarget(render_target RID, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_render_target")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(render_target.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RenderSceneBuffersConfiguration) GetInternalSize() Vector2i {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_internal_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetInternalSize(internal_size Vector2i, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_internal_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(internal_size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RenderSceneBuffersConfiguration) GetTargetSize() Vector2i {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetTargetSize(target_size Vector2i, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(target_size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RenderSceneBuffersConfiguration) GetViewCount() int {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_view_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetViewCount(view_count int, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_view_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&view_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RenderSceneBuffersConfiguration) GetScaling3DMode() RenderingServerViewportScaling3DMode {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scaling_3d_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 976778074) // FIXME: should cache?
  var ret RenderingServerViewportScaling3DMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetScaling3DMode(scaling_3d_mode RenderingServerViewportScaling3DMode, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scaling_3d_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 447477857) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scaling_3d_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RenderSceneBuffersConfiguration) GetMsaa3D() RenderingServerViewportMSAA {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_msaa_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3109158617) // FIXME: should cache?
  var ret RenderingServerViewportMSAA
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetMsaa3D(msaa_3d RenderingServerViewportMSAA, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_msaa_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3952630748) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msaa_3d), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RenderSceneBuffersConfiguration) GetScreenSpaceAa() RenderingServerViewportScreenSpaceAA {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_screen_space_aa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 641513172) // FIXME: should cache?
  var ret RenderingServerViewportScreenSpaceAA
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetScreenSpaceAa(screen_space_aa RenderingServerViewportScreenSpaceAA, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_screen_space_aa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 139543108) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen_space_aa), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RenderSceneBuffersConfiguration) GetFsrSharpness() float32 {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fsr_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetFsrSharpness(fsr_sharpness float32, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fsr_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fsr_sharpness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RenderSceneBuffersConfiguration) GetTextureMipmapBias() float32 {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_mipmap_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetTextureMipmapBias(texture_mipmap_bias float32, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_mipmap_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_mipmap_bias), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
