// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VSlider struct {
  obj gdc.ObjectPtr
}

func createVSlider(obj gdc.ObjectPtr) *VSlider {
  return &VSlider{
    obj: obj,
  }
}

func (me *VSlider) BaseClass() string {
  return "VSlider"
}
