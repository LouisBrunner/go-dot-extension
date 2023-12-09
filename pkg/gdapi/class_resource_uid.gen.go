// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourceUID struct {
  obj gdc.ObjectPtr
}

func (me *ResourceUID) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceUID) BaseClass() string {
  return "ResourceUID"
}



// Constants

var (
  ResourceUIDInvalidId = "-1" // TODO: construct correctly
)

// Enums

func (me *ResourceUID) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceUID) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceUID) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ResourceUID) IdToText(id int, )  {
  panic("TODO: implement")
}

func  (me *ResourceUID) TextToId(text_id String, )  {
  panic("TODO: implement")
}

func  (me *ResourceUID) CreateId()  {
  panic("TODO: implement")
}

func  (me *ResourceUID) HasId(id int, )  {
  panic("TODO: implement")
}

func  (me *ResourceUID) AddId(id int, path String, )  {
  panic("TODO: implement")
}

func  (me *ResourceUID) SetId(id int, path String, )  {
  panic("TODO: implement")
}

func  (me *ResourceUID) GetIdPath(id int, )  {
  panic("TODO: implement")
}

func  (me *ResourceUID) RemoveId(id int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
