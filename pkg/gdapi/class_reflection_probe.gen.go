// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ReflectionProbe struct {
  obj gdc.ObjectPtr
}

func (me *ReflectionProbe) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ReflectionProbe) BaseClass() string {
  return "ReflectionProbe"
}
