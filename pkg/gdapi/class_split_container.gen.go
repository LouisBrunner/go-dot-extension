// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SplitContainer struct {
  obj gdc.ObjectPtr
}

func (me *SplitContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SplitContainer) BaseClass() string {
  return "SplitContainer"
}

type SplitContainerDraggerVisibility int
const (
  SplitContainerDraggerVisible SplitContainerDraggerVisibility = 0
  SplitContainerDraggerHidden SplitContainerDraggerVisibility = 1
  SplitContainerDraggerHiddenCollapsed SplitContainerDraggerVisibility = 2
)
