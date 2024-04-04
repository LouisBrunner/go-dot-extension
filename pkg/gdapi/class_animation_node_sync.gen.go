// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type AnimationNodeSync struct {
  AnimationNode
}

func (me *AnimationNodeSync) BaseClass() string {
  return "AnimationNodeSync"
}

func NewAnimationNodeSync() *AnimationNodeSync {
  str := StringNameFromStr("AnimationNodeSync") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeSync{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AnimationNodeSync) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeSync) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeSync) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNodeSync) SetUseSync(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeSync")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeSync) IsUsingSync() bool {
  classNameV := StringNameFromStr("AnimationNodeSync")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
