// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedScene struct {
  obj gdc.ObjectPtr
}

func (me *PackedScene) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PackedScene) BaseClass() string {
  return "PackedScene"
}

type PackedSceneGenEditState int
const (
  PackedSceneGenEditStateGenEditStateDisabled PackedSceneGenEditState = 0
  PackedSceneGenEditStateGenEditStateInstance PackedSceneGenEditState = 1
  PackedSceneGenEditStateGenEditStateMain PackedSceneGenEditState = 2
  PackedSceneGenEditStateGenEditStateMainInherited PackedSceneGenEditState = 3
)

func  (me *PackedScene) Pack(path Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *PackedScene) Instantiate(edit_state PackedSceneGenEditState, ) { // TODO: return value
  // TODO: implement
}

func  (me *PackedScene) CanInstantiate() { // TODO: return value
  // TODO: implement
}

func  (me *PackedScene) GetState() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
