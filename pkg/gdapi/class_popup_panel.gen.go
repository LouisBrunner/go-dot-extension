// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PopupPanel struct {
  Popup
}

func (me *PopupPanel) BaseClass() string {
  return "PopupPanel"
}

func NewPopupPanel() *PopupPanel {
  str := StringNameFromStr("PopupPanel") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PopupPanel{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PopupPanel) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PopupPanel) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PopupPanel) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
