// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFBufferView struct {
  obj gdc.ObjectPtr
}

func createGLTFBufferView(obj gdc.ObjectPtr) *GLTFBufferView {
  return &GLTFBufferView{
    obj: obj,
  }
}

func (me *GLTFBufferView) BaseClass() string {
  return "GLTFBufferView"
}
