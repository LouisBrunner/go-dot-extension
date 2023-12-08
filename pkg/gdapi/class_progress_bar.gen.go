// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ProgressBar struct {
  obj gdc.ObjectPtr
}

func (me *ProgressBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ProgressBar) BaseClass() string {
  return "ProgressBar"
}

type ProgressBarFillMode int
const (
  ProgressBarFillModeFillBeginToEnd ProgressBarFillMode = 0
  ProgressBarFillModeFillEndToBegin ProgressBarFillMode = 1
  ProgressBarFillModeFillTopToBottom ProgressBarFillMode = 2
  ProgressBarFillModeFillBottomToTop ProgressBarFillMode = 3
)

func  (me *ProgressBar) SetFillMode(mode int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ProgressBar) GetFillMode() { // TODO: return value
  // TODO: implement
}

func  (me *ProgressBar) SetShowPercentage(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ProgressBar) IsPercentageShown() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
