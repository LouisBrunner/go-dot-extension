// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("ProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fill_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProgressBar) GetFillMode() int {
  classNameV := StringNameFromStr("ProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fill_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProgressBar) SetShowPercentage(visible bool, )  {
  classNameV := StringNameFromStr("ProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_percentage")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProgressBar) IsPercentageShown() bool {
  classNameV := StringNameFromStr("ProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_percentage_shown")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ProgressBar) GetPropFillMode() int {
  panic("TODO: implement")
}

func (me *ProgressBar) SetPropFillMode(value int) {
  panic("TODO: implement")
}

func (me *ProgressBar) GetPropShowPercentage() bool {
  panic("TODO: implement")
}

func (me *ProgressBar) SetPropShowPercentage(value bool) {
  panic("TODO: implement")
}