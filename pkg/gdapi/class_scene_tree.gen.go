// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneTree struct {
  obj gdc.ObjectPtr
}

func createSceneTree(obj gdc.ObjectPtr) *SceneTree {
  return &SceneTree{
    obj: obj,
  }
}

func (me *SceneTree) BaseClass() string {
  return "SceneTree"
}
