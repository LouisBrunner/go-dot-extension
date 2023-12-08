// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  BoxContainerAlignmentBegin BoxContainerAlignmentMode = 0
  BoxContainerAlignmentCenter BoxContainerAlignmentMode = 1
  BoxContainerAlignmentEnd BoxContainerAlignmentMode = 2
)
