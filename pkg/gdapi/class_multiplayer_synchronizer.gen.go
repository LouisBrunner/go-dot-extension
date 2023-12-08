// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerSynchronizer struct {
  obj gdc.ObjectPtr
}

func (me *MultiplayerSynchronizer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiplayerSynchronizer) BaseClass() string {
  return "MultiplayerSynchronizer"
}

type MultiplayerSynchronizerVisibilityUpdateMode int
const (
  MultiplayerSynchronizerVisibilityProcessIdle MultiplayerSynchronizerVisibilityUpdateMode = 0
  MultiplayerSynchronizerVisibilityProcessPhysics MultiplayerSynchronizerVisibilityUpdateMode = 1
  MultiplayerSynchronizerVisibilityProcessNone MultiplayerSynchronizerVisibilityUpdateMode = 2
)
