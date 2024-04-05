// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForSceneTreeList struct {
  fnGetRoot gdc.MethodBindPtr
  fnHasGroup gdc.MethodBindPtr
  fnIsAutoAcceptQuit gdc.MethodBindPtr
  fnSetAutoAcceptQuit gdc.MethodBindPtr
  fnIsQuitOnGoBack gdc.MethodBindPtr
  fnSetQuitOnGoBack gdc.MethodBindPtr
  fnSetDebugCollisionsHint gdc.MethodBindPtr
  fnIsDebuggingCollisionsHint gdc.MethodBindPtr
  fnSetDebugPathsHint gdc.MethodBindPtr
  fnIsDebuggingPathsHint gdc.MethodBindPtr
  fnSetDebugNavigationHint gdc.MethodBindPtr
  fnIsDebuggingNavigationHint gdc.MethodBindPtr
  fnSetEditedSceneRoot gdc.MethodBindPtr
  fnGetEditedSceneRoot gdc.MethodBindPtr
  fnSetPause gdc.MethodBindPtr
  fnIsPaused gdc.MethodBindPtr
  fnCreateTimer gdc.MethodBindPtr
  fnCreateTween gdc.MethodBindPtr
  fnGetProcessedTweens gdc.MethodBindPtr
  fnGetNodeCount gdc.MethodBindPtr
  fnGetFrame gdc.MethodBindPtr
  fnQuit gdc.MethodBindPtr
  fnQueueDelete gdc.MethodBindPtr
  fnCallGroupFlags gdc.MethodBindPtr
  fnNotifyGroupFlags gdc.MethodBindPtr
  fnSetGroupFlags gdc.MethodBindPtr
  fnCallGroup gdc.MethodBindPtr
  fnNotifyGroup gdc.MethodBindPtr
  fnSetGroup gdc.MethodBindPtr
  fnGetNodesInGroup gdc.MethodBindPtr
  fnGetFirstNodeInGroup gdc.MethodBindPtr
  fnSetCurrentScene gdc.MethodBindPtr
  fnGetCurrentScene gdc.MethodBindPtr
  fnChangeSceneToFile gdc.MethodBindPtr
  fnChangeSceneToPacked gdc.MethodBindPtr
  fnReloadCurrentScene gdc.MethodBindPtr
  fnUnloadCurrentScene gdc.MethodBindPtr
  fnSetMultiplayer gdc.MethodBindPtr
  fnGetMultiplayer gdc.MethodBindPtr
  fnSetMultiplayerPollEnabled gdc.MethodBindPtr
  fnIsMultiplayerPollEnabled gdc.MethodBindPtr
}

var ptrsForSceneTree ptrsForSceneTreeList

