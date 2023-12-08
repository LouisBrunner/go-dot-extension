// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNode struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNode) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNode) BaseClass() string {
  return "AnimationNode"
}

type AnimationNodeFilterAction int
const (
  AnimationNodeFilterActionFilterIgnore AnimationNodeFilterAction = 0
  AnimationNodeFilterActionFilterPass AnimationNodeFilterAction = 1
  AnimationNodeFilterActionFilterStop AnimationNodeFilterAction = 2
  AnimationNodeFilterActionFilterBlend AnimationNodeFilterAction = 3
)

func  (me *AnimationNode) XGetChildNodes() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) XGetParameterList() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) XGetChildByName(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) XGetParameterDefaultValue(parameter StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) XIsParameterReadOnly(parameter StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) XProcess(time float32, seek bool, is_external_seeking bool, test_only bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) XGetCaption() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) XHasFilter() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) AddInput(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) RemoveInput(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) SetInputName(input int, name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) GetInputName(input int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) GetInputCount() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) FindInput(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) SetFilterPath(path NodePath, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) IsPathFiltered(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) SetFilterEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) IsFilterEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) BlendAnimation(animation StringName, time float32, delta float32, seeked bool, is_external_seeking bool, blend float32, looped_flag AnimationLoopedFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) BlendNode(name StringName, node AnimationNode, time float32, seek bool, is_external_seeking bool, blend float32, filter AnimationNodeFilterAction, sync bool, test_only bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) BlendInput(input_index int, time float32, seek bool, is_external_seeking bool, blend float32, filter AnimationNodeFilterAction, sync bool, test_only bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) SetParameter(name StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNode) GetParameter(name StringName, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
