// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeStateMachinePlayback struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeStateMachinePlayback) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeStateMachinePlayback) BaseClass() string {
  return "AnimationNodeStateMachinePlayback"
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
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3683006648) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(to_node.AsCTypePtr()), gdc.ConstTypePtr(&reset_on_teleport), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachinePlayback) Start(node StringName, reset bool, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3683006648) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(&reset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachinePlayback) Next()  {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachinePlayback) Stop()  {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachinePlayback) IsPlaying() bool {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachinePlayback) GetCurrentNode() StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachinePlayback) GetCurrentPlayPosition() float32 {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_play_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachinePlayback) GetCurrentLength() float32 {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachinePlayback) GetFadingFromNode() StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fading_from_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachinePlayback) GetTravelPath() StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachinePlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_travel_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
