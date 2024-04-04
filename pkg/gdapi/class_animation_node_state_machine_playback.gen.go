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

type AnimationNodeStateMachinePlayback struct {
  Resource
}

func (me *AnimationNodeStateMachinePlayback) BaseClass() string {
  return "AnimationNodeStateMachinePlayback"
}

func NewAnimationNodeStateMachinePlayback() *AnimationNodeStateMachinePlayback {
  str := StringNameFromStr("AnimationNodeStateMachinePlayback") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeStateMachinePlayback{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AnimationNodeStateMachinePlayback) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeStateMachinePlayback) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeStateMachinePlayback) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNodeStateMachinePlayback) Travel(to_node StringName, reset_on_teleport bool, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("travel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3823612587) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{to_node.AsCTypePtr(), gdc.ConstTypePtr(&reset_on_teleport) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachinePlayback) Start(node StringName, reset bool, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3823612587) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), gdc.ConstTypePtr(&reset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachinePlayback) Next()  {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachinePlayback) Stop()  {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachinePlayback) IsPlaying() bool {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachinePlayback) GetCurrentNode() StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachinePlayback) GetCurrentPlayPosition() float64 {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_play_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachinePlayback) GetCurrentLength() float64 {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachinePlayback) GetFadingFromNode() StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fading_from_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachinePlayback) GetTravelPath() []StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_travel_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[StringName](ret)
}

// Signals
