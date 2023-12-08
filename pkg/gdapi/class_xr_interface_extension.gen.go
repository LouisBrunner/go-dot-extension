// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRInterfaceExtension struct {
  obj gdc.ObjectPtr
}

func createXRInterfaceExtension(obj gdc.ObjectPtr) *XRInterfaceExtension {
  return &XRInterfaceExtension{
    obj: obj,
  }
}

func (me *XRInterfaceExtension) BaseClass() string {
  return "XRInterfaceExtension"
}
