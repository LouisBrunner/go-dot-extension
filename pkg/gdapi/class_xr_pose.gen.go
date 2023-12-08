// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRPose struct {
  obj gdc.ObjectPtr
}

func createXRPose(obj gdc.ObjectPtr) *XRPose {
  return &XRPose{
    obj: obj,
  }
}

func (me *XRPose) BaseClass() string {
  return "XRPose"
}
