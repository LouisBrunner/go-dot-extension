// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HSlider struct {
  obj gdc.ObjectPtr
}

func createHSlider(obj gdc.ObjectPtr) *HSlider {
  return &HSlider{
    obj: obj,
  }
}

func (me *HSlider) BaseClass() string {
  return "HSlider"
}
