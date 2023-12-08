// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFDocument struct {
  obj gdc.ObjectPtr
}

func createGLTFDocument(obj gdc.ObjectPtr) *GLTFDocument {
  return &GLTFDocument{
    obj: obj,
  }
}

func (me *GLTFDocument) BaseClass() string {
  return "GLTFDocument"
}
