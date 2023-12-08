// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type IPUnix struct {
  obj gdc.ObjectPtr
}

func createIPUnix(obj gdc.ObjectPtr) *IPUnix {
  return &IPUnix{
    obj: obj,
  }
}

func (me *IPUnix) BaseClass() string {
  return "IPUnix"
}
