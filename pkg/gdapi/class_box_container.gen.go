// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type BoxContainer struct {
  Container
}

func (me *BoxContainer) BaseClass() string {
  return "BoxContainer"
}



// Enums

type BoxContainerAlignmentMode int
const (
  BoxContainerAlignmentModeAlignmentBegin BoxContainerAlignmentMode = 0
  BoxContainerAlignmentModeAlignmentCenter BoxContainerAlignmentMode = 1
  BoxContainerAlignmentModeAlignmentEnd BoxContainerAlignmentMode = 2
)

func (me *BoxContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BoxContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BoxContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *BoxContainer) AddSpacer(begin bool, ) Control {
  classNameV := StringNameFromStr("BoxContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_spacer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1326660695) // FIXME: should cache?
  var ret Control
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&begin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BoxContainer) SetAlignment(alignment BoxContainerAlignmentMode, )  {
  classNameV := StringNameFromStr("BoxContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2456745134) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BoxContainer) GetAlignment() BoxContainerAlignmentMode {
  classNameV := StringNameFromStr("BoxContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1915476527) // FIXME: should cache?
  var ret BoxContainerAlignmentMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BoxContainer) SetVertical(vertical bool, )  {
  classNameV := StringNameFromStr("BoxContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BoxContainer) IsVertical() bool {
  classNameV := StringNameFromStr("BoxContainer")
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
// FIXME: can't seem to be able to use those from this side of the API

// Signals
