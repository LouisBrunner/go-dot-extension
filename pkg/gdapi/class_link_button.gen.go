// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LinkButton struct {
  obj gdc.ObjectPtr
}

func createLinkButton(obj gdc.ObjectPtr) *LinkButton {
  return &LinkButton{
    obj: obj,
  }
}

func (me *LinkButton) BaseClass() string {
  return "LinkButton"
}
