// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CenterContainer struct {
  obj gdc.ObjectPtr
}

func (me *CenterContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CenterContainer) BaseClass() string {
  return "CenterContainer"
}



// Enums

func (me *CenterContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CenterContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CenterContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CenterContainer) SetUseTopLeft(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CenterContainer) IsUsingTopLeft()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
