// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneState struct {
  obj gdc.ObjectPtr
}

func (me *SceneState) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SceneState) BaseClass() string {
  return "SceneState"
}

type SceneStateGenEditState int
const (
  SceneStateGenEditStateDisabled SceneStateGenEditState = 0
  SceneStateGenEditStateInstance SceneStateGenEditState = 1
  SceneStateGenEditStateMain SceneStateGenEditState = 2
  SceneStateGenEditStateMainInherited SceneStateGenEditState = 3
)
