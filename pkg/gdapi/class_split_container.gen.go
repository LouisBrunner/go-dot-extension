// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SplitContainer struct {
  obj gdc.ObjectPtr
}

func createSplitContainer(obj gdc.ObjectPtr) *SplitContainer {
  return &SplitContainer{
    obj: obj,
  }
}

func (me *SplitContainer) BaseClass() string {
  return "SplitContainer"
}
