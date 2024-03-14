// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type DirectionalLight3D struct {
  Light3D
}

func (me *DirectionalLight3D) BaseClass() string {
  return "DirectionalLight3D"
}



// Enums

type DirectionalLight3DShadowMode int
const (
  DirectionalLight3DShadowModeShadowOrthogonal DirectionalLight3DShadowMode = 0
  DirectionalLight3DShadowModeShadowParallel2Splits DirectionalLight3DShadowMode = 1
  DirectionalLight3DShadowModeShadowParallel4Splits DirectionalLight3DShadowMode = 2
)

type DirectionalLight3DSkyMode int
const (
  DirectionalLight3DSkyModeSkyModeLightAndSky DirectionalLight3DSkyMode = 0
  DirectionalLight3DSkyModeSkyModeLightOnly DirectionalLight3DSkyMode = 1
  DirectionalLight3DSkyModeSkyModeSkyOnly DirectionalLight3DSkyMode = 2
)

func (me *DirectionalLight3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *DirectionalLight3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *DirectionalLight3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *DirectionalLight3D) SetShadowMode(mode DirectionalLight3DShadowMode, )  {
  classNameV := StringNameFromStr("DirectionalLight3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1261211726) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DirectionalLight3D) GetShadowMode() DirectionalLight3DShadowMode {
  classNameV := StringNameFromStr("DirectionalLight3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2765228544) // FIXME: should cache?
  var ret DirectionalLight3DShadowMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirectionalLight3D) SetBlendSplits(enabled bool, )  {
  classNameV := StringNameFromStr("DirectionalLight3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_splits")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DirectionalLight3D) IsBlendSplitsEnabled() bool {
  classNameV := StringNameFromStr("DirectionalLight3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_blend_splits_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirectionalLight3D) SetSkyMode(mode DirectionalLight3DSkyMode, )  {
  classNameV := StringNameFromStr("DirectionalLight3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2691194817) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DirectionalLight3D) GetSkyMode() DirectionalLight3DSkyMode {
  classNameV := StringNameFromStr("DirectionalLight3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3819982774) // FIXME: should cache?
  var ret DirectionalLight3DSkyMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
