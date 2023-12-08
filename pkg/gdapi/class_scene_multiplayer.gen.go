// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneMultiplayer struct {
  obj gdc.ObjectPtr
}

func createSceneMultiplayer(obj gdc.ObjectPtr) *SceneMultiplayer {
  return &SceneMultiplayer{
    obj: obj,
  }
}

func (me *SceneMultiplayer) BaseClass() string {
  return "SceneMultiplayer"
}
