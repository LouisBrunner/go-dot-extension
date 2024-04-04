// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type XRServer struct {
  Object
}

func (me *XRServer) BaseClass() string {
  return "XRServer"
}

func NewXRServer() *XRServer {
  str := StringNameFromStr("XRServer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &XRServer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type XRServerTrackerType int
const (
  XRServerTrackerTypeTrackerHead XRServerTrackerType = 1
  XRServerTrackerTypeTrackerController XRServerTrackerType = 2
  XRServerTrackerTypeTrackerBasestation XRServerTrackerType = 4
  XRServerTrackerTypeTrackerAnchor XRServerTrackerType = 8
  XRServerTrackerTypeTrackerAnyKnown XRServerTrackerType = 127
  XRServerTrackerTypeTrackerUnknown XRServerTrackerType = 128
  XRServerTrackerTypeTrackerAny XRServerTrackerType = 255
)

type XRServerRotationMode int
const (
  XRServerRotationModeResetFullRotation XRServerRotationMode = 0
  XRServerRotationModeResetButKeepTilt XRServerRotationMode = 1
  XRServerRotationModeDontResetRotation XRServerRotationMode = 2
)

func (me *XRServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *XRServer) GetWorldScale() float64 {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_world_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRServer) SetWorldScale(scale float64, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_world_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRServer) GetWorldOrigin() Transform3D {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_world_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRServer) SetWorldOrigin(world_origin Transform3D, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_world_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2952846383) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{world_origin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRServer) GetReferenceFrame() Transform3D {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reference_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRServer) CenterOnHmd(rotation_mode XRServerRotationMode, keep_height bool, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("center_on_hmd")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1450904707) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation_mode) , gdc.ConstTypePtr(&keep_height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRServer) GetHmdTransform() Transform3D {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hmd_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4183770049) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRServer) AddInterface(interface_ XRInterface, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1898711491) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{interface_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRServer) GetInterfaceCount() int64 {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interface_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRServer) RemoveInterface(interface_ XRInterface, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1898711491) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{interface_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRServer) GetInterface(idx int64, ) XRInterface {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4237347919) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewXRInterface()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRServer) GetInterfaces() []Dictionary {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Dictionary](ret)
}

func  (me *XRServer) FindInterface(name String, ) XRInterface {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1395192955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewXRInterface()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRServer) AddTracker(tracker XRPositionalTracker, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2692800323) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{tracker.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRServer) RemoveTracker(tracker XRPositionalTracker, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2692800323) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{tracker.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRServer) GetTrackers(tracker_types int64, ) Dictionary {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_trackers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3554694381) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tracker_types) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&tracker_types)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRServer) GetTracker(tracker_name StringName, ) XRPositionalTracker {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2742084544) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{tracker_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewXRPositionalTracker()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRServer) GetPrimaryInterface() XRInterface {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_primary_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2143545064) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewXRInterface()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRServer) SetPrimaryInterface(interface_ XRInterface, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_primary_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1898711491) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{interface_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type XRServerInterfaceAddedSignalFn func(interface_name StringName, )

func (me *XRServer) ConnectInterfaceAdded(subs SignalSubscribers, fn XRServerInterfaceAddedSignalFn) {
  sig := StringNameFromStr("interface_added")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRServer) DisconnectInterfaceAdded(subs SignalSubscribers, fn XRServerInterfaceAddedSignalFn) {
  sig := StringNameFromStr("interface_added")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type XRServerInterfaceRemovedSignalFn func(interface_name StringName, )

func (me *XRServer) ConnectInterfaceRemoved(subs SignalSubscribers, fn XRServerInterfaceRemovedSignalFn) {
  sig := StringNameFromStr("interface_removed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRServer) DisconnectInterfaceRemoved(subs SignalSubscribers, fn XRServerInterfaceRemovedSignalFn) {
  sig := StringNameFromStr("interface_removed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type XRServerTrackerAddedSignalFn func(tracker_name StringName, type_ int, )

func (me *XRServer) ConnectTrackerAdded(subs SignalSubscribers, fn XRServerTrackerAddedSignalFn) {
  sig := StringNameFromStr("tracker_added")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRServer) DisconnectTrackerAdded(subs SignalSubscribers, fn XRServerTrackerAddedSignalFn) {
  sig := StringNameFromStr("tracker_added")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type XRServerTrackerUpdatedSignalFn func(tracker_name StringName, type_ int, )

func (me *XRServer) ConnectTrackerUpdated(subs SignalSubscribers, fn XRServerTrackerUpdatedSignalFn) {
  sig := StringNameFromStr("tracker_updated")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRServer) DisconnectTrackerUpdated(subs SignalSubscribers, fn XRServerTrackerUpdatedSignalFn) {
  sig := StringNameFromStr("tracker_updated")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type XRServerTrackerRemovedSignalFn func(tracker_name StringName, type_ int, )

func (me *XRServer) ConnectTrackerRemoved(subs SignalSubscribers, fn XRServerTrackerRemovedSignalFn) {
  sig := StringNameFromStr("tracker_removed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRServer) DisconnectTrackerRemoved(subs SignalSubscribers, fn XRServerTrackerRemovedSignalFn) {
  sig := StringNameFromStr("tracker_removed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
