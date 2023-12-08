// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFCamera struct {
  obj gdc.ObjectPtr
}

func (me *GLTFCamera) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFCamera) BaseClass() string {
  return "GLTFCamera"
}
