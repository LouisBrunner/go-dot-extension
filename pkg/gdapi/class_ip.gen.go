// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type IP struct {
  obj gdc.ObjectPtr
}

func createIP(obj gdc.ObjectPtr) *IP {
  return &IP{
    obj: obj,
  }
}

func (me *IP) BaseClass() string {
  return "IP"
}
