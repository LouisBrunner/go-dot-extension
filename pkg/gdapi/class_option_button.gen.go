// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OptionButton struct {
  obj gdc.ObjectPtr
}

func createOptionButton(obj gdc.ObjectPtr) *OptionButton {
  return &OptionButton{
    obj: obj,
  }
}

func (me *OptionButton) BaseClass() string {
  return "OptionButton"
}
