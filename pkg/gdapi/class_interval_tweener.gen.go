// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type IntervalTweener struct {
  obj gdc.ObjectPtr
}

func createIntervalTweener(obj gdc.ObjectPtr) *IntervalTweener {
  return &IntervalTweener{
    obj: obj,
  }
}

func (me *IntervalTweener) BaseClass() string {
  return "IntervalTweener"
}
