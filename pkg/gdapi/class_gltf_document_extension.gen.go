// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFDocumentExtension struct {
  obj gdc.ObjectPtr
}

func createGLTFDocumentExtension(obj gdc.ObjectPtr) *GLTFDocumentExtension {
  return &GLTFDocumentExtension{
    obj: obj,
  }
}

func (me *GLTFDocumentExtension) BaseClass() string {
  return "GLTFDocumentExtension"
}
