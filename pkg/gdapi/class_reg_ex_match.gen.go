// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RegExMatch struct {
  obj gdc.ObjectPtr
}

func createRegExMatch(obj gdc.ObjectPtr) *RegExMatch {
  return &RegExMatch{
    obj: obj,
  }
}

func (me *RegExMatch) BaseClass() string {
  return "RegExMatch"
}
