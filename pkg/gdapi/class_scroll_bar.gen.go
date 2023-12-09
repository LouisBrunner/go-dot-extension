// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScrollBar struct {
  obj gdc.ObjectPtr
}

func (me *ScrollBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ScrollBar) BaseClass() string {
  return "ScrollBar"
}



// Enums

func (me *ScrollBar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScrollBar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScrollBar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ScrollBar) SetCustomStep(step float32, )  {
  panic("TODO: implement")
}

func  (me *ScrollBar) GetCustomStep()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
