// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TLSOptions struct {
  obj gdc.ObjectPtr
}

func createTLSOptions(obj gdc.ObjectPtr) *TLSOptions {
  return &TLSOptions{
    obj: obj,
  }
}

func (me *TLSOptions) BaseClass() string {
  return "TLSOptions"
}
