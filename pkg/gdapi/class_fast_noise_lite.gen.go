// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FastNoiseLite struct {
  obj gdc.ObjectPtr
}

func createFastNoiseLite(obj gdc.ObjectPtr) *FastNoiseLite {
  return &FastNoiseLite{
    obj: obj,
  }
}

func (me *FastNoiseLite) BaseClass() string {
  return "FastNoiseLite"
}
