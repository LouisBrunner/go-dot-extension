// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SubViewport struct {
  obj gdc.ObjectPtr
}

func (me *SubViewport) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SubViewport) BaseClass() string {
  return "SubViewport"
}

type SubViewportClearMode int
const (
  SubViewportClearModeAlways SubViewportClearMode = 0
  SubViewportClearModeNever SubViewportClearMode = 1
  SubViewportClearModeOnce SubViewportClearMode = 2
)

type SubViewportUpdateMode int
const (
  SubViewportUpdateDisabled SubViewportUpdateMode = 0
  SubViewportUpdateOnce SubViewportUpdateMode = 1
  SubViewportUpdateWhenVisible SubViewportUpdateMode = 2
  SubViewportUpdateWhenParentVisible SubViewportUpdateMode = 3
  SubViewportUpdateAlways SubViewportUpdateMode = 4
)
