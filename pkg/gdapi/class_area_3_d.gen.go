// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Area3D struct {
  obj gdc.ObjectPtr
}

func (me *Area3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Area3D) BaseClass() string {
  return "Area3D"
}

type Area3DSpaceOverride int
const (
  Area3DSpaceOverrideDisabled Area3DSpaceOverride = 0
  Area3DSpaceOverrideCombine Area3DSpaceOverride = 1
  Area3DSpaceOverrideCombineReplace Area3DSpaceOverride = 2
  Area3DSpaceOverrideReplace Area3DSpaceOverride = 3
  Area3DSpaceOverrideReplaceCombine Area3DSpaceOverride = 4
)
