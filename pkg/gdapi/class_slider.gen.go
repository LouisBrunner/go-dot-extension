// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Slider struct {
  obj gdc.ObjectPtr
}

func createSlider(obj gdc.ObjectPtr) *Slider {
  return &Slider{
    obj: obj,
  }
}

func (me *Slider) BaseClass() string {
  return "Slider"
}
