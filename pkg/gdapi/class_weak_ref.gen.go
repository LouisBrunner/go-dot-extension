// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WeakRef struct {
  obj gdc.ObjectPtr
}

func (me *WeakRef) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WeakRef) BaseClass() string {
  return "WeakRef"
}
