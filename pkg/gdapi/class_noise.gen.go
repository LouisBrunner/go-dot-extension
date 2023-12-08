// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Noise struct {
  obj gdc.ObjectPtr
}

func createNoise(obj gdc.ObjectPtr) *Noise {
  return &Noise{
    obj: obj,
  }
}

func (me *Noise) BaseClass() string {
  return "Noise"
}
