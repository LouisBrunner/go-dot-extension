// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ReflectionProbe struct {
  obj gdc.ObjectPtr
}

func createReflectionProbe(obj gdc.ObjectPtr) *ReflectionProbe {
  return &ReflectionProbe{
    obj: obj,
  }
}

func (me *ReflectionProbe) BaseClass() string {
  return "ReflectionProbe"
}
