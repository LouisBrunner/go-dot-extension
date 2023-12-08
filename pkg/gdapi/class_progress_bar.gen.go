// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  ProgressBarFillBeginToEnd ProgressBarFillMode = 0
  ProgressBarFillEndToBegin ProgressBarFillMode = 1
  ProgressBarFillTopToBottom ProgressBarFillMode = 2
  ProgressBarFillBottomToTop ProgressBarFillMode = 3
)
