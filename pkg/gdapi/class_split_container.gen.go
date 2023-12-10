// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SplitContainer struct {
  obj gdc.ObjectPtr
}

func (me *SplitContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SplitContainer) BaseClass() string {
  return "SplitContainer"
}



// Enums

type SplitContainerDraggerVisibility int
const (
  SplitContainerDraggerVisibilityDraggerVisible SplitContainerDraggerVisibility = 0
  SplitContainerDraggerVisibilityDraggerHidden SplitContainerDraggerVisibility = 1
  SplitContainerDraggerVisibilityDraggerHiddenCollapsed SplitContainerDraggerVisibility = 2
)

func (me *SplitContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SplitContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SplitContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SplitContainer) SetSplitOffset(offset int, )  {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_split_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SplitContainer) GetSplitOffset() int {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_split_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SplitContainer) ClampSplitOffset()  {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clamp_split_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SplitContainer) SetCollapsed(collapsed bool, )  {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collapsed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collapsed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SplitContainer) IsCollapsed() bool {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collapsed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SplitContainer) SetDraggerVisibility(mode SplitContainerDraggerVisibility, )  {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dragger_visibility")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1168273952) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SplitContainer) GetDraggerVisibility() SplitContainerDraggerVisibility {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dragger_visibility")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 967297479) // FIXME: should cache?
  var ret SplitContainerDraggerVisibility
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SplitContainer) SetVertical(vertical bool, )  {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SplitContainer) IsVertical() bool {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SplitContainer) GetPropSplitOffset() int {
  panic("TODO: implement")
}

func (me *SplitContainer) SetPropSplitOffset(value int) {
  panic("TODO: implement")
}

func (me *SplitContainer) GetPropCollapsed() bool {
  panic("TODO: implement")
}

func (me *SplitContainer) SetPropCollapsed(value bool) {
  panic("TODO: implement")
}

func (me *SplitContainer) GetPropDraggerVisibility() int {
  panic("TODO: implement")
}

func (me *SplitContainer) SetPropDraggerVisibility(value int) {
  panic("TODO: implement")
}

func (me *SplitContainer) GetPropVertical() bool {
  panic("TODO: implement")
}

func (me *SplitContainer) SetPropVertical(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API