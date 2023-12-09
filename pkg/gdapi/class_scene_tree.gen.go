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



// Enums

type SceneTreeGroupCallFlags int
const (
  SceneTreeGroupCallFlagsGroupCallDefault SceneTreeGroupCallFlags = 0
  SceneTreeGroupCallFlagsGroupCallReverse SceneTreeGroupCallFlags = 1
  SceneTreeGroupCallFlagsGroupCallDeferred SceneTreeGroupCallFlags = 2
  SceneTreeGroupCallFlagsGroupCallUnique SceneTreeGroupCallFlags = 4
)

func (me *SceneTree) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SceneTree) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SceneTree) GetRoot()  {
  panic("TODO: implement")
}

func  (me *SceneTree) HasGroup(name StringName, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) IsAutoAcceptQuit()  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetAutoAcceptQuit(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) IsQuitOnGoBack()  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetQuitOnGoBack(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetDebugCollisionsHint(enable bool, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) IsDebuggingCollisionsHint()  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetDebugPathsHint(enable bool, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) IsDebuggingPathsHint()  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetDebugNavigationHint(enable bool, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) IsDebuggingNavigationHint()  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetEditedSceneRoot(scene Node, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) GetEditedSceneRoot()  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetPause(enable bool, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) IsPaused()  {
  panic("TODO: implement")
}

func  (me *SceneTree) CreateTimer(time_sec float32, process_always bool, process_in_physics bool, ignore_time_scale bool, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) CreateTween()  {
  panic("TODO: implement")
}

func  (me *SceneTree) GetProcessedTweens()  {
  panic("TODO: implement")
}

func  (me *SceneTree) GetNodeCount()  {
  panic("TODO: implement")
}

func  (me *SceneTree) GetFrame()  {
  panic("TODO: implement")
}

func  (me *SceneTree) Quit(exit_code int, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) QueueDelete(obj Object, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) CallGroupFlags(flags int, group StringName, method StringName, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) NotifyGroupFlags(call_flags int, group StringName, notification int, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetGroupFlags(call_flags int, group StringName, property String, value Variant, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) CallGroup(group StringName, method StringName, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) NotifyGroup(group StringName, notification int, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetGroup(group StringName, property String, value Variant, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) GetNodesInGroup(group StringName, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) GetFirstNodeInGroup(group StringName, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetCurrentScene(child_node Node, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) GetCurrentScene()  {
  panic("TODO: implement")
}

func  (me *SceneTree) ChangeSceneToFile(path String, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) ChangeSceneToPacked(packed_scene PackedScene, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) ReloadCurrentScene()  {
  panic("TODO: implement")
}

func  (me *SceneTree) UnloadCurrentScene()  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetMultiplayer(multiplayer MultiplayerAPI, root_path NodePath, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) GetMultiplayer(for_path NodePath, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) SetMultiplayerPollEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *SceneTree) IsMultiplayerPollEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
