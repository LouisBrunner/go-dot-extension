// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MultiplayerSpawner struct {
  Node
}

func (me *MultiplayerSpawner) BaseClass() string {
  return "MultiplayerSpawner"
}



// Enums

func (me *MultiplayerSpawner) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiplayerSpawner) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerSpawner) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MultiplayerSpawner) AddSpawnableScene(path String, )  {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_spawnable_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSpawner) GetSpawnableSceneCount() int {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spawnable_scene_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSpawner) GetSpawnableScene(index int, ) String {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spawnable_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSpawner) ClearSpawnableScenes()  {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_spawnable_scenes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSpawner) Spawn(data Variant, ) Node {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("spawn")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1991184589) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSpawner) GetSpawnPath() NodePath {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spawn_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSpawner) SetSpawnPath(path NodePath, )  {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spawn_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSpawner) GetSpawnLimit() int {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spawn_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSpawner) SetSpawnLimit(limit int, )  {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spawn_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSpawner) GetSpawnFunction() Callable {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spawn_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1307783378) // FIXME: should cache?
  var ret Callable
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSpawner) SetSpawnFunction(spawn_function Callable, )  {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spawn_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(spawn_function.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiplayerSpawnerDespawnedSignalFn func(node Node, )

func (me *MultiplayerSpawner) ConnectDespawned(subs SignalSubscribers, fn MultiplayerSpawnerDespawnedSignalFn) {
  sig := StringNameFromStr("despawned")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MultiplayerSpawner) DisconnectDespawned(subs SignalSubscribers, fn MultiplayerSpawnerDespawnedSignalFn) {
  sig := StringNameFromStr("despawned")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type MultiplayerSpawnerSpawnedSignalFn func(node Node, )

func (me *MultiplayerSpawner) ConnectSpawned(subs SignalSubscribers, fn MultiplayerSpawnerSpawnedSignalFn) {
  sig := StringNameFromStr("spawned")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MultiplayerSpawner) DisconnectSpawned(subs SignalSubscribers, fn MultiplayerSpawnerSpawnedSignalFn) {
  sig := StringNameFromStr("spawned")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
