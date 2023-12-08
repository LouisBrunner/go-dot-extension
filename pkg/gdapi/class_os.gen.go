// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OS struct {
  obj gdc.ObjectPtr
}

func createOS(obj gdc.ObjectPtr) *OS {
  return &OS{
    obj: obj,
  }
}

func (me *OS) BaseClass() string {
  return "OS"
}
