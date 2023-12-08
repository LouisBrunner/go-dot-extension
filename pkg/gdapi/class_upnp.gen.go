// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UPNP struct {
  obj gdc.ObjectPtr
}

func (me *UPNP) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *UPNP) BaseClass() string {
  return "UPNP"
}
