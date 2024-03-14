// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OmniLight3D struct {
  Light3D
}

func (me *OmniLight3D) BaseClass() string {
  return "OmniLight3D"
}



// Enums

type OmniLight3DShadowMode int
const (
  OmniLight3DShadowModeShadowDualParaboloid OmniLight3DShadowMode = 0
  OmniLight3DShadowModeShadowCube OmniLight3DShadowMode = 1
)

func (me *OmniLight3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OmniLight3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OmniLight3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OmniLight3D) SetShadowMode(mode OmniLight3DShadowMode, )  {
  classNameV := StringNameFromStr("OmniLight3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121862228) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OmniLight3D) GetShadowMode() OmniLight3DShadowMode {
  classNameV := StringNameFromStr("OmniLight3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4181586331) // FIXME: should cache?
  var ret OmniLight3DShadowMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
