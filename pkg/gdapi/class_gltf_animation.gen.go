// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFAnimation struct {
  obj gdc.ObjectPtr
}

func (me *GLTFAnimation) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFAnimation) BaseClass() string {
  return "GLTFAnimation"
}
