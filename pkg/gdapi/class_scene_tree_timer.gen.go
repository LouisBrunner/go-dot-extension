// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneTreeTimer struct {
  obj gdc.ObjectPtr
}

func createSceneTreeTimer(obj gdc.ObjectPtr) *SceneTreeTimer {
  return &SceneTreeTimer{
    obj: obj,
  }
}

func (me *SceneTreeTimer) BaseClass() string {
  return "SceneTreeTimer"
}
