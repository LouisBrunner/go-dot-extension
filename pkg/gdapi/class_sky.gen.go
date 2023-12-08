// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Sky struct {
  obj gdc.ObjectPtr
}

func createSky(obj gdc.ObjectPtr) *Sky {
  return &Sky{
    obj: obj,
  }
}

func (me *Sky) BaseClass() string {
  return "Sky"
}
