// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Area2D struct {
  obj gdc.ObjectPtr
}

func (me *Area2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Area2D) BaseClass() string {
  return "Area2D"
}

type Area2DSpaceOverride int
const (
  Area2DSpaceOverrideDisabled Area2DSpaceOverride = 0
  Area2DSpaceOverrideCombine Area2DSpaceOverride = 1
  Area2DSpaceOverrideCombineReplace Area2DSpaceOverride = 2
  Area2DSpaceOverrideReplace Area2DSpaceOverride = 3
  Area2DSpaceOverrideReplaceCombine Area2DSpaceOverride = 4
)
