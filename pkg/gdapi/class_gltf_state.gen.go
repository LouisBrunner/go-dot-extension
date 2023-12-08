// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFState struct {
  obj gdc.ObjectPtr
}

func createGLTFState(obj gdc.ObjectPtr) *GLTFState {
  return &GLTFState{
    obj: obj,
  }
}

func (me *GLTFState) BaseClass() string {
  return "GLTFState"
}
