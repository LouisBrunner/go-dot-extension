// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type PackedScene struct {
  Resource
}

func (me *PackedScene) BaseClass() string {
  return "PackedScene"
}

func NewPackedScene() *PackedScene {
  str := StringNameFromStr("PackedScene") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PackedScene{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PackedScene) Instantiate(edit_state PackedSceneGenEditState, ) Node {
  classNameV := StringNameFromStr("PackedScene")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("instantiate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2628778455) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edit_state) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()
  pinner.Pin(&edit_state)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PackedScene) CanInstantiate() bool {
  classNameV := StringNameFromStr("PackedScene")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_instantiate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PackedScene) GetState() SceneState {
  classNameV := StringNameFromStr("PackedScene")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3479783971) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSceneState()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
