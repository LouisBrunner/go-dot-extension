// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RandomNumberGenerator struct {
  obj gdc.ObjectPtr
}

func createRandomNumberGenerator(obj gdc.ObjectPtr) *RandomNumberGenerator {
  return &RandomNumberGenerator{
    obj: obj,
  }
}

func (me *RandomNumberGenerator) BaseClass() string {
  return "RandomNumberGenerator"
}
