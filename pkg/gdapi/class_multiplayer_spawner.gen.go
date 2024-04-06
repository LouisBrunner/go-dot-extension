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

type ptrsForMultiplayerSpawnerList struct {
	fnAddSpawnableScene      gdc.MethodBindPtr
	fnGetSpawnableSceneCount gdc.MethodBindPtr
	fnGetSpawnableScene      gdc.MethodBindPtr
	fnClearSpawnableScenes   gdc.MethodBindPtr
	fnSpawn                  gdc.MethodBindPtr
	fnGetSpawnPath           gdc.MethodBindPtr
	fnSetSpawnPath           gdc.MethodBindPtr
	fnGetSpawnLimit          gdc.MethodBindPtr
	fnSetSpawnLimit          gdc.MethodBindPtr
	fnGetSpawnFunction       gdc.MethodBindPtr
	fnSetSpawnFunction       gdc.MethodBindPtr
}

var ptrsForMultiplayerSpawner ptrsForMultiplayerSpawnerList

func initMultiplayerSpawnerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MultiplayerSpawner")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_spawnable_scene")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnAddSpawnableScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_spawnable_scene_count")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnGetSpawnableSceneCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_spawnable_scene")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnGetSpawnableScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("clear_spawnable_scenes")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnClearSpawnableScenes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("spawn")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnSpawn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1991184589))
	}
	{
		methodName := StringNameFromStr("get_spawn_path")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnGetSpawnPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_spawn_path")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnSetSpawnPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_spawn_limit")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnGetSpawnLimit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_spawn_limit")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnSetSpawnLimit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_spawn_function")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnGetSpawnFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1307783378))
	}
	{
		methodName := StringNameFromStr("set_spawn_function")
		defer methodName.Destroy()
		ptrsForMultiplayerSpawner.fnSetSpawnFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
	}
}

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

func (me *MultiplayerSpawner) AddSpawnableScene(path String) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnAddSpawnableScene), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiplayerSpawner) GetSpawnableSceneCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnGetSpawnableSceneCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MultiplayerSpawner) GetSpawnableScene(index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnGetSpawnableScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiplayerSpawner) ClearSpawnableScenes() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnClearSpawnableScenes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiplayerSpawner) Spawn(data Variant) Node {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnSpawn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiplayerSpawner) GetSpawnPath() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnGetSpawnPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiplayerSpawner) SetSpawnPath(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnSetSpawnPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiplayerSpawner) GetSpawnLimit() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnGetSpawnLimit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MultiplayerSpawner) SetSpawnLimit(limit int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnSetSpawnLimit), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiplayerSpawner) GetSpawnFunction() Callable {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCallable()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnGetSpawnFunction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiplayerSpawner) SetSpawnFunction(spawn_function Callable) {
	cargs := []gdc.ConstTypePtr{spawn_function.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSpawner.fnSetSpawnFunction), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiplayerSpawnerDespawnedSignalFn func(node Node)

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

type MultiplayerSpawnerSpawnedSignalFn func(node Node)

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
