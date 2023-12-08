// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RegEx struct {
  obj gdc.ObjectPtr
}

func createRegEx(obj gdc.ObjectPtr) *RegEx {
  return &RegEx{
    obj: obj,
  }
}

func (me *RegEx) BaseClass() string {
  return "RegEx"
}
