// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScrollContainer struct {
  obj gdc.ObjectPtr
}

func (me *ScrollContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ScrollContainer) BaseClass() string {
  return "ScrollContainer"
}

type ScrollContainerScrollMode int
const (
  ScrollContainerScrollModeDisabled ScrollContainerScrollMode = 0
  ScrollContainerScrollModeAuto ScrollContainerScrollMode = 1
  ScrollContainerScrollModeShowAlways ScrollContainerScrollMode = 2
  ScrollContainerScrollModeShowNever ScrollContainerScrollMode = 3
)
