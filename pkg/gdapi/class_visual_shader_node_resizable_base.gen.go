// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeResizableBase struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeResizableBase) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeResizableBase) BaseClass() string {
  return "VisualShaderNodeResizableBase"
}



// Enums

func (me *VisualShaderNodeResizableBase) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeResizableBase) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeResizableBase) SetSize(size Vector2, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeResizableBase) GetSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
