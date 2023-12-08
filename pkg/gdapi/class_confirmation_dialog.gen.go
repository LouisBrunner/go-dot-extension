// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConfirmationDialog struct {
  obj gdc.ObjectPtr
}

func createConfirmationDialog(obj gdc.ObjectPtr) *ConfirmationDialog {
  return &ConfirmationDialog{
    obj: obj,
  }
}

func (me *ConfirmationDialog) BaseClass() string {
  return "ConfirmationDialog"
}
