// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InstancePlaceholder struct {
  obj gdc.ObjectPtr
}

func (me *InstancePlaceholder) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InstancePlaceholder) BaseClass() string {
  return "InstancePlaceholder"
}



// Enums

func (me *InstancePlaceholder) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InstancePlaceholder) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InstancePlaceholder) GetStoredValues(with_order bool, )  {
  panic("TODO: implement")
}

func  (me *InstancePlaceholder) CreateInstance(replace bool, custom_scene PackedScene, )  {
  panic("TODO: implement")
}

func  (me *InstancePlaceholder) GetInstancePath()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
