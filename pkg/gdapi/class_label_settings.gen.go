// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LabelSettings struct {
  obj gdc.ObjectPtr
}

func createLabelSettings(obj gdc.ObjectPtr) *LabelSettings {
  return &LabelSettings{
    obj: obj,
  }
}

func (me *LabelSettings) BaseClass() string {
  return "LabelSettings"
}
