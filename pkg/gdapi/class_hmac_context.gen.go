// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HMACContext struct {
  obj gdc.ObjectPtr
}

func createHMACContext(obj gdc.ObjectPtr) *HMACContext {
  return &HMACContext{
    obj: obj,
  }
}

func (me *HMACContext) BaseClass() string {
  return "HMACContext"
}
