// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Resource struct {
  obj gdc.ObjectPtr
}

func (me *Resource) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Resource) BaseClass() string {
  return "Resource"
}



// Enums

func (me *Resource) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Resource) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Resource) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Resource) SetPath(path String, )  {
  panic("TODO: implement")
}

func  (me *Resource) TakeOverPath(path String, )  {
  panic("TODO: implement")
}

func  (me *Resource) GetPath()  {
  panic("TODO: implement")
}

func  (me *Resource) SetName(name String, )  {
  panic("TODO: implement")
}

func  (me *Resource) GetName()  {
  panic("TODO: implement")
}

func  (me *Resource) GetRid()  {
  panic("TODO: implement")
}

func  (me *Resource) SetLocalToScene(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Resource) IsLocalToScene()  {
  panic("TODO: implement")
}

func  (me *Resource) GetLocalScene()  {
  panic("TODO: implement")
}

func  (me *Resource) SetupLocalToScene()  {
  panic("TODO: implement")
}

func  (me *Resource) EmitChanged()  {
  panic("TODO: implement")
}

func  (me *Resource) Duplicate(subresources bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
