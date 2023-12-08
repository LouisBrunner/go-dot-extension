// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Thread struct {
  obj gdc.ObjectPtr
}

func (me *Thread) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Thread) BaseClass() string {
  return "Thread"
}
