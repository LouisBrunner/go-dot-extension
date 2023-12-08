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

type AnimationNodeStateMachineStateMachineType int
const (
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeRoot AnimationNodeStateMachineStateMachineType = 0
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeNested AnimationNodeStateMachineStateMachineType = 1
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeGrouped AnimationNodeStateMachineStateMachineType = 2
)

func  (me *AnimationNodeStateMachine) AddNode(name StringName, node AnimationNode, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) ReplaceNode(name StringName, node AnimationNode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) GetNode(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) RemoveNode(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) RenameNode(name StringName, new_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) HasNode(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) GetNodeName(node AnimationNode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) SetNodePosition(name StringName, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) GetNodePosition(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) HasTransition(from StringName, to StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) AddTransition(from StringName, to StringName, transition AnimationNodeStateMachineTransition, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) GetTransition(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) GetTransitionFrom(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) GetTransitionTo(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) GetTransitionCount() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) RemoveTransitionByIndex(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) RemoveTransition(from StringName, to StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) SetGraphOffset(offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) GetGraphOffset() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) SetStateMachineType(state_machine_type AnimationNodeStateMachineStateMachineType, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) GetStateMachineType() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) SetAllowTransitionToSelf(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) IsAllowTransitionToSelf() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) SetResetEnds(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeStateMachine) AreEndsReset() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
