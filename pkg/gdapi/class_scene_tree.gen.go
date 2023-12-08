// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneTree struct {
  obj gdc.ObjectPtr
}

func (me *SceneTree) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SceneTree) BaseClass() string {
  return "SceneTree"
}
