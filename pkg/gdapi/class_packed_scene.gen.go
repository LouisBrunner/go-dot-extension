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



// Enums

type PackedSceneGenEditState int
const (
  PackedSceneGenEditStateGenEditStateDisabled PackedSceneGenEditState = 0
  PackedSceneGenEditStateGenEditStateInstance PackedSceneGenEditState = 1
  PackedSceneGenEditStateGenEditStateMain PackedSceneGenEditState = 2
  PackedSceneGenEditStateGenEditStateMainInherited PackedSceneGenEditState = 3
)

func (me *PackedScene) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PackedScene) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PackedScene) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PackedScene) Pack(path Node, )  {
  panic("TODO: implement")
}

func  (me *PackedScene) Instantiate(edit_state PackedSceneGenEditState, )  {
  panic("TODO: implement")
}

func  (me *PackedScene) CanInstantiate()  {
  panic("TODO: implement")
}

func  (me *PackedScene) GetState()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
