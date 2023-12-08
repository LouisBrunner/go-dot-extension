// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BackBufferCopy struct {
  obj gdc.ObjectPtr
}

func createBackBufferCopy(obj gdc.ObjectPtr) *BackBufferCopy {
  return &BackBufferCopy{
    obj: obj,
  }
}

func (me *BackBufferCopy) BaseClass() string {
  return "BackBufferCopy"
}
