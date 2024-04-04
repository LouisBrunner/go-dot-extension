// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type RenderSceneBuffersConfiguration struct {
  RefCounted
}

func (me *RenderSceneBuffersConfiguration) BaseClass() string {
  return "RenderSceneBuffersConfiguration"
}

func NewRenderSceneBuffersConfiguration() *RenderSceneBuffersConfiguration {
  str := StringNameFromStr("RenderSceneBuffersConfiguration") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RenderSceneBuffersConfiguration{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersConfiguration) SetRenderTarget(render_target RID, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_render_target")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{render_target.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetInternalSize() Vector2i {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
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

func  (me *RenderSceneBuffersConfiguration) SetInternalSize(internal_size Vector2i, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_internal_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{internal_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetTargetSize() Vector2i {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersConfiguration) SetTargetSize(target_size Vector2i, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{target_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetViewCount() int64 {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
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

func  (me *RenderSceneBuffersConfiguration) SetViewCount(view_count int64, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_view_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&view_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetScaling3DMode() RenderingServerViewportScaling3DMode {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scaling_3d_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 976778074) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingServerViewportScaling3DMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetScaling3DMode(scaling_3d_mode RenderingServerViewportScaling3DMode, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scaling_3d_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 447477857) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scaling_3d_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetMsaa3D() RenderingServerViewportMSAA {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_msaa_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3109158617) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingServerViewportMSAA

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetMsaa3D(msaa_3d RenderingServerViewportMSAA, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_msaa_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3952630748) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msaa_3d) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetScreenSpaceAa() RenderingServerViewportScreenSpaceAA {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_screen_space_aa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 641513172) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingServerViewportScreenSpaceAA

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetScreenSpaceAa(screen_space_aa RenderingServerViewportScreenSpaceAA, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_screen_space_aa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 139543108) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen_space_aa) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetFsrSharpness() float64 {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fsr_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderSceneBuffersConfiguration) SetFsrSharpness(fsr_sharpness float64, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fsr_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fsr_sharpness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetTextureMipmapBias() float64 {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_mipmap_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderSceneBuffersConfiguration) SetTextureMipmapBias(texture_mipmap_bias float64, )  {
  classNameV := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_mipmap_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_mipmap_bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
