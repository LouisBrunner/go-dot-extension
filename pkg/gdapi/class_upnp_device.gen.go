// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UPNPDevice struct {
  obj gdc.ObjectPtr
}

func (me *UPNPDevice) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *UPNPDevice) BaseClass() string {
  return "UPNPDevice"
}
