// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JavaClassWrapper struct {
  obj gdc.ObjectPtr
}

func (me *JavaClassWrapper) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *JavaClassWrapper) BaseClass() string {
  return "JavaClassWrapper"
}
