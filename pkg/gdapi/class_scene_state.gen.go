// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneState struct {
  obj gdc.ObjectPtr
}

func createSceneState(obj gdc.ObjectPtr) *SceneState {
  return &SceneState{
    obj: obj,
  }
}

func (me *SceneState) BaseClass() string {
  return "SceneState"
}
