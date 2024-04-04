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
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1757182445) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWindow()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) HasGroup(name StringName, ) bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) IsAutoAcceptQuit() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_auto_accept_quit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) SetAutoAcceptQuit(enabled bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_accept_quit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsQuitOnGoBack() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_quit_on_go_back")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) SetQuitOnGoBack(enabled bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_quit_on_go_back")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) SetDebugCollisionsHint(enable bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_collisions_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsDebuggingCollisionsHint() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_debugging_collisions_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) SetDebugPathsHint(enable bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_paths_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsDebuggingPathsHint() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_debugging_paths_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) SetDebugNavigationHint(enable bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_navigation_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsDebuggingNavigationHint() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_debugging_navigation_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) SetEditedSceneRoot(scene Node, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edited_scene_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scene.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) GetEditedSceneRoot() Node {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited_scene_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) SetPause(enable bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pause")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsPaused() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) CreateTimer(time_sec float64, process_always bool, process_in_physics bool, ignore_time_scale bool, ) SceneTreeTimer {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_timer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2709170273) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec) , gdc.ConstTypePtr(&process_always) , gdc.ConstTypePtr(&process_in_physics) , gdc.ConstTypePtr(&ignore_time_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSceneTreeTimer()
  pinner.Pin(&time_sec)
  pinner.Pin(&process_always)
  pinner.Pin(&process_in_physics)
  pinner.Pin(&ignore_time_scale)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) CreateTween() Tween {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_tween")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3426978995) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) GetProcessedTweens() []Tween {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_processed_tweens")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Tween](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *SceneTree) GetNodeCount() int64 {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) GetFrame() int64 {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneTree) Quit(exit_code int64, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("quit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exit_code) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) QueueDelete(obj Object, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue_delete")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3975164845) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obj.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) CallGroupFlags(flags int64, group StringName, method StringName, varargs ...Variant)  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("call_group_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1527739229) // FIXME: should cache?
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
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return
  }

}

func  (me *SceneTree) NotifyGroupFlags(call_flags int64, group StringName, notification int64, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("notify_group_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1245489420) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&call_flags) , group.AsCTypePtr(), gdc.ConstTypePtr(&notification) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) SetGroupFlags(call_flags int64, group StringName, property String, value Variant, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_group_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3497599527) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&call_flags) , group.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) CallGroup(group StringName, method StringName, varargs ...Variant)  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("call_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1257962832) // FIXME: should cache?
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
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return
  }

}

func  (me *SceneTree) NotifyGroup(group StringName, notification int64, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("notify_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2415702435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{group.AsCTypePtr(), gdc.ConstTypePtr(&notification) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) SetGroup(group StringName, property String, value Variant, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1279312029) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{group.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) GetNodesInGroup(group StringName, ) []Node {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_nodes_in_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 689397652) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{group.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Node](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *SceneTree) GetFirstNodeInGroup(group StringName, ) Node {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_first_node_in_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4071044623) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{group.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) SetCurrentScene(child_node Node, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{child_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) GetCurrentScene() Node {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) ChangeSceneToFile(path String, ) Error {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("change_scene_to_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SceneTree) ChangeSceneToPacked(packed_scene PackedScene, ) Error {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("change_scene_to_packed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 107349098) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{packed_scene.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SceneTree) ReloadCurrentScene() Error {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reload_current_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SceneTree) UnloadCurrentScene()  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unload_current_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) SetMultiplayer(multiplayer MultiplayerAPI, root_path NodePath, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multiplayer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2385607013) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{multiplayer.AsCTypePtr(), root_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) GetMultiplayer(for_path NodePath, ) MultiplayerAPI {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_multiplayer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3453401404) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{for_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMultiplayerAPI()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneTree) SetMultiplayerPollEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multiplayer_poll_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTree) IsMultiplayerPollEnabled() bool {
  classNameV := StringNameFromStr("SceneTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_multiplayer_poll_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
