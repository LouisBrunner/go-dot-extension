// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MultiplayerSynchronizer struct {
  obj gdc.ObjectPtr
}

func (me *MultiplayerSynchronizer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiplayerSynchronizer) BaseClass() string {
  return "MultiplayerSynchronizer"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSynchronizer) GetRootPath() NodePath {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSynchronizer) SetReplicationInterval(milliseconds float32, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_replication_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&milliseconds), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSynchronizer) GetReplicationInterval() float32 {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_replication_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSynchronizer) SetDeltaInterval(milliseconds float32, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_delta_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&milliseconds), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSynchronizer) GetDeltaInterval() float32 {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_delta_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSynchronizer) SetReplicationConfig(config SceneReplicationConfig, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_replication_config")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3889206742) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(config.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSynchronizer) GetReplicationConfig() SceneReplicationConfig {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_replication_config")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200254614) // FIXME: should cache?
  var ret SceneReplicationConfig
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSynchronizer) SetVisibilityUpdateMode(mode MultiplayerSynchronizerVisibilityUpdateMode, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3494860300) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSynchronizer) GetVisibilityUpdateMode() MultiplayerSynchronizerVisibilityUpdateMode {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3352241418) // FIXME: should cache?
  var ret MultiplayerSynchronizerVisibilityUpdateMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSynchronizer) UpdateVisibility(for_peer int, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_visibility")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&for_peer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSynchronizer) SetVisibilityPublic(visible bool, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_public")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSynchronizer) IsVisibilityPublic() bool {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visibility_public")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerSynchronizer) AddVisibilityFilter(filter Callable, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_visibility_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(filter.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSynchronizer) RemoveVisibilityFilter(filter Callable, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_visibility_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(filter.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSynchronizer) SetVisibilityFor(peer int, visible bool, )  {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_for")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer), gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerSynchronizer) GetVisibilityFor(peer int, ) bool {
  classNameV := StringNameFromStr("MultiplayerSynchronizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_for")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiplayerSynchronizerSynchronizedSignalFn func()

func (me *MultiplayerSynchronizer) ConnectSynchronized(subs SignalSubscribers, fn MultiplayerSynchronizerSynchronizedSignalFn) {
  sig := StringNameFromStr("synchronized")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MultiplayerSynchronizer) DisconnectSynchronized(subs SignalSubscribers, fn MultiplayerSynchronizerSynchronizedSignalFn) {
  sig := StringNameFromStr("synchronized")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type MultiplayerSynchronizerDeltaSynchronizedSignalFn func()

func (me *MultiplayerSynchronizer) ConnectDeltaSynchronized(subs SignalSubscribers, fn MultiplayerSynchronizerDeltaSynchronizedSignalFn) {
  sig := StringNameFromStr("delta_synchronized")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MultiplayerSynchronizer) DisconnectDeltaSynchronized(subs SignalSubscribers, fn MultiplayerSynchronizerDeltaSynchronizedSignalFn) {
  sig := StringNameFromStr("delta_synchronized")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type MultiplayerSynchronizerVisibilityChangedSignalFn func(for_peer int, )

func (me *MultiplayerSynchronizer) ConnectVisibilityChanged(subs SignalSubscribers, fn MultiplayerSynchronizerVisibilityChangedSignalFn) {
  sig := StringNameFromStr("visibility_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MultiplayerSynchronizer) DisconnectVisibilityChanged(subs SignalSubscribers, fn MultiplayerSynchronizerVisibilityChangedSignalFn) {
  sig := StringNameFromStr("visibility_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
