// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SubViewportContainer struct {
  obj gdc.ObjectPtr
}

func (me *SubViewportContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SubViewportContainer) BaseClass() string {
  return "SubViewportContainer"
}



// Enums

func (me *SubViewportContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SubViewportContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SubViewportContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SubViewportContainer) SetStretch(enable bool, )  {
  panic("TODO: implement")
}

func  (me *SubViewportContainer) IsStretchEnabled()  {
  panic("TODO: implement")
}

func  (me *SubViewportContainer) SetStretchShrink(amount int, )  {
  panic("TODO: implement")
}

func  (me *SubViewportContainer) GetStretchShrink()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
