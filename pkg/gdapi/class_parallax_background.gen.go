// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ParallaxBackground struct {
  obj gdc.ObjectPtr
}

func createParallaxBackground(obj gdc.ObjectPtr) *ParallaxBackground {
  return &ParallaxBackground{
    obj: obj,
  }
}

func (me *ParallaxBackground) BaseClass() string {
  return "ParallaxBackground"
}
