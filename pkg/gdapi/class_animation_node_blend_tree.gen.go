// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeBlendTree struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeBlendTree) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeBlendTree) BaseClass() string {
  return "AnimationNodeBlendTree"
}



// Constants

var (
  AnimationNodeBlendTreeConnectionOk = "0" // TODO: construct correctly
  AnimationNodeBlendTreeConnectionErrorNoInput = "1" // TODO: construct correctly
  AnimationNodeBlendTreeConnectionErrorNoInputIndex = "2" // TODO: construct correctly
  AnimationNodeBlendTreeConnectionErrorNoOutput = "3" // TODO: construct correctly
  AnimationNodeBlendTreeConnectionErrorSameNode = "4" // TODO: construct correctly
  AnimationNodeBlendTreeConnectionErrorConnectionExists = "5" // TODO: construct correctly
)

// Enums

func (me *AnimationNodeBlendTree) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeBlendTree) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeBlendTree) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimationNodeBlendTree) AddNode(name StringName, node AnimationNode, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendTree) GetNode(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendTree) RemoveNode(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendTree) RenameNode(name StringName, new_name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendTree) HasNode(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendTree) ConnectNode(input_node StringName, input_index int, output_node StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendTree) DisconnectNode(input_node StringName, input_index int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendTree) SetNodePosition(name StringName, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendTree) GetNodePosition(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendTree) SetGraphOffset(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendTree) GetGraphOffset()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
