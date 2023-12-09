// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeStateMachine struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeStateMachine) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeStateMachine) BaseClass() string {
  return "AnimationNodeStateMachine"
}



// Enums

type AnimationNodeStateMachineStateMachineType int
const (
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeRoot AnimationNodeStateMachineStateMachineType = 0
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeNested AnimationNodeStateMachineStateMachineType = 1
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeGrouped AnimationNodeStateMachineStateMachineType = 2
)

func (me *AnimationNodeStateMachine) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeStateMachine) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimationNodeStateMachine) AddNode(name StringName, node AnimationNode, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) ReplaceNode(name StringName, node AnimationNode, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) GetNode(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) RemoveNode(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) RenameNode(name StringName, new_name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) HasNode(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) GetNodeName(node AnimationNode, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) SetNodePosition(name StringName, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) GetNodePosition(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) HasTransition(from StringName, to StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) AddTransition(from StringName, to StringName, transition AnimationNodeStateMachineTransition, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) GetTransition(idx int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) GetTransitionFrom(idx int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) GetTransitionTo(idx int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) GetTransitionCount()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) RemoveTransitionByIndex(idx int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) RemoveTransition(from StringName, to StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) SetGraphOffset(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) GetGraphOffset()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) SetStateMachineType(state_machine_type AnimationNodeStateMachineStateMachineType, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) GetStateMachineType()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) SetAllowTransitionToSelf(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) IsAllowTransitionToSelf()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) SetResetEnds(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachine) AreEndsReset()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
