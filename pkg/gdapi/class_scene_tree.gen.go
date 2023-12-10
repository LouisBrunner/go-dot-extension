// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *SceneTree) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SceneTree) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SceneTree) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SceneTree) GetRoot() Window {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1757182445) // FIXME: should cache?
  var ret Window
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) HasGroup(name StringName, ) bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) IsAutoAcceptQuit() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_auto_accept_quit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) SetAutoAcceptQuit(enabled bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_accept_quit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) IsQuitOnGoBack() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_quit_on_go_back")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) SetQuitOnGoBack(enabled bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_quit_on_go_back")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) SetDebugCollisionsHint(enable bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_collisions_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) IsDebuggingCollisionsHint() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_debugging_collisions_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) SetDebugPathsHint(enable bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_paths_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) IsDebuggingPathsHint() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_debugging_paths_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) SetDebugNavigationHint(enable bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_navigation_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) IsDebuggingNavigationHint() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_debugging_navigation_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) SetEditedSceneRoot(scene Node, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edited_scene_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scene.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) GetEditedSceneRoot() Node {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited_scene_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) SetPause(enable bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pause")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) IsPaused() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) CreateTimer(time_sec float32, process_always bool, process_in_physics bool, ignore_time_scale bool, ) SceneTreeTimer {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_timer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1780978058) // FIXME: should cache?
  var ret SceneTreeTimer
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec), gdc.ConstTypePtr(&process_always), gdc.ConstTypePtr(&process_in_physics), gdc.ConstTypePtr(&ignore_time_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) CreateTween() Tween {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_tween")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3426978995) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) GetProcessedTweens() Tween {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_processed_tweens")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) GetNodeCount() int {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) GetFrame() int {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) Quit(exit_code int, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("quit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exit_code), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) QueueDelete(obj Object, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue_delete")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3975164845) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obj.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) CallGroupFlags(flags int, group StringName, method StringName, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("call_group_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1527739229) // FIXME: should cache?
  cargs := []gdc.ConstVariantPtr{gdc.ConstVariantPtr(&flags), gdc.ConstVariantPtr(group.AsCTypePtr()), gdc.ConstVariantPtr(method.AsCTypePtr()), }
  err := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, err)
  if err.Error != gdc.CallOk {
    panic(err) // TODO: return `err`?
  }

}

func  (me *SceneTree) NotifyGroupFlags(call_flags int, group StringName, notification int, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("notify_group_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1245489420) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&call_flags), gdc.ConstTypePtr(group.AsCTypePtr()), gdc.ConstTypePtr(&notification), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) SetGroupFlags(call_flags int, group StringName, property String, value Variant, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_group_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3497599527) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&call_flags), gdc.ConstTypePtr(group.AsCTypePtr()), gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) CallGroup(group StringName, method StringName, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("call_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1257962832) // FIXME: should cache?
  cargs := []gdc.ConstVariantPtr{gdc.ConstVariantPtr(group.AsCTypePtr()), gdc.ConstVariantPtr(method.AsCTypePtr()), }
  err := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, err)
  if err.Error != gdc.CallOk {
    panic(err) // TODO: return `err`?
  }

}

func  (me *SceneTree) NotifyGroup(group StringName, notification int, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("notify_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2415702435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(group.AsCTypePtr()), gdc.ConstTypePtr(&notification), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) SetGroup(group StringName, property String, value Variant, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1279312029) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(group.AsCTypePtr()), gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) GetNodesInGroup(group StringName, ) Node {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_nodes_in_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 689397652) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(group.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) GetFirstNodeInGroup(group StringName, ) Node {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_first_node_in_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4071044623) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(group.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) SetCurrentScene(child_node Node, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(child_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) GetCurrentScene() Node {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) ChangeSceneToFile(path String, ) Error {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("change_scene_to_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) ChangeSceneToPacked(packed_scene PackedScene, ) Error {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("change_scene_to_packed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 107349098) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(packed_scene.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) ReloadCurrentScene() Error {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reload_current_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) UnloadCurrentScene()  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unload_current_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) SetMultiplayer(multiplayer MultiplayerAPI, root_path NodePath, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multiplayer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2385607013) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(multiplayer.AsCTypePtr()), gdc.ConstTypePtr(root_path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) GetMultiplayer(for_path NodePath, ) MultiplayerAPI {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_multiplayer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3453401404) // FIXME: should cache?
  var ret MultiplayerAPI
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(for_path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneTree) SetMultiplayerPollEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multiplayer_poll_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTree) IsMultiplayerPollEnabled() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_multiplayer_poll_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SceneTree) GetPropAutoAcceptQuit() bool {
  panic("TODO: implement")
}

func (me *SceneTree) SetPropAutoAcceptQuit(value bool) {
  panic("TODO: implement")
}

func (me *SceneTree) GetPropQuitOnGoBack() bool {
  panic("TODO: implement")
}

func (me *SceneTree) SetPropQuitOnGoBack(value bool) {
  panic("TODO: implement")
}

func (me *SceneTree) GetPropDebugCollisionsHint() bool {
  panic("TODO: implement")
}

func (me *SceneTree) SetPropDebugCollisionsHint(value bool) {
  panic("TODO: implement")
}

func (me *SceneTree) GetPropDebugPathsHint() bool {
  panic("TODO: implement")
}

func (me *SceneTree) SetPropDebugPathsHint(value bool) {
  panic("TODO: implement")
}

func (me *SceneTree) GetPropDebugNavigationHint() bool {
  panic("TODO: implement")
}

func (me *SceneTree) SetPropDebugNavigationHint(value bool) {
  panic("TODO: implement")
}

func (me *SceneTree) GetPropPaused() bool {
  panic("TODO: implement")
}

func (me *SceneTree) SetPropPaused(value bool) {
  panic("TODO: implement")
}

func (me *SceneTree) GetPropEditedSceneRoot() Node {
  panic("TODO: implement")
}

func (me *SceneTree) SetPropEditedSceneRoot(value Node) {
  panic("TODO: implement")
}

func (me *SceneTree) GetPropCurrentScene() Node {
  panic("TODO: implement")
}

func (me *SceneTree) SetPropCurrentScene(value Node) {
  panic("TODO: implement")
}

func (me *SceneTree) GetPropRoot() Node {
  panic("TODO: implement")
}

func (me *SceneTree) SetPropRoot(value Node) {
  panic("TODO: implement")
}

func (me *SceneTree) GetPropMultiplayerPoll() bool {
  panic("TODO: implement")
}

func (me *SceneTree) SetPropMultiplayerPoll(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API