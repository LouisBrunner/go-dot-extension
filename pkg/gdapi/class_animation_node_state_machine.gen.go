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

type AnimationNodeStateMachine struct {
  AnimationRootNode
}

func (me *AnimationNodeStateMachine) BaseClass() string {
  return "AnimationNodeStateMachine"
}

func NewAnimationNodeStateMachine() *AnimationNodeStateMachine {
  str := StringNameFromStr("AnimationNodeStateMachine") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeStateMachine{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AnimationNodeStateMachineStateMachineType int
const (
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeRoot AnimationNodeStateMachineStateMachineType = 0
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeNested AnimationNodeStateMachineStateMachineType = 1
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeGrouped AnimationNodeStateMachineStateMachineType = 2
)

func (me *AnimationNodeStateMachine) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeStateMachine) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeStateMachine) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNodeStateMachine) AddNode(name StringName, node AnimationNode, position Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1980270704) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), node.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) ReplaceNode(name StringName, node AnimationNode, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("replace_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2559412862) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) GetNode(name StringName, ) AnimationNode {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 625644256) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) RemoveNode(name StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) RenameNode(name StringName, new_name StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), new_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) HasNode(name StringName, ) bool {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachine) GetNodeName(node AnimationNode, ) StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 739213945) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) SetNodePosition(name StringName, position Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_node_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1999414630) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) GetNodePosition(name StringName, ) Vector2 {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3100822709) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) HasTransition(from StringName, to StringName, ) bool {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachine) AddTransition(from StringName, to StringName, transition AnimationNodeStateMachineTransition, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 795486887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), transition.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) GetTransition(idx int64, ) AnimationNodeStateMachineTransition {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4192381260) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationNodeStateMachineTransition()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) GetTransitionFrom(idx int64, ) StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transition_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) GetTransitionTo(idx int64, ) StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transition_to")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) GetTransitionCount() int64 {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transition_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachine) RemoveTransitionByIndex(idx int64, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_transition_by_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) RemoveTransition(from StringName, to StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) SetGraphOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_graph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) GetGraphOffset() Vector2 {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_graph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) SetStateMachineType(state_machine_type AnimationNodeStateMachineStateMachineType, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_state_machine_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2584759088) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&state_machine_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) GetStateMachineType() AnimationNodeStateMachineStateMachineType {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_state_machine_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1140726469) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationNodeStateMachineStateMachineType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationNodeStateMachine) SetAllowTransitionToSelf(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_transition_to_self")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) IsAllowTransitionToSelf() bool {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_allow_transition_to_self")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachine) SetResetEnds(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reset_ends")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) AreEndsReset() bool {
  classNameV := StringNameFromStr("AnimationNodeStateMachine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("are_ends_reset")
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
