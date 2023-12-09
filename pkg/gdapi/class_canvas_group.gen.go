// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasGroup struct {
  obj gdc.ObjectPtr
}

func (me *CanvasGroup) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasGroup) BaseClass() string {
  return "CanvasGroup"
}



// Enums

func (me *CanvasGroup) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CanvasGroup) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CanvasGroup) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CanvasGroup) SetFitMargin(fit_margin float32, )  {
  panic("TODO: implement")
}

func  (me *CanvasGroup) GetFitMargin()  {
  panic("TODO: implement")
}

func  (me *CanvasGroup) SetClearMargin(clear_margin float32, )  {
  panic("TODO: implement")
}

func  (me *CanvasGroup) GetClearMargin()  {
  panic("TODO: implement")
}

func  (me *CanvasGroup) SetUseMipmaps(use_mipmaps bool, )  {
  panic("TODO: implement")
}

func  (me *CanvasGroup) IsUsingMipmaps()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
