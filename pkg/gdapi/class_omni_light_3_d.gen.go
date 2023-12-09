// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OmniLight3D struct {
  obj gdc.ObjectPtr
}

func (me *OmniLight3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func (me *OmniLight3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OmniLight3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *OmniLight3D) SetShadowMode(mode OmniLight3DShadowMode, )  {
  panic("TODO: implement")
}

func  (me *OmniLight3D) GetShadowMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
