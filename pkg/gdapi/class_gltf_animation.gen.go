// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFAnimation struct {
  obj gdc.ObjectPtr
}

func createGLTFAnimation(obj gdc.ObjectPtr) *GLTFAnimation {
  return &GLTFAnimation{
    obj: obj,
  }
}

func (me *GLTFAnimation) BaseClass() string {
  return "GLTFAnimation"
}
