// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedDataContainer struct {
  obj gdc.ObjectPtr
}

func (me *PackedDataContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PackedDataContainer) BaseClass() string {
  return "PackedDataContainer"
}



// Enums

func (me *PackedDataContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PackedDataContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PackedDataContainer) Pack(value Variant, )  {
  panic("TODO: implement")
}

func  (me *PackedDataContainer) Size()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
