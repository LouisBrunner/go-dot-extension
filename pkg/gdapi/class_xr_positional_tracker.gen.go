// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRPositionalTracker struct {
  obj gdc.ObjectPtr
}

func createXRPositionalTracker(obj gdc.ObjectPtr) *XRPositionalTracker {
  return &XRPositionalTracker{
    obj: obj,
  }
}

func (me *XRPositionalTracker) BaseClass() string {
  return "XRPositionalTracker"
}
