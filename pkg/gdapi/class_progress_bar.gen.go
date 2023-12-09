// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ProgressBar struct {
  obj gdc.ObjectPtr
}

func (me *ProgressBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ProgressBar) BaseClass() string {
  return "ProgressBar"
}



// Enums

type ProgressBarFillMode int
const (
  ProgressBarFillModeFillBeginToEnd ProgressBarFillMode = 0
  ProgressBarFillModeFillEndToBegin ProgressBarFillMode = 1
  ProgressBarFillModeFillTopToBottom ProgressBarFillMode = 2
  ProgressBarFillModeFillBottomToTop ProgressBarFillMode = 3
)

func (me *ProgressBar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ProgressBar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ProgressBar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ProgressBar) SetFillMode(mode int, )  {
  panic("TODO: implement")
}

func  (me *ProgressBar) GetFillMode()  {
  panic("TODO: implement")
}

func  (me *ProgressBar) SetShowPercentage(visible bool, )  {
  panic("TODO: implement")
}

func  (me *ProgressBar) IsPercentageShown()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
