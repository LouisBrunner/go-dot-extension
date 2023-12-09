// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConfirmationDialog struct {
  obj gdc.ObjectPtr
}

func (me *ConfirmationDialog) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConfirmationDialog) BaseClass() string {
  return "ConfirmationDialog"
}



// Enums

func (me *ConfirmationDialog) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConfirmationDialog) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ConfirmationDialog) GetCancelButton()  {
  panic("TODO: implement")
}

func  (me *ConfirmationDialog) SetCancelButtonText(text String, )  {
  panic("TODO: implement")
}

func  (me *ConfirmationDialog) GetCancelButtonText()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
