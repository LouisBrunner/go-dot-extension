// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedDataContainerRef struct {
  obj gdc.ObjectPtr
}

func (me *PackedDataContainerRef) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PackedDataContainerRef) BaseClass() string {
  return "PackedDataContainerRef"
}



// Enums

func (me *PackedDataContainerRef) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PackedDataContainerRef) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PackedDataContainerRef) Size()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
