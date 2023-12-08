// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SubViewport struct {
  obj gdc.ObjectPtr
}

func (me *SubViewport) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SubViewport) BaseClass() string {
  return "SubViewport"
}
