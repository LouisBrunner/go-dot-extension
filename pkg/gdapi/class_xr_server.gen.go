// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type XRServer struct {
  obj gdc.ObjectPtr
}

func (me *XRServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRServer) BaseClass() string {
  return "XRServer"
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

func  (me *XRServer) GetWorldScale() float32 {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_world_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) SetWorldScale(scale float32, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_world_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRServer) GetWorldOrigin() Transform3D {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_world_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) SetWorldOrigin(world_origin Transform3D, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_world_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2952846383) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(world_origin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRServer) GetReferenceFrame() Transform3D {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reference_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) CenterOnHmd(rotation_mode XRServerRotationMode, keep_height bool, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("center_on_hmd")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1450904707) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation_mode), gdc.ConstTypePtr(&keep_height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRServer) GetHmdTransform() Transform3D {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hmd_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4183770049) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) AddInterface(interface_ XRInterface, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1898711491) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(interface_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRServer) GetInterfaceCount() int {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interface_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) RemoveInterface(interface_ XRInterface, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1898711491) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(interface_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRServer) GetInterface(idx int, ) XRInterface {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4237347919) // FIXME: should cache?
  var ret XRInterface
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) GetInterfaces() Dictionary {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) FindInterface(name String, ) XRInterface {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1395192955) // FIXME: should cache?
  var ret XRInterface
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) AddTracker(tracker XRPositionalTracker, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2692800323) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tracker.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRServer) RemoveTracker(tracker XRPositionalTracker, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2692800323) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tracker.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRServer) GetTrackers(tracker_types int, ) Dictionary {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_trackers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3554694381) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tracker_types), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) GetTracker(tracker_name StringName, ) XRPositionalTracker {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2742084544) // FIXME: should cache?
  var ret XRPositionalTracker
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tracker_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) GetPrimaryInterface() XRInterface {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_primary_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2143545064) // FIXME: should cache?
  var ret XRInterface
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRServer) SetPrimaryInterface(interface_ XRInterface, )  {
  classNameV := StringNameFromStr("XRServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_primary_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1898711491) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(interface_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *XRServer) GetPropWorldScale() float32 {
  panic("TODO: implement")
}

func (me *XRServer) SetPropWorldScale(value float32) {
  panic("TODO: implement")
}

func (me *XRServer) GetPropWorldOrigin() Vector3 {
  panic("TODO: implement")
}

func (me *XRServer) SetPropWorldOrigin(value Vector3) {
  panic("TODO: implement")
}

func (me *XRServer) GetPropPrimaryInterface() Object {
  panic("TODO: implement")
}

func (me *XRServer) SetPropPrimaryInterface(value Object) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API