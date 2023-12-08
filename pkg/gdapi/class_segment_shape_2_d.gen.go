// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SegmentShape2D struct {
  obj gdc.ObjectPtr
}

func (me *SegmentShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SegmentShape2D) BaseClass() string {
  return "SegmentShape2D"
}
