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



// Enums

type AnimationNodeFilterAction int
const (
  AnimationNodeFilterActionFilterIgnore AnimationNodeFilterAction = 0
  AnimationNodeFilterActionFilterPass AnimationNodeFilterAction = 1
  AnimationNodeFilterActionFilterStop AnimationNodeFilterAction = 2
  AnimationNodeFilterActionFilterBlend AnimationNodeFilterAction = 3
)

func (me *AnimationNode) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNode) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimationNode) XGetChildNodes()  {
  panic("TODO: implement")
}

func  (me *AnimationNode) XGetParameterList()  {
  panic("TODO: implement")
}

func  (me *AnimationNode) XGetChildByName(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) XGetParameterDefaultValue(parameter StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) XIsParameterReadOnly(parameter StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) XProcess(time float32, seek bool, is_external_seeking bool, test_only bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) XGetCaption()  {
  panic("TODO: implement")
}

func  (me *AnimationNode) XHasFilter()  {
  panic("TODO: implement")
}

func  (me *AnimationNode) AddInput(name String, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) RemoveInput(index int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) SetInputName(input int, name String, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) GetInputName(input int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) GetInputCount()  {
  panic("TODO: implement")
}

func  (me *AnimationNode) FindInput(name String, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) SetFilterPath(path NodePath, enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) IsPathFiltered(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) SetFilterEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) IsFilterEnabled()  {
  panic("TODO: implement")
}

func  (me *AnimationNode) BlendAnimation(animation StringName, time float32, delta float32, seeked bool, is_external_seeking bool, blend float32, looped_flag AnimationLoopedFlag, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) BlendNode(name StringName, node AnimationNode, time float32, seek bool, is_external_seeking bool, blend float32, filter AnimationNodeFilterAction, sync bool, test_only bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) BlendInput(input_index int, time float32, seek bool, is_external_seeking bool, blend float32, filter AnimationNodeFilterAction, sync bool, test_only bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) SetParameter(name StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *AnimationNode) GetParameter(name StringName, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
