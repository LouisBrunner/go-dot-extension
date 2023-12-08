// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImporterMesh struct {
  obj gdc.ObjectPtr
}

func createImporterMesh(obj gdc.ObjectPtr) *ImporterMesh {
  return &ImporterMesh{
    obj: obj,
  }
}

func (me *ImporterMesh) BaseClass() string {
  return "ImporterMesh"
}
