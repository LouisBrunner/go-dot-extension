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

type ptrsForRenderSceneBuffersConfigurationList struct {
  fnGetRenderTarget gdc.MethodBindPtr
  fnSetRenderTarget gdc.MethodBindPtr
  fnGetInternalSize gdc.MethodBindPtr
  fnSetInternalSize gdc.MethodBindPtr
  fnGetTargetSize gdc.MethodBindPtr
  fnSetTargetSize gdc.MethodBindPtr
  fnGetViewCount gdc.MethodBindPtr
  fnSetViewCount gdc.MethodBindPtr
  fnGetScaling3DMode gdc.MethodBindPtr
  fnSetScaling3DMode gdc.MethodBindPtr
  fnGetMsaa3D gdc.MethodBindPtr
  fnSetMsaa3D gdc.MethodBindPtr
  fnGetScreenSpaceAa gdc.MethodBindPtr
  fnSetScreenSpaceAa gdc.MethodBindPtr
  fnGetFsrSharpness gdc.MethodBindPtr
  fnSetFsrSharpness gdc.MethodBindPtr
  fnGetTextureMipmapBias gdc.MethodBindPtr
  fnSetTextureMipmapBias gdc.MethodBindPtr
}

var ptrsForRenderSceneBuffersConfiguration ptrsForRenderSceneBuffersConfigurationList

func initRenderSceneBuffersConfigurationPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RenderSceneBuffersConfiguration")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_render_target")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnGetRenderTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
  {
    methodName := StringNameFromStr("set_render_target")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnSetRenderTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("get_internal_size")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnGetInternalSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_internal_size")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnSetInternalSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_target_size")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnGetTargetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_target_size")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnSetTargetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_view_count")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnGetViewCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_view_count")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnSetViewCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_scaling_3d_mode")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnGetScaling3DMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 976778074))
  }
  {
    methodName := StringNameFromStr("set_scaling_3d_mode")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnSetScaling3DMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 447477857))
  }
  {
    methodName := StringNameFromStr("get_msaa_3d")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnGetMsaa3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3109158617))
  }
  {
    methodName := StringNameFromStr("set_msaa_3d")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnSetMsaa3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3952630748))
  }
  {
    methodName := StringNameFromStr("get_screen_space_aa")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnGetScreenSpaceAa = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 641513172))
  }
  {
    methodName := StringNameFromStr("set_screen_space_aa")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnSetScreenSpaceAa = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 139543108))
  }
  {
    methodName := StringNameFromStr("get_fsr_sharpness")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnGetFsrSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_fsr_sharpness")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnSetFsrSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_texture_mipmap_bias")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnGetTextureMipmapBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_texture_mipmap_bias")
    defer methodName.Destroy()
    ptrsForRenderSceneBuffersConfiguration.fnSetTextureMipmapBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnGetRenderTarget), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersConfiguration) SetRenderTarget(render_target RID, )  {
  cargs := []gdc.ConstTypePtr{render_target.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnSetRenderTarget), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetInternalSize() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnGetInternalSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersConfiguration) SetInternalSize(internal_size Vector2i, )  {
  cargs := []gdc.ConstTypePtr{internal_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnSetInternalSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetTargetSize() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnGetTargetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderSceneBuffersConfiguration) SetTargetSize(target_size Vector2i, )  {
  cargs := []gdc.ConstTypePtr{target_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnSetTargetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetViewCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnGetViewCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderSceneBuffersConfiguration) SetViewCount(view_count int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&view_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnSetViewCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetScaling3DMode() RenderingServerViewportScaling3DMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingServerViewportScaling3DMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnGetScaling3DMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetScaling3DMode(scaling_3d_mode RenderingServerViewportScaling3DMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scaling_3d_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnSetScaling3DMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetMsaa3D() RenderingServerViewportMSAA {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingServerViewportMSAA

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnGetMsaa3D), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetMsaa3D(msaa_3d RenderingServerViewportMSAA, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msaa_3d) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnSetMsaa3D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetScreenSpaceAa() RenderingServerViewportScreenSpaceAA {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingServerViewportScreenSpaceAA

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnGetScreenSpaceAa), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RenderSceneBuffersConfiguration) SetScreenSpaceAa(screen_space_aa RenderingServerViewportScreenSpaceAA, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen_space_aa) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnSetScreenSpaceAa), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetFsrSharpness() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnGetFsrSharpness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderSceneBuffersConfiguration) SetFsrSharpness(fsr_sharpness float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fsr_sharpness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnSetFsrSharpness), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderSceneBuffersConfiguration) GetTextureMipmapBias() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnGetTextureMipmapBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderSceneBuffersConfiguration) SetTextureMipmapBias(texture_mipmap_bias float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_mipmap_bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffersConfiguration.fnSetTextureMipmapBias), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
