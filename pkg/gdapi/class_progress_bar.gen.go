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

type ptrsForProgressBarList struct {
  fnSetFillMode gdc.MethodBindPtr
  fnGetFillMode gdc.MethodBindPtr
  fnSetShowPercentage gdc.MethodBindPtr
  fnIsPercentageShown gdc.MethodBindPtr
}

var ptrsForProgressBar ptrsForProgressBarList

func initProgressBarPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ProgressBar")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_fill_mode")
    defer methodName.Destroy()
    ptrsForProgressBar.fnSetFillMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_fill_mode")
    defer methodName.Destroy()
    ptrsForProgressBar.fnGetFillMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_show_percentage")
    defer methodName.Destroy()
    ptrsForProgressBar.fnSetShowPercentage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_percentage_shown")
    defer methodName.Destroy()
    ptrsForProgressBar.fnIsPercentageShown = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type ProgressBar struct {
  Range
}

func (me *ProgressBar) BaseClass() string {
  return "ProgressBar"
}

func NewProgressBar() *ProgressBar {
  str := StringNameFromStr("ProgressBar") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ProgressBar{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *ProgressBar) SetFillMode(mode int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProgressBar.fnSetFillMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProgressBar) GetFillMode() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProgressBar.fnGetFillMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProgressBar) SetShowPercentage(visible bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProgressBar.fnSetShowPercentage), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProgressBar) IsPercentageShown() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProgressBar.fnIsPercentageShown), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
