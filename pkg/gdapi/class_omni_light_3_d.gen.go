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

type OmniLight3DShadowMode int
const (
  OmniLight3DShadowModeShadowDualParaboloid OmniLight3DShadowMode = 0
  OmniLight3DShadowModeShadowCube OmniLight3DShadowMode = 1
)

func  (me *OmniLight3D) SetShadowMode(mode OmniLight3DShadowMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *OmniLight3D) GetShadowMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
