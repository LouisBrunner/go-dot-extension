// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type FlowContainer struct {
  obj gdc.ObjectPtr
}

func (me *FlowContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FlowContainer) BaseClass() string {
  return "FlowContainer"
}



// Enums

type FlowContainerAlignmentMode int
const (
  FlowContainerAlignmentModeAlignmentBegin FlowContainerAlignmentMode = 0
  FlowContainerAlignmentModeAlignmentCenter FlowContainerAlignmentMode = 1
  FlowContainerAlignmentModeAlignmentEnd FlowContainerAlignmentMode = 2
)

func (me *FlowContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *FlowContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FlowContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *FlowContainer) GetLineCount() int {
  classNameV := StringNameFromStr("FlowContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FlowContainer) SetAlignment(alignment FlowContainerAlignmentMode, )  {
  classNameV := StringNameFromStr("FlowContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 575250951) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FlowContainer) GetAlignment() FlowContainerAlignmentMode {
  classNameV := StringNameFromStr("FlowContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3749743559) // FIXME: should cache?
  var ret FlowContainerAlignmentMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FlowContainer) SetVertical(vertical bool, )  {
  classNameV := StringNameFromStr("FlowContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FlowContainer) IsVertical() bool {
  classNameV := StringNameFromStr("FlowContainer")
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
