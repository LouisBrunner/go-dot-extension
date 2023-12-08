// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneTree struct {
  obj gdc.ObjectPtr
}

func (me *SceneTree) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SceneTree) BaseClass() string {
  return "SceneTree"
}

type SceneTreeGroupCallFlags int
const (
  SceneTreeGroupCallFlagsGroupCallDefault SceneTreeGroupCallFlags = 0
  SceneTreeGroupCallFlagsGroupCallReverse SceneTreeGroupCallFlags = 1
  SceneTreeGroupCallFlagsGroupCallDeferred SceneTreeGroupCallFlags = 2
  SceneTreeGroupCallFlagsGroupCallUnique SceneTreeGroupCallFlags = 4
)

func  (me *SceneTree) GetRoot() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) HasGroup(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) IsAutoAcceptQuit() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetAutoAcceptQuit(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) IsQuitOnGoBack() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetQuitOnGoBack(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetDebugCollisionsHint(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) IsDebuggingCollisionsHint() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetDebugPathsHint(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) IsDebuggingPathsHint() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetDebugNavigationHint(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) IsDebuggingNavigationHint() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetEditedSceneRoot(scene Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) GetEditedSceneRoot() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetPause(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) IsPaused() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) CreateTimer(time_sec float32, process_always bool, process_in_physics bool, ignore_time_scale bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) CreateTween() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) GetProcessedTweens() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) GetNodeCount() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) GetFrame() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) Quit(exit_code int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) QueueDelete(obj Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) CallGroupFlags(flags int, group StringName, method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) NotifyGroupFlags(call_flags int, group StringName, notification int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetGroupFlags(call_flags int, group StringName, property String, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) CallGroup(group StringName, method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) NotifyGroup(group StringName, notification int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetGroup(group StringName, property String, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) GetNodesInGroup(group StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) GetFirstNodeInGroup(group StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetCurrentScene(child_node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) GetCurrentScene() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) ChangeSceneToFile(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) ChangeSceneToPacked(packed_scene PackedScene, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) ReloadCurrentScene() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) UnloadCurrentScene() { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetMultiplayer(multiplayer MultiplayerAPI, root_path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) GetMultiplayer(for_path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) SetMultiplayerPollEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SceneTree) IsMultiplayerPollEnabled() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