func initSceneTreePtrs(iface gdc.Interface) {

  className := StringNameFromStr("SceneTree")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_root")
    defer methodName.Destroy()
    ptrsForSceneTree.fnGetRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1757182445))
  }
  {
    methodName := StringNameFromStr("has_group")
    defer methodName.Destroy()
    ptrsForSceneTree.fnHasGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
  }
  {
    methodName := StringNameFromStr("is_auto_accept_quit")
    defer methodName.Destroy()
    ptrsForSceneTree.fnIsAutoAcceptQuit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_auto_accept_quit")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetAutoAcceptQuit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_quit_on_go_back")
    defer methodName.Destroy()
    ptrsForSceneTree.fnIsQuitOnGoBack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_quit_on_go_back")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetQuitOnGoBack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_debug_collisions_hint")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetDebugCollisionsHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_debugging_collisions_hint")
    defer methodName.Destroy()
    ptrsForSceneTree.fnIsDebuggingCollisionsHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_debug_paths_hint")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetDebugPathsHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_debugging_paths_hint")
    defer methodName.Destroy()
    ptrsForSceneTree.fnIsDebuggingPathsHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_debug_navigation_hint")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetDebugNavigationHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_debugging_navigation_hint")
    defer methodName.Destroy()
    ptrsForSceneTree.fnIsDebuggingNavigationHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_edited_scene_root")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetEditedSceneRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
  }
  {
    methodName := StringNameFromStr("get_edited_scene_root")
    defer methodName.Destroy()
    ptrsForSceneTree.fnGetEditedSceneRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
  }
  {
    methodName := StringNameFromStr("set_pause")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetPause = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_paused")
    defer methodName.Destroy()
    ptrsForSceneTree.fnIsPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("create_timer")
    defer methodName.Destroy()
    ptrsForSceneTree.fnCreateTimer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2709170273))
  }
  {
    methodName := StringNameFromStr("create_tween")
    defer methodName.Destroy()
    ptrsForSceneTree.fnCreateTween = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3426978995))
  }
  {
    methodName := StringNameFromStr("get_processed_tweens")
    defer methodName.Destroy()
    ptrsForSceneTree.fnGetProcessedTweens = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("get_node_count")
    defer methodName.Destroy()
    ptrsForSceneTree.fnGetNodeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_frame")
    defer methodName.Destroy()
    ptrsForSceneTree.fnGetFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("quit")
    defer methodName.Destroy()
    ptrsForSceneTree.fnQuit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1995695955))
  }
  {
    methodName := StringNameFromStr("queue_delete")
    defer methodName.Destroy()
    ptrsForSceneTree.fnQueueDelete = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3975164845))
  }
  {
    methodName := StringNameFromStr("call_group_flags")
    defer methodName.Destroy()
    ptrsForSceneTree.fnCallGroupFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1527739229))
  }
  {
    methodName := StringNameFromStr("notify_group_flags")
    defer methodName.Destroy()
    ptrsForSceneTree.fnNotifyGroupFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1245489420))
  }
  {
    methodName := StringNameFromStr("set_group_flags")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetGroupFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3497599527))
  }
  {
    methodName := StringNameFromStr("call_group")
    defer methodName.Destroy()
    ptrsForSceneTree.fnCallGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1257962832))
  }
  {
    methodName := StringNameFromStr("notify_group")
    defer methodName.Destroy()
    ptrsForSceneTree.fnNotifyGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2415702435))
  }
  {
    methodName := StringNameFromStr("set_group")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1279312029))
  }
  {
    methodName := StringNameFromStr("get_nodes_in_group")
    defer methodName.Destroy()
    ptrsForSceneTree.fnGetNodesInGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 689397652))
  }
  {
    methodName := StringNameFromStr("get_first_node_in_group")
    defer methodName.Destroy()
    ptrsForSceneTree.fnGetFirstNodeInGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4071044623))
  }
  {
    methodName := StringNameFromStr("set_current_scene")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetCurrentScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
  }
  {
    methodName := StringNameFromStr("get_current_scene")
    defer methodName.Destroy()
    ptrsForSceneTree.fnGetCurrentScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
  }
  {
    methodName := StringNameFromStr("change_scene_to_file")
    defer methodName.Destroy()
    ptrsForSceneTree.fnChangeSceneToFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
  }
  {
    methodName := StringNameFromStr("change_scene_to_packed")
    defer methodName.Destroy()
    ptrsForSceneTree.fnChangeSceneToPacked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 107349098))
  }
  {
    methodName := StringNameFromStr("reload_current_scene")
    defer methodName.Destroy()
    ptrsForSceneTree.fnReloadCurrentScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
  }
  {
    methodName := StringNameFromStr("unload_current_scene")
    defer methodName.Destroy()
    ptrsForSceneTree.fnUnloadCurrentScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_multiplayer")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetMultiplayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2385607013))
  }
  {
    methodName := StringNameFromStr("get_multiplayer")
    defer methodName.Destroy()
    ptrsForSceneTree.fnGetMultiplayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3453401404))
  }
  {
    methodName := StringNameFromStr("set_multiplayer_poll_enabled")
    defer methodName.Destroy()
    ptrsForSceneTree.fnSetMultiplayerPollEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_multiplayer_poll_enabled")
    defer methodName.Destroy()
    ptrsForSceneTree.fnIsMultiplayerPollEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type SceneTree struct {
  MainLoop
}

func (me *SceneTree) BaseClass() string {
  return "SceneTree"
}

func NewSceneTree() *SceneTree {
  str := StringNameFromStr("SceneTree") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SceneTree{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWindow()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnGetRoot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) HasGroup(name StringName, ) bool {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnHasGroup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) IsAutoAcceptQuit() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnIsAutoAcceptQuit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) SetAutoAcceptQuit(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetAutoAcceptQuit), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsQuitOnGoBack() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnIsQuitOnGoBack), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) SetQuitOnGoBack(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetQuitOnGoBack), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) SetDebugCollisionsHint(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetDebugCollisionsHint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsDebuggingCollisionsHint() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnIsDebuggingCollisionsHint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) SetDebugPathsHint(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetDebugPathsHint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsDebuggingPathsHint() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnIsDebuggingPathsHint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) SetDebugNavigationHint(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetDebugNavigationHint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsDebuggingNavigationHint() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnIsDebuggingNavigationHint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) SetEditedSceneRoot(scene Node, )  {
  cargs := []gdc.ConstTypePtr{scene.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetEditedSceneRoot), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) GetEditedSceneRoot() Node {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnGetEditedSceneRoot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) SetPause(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetPause), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsPaused() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnIsPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) CreateTimer(time_sec float64, process_always bool, process_in_physics bool, ignore_time_scale bool, ) SceneTreeTimer {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec) , gdc.ConstTypePtr(&process_always) , gdc.ConstTypePtr(&process_in_physics) , gdc.ConstTypePtr(&ignore_time_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSceneTreeTimer()
  pinner.Pin(&time_sec)
  pinner.Pin(&process_always)
  pinner.Pin(&process_in_physics)
  pinner.Pin(&ignore_time_scale)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnCreateTimer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) CreateTween() Tween {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnCreateTween), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) GetProcessedTweens() []Tween {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnGetProcessedTweens), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Tween](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *SceneTree) GetNodeCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnGetNodeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) GetFrame() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnGetFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) Quit(exit_code int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exit_code) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnQuit), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) QueueDelete(obj Object, )  {
  cargs := []gdc.ConstTypePtr{obj.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnQueueDelete), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) CallGroupFlags(flags int64, group StringName, method StringName, varargs ...Variant)  {
  cargs := make([]gdc.ConstVariantPtr, 0, 3 + len(varargs))
  intVar0 := NewIntFromInt(flags)
  defer intVar0.Destroy()
  var0 := intVar0.AsVariant()
  defer var0.Destroy()
  cargs = append(cargs, var0.AsCPtr())
  var1 := group.AsVariant()
  defer var1.Destroy()
  cargs = append(cargs, var1.AsCPtr())
  var2 := method.AsVariant()
  defer var2.Destroy()
  cargs = append(cargs, var2.AsCPtr())
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }

  cerr := &gdc.CallError{}
  giface.ObjectMethodBindCall(ensurePtr(ptrsForSceneTree.fnCallGroupFlags), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return
  }

}

