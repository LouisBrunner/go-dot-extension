// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeComment struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeComment) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeComment) BaseClass() string {
  return "VisualShaderNodeComment"
}



// Enums

func (me *VisualShaderNodeComment) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeComment) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeComment) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeComment) SetTitle(title String, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeComment) GetTitle()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeComment) SetDescription(description String, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeComment) GetDescription()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
