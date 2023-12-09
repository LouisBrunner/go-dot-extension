// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDUniform struct {
  obj gdc.ObjectPtr
}

func (me *RDUniform) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDUniform) BaseClass() string {
  return "RDUniform"
}



// Enums

func (me *RDUniform) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDUniform) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDUniform) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RDUniform) SetUniformType(p_member RenderingDeviceUniformType, )  {
  panic("TODO: implement")
}

func  (me *RDUniform) GetUniformType()  {
  panic("TODO: implement")
}

func  (me *RDUniform) SetBinding(p_member int, )  {
  panic("TODO: implement")
}

func  (me *RDUniform) GetBinding()  {
  panic("TODO: implement")
}

func  (me *RDUniform) AddId(id RID, )  {
  panic("TODO: implement")
}

func  (me *RDUniform) ClearIds()  {
  panic("TODO: implement")
}

func  (me *RDUniform) GetIds()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
