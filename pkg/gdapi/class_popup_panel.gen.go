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

type ptrsForPopupPanelList struct {
}

var ptrsForPopupPanel ptrsForPopupPanelList

func initPopupPanelPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PopupPanel")
  defer className.Destroy()
}

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
