// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AESContext struct {
  obj gdc.ObjectPtr
}

func createAESContext(obj gdc.ObjectPtr) *AESContext {
  return &AESContext{
    obj: obj,
  }
}

func (me *AESContext) BaseClass() string {
  return "AESContext"
}
