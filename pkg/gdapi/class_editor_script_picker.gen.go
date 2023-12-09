// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorScriptPicker struct {
  obj gdc.ObjectPtr
}

func (me *EditorScriptPicker) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorScriptPicker) BaseClass() string {
  return "EditorScriptPicker"
}



// Enums

func (me *EditorScriptPicker) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorScriptPicker) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorScriptPicker) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorScriptPicker) SetScriptOwner(owner_node Node, )  {
  panic("TODO: implement")
}

func  (me *EditorScriptPicker) GetScriptOwner()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
