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

type ptrsForMultiplayerSynchronizerList struct {
  fnSetRootPath gdc.MethodBindPtr
  fnGetRootPath gdc.MethodBindPtr
  fnSetReplicationInterval gdc.MethodBindPtr
  fnGetReplicationInterval gdc.MethodBindPtr
  fnSetDeltaInterval gdc.MethodBindPtr
  fnGetDeltaInterval gdc.MethodBindPtr
  fnSetReplicationConfig gdc.MethodBindPtr
  fnGetReplicationConfig gdc.MethodBindPtr
  fnSetVisibilityUpdateMode gdc.MethodBindPtr
  fnGetVisibilityUpdateMode gdc.MethodBindPtr
  fnUpdateVisibility gdc.MethodBindPtr
  fnSetVisibilityPublic gdc.MethodBindPtr
  fnIsVisibilityPublic gdc.MethodBindPtr
  fnAddVisibilityFilter gdc.MethodBindPtr
  fnRemoveVisibilityFilter gdc.MethodBindPtr
  fnSetVisibilityFor gdc.MethodBindPtr
  fnGetVisibilityFor gdc.MethodBindPtr
}

var ptrsForMultiplayerSynchronizer ptrsForMultiplayerSynchronizerList

func initMultiplayerSynchronizerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("MultiplayerSynchronizer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_root_path")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnSetRootPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_root_path")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnGetRootPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("set_replication_interval")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnSetReplicationInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_replication_interval")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnGetReplicationInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_delta_interval")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnSetDeltaInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_delta_interval")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnGetDeltaInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_replication_config")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnSetReplicationConfig = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3889206742))
  }
  {
    methodName := StringNameFromStr("get_replication_config")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnGetReplicationConfig = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3200254614))
  }
  {
    methodName := StringNameFromStr("set_visibility_update_mode")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnSetVisibilityUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3494860300))
  }
  {
    methodName := StringNameFromStr("get_visibility_update_mode")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnGetVisibilityUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3352241418))
  }
  {
    methodName := StringNameFromStr("update_visibility")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnUpdateVisibility = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1995695955))
  }
  {
    methodName := StringNameFromStr("set_visibility_public")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnSetVisibilityPublic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_visibility_public")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnIsVisibilityPublic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("add_visibility_filter")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnAddVisibilityFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
  }
  {
    methodName := StringNameFromStr("remove_visibility_filter")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnRemoveVisibilityFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
  }
  {
    methodName := StringNameFromStr("set_visibility_for")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnSetVisibilityFor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_visibility_for")
    defer methodName.Destroy()
    ptrsForMultiplayerSynchronizer.fnGetVisibilityFor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
}

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
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnSetRootPath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetRootPath() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnGetRootPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiplayerSynchronizer) SetReplicationInterval(milliseconds float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&milliseconds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnSetReplicationInterval), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetReplicationInterval() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnGetReplicationInterval), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerSynchronizer) SetDeltaInterval(milliseconds float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&milliseconds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnSetDeltaInterval), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetDeltaInterval() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnGetDeltaInterval), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerSynchronizer) SetReplicationConfig(config SceneReplicationConfig, )  {
  cargs := []gdc.ConstTypePtr{config.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnSetReplicationConfig), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetReplicationConfig() SceneReplicationConfig {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSceneReplicationConfig()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnGetReplicationConfig), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiplayerSynchronizer) SetVisibilityUpdateMode(mode MultiplayerSynchronizerVisibilityUpdateMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnSetVisibilityUpdateMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetVisibilityUpdateMode() MultiplayerSynchronizerVisibilityUpdateMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MultiplayerSynchronizerVisibilityUpdateMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnGetVisibilityUpdateMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MultiplayerSynchronizer) UpdateVisibility(for_peer int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&for_peer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnUpdateVisibility), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) SetVisibilityPublic(visible bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnSetVisibilityPublic), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) IsVisibilityPublic() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnIsVisibilityPublic), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerSynchronizer) AddVisibilityFilter(filter Callable, )  {
  cargs := []gdc.ConstTypePtr{filter.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnAddVisibilityFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) RemoveVisibilityFilter(filter Callable, )  {
  cargs := []gdc.ConstTypePtr{filter.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnRemoveVisibilityFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) SetVisibilityFor(peer int64, visible bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer) , gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnSetVisibilityFor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerSynchronizer) GetVisibilityFor(peer int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&peer)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerSynchronizer.fnGetVisibilityFor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
