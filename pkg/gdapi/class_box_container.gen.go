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

type BoxContainerAlignmentMode int
const (
  BoxContainerAlignmentModeAlignmentBegin BoxContainerAlignmentMode = 0
  BoxContainerAlignmentModeAlignmentCenter BoxContainerAlignmentMode = 1
  BoxContainerAlignmentModeAlignmentEnd BoxContainerAlignmentMode = 2
)

func  (me *BoxContainer) AddSpacer(begin bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BoxContainer) SetAlignment(alignment BoxContainerAlignmentMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BoxContainer) GetAlignment() { // TODO: return value
  // TODO: implement
}

func  (me *BoxContainer) SetVertical(vertical bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BoxContainer) IsVertical() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
