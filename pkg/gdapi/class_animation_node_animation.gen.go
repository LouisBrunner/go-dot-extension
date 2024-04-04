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

type AnimationNodeAnimation struct {
  AnimationRootNode
}

func (me *AnimationNodeAnimation) BaseClass() string {
  return "AnimationNodeAnimation"
}

func NewAnimationNodeAnimation() *AnimationNodeAnimation {
  str := StringNameFromStr("AnimationNodeAnimation") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeAnimation{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AnimationNodeAnimationPlayMode int
const (
  AnimationNodeAnimationPlayModePlayModeForward AnimationNodeAnimationPlayMode = 0
  AnimationNodeAnimationPlayModePlayModeBackward AnimationNodeAnimationPlayMode = 1
)

func (me *AnimationNodeAnimation) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeAnimation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeAnimation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNodeAnimation) SetAnimation(name StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeAnimation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeAnimation) GetAnimation() StringName {
  classNameV := StringNameFromStr("AnimationNodeAnimation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeAnimation) SetPlayMode(mode AnimationNodeAnimationPlayMode, )  {
  classNameV := StringNameFromStr("AnimationNodeAnimation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_play_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3347718873) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeAnimation) GetPlayMode() AnimationNodeAnimationPlayMode {
  classNameV := StringNameFromStr("AnimationNodeAnimation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_play_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2061244637) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationNodeAnimationPlayMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
