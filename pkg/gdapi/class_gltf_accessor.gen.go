// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFAccessor struct {
  obj gdc.ObjectPtr
}

func createGLTFAccessor(obj gdc.ObjectPtr) *GLTFAccessor {
  return &GLTFAccessor{
    obj: obj,
  }
}

func (me *GLTFAccessor) BaseClass() string {
  return "GLTFAccessor"
}
