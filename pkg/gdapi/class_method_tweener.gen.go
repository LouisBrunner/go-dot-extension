// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MethodTweener struct {
  obj gdc.ObjectPtr
}

func createMethodTweener(obj gdc.ObjectPtr) *MethodTweener {
  return &MethodTweener{
    obj: obj,
  }
}

func (me *MethodTweener) BaseClass() string {
  return "MethodTweener"
}
