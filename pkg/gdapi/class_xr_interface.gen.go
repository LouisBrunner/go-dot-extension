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

type XRInterface struct {
  RefCounted
}

func (me *XRInterface) BaseClass() string {
  return "XRInterface"
}

func NewXRInterface() *XRInterface {
  str := StringNameFromStr("XRInterface") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &XRInterface{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type XRInterfaceCapabilities int
const (
  XRInterfaceCapabilitiesXrNone XRInterfaceCapabilities = 0
  XRInterfaceCapabilitiesXrMono XRInterfaceCapabilities = 1
  XRInterfaceCapabilitiesXrStereo XRInterfaceCapabilities = 2
  XRInterfaceCapabilitiesXrQuad XRInterfaceCapabilities = 4
  XRInterfaceCapabilitiesXrVr XRInterfaceCapabilities = 8
  XRInterfaceCapabilitiesXrAr XRInterfaceCapabilities = 16
  XRInterfaceCapabilitiesXrExternal XRInterfaceCapabilities = 32
)

type XRInterfaceTrackingStatus int
const (
  XRInterfaceTrackingStatusXrNormalTracking XRInterfaceTrackingStatus = 0
  XRInterfaceTrackingStatusXrExcessiveMotion XRInterfaceTrackingStatus = 1
  XRInterfaceTrackingStatusXrInsufficientFeatures XRInterfaceTrackingStatus = 2
  XRInterfaceTrackingStatusXrUnknownTracking XRInterfaceTrackingStatus = 3
  XRInterfaceTrackingStatusXrNotTracking XRInterfaceTrackingStatus = 4
)

type XRInterfacePlayAreaMode int
const (
  XRInterfacePlayAreaModeXrPlayAreaUnknown XRInterfacePlayAreaMode = 0
  XRInterfacePlayAreaModeXrPlayArea3Dof XRInterfacePlayAreaMode = 1
  XRInterfacePlayAreaModeXrPlayAreaSitting XRInterfacePlayAreaMode = 2
  XRInterfacePlayAreaModeXrPlayAreaRoomscale XRInterfacePlayAreaMode = 3
  XRInterfacePlayAreaModeXrPlayAreaStage XRInterfacePlayAreaMode = 4
)

type XRInterfaceEnvironmentBlendMode int
const (
  XRInterfaceEnvironmentBlendModeXrEnvBlendModeOpaque XRInterfaceEnvironmentBlendMode = 0
  XRInterfaceEnvironmentBlendModeXrEnvBlendModeAdditive XRInterfaceEnvironmentBlendMode = 1
  XRInterfaceEnvironmentBlendModeXrEnvBlendModeAlphaBlend XRInterfaceEnvironmentBlendMode = 2
)

func (me *XRInterface) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRInterface) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRInterface) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *XRInterface) GetName() StringName {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRInterface) GetCapabilities() int64 {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_capabilities")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) IsPrimary() bool {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_primary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) SetPrimary(primary bool, )  {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_primary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primary) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRInterface) IsInitialized() bool {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_initialized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) Initialize() bool {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("initialize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) Uninitialize()  {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("uninitialize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRInterface) GetSystemInfo() Dictionary {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_system_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2382534195) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRInterface) GetTrackingStatus() XRInterfaceTrackingStatus {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracking_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 167423259) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret XRInterfaceTrackingStatus

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *XRInterface) GetRenderTargetSize() Vector2 {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_render_target_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1497962370) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRInterface) GetViewCount() int64 {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_view_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) TriggerHapticPulse(action_name String, tracker_name StringName, frequency float64, amplitude float64, duration_sec float64, delay_sec float64, )  {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("trigger_haptic_pulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3752640163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action_name.AsCTypePtr(), tracker_name.AsCTypePtr(), gdc.ConstTypePtr(&frequency) , gdc.ConstTypePtr(&amplitude) , gdc.ConstTypePtr(&duration_sec) , gdc.ConstTypePtr(&delay_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRInterface) SupportsPlayAreaMode(mode XRInterfacePlayAreaMode, ) bool {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("supports_play_area_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3429955281) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&mode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) GetPlayAreaMode() XRInterfacePlayAreaMode {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_play_area_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1615132885) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret XRInterfacePlayAreaMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *XRInterface) SetPlayAreaMode(mode XRInterfacePlayAreaMode, ) bool {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_play_area_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3429955281) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&mode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) GetPlayArea() PackedVector3Array {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_play_area")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRInterface) GetAnchorDetectionIsEnabled() bool {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_anchor_detection_is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) SetAnchorDetectionIsEnabled(enable bool, )  {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_anchor_detection_is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRInterface) GetCameraFeedId() int64 {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_feed_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) IsPassthroughSupported() bool {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_passthrough_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) IsPassthroughEnabled() bool {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_passthrough_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) StartPassthrough() bool {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start_passthrough")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) StopPassthrough()  {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop_passthrough")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRInterface) GetTransformForView(view int64, cam_transform Transform3D, ) Transform3D {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform_for_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 518934792) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&view) , cam_transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&view)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRInterface) GetProjectionForView(view int64, aspect float64, near float64, far float64, ) Projection {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_projection_for_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3766090294) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&view) , gdc.ConstTypePtr(&aspect) , gdc.ConstTypePtr(&near) , gdc.ConstTypePtr(&far) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewProjection()
  pinner.Pin(&view)
  pinner.Pin(&aspect)
  pinner.Pin(&near)
  pinner.Pin(&far)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRInterface) GetSupportedEnvironmentBlendModes() Array {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_supported_environment_blend_modes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRInterface) SetEnvironmentBlendMode(mode XRInterfaceEnvironmentBlendMode, ) bool {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_environment_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 551152418) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&mode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRInterface) GetEnvironmentBlendMode() XRInterfaceEnvironmentBlendMode {
  classNameV := StringNameFromStr("XRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1984334071) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret XRInterfaceEnvironmentBlendMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type XRInterfacePlayAreaChangedSignalFn func(mode int, )

func (me *XRInterface) ConnectPlayAreaChanged(subs SignalSubscribers, fn XRInterfacePlayAreaChangedSignalFn) {
  sig := StringNameFromStr("play_area_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRInterface) DisconnectPlayAreaChanged(subs SignalSubscribers, fn XRInterfacePlayAreaChangedSignalFn) {
  sig := StringNameFromStr("play_area_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
