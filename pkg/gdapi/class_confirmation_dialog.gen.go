// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *ConfirmationDialog) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ConfirmationDialog) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConfirmationDialog) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ConfirmationDialog) GetCancelButton() Button {
  classNameV := StringNameFromStr("ConfirmationDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cancel_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1856205918) // FIXME: should cache?
  var ret Button
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ConfirmationDialog) SetCancelButtonText(text String, )  {
  classNameV := StringNameFromStr("ConfirmationDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cancel_button_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ConfirmationDialog) GetCancelButtonText() String {
  classNameV := StringNameFromStr("ConfirmationDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cancel_button_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ConfirmationDialog) GetPropCancelButtonText() String {
  panic("TODO: implement")
}

func (me *ConfirmationDialog) SetPropCancelButtonText(value String) {
  panic("TODO: implement")
}