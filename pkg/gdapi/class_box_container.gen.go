// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoxContainer struct {
  obj gdc.ObjectPtr
}

func (me *BoxContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func  (me *BoxContainer) AddSpacer(begin bool, )  {
  panic("TODO: implement")
}

func  (me *BoxContainer) SetAlignment(alignment BoxContainerAlignmentMode, )  {
  panic("TODO: implement")
}

func  (me *BoxContainer) GetAlignment()  {
  panic("TODO: implement")
}

func  (me *BoxContainer) SetVertical(vertical bool, )  {
  panic("TODO: implement")
}

func  (me *BoxContainer) IsVertical()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
