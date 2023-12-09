// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DirectionalLight3D struct {
  obj gdc.ObjectPtr
}

func (me *DirectionalLight3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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
  panic("TODO: implement")
}

func  (me *DirectionalLight3D) GetShadowMode()  {
  panic("TODO: implement")
}

func  (me *DirectionalLight3D) SetBlendSplits(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *DirectionalLight3D) IsBlendSplitsEnabled()  {
  panic("TODO: implement")
}

func  (me *DirectionalLight3D) SetSkyMode(mode DirectionalLight3DSkyMode, )  {
  panic("TODO: implement")
}

func  (me *DirectionalLight3D) GetSkyMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
