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

type MultiplayerSpawner struct {
  Node
}

func (me *MultiplayerSpawner) BaseClass() string {
  return "MultiplayerSpawner"
}

func NewMultiplayerSpawner() *MultiplayerSpawner {
  str := StringNameFromStr("MultiplayerSpawner") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MultiplayerSpawner{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSpawner) GetSpawnableSceneCount() int64 {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spawnable_scene_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerSpawner) GetSpawnableScene(index int64, ) String {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spawnable_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiplayerSpawner) ClearSpawnableScenes()  {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_spawnable_scenes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSpawner) Spawn(data Variant, ) Node {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("spawn")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1991184589) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiplayerSpawner) GetSpawnPath() NodePath {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spawn_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiplayerSpawner) SetSpawnPath(path NodePath, )  {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spawn_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSpawner) GetSpawnLimit() int64 {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spawn_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerSpawner) SetSpawnLimit(limit int64, )  {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spawn_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSpawner) GetSpawnFunction() Callable {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spawn_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1307783378) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCallable()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiplayerSpawner) SetSpawnFunction(spawn_function Callable, )  {
  classNameV := StringNameFromStr("MultiplayerSpawner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spawn_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{spawn_function.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiplayerSpawnerDespawnedSignalFn func(node Node, )

func (me *MultiplayerSpawner) ConnectDespawned(subs SignalSubscribers, fn MultiplayerSpawnerDespawnedSignalFn) {
  sig := StringNameFromStr("despawned")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerSpawner) DisconnectDespawned(subs SignalSubscribers, fn MultiplayerSpawnerDespawnedSignalFn) {
  sig := StringNameFromStr("despawned")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type MultiplayerSpawnerSpawnedSignalFn func(node Node, )

func (me *MultiplayerSpawner) ConnectSpawned(subs SignalSubscribers, fn MultiplayerSpawnerSpawnedSignalFn) {
  sig := StringNameFromStr("spawned")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerSpawner) DisconnectSpawned(subs SignalSubscribers, fn MultiplayerSpawnerSpawnedSignalFn) {
  sig := StringNameFromStr("spawned")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
