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

type SceneTreeGroupCallFlags int
const (
  SceneTreeGroupCallDefault SceneTreeGroupCallFlags = 0
  SceneTreeGroupCallReverse SceneTreeGroupCallFlags = 1
  SceneTreeGroupCallDeferred SceneTreeGroupCallFlags = 2
  SceneTreeGroupCallUnique SceneTreeGroupCallFlags = 4
)
