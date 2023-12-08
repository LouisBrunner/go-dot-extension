// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRInterface struct {
  obj gdc.ObjectPtr
}

func createXRInterface(obj gdc.ObjectPtr) *XRInterface {
  return &XRInterface{
    obj: obj,
  }
}

func (me *XRInterface) BaseClass() string {
  return "XRInterface"
}
