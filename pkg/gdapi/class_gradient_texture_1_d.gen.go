// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GradientTexture1D struct {
  obj gdc.ObjectPtr
}

func (me *GradientTexture1D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GradientTexture1D) BaseClass() string {
  return "GradientTexture1D"
}



// Enums

func (me *GradientTexture1D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GradientTexture1D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GradientTexture1D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GradientTexture1D) SetGradient(gradient Gradient, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture1D) GetGradient()  {
  panic("TODO: implement")
}

func  (me *GradientTexture1D) SetWidth(width int, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture1D) SetUseHdr(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture1D) IsUsingHdr()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
