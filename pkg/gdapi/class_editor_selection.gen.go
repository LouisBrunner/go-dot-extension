// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorSelection struct {
  obj gdc.ObjectPtr
}

func (me *EditorSelection) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorSelection) BaseClass() string {
  return "EditorSelection"
}



// Enums

func (me *EditorSelection) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorSelection) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorSelection) Clear()  {
  panic("TODO: implement")
}

func  (me *EditorSelection) AddNode(node Node, )  {
  panic("TODO: implement")
}

func  (me *EditorSelection) RemoveNode(node Node, )  {
  panic("TODO: implement")
}

func  (me *EditorSelection) GetSelectedNodes()  {
  panic("TODO: implement")
}

func  (me *EditorSelection) GetTransformableSelectedNodes()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
