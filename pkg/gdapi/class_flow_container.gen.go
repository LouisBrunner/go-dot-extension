// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *FlowContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FlowContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *FlowContainer) GetLineCount()  {
  panic("TODO: implement")
}

func  (me *FlowContainer) SetAlignment(alignment FlowContainerAlignmentMode, )  {
  panic("TODO: implement")
}

func  (me *FlowContainer) GetAlignment()  {
  panic("TODO: implement")
}

func  (me *FlowContainer) SetVertical(vertical bool, )  {
  panic("TODO: implement")
}

func  (me *FlowContainer) IsVertical()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
