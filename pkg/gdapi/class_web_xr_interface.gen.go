// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type WebXRInterface struct {
  obj gdc.ObjectPtr
}

func (me *WebXRInterface) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebXRInterface) BaseClass() string {
  return "WebXRInterface"
}



// Enums

type WebXRInterfaceTargetRayMode int
const (
  WebXRInterfaceTargetRayModeTargetRayModeUnknown WebXRInterfaceTargetRayMode = 0
  WebXRInterfaceTargetRayModeTargetRayModeGaze WebXRInterfaceTargetRayMode = 1
  WebXRInterfaceTargetRayModeTargetRayModeTrackedPointer WebXRInterfaceTargetRayMode = 2
  WebXRInterfaceTargetRayModeTargetRayModeScreen WebXRInterfaceTargetRayMode = 3
)

func (me *WebXRInterface) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WebXRInterface) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebXRInterface) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WebXRInterface) IsSessionSupported(session_mode String, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_session_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(session_mode.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) SetSessionMode(session_mode String, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_session_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(session_mode.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) GetSessionMode() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_session_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) SetRequiredFeatures(required_features String, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_required_features")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(required_features.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) GetRequiredFeatures() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_required_features")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) SetOptionalFeatures(optional_features String, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_optional_features")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(optional_features.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) GetOptionalFeatures() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_optional_features")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) GetReferenceSpaceType() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reference_space_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) SetRequestedReferenceSpaceTypes(requested_reference_space_types String, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_requested_reference_space_types")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(requested_reference_space_types.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) GetRequestedReferenceSpaceTypes() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_requested_reference_space_types")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) IsInputSourceActive(input_source_id int, ) bool {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_input_source_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_source_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) GetInputSourceTracker(input_source_id int, ) XRPositionalTracker {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_source_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 636011756) // FIXME: should cache?
  var ret XRPositionalTracker
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_source_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) GetInputSourceTargetRayMode(input_source_id int, ) WebXRInterfaceTargetRayMode {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_source_target_ray_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2852387453) // FIXME: should cache?
  var ret WebXRInterfaceTargetRayMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_source_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) GetVisibilityState() String {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) GetDisplayRefreshRate() float32 {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_display_refresh_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebXRInterface) SetDisplayRefreshRate(refresh_rate float32, )  {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_display_refresh_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refresh_rate), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebXRInterface) GetAvailableDisplayRefreshRates() Array {
  classNameV := StringNameFromStr("WebXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_available_display_refresh_rates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type WebXRInterfaceSessionSupportedSignalFn func(session_mode String, supported bool, )

func (me *WebXRInterface) ConnectSessionSupported(subs SignalSubscribers, fn WebXRInterfaceSessionSupportedSignalFn) {
  sig := StringNameFromStr("session_supported")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSessionSupported(subs SignalSubscribers, fn WebXRInterfaceSessionSupportedSignalFn) {
  sig := StringNameFromStr("session_supported")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceSessionStartedSignalFn func()

func (me *WebXRInterface) ConnectSessionStarted(subs SignalSubscribers, fn WebXRInterfaceSessionStartedSignalFn) {
  sig := StringNameFromStr("session_started")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSessionStarted(subs SignalSubscribers, fn WebXRInterfaceSessionStartedSignalFn) {
  sig := StringNameFromStr("session_started")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceSessionEndedSignalFn func()

func (me *WebXRInterface) ConnectSessionEnded(subs SignalSubscribers, fn WebXRInterfaceSessionEndedSignalFn) {
  sig := StringNameFromStr("session_ended")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSessionEnded(subs SignalSubscribers, fn WebXRInterfaceSessionEndedSignalFn) {
  sig := StringNameFromStr("session_ended")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceSessionFailedSignalFn func(message String, )

func (me *WebXRInterface) ConnectSessionFailed(subs SignalSubscribers, fn WebXRInterfaceSessionFailedSignalFn) {
  sig := StringNameFromStr("session_failed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSessionFailed(subs SignalSubscribers, fn WebXRInterfaceSessionFailedSignalFn) {
  sig := StringNameFromStr("session_failed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceSelectstartSignalFn func(input_source_id int, )

func (me *WebXRInterface) ConnectSelectstart(subs SignalSubscribers, fn WebXRInterfaceSelectstartSignalFn) {
  sig := StringNameFromStr("selectstart")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSelectstart(subs SignalSubscribers, fn WebXRInterfaceSelectstartSignalFn) {
  sig := StringNameFromStr("selectstart")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceSelectSignalFn func(input_source_id int, )

func (me *WebXRInterface) ConnectSelect(subs SignalSubscribers, fn WebXRInterfaceSelectSignalFn) {
  sig := StringNameFromStr("select")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSelect(subs SignalSubscribers, fn WebXRInterfaceSelectSignalFn) {
  sig := StringNameFromStr("select")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceSelectendSignalFn func(input_source_id int, )

func (me *WebXRInterface) ConnectSelectend(subs SignalSubscribers, fn WebXRInterfaceSelectendSignalFn) {
  sig := StringNameFromStr("selectend")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSelectend(subs SignalSubscribers, fn WebXRInterfaceSelectendSignalFn) {
  sig := StringNameFromStr("selectend")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceSqueezestartSignalFn func(input_source_id int, )

func (me *WebXRInterface) ConnectSqueezestart(subs SignalSubscribers, fn WebXRInterfaceSqueezestartSignalFn) {
  sig := StringNameFromStr("squeezestart")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSqueezestart(subs SignalSubscribers, fn WebXRInterfaceSqueezestartSignalFn) {
  sig := StringNameFromStr("squeezestart")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceSqueezeSignalFn func(input_source_id int, )

func (me *WebXRInterface) ConnectSqueeze(subs SignalSubscribers, fn WebXRInterfaceSqueezeSignalFn) {
  sig := StringNameFromStr("squeeze")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSqueeze(subs SignalSubscribers, fn WebXRInterfaceSqueezeSignalFn) {
  sig := StringNameFromStr("squeeze")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceSqueezeendSignalFn func(input_source_id int, )

func (me *WebXRInterface) ConnectSqueezeend(subs SignalSubscribers, fn WebXRInterfaceSqueezeendSignalFn) {
  sig := StringNameFromStr("squeezeend")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSqueezeend(subs SignalSubscribers, fn WebXRInterfaceSqueezeendSignalFn) {
  sig := StringNameFromStr("squeezeend")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceVisibilityStateChangedSignalFn func()

func (me *WebXRInterface) ConnectVisibilityStateChanged(subs SignalSubscribers, fn WebXRInterfaceVisibilityStateChangedSignalFn) {
  sig := StringNameFromStr("visibility_state_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectVisibilityStateChanged(subs SignalSubscribers, fn WebXRInterfaceVisibilityStateChangedSignalFn) {
  sig := StringNameFromStr("visibility_state_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceReferenceSpaceResetSignalFn func()

func (me *WebXRInterface) ConnectReferenceSpaceReset(subs SignalSubscribers, fn WebXRInterfaceReferenceSpaceResetSignalFn) {
  sig := StringNameFromStr("reference_space_reset")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectReferenceSpaceReset(subs SignalSubscribers, fn WebXRInterfaceReferenceSpaceResetSignalFn) {
  sig := StringNameFromStr("reference_space_reset")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type WebXRInterfaceDisplayRefreshRateChangedSignalFn func()

func (me *WebXRInterface) ConnectDisplayRefreshRateChanged(subs SignalSubscribers, fn WebXRInterfaceDisplayRefreshRateChangedSignalFn) {
  sig := StringNameFromStr("display_refresh_rate_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectDisplayRefreshRateChanged(subs SignalSubscribers, fn WebXRInterfaceDisplayRefreshRateChangedSignalFn) {
  sig := StringNameFromStr("display_refresh_rate_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
