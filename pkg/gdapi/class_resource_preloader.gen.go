// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourcePreloader struct {
  obj gdc.ObjectPtr
}

func (me *ResourcePreloader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourcePreloader) BaseClass() string {
  return "ResourcePreloader"
}



// Enums

func (me *ResourcePreloader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourcePreloader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ResourcePreloader) AddResource(name StringName, resource Resource, )  {
  panic("TODO: implement")
}

func  (me *ResourcePreloader) RemoveResource(name StringName, )  {
  panic("TODO: implement")
}

func  (me *ResourcePreloader) RenameResource(name StringName, newname StringName, )  {
  panic("TODO: implement")
}

func  (me *ResourcePreloader) HasResource(name StringName, )  {
  panic("TODO: implement")
}

func  (me *ResourcePreloader) GetResource(name StringName, )  {
  panic("TODO: implement")
}

func  (me *ResourcePreloader) GetResourceList()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
