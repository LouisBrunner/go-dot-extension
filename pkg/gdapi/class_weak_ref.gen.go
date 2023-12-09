// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WeakRef struct {
  obj gdc.ObjectPtr
}

func (me *WeakRef) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WeakRef) BaseClass() string {
  return "WeakRef"
}



// Enums

func (me *WeakRef) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WeakRef) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WeakRef) GetRef()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
