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

type ptrsForAnimationNodeAnimationList struct {
  fnSetAnimation gdc.MethodBindPtr
  fnGetAnimation gdc.MethodBindPtr
  fnSetPlayMode gdc.MethodBindPtr
  fnGetPlayMode gdc.MethodBindPtr
}

var ptrsForAnimationNodeAnimation ptrsForAnimationNodeAnimationList

func initAnimationNodeAnimationPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AnimationNodeAnimation")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_animation")
    defer methodName.Destroy()
    ptrsForAnimationNodeAnimation.fnSetAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_animation")
    defer methodName.Destroy()
    ptrsForAnimationNodeAnimation.fnGetAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_play_mode")
    defer methodName.Destroy()
    ptrsForAnimationNodeAnimation.fnSetPlayMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3347718873))
  }
  {
    methodName := StringNameFromStr("get_play_mode")
    defer methodName.Destroy()
    ptrsForAnimationNodeAnimation.fnGetPlayMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2061244637))
  }
}

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
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnSetAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeAnimation) GetAnimation() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnGetAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeAnimation) SetPlayMode(mode AnimationNodeAnimationPlayMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnSetPlayMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeAnimation) GetPlayMode() AnimationNodeAnimationPlayMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationNodeAnimationPlayMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnGetPlayMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
