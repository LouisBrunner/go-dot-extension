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

type MultiplayerSynchronizerVisibilityUpdateMode int
const (
  MultiplayerSynchronizerVisibilityUpdateModeVisibilityProcessIdle MultiplayerSynchronizerVisibilityUpdateMode = 0
  MultiplayerSynchronizerVisibilityUpdateModeVisibilityProcessPhysics MultiplayerSynchronizerVisibilityUpdateMode = 1
  MultiplayerSynchronizerVisibilityUpdateModeVisibilityProcessNone MultiplayerSynchronizerVisibilityUpdateMode = 2
)

func  (me *MultiplayerSynchronizer) SetRootPath(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) GetRootPath() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) SetReplicationInterval(milliseconds float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) GetReplicationInterval() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) SetDeltaInterval(milliseconds float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) GetDeltaInterval() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) SetReplicationConfig(config SceneReplicationConfig, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) GetReplicationConfig() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) SetVisibilityUpdateMode(mode MultiplayerSynchronizerVisibilityUpdateMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) GetVisibilityUpdateMode() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) UpdateVisibility(for_peer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) SetVisibilityPublic(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) IsVisibilityPublic() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) AddVisibilityFilter(filter Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) RemoveVisibilityFilter(filter Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) SetVisibilityFor(peer int, visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerSynchronizer) GetVisibilityFor(peer int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
