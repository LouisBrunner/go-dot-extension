// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ParallaxLayer struct {
  obj gdc.ObjectPtr
}

func createParallaxLayer(obj gdc.ObjectPtr) *ParallaxLayer {
  return &ParallaxLayer{
    obj: obj,
  }
}

func (me *ParallaxLayer) BaseClass() string {
  return "ParallaxLayer"
}
