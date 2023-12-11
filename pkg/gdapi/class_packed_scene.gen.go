// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *PackedScene) Pack(path Node, ) Error {
  classNameV := StringNameFromStr("PackedScene")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2584678054) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PackedScene) Instantiate(edit_state PackedSceneGenEditState, ) Node {
  classNameV := StringNameFromStr("PackedScene")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("instantiate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2628778455) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edit_state), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PackedScene) CanInstantiate() bool {
  classNameV := StringNameFromStr("PackedScene")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_instantiate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PackedScene) GetState() SceneState {
  classNameV := StringNameFromStr("PackedScene")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3479783971) // FIXME: should cache?
  var ret SceneState
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
