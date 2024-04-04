// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ConfirmationDialog struct {
  AcceptDialog
}

func (me *ConfirmationDialog) BaseClass() string {
  return "ConfirmationDialog"
}

func NewConfirmationDialog() *ConfirmationDialog {
  str := StringNameFromStr("ConfirmationDialog") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ConfirmationDialog{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewButton()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ConfirmationDialog) SetCancelButtonText(text String, )  {
  classNameV := StringNameFromStr("ConfirmationDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cancel_button_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ConfirmationDialog) GetCancelButtonText() String {
  classNameV := StringNameFromStr("ConfirmationDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cancel_button_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
