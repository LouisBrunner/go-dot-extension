// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ProjectSettings struct {
  obj gdc.ObjectPtr
}

func createProjectSettings(obj gdc.ObjectPtr) *ProjectSettings {
  return &ProjectSettings{
    obj: obj,
  }
}

func (me *ProjectSettings) BaseClass() string {
  return "ProjectSettings"
}
