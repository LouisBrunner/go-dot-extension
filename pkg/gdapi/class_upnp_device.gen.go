// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UPNPDevice struct {
  obj gdc.ObjectPtr
}

func createUPNPDevice(obj gdc.ObjectPtr) *UPNPDevice {
  return &UPNPDevice{
    obj: obj,
  }
}

func (me *UPNPDevice) BaseClass() string {
  return "UPNPDevice"
}
