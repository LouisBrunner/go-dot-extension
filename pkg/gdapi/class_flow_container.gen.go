// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type FlowContainerAlignmentMode int
const (
  FlowContainerAlignmentBegin FlowContainerAlignmentMode = 0
  FlowContainerAlignmentCenter FlowContainerAlignmentMode = 1
  FlowContainerAlignmentEnd FlowContainerAlignmentMode = 2
)