func  (me *SceneTree) NotifyGroupFlags(call_flags int64, group StringName, notification int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&call_flags) , group.AsCTypePtr(), gdc.ConstTypePtr(&notification) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnNotifyGroupFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) SetGroupFlags(call_flags int64, group StringName, property String, value Variant, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&call_flags) , group.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetGroupFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) CallGroup(group StringName, method StringName, varargs ...Variant)  {
  cargs := make([]gdc.ConstVariantPtr, 0, 2 + len(varargs))
  var0 := group.AsVariant()
  defer var0.Destroy()
  cargs = append(cargs, var0.AsCPtr())
  var1 := method.AsVariant()
  defer var1.Destroy()
  cargs = append(cargs, var1.AsCPtr())
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }

  cerr := &gdc.CallError{}
  giface.ObjectMethodBindCall(ensurePtr(ptrsForSceneTree.fnCallGroup), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return
  }

}

func  (me *SceneTree) NotifyGroup(group StringName, notification int64, )  {
  cargs := []gdc.ConstTypePtr{group.AsCTypePtr(), gdc.ConstTypePtr(&notification) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnNotifyGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) SetGroup(group StringName, property String, value Variant, )  {
  cargs := []gdc.ConstTypePtr{group.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) GetNodesInGroup(group StringName, ) []Node {
  cargs := []gdc.ConstTypePtr{group.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnGetNodesInGroup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Node](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *SceneTree) GetFirstNodeInGroup(group StringName, ) Node {
  cargs := []gdc.ConstTypePtr{group.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnGetFirstNodeInGroup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) SetCurrentScene(child_node Node, )  {
  cargs := []gdc.ConstTypePtr{child_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetCurrentScene), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) GetCurrentScene() Node {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnGetCurrentScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) ChangeSceneToFile(path String, ) Error {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnChangeSceneToFile), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SceneTree) ChangeSceneToPacked(packed_scene PackedScene, ) Error {
  cargs := []gdc.ConstTypePtr{packed_scene.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnChangeSceneToPacked), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SceneTree) ReloadCurrentScene() Error {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnReloadCurrentScene), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SceneTree) UnloadCurrentScene()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnUnloadCurrentScene), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) SetMultiplayer(multiplayer MultiplayerAPI, root_path NodePath, )  {
  cargs := []gdc.ConstTypePtr{multiplayer.AsCTypePtr(), root_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetMultiplayer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) GetMultiplayer(for_path NodePath, ) MultiplayerAPI {
  cargs := []gdc.ConstTypePtr{for_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMultiplayerAPI()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnGetMultiplayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) SetMultiplayerPollEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnSetMultiplayerPollEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsMultiplayerPollEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTree.fnIsMultiplayerPollEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SceneTreeTreeChangedSignalFn func()

func (me *SceneTree) ConnectTreeChanged(subs SignalSubscribers, fn SceneTreeTreeChangedSignalFn) {
  sig := StringNameFromStr("tree_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneTree) DisconnectTreeChanged(subs SignalSubscribers, fn SceneTreeTreeChangedSignalFn) {
  sig := StringNameFromStr("tree_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type SceneTreeTreeProcessModeChangedSignalFn func()

func (me *SceneTree) ConnectTreeProcessModeChanged(subs SignalSubscribers, fn SceneTreeTreeProcessModeChangedSignalFn) {
  sig := StringNameFromStr("tree_process_mode_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneTree) DisconnectTreeProcessModeChanged(subs SignalSubscribers, fn SceneTreeTreeProcessModeChangedSignalFn) {
  sig := StringNameFromStr("tree_process_mode_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type SceneTreeNodeAddedSignalFn func(node Node, )

func (me *SceneTree) ConnectNodeAdded(subs SignalSubscribers, fn SceneTreeNodeAddedSignalFn) {
  sig := StringNameFromStr("node_added")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneTree) DisconnectNodeAdded(subs SignalSubscribers, fn SceneTreeNodeAddedSignalFn) {
  sig := StringNameFromStr("node_added")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type SceneTreeNodeRemovedSignalFn func(node Node, )

func (me *SceneTree) ConnectNodeRemoved(subs SignalSubscribers, fn SceneTreeNodeRemovedSignalFn) {
  sig := StringNameFromStr("node_removed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneTree) DisconnectNodeRemoved(subs SignalSubscribers, fn SceneTreeNodeRemovedSignalFn) {
  sig := StringNameFromStr("node_removed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type SceneTreeNodeRenamedSignalFn func(node Node, )

func (me *SceneTree) ConnectNodeRenamed(subs SignalSubscribers, fn SceneTreeNodeRenamedSignalFn) {
  sig := StringNameFromStr("node_renamed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneTree) DisconnectNodeRenamed(subs SignalSubscribers, fn SceneTreeNodeRenamedSignalFn) {
  sig := StringNameFromStr("node_renamed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type SceneTreeNodeConfigurationWarningChangedSignalFn func(node Node, )

func (me *SceneTree) ConnectNodeConfigurationWarningChanged(subs SignalSubscribers, fn SceneTreeNodeConfigurationWarningChangedSignalFn) {
  sig := StringNameFromStr("node_configuration_warning_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneTree) DisconnectNodeConfigurationWarningChanged(subs SignalSubscribers, fn SceneTreeNodeConfigurationWarningChangedSignalFn) {
  sig := StringNameFromStr("node_configuration_warning_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type SceneTreeProcessFrameSignalFn func()

func (me *SceneTree) ConnectProcessFrame(subs SignalSubscribers, fn SceneTreeProcessFrameSignalFn) {
  sig := StringNameFromStr("process_frame")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneTree) DisconnectProcessFrame(subs SignalSubscribers, fn SceneTreeProcessFrameSignalFn) {
  sig := StringNameFromStr("process_frame")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type SceneTreePhysicsFrameSignalFn func()

func (me *SceneTree) ConnectPhysicsFrame(subs SignalSubscribers, fn SceneTreePhysicsFrameSignalFn) {
  sig := StringNameFromStr("physics_frame")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneTree) DisconnectPhysicsFrame(subs SignalSubscribers, fn SceneTreePhysicsFrameSignalFn) {
  sig := StringNameFromStr("physics_frame")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
