// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImporterMeshInstance3D struct {
  obj gdc.ObjectPtr
}

func createImporterMeshInstance3D(obj gdc.ObjectPtr) *ImporterMeshInstance3D {
  return &ImporterMeshInstance3D{
    obj: obj,
  }
}

func (me *ImporterMeshInstance3D) BaseClass() string {
  return "ImporterMeshInstance3D"
}
