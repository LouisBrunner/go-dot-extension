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

type MultiplayerSynchronizer struct {
  Node
}

func (me *MultiplayerSynchronizer) BaseClass() string {
  return "MultiplayerSynchronizer"
}

func NewMultiplayerSynchronizer() *MultiplayerSynchronizer {
  str := StringNameFromStr("MultiplayerSynchronizer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MultiplayerSynchronizer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type MultiplayerSynchronizerVisibilityUpdateMode int
const (
  MultiplayerSynchronizerVisibilityUpdateModeVisibilityProcessIdle MultiplayerSynchronizerVisibilityUpdateMode = 0
  MultiplayerSynchronizerVisibilityUpdateModeVisibilityProcessPhysics MultiplayerSynchronizerVisibilityUpdateMode = 1
  MultiplayerSynchronizerVisibilityUpdateModeVisibilityProcessNone MultiplayerSynchronizerVisibilityUpdateMode = 2
)

func (me *MultiplayerSynchronizer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiplayerSynchronizer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerSynchronizer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MultiplayerSynchronizer) SetRootPath(path NodePath, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetRootPath() NodePath {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiplayerSynchronizer) SetReplicationInterval(milliseconds float64, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_replication_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&milliseconds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetReplicationInterval() float64 {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_replication_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerSynchronizer) SetDeltaInterval(milliseconds float64, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_delta_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&milliseconds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetDeltaInterval() float64 {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_delta_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerSynchronizer) SetReplicationConfig(config SceneReplicationConfig, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_replication_config")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3889206742) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{config.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetReplicationConfig() SceneReplicationConfig {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_replication_config")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200254614) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSceneReplicationConfig()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiplayerSynchronizer) SetVisibilityUpdateMode(mode MultiplayerSynchronizerVisibilityUpdateMode, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3494860300) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetVisibilityUpdateMode() MultiplayerSynchronizerVisibilityUpdateMode {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3352241418) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MultiplayerSynchronizerVisibilityUpdateMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MultiplayerSynchronizer) UpdateVisibility(for_peer int64, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_visibility")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&for_peer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) SetVisibilityPublic(visible bool, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_public")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) IsVisibilityPublic() bool {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visibility_public")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerSynchronizer) AddVisibilityFilter(filter Callable, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_visibility_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{filter.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) RemoveVisibilityFilter(filter Callable, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_visibility_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{filter.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) SetVisibilityFor(peer int64, visible bool, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_for")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer) , gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetVisibilityFor(peer int64, ) bool {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_for")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&peer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiplayerSynchronizerSynchronizedSignalFn func()

func (me *MultiplayerSynchronizer) ConnectSynchronized(subs SignalSubscribers, fn MultiplayerSynchronizerSynchronizedSignalFn) {
  sig := StringNameFromStr("synchronized")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerSynchronizer) DisconnectSynchronized(subs SignalSubscribers, fn MultiplayerSynchronizerSynchronizedSignalFn) {
  sig := StringNameFromStr("synchronized")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type MultiplayerSynchronizerDeltaSynchronizedSignalFn func()

func (me *MultiplayerSynchronizer) ConnectDeltaSynchronized(subs SignalSubscribers, fn MultiplayerSynchronizerDeltaSynchronizedSignalFn) {
  sig := StringNameFromStr("delta_synchronized")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerSynchronizer) DisconnectDeltaSynchronized(subs SignalSubscribers, fn MultiplayerSynchronizerDeltaSynchronizedSignalFn) {
  sig := StringNameFromStr("delta_synchronized")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type MultiplayerSynchronizerVisibilityChangedSignalFn func(for_peer int, )

func (me *MultiplayerSynchronizer) ConnectVisibilityChanged(subs SignalSubscribers, fn MultiplayerSynchronizerVisibilityChangedSignalFn) {
  sig := StringNameFromStr("visibility_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerSynchronizer) DisconnectVisibilityChanged(subs SignalSubscribers, fn MultiplayerSynchronizerVisibilityChangedSignalFn) {
  sig := StringNameFromStr("visibility_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
