// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ReferenceRect struct {
  obj gdc.ObjectPtr
}

func (me *ReferenceRect) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ReferenceRect) BaseClass() string {
  return "ReferenceRect"
}



// Enums

func (me *ReferenceRect) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ReferenceRect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ReferenceRect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ReferenceRect) GetBorderColor()  {
  panic("TODO: implement")
}

func  (me *ReferenceRect) SetBorderColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *ReferenceRect) GetBorderWidth()  {
  panic("TODO: implement")
}

func  (me *ReferenceRect) SetBorderWidth(width float32, )  {
  panic("TODO: implement")
}

func  (me *ReferenceRect) GetEditorOnly()  {
  panic("TODO: implement")
}

func  (me *ReferenceRect) SetEditorOnly(enabled bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
