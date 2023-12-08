// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ProgressBar struct {
  obj gdc.ObjectPtr
}

func createProgressBar(obj gdc.ObjectPtr) *ProgressBar {
  return &ProgressBar{
    obj: obj,
  }
}

func (me *ProgressBar) BaseClass() string {
  return "ProgressBar"
}
