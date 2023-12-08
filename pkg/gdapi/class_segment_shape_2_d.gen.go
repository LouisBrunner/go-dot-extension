// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SegmentShape2D struct {
  obj gdc.ObjectPtr
}

func createSegmentShape2D(obj gdc.ObjectPtr) *SegmentShape2D {
  return &SegmentShape2D{
    obj: obj,
  }
}

func (me *SegmentShape2D) BaseClass() string {
  return "SegmentShape2D"
}
