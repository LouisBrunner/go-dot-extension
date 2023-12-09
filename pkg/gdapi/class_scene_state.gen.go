// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneState struct {
  obj gdc.ObjectPtr
}

func (me *SceneState) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SceneState) BaseClass() string {
  return "SceneState"
}



// Enums

type SceneStateGenEditState int
const (
  SceneStateGenEditStateGenEditStateDisabled SceneStateGenEditState = 0
  SceneStateGenEditStateGenEditStateInstance SceneStateGenEditState = 1
  SceneStateGenEditStateGenEditStateMain SceneStateGenEditState = 2
  SceneStateGenEditStateGenEditStateMainInherited SceneStateGenEditState = 3
)

func (me *SceneState) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SceneState) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SceneState) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SceneState) GetNodeCount()  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodeType(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodeName(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodePath(idx int, for_parent bool, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodeOwnerPath(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) IsNodeInstancePlaceholder(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodeInstancePlaceholder(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodeInstance(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodeGroups(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodeIndex(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodePropertyCount(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodePropertyName(idx int, prop_idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetNodePropertyValue(idx int, prop_idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetConnectionCount()  {
  panic("TODO: implement")
}

func  (me *SceneState) GetConnectionSource(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetConnectionSignal(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetConnectionTarget(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetConnectionMethod(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetConnectionFlags(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetConnectionBinds(idx int, )  {
  panic("TODO: implement")
}

func  (me *SceneState) GetConnectionUnbinds(idx int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
