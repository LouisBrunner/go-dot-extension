// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AcceptDialog struct {
  obj gdc.ObjectPtr
}

func createAcceptDialog(obj gdc.ObjectPtr) *AcceptDialog {
  return &AcceptDialog{
    obj: obj,
  }
}

func (me *AcceptDialog) BaseClass() string {
  return "AcceptDialog"
}
