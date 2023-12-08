// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  SubViewportClearModeClearModeAlways SubViewportClearMode = 0
  SubViewportClearModeClearModeNever SubViewportClearMode = 1
  SubViewportClearModeClearModeOnce SubViewportClearMode = 2
)

type SubViewportUpdateMode int
const (
  SubViewportUpdateModeUpdateDisabled SubViewportUpdateMode = 0
  SubViewportUpdateModeUpdateOnce SubViewportUpdateMode = 1
  SubViewportUpdateModeUpdateWhenVisible SubViewportUpdateMode = 2
  SubViewportUpdateModeUpdateWhenParentVisible SubViewportUpdateMode = 3
  SubViewportUpdateModeUpdateAlways SubViewportUpdateMode = 4
)

func  (me *SubViewport) SetSize(size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *SubViewport) GetSize() { // TODO: return value
  // TODO: implement
}

func  (me *SubViewport) SetSize2DOverride(size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *SubViewport) GetSize2DOverride() { // TODO: return value
  // TODO: implement
}

func  (me *SubViewport) SetSize2DOverrideStretch(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SubViewport) IsSize2DOverrideStretchEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *SubViewport) SetUpdateMode(mode SubViewportUpdateMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *SubViewport) GetUpdateMode() { // TODO: return value
  // TODO: implement
}

func  (me *SubViewport) SetClearMode(mode SubViewportClearMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *SubViewport) GetClearMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
