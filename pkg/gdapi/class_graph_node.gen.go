// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GraphNode struct {
  obj gdc.ObjectPtr
}

func (me *GraphNode) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GraphNode) BaseClass() string {
  return "GraphNode"
}

type GraphNodeOverlay int
const (
  GraphNodeOverlayDisabled GraphNodeOverlay = 0
  GraphNodeOverlayBreakpoint GraphNodeOverlay = 1
  GraphNodeOverlayPosition GraphNodeOverlay = 2
)
