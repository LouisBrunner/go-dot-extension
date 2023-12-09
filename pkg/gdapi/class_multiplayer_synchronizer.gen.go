// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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



// Enums

type MultiplayerSynchronizerVisibilityUpdateMode int
const (
  MultiplayerSynchronizerVisibilityUpdateModeVisibilityProcessIdle MultiplayerSynchronizerVisibilityUpdateMode = 0
  MultiplayerSynchronizerVisibilityUpdateModeVisibilityProcessPhysics MultiplayerSynchronizerVisibilityUpdateMode = 1
  MultiplayerSynchronizerVisibilityUpdateModeVisibilityProcessNone MultiplayerSynchronizerVisibilityUpdateMode = 2
)

func (me *MultiplayerSynchronizer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerSynchronizer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MultiplayerSynchronizer) SetRootPath(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) GetRootPath()  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) SetReplicationInterval(milliseconds float32, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) GetReplicationInterval()  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) SetDeltaInterval(milliseconds float32, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) GetDeltaInterval()  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) SetReplicationConfig(config SceneReplicationConfig, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) GetReplicationConfig()  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) SetVisibilityUpdateMode(mode MultiplayerSynchronizerVisibilityUpdateMode, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) GetVisibilityUpdateMode()  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) UpdateVisibility(for_peer int, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) SetVisibilityPublic(visible bool, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) IsVisibilityPublic()  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) AddVisibilityFilter(filter Callable, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) RemoveVisibilityFilter(filter Callable, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) SetVisibilityFor(peer int, visible bool, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerSynchronizer) GetVisibilityFor(peer int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
