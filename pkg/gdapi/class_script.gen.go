// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Script struct {
  obj gdc.ObjectPtr
}

func createScript(obj gdc.ObjectPtr) *Script {
  return &Script{
    obj: obj,
  }
}

func (me *Script) BaseClass() string {
  return "Script"
}
