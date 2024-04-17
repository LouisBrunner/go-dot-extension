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

type ptrsForWebXRInterfaceList struct {
	fnIsSessionSupported              gdc.MethodBindPtr
	fnSetSessionMode                  gdc.MethodBindPtr
	fnGetSessionMode                  gdc.MethodBindPtr
	fnSetRequiredFeatures             gdc.MethodBindPtr
	fnGetRequiredFeatures             gdc.MethodBindPtr
	fnSetOptionalFeatures             gdc.MethodBindPtr
	fnGetOptionalFeatures             gdc.MethodBindPtr
	fnGetReferenceSpaceType           gdc.MethodBindPtr
	fnGetEnabledFeatures              gdc.MethodBindPtr
	fnSetRequestedReferenceSpaceTypes gdc.MethodBindPtr
	fnGetRequestedReferenceSpaceTypes gdc.MethodBindPtr
	fnIsInputSourceActive             gdc.MethodBindPtr
	fnGetInputSourceTracker           gdc.MethodBindPtr
	fnGetInputSourceTargetRayMode     gdc.MethodBindPtr
	fnGetVisibilityState              gdc.MethodBindPtr
	fnGetDisplayRefreshRate           gdc.MethodBindPtr
	fnSetDisplayRefreshRate           gdc.MethodBindPtr
	fnGetAvailableDisplayRefreshRates gdc.MethodBindPtr
}

var ptrsForWebXRInterface ptrsForWebXRInterfaceList

func initWebXRInterfacePtrs(iface gdc.Interface) {

	className := StringNameFromStr("WebXRInterface")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("is_session_supported")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnIsSessionSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_session_mode")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnSetSessionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_session_mode")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetSessionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_required_features")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnSetRequiredFeatures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_required_features")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetRequiredFeatures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_optional_features")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnSetOptionalFeatures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_optional_features")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetOptionalFeatures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_reference_space_type")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetReferenceSpaceType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_enabled_features")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetEnabledFeatures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_requested_reference_space_types")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnSetRequestedReferenceSpaceTypes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_requested_reference_space_types")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetRequestedReferenceSpaceTypes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_input_source_active")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnIsInputSourceActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_input_source_tracker")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetInputSourceTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 636011756))
	}
	{
		methodName := StringNameFromStr("get_input_source_target_ray_mode")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetInputSourceTargetRayMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2852387453))
	}
	{
		methodName := StringNameFromStr("get_visibility_state")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetVisibilityState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_display_refresh_rate")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetDisplayRefreshRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_display_refresh_rate")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnSetDisplayRefreshRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_available_display_refresh_rates")
		defer methodName.Destroy()
		ptrsForWebXRInterface.fnGetAvailableDisplayRefreshRates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}

}

type WebXRInterface struct {
	XRInterface
}

func (me *WebXRInterface) BaseClass() string {
	return "WebXRInterface"
}

func NewWebXRInterface() *WebXRInterface {
	str := StringNameFromStr("WebXRInterface") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &WebXRInterface{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type WebXRInterfaceTargetRayMode int

const (
	WebXRInterfaceTargetRayModeTargetRayModeUnknown        WebXRInterfaceTargetRayMode = 0
	WebXRInterfaceTargetRayModeTargetRayModeGaze           WebXRInterfaceTargetRayMode = 1
	WebXRInterfaceTargetRayModeTargetRayModeTrackedPointer WebXRInterfaceTargetRayMode = 2
	WebXRInterfaceTargetRayModeTargetRayModeScreen         WebXRInterfaceTargetRayMode = 3
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

func (me *WebXRInterface) IsSessionSupported(session_mode String) {
	cargs := []gdc.ConstTypePtr{session_mode.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnIsSessionSupported), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WebXRInterface) SetSessionMode(session_mode String) {
	cargs := []gdc.ConstTypePtr{session_mode.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnSetSessionMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WebXRInterface) GetSessionMode() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetSessionMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebXRInterface) SetRequiredFeatures(required_features String) {
	cargs := []gdc.ConstTypePtr{required_features.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnSetRequiredFeatures), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WebXRInterface) GetRequiredFeatures() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetRequiredFeatures), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebXRInterface) SetOptionalFeatures(optional_features String) {
	cargs := []gdc.ConstTypePtr{optional_features.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnSetOptionalFeatures), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WebXRInterface) GetOptionalFeatures() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetOptionalFeatures), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebXRInterface) GetReferenceSpaceType() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetReferenceSpaceType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebXRInterface) GetEnabledFeatures() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetEnabledFeatures), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebXRInterface) SetRequestedReferenceSpaceTypes(requested_reference_space_types String) {
	cargs := []gdc.ConstTypePtr{requested_reference_space_types.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnSetRequestedReferenceSpaceTypes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WebXRInterface) GetRequestedReferenceSpaceTypes() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetRequestedReferenceSpaceTypes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebXRInterface) IsInputSourceActive(input_source_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_source_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&input_source_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnIsInputSourceActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WebXRInterface) GetInputSourceTracker(input_source_id int64) XRPositionalTracker {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_source_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewXRPositionalTracker()
	pinner.Pin(&input_source_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetInputSourceTracker), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebXRInterface) GetInputSourceTargetRayMode(input_source_id int64) WebXRInterfaceTargetRayMode {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_source_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WebXRInterfaceTargetRayMode
	pinner.Pin(&input_source_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetInputSourceTargetRayMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebXRInterface) GetVisibilityState() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetVisibilityState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebXRInterface) GetDisplayRefreshRate() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetDisplayRefreshRate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WebXRInterface) SetDisplayRefreshRate(refresh_rate float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refresh_rate)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnSetDisplayRefreshRate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WebXRInterface) GetAvailableDisplayRefreshRates() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebXRInterface.fnGetAvailableDisplayRefreshRates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type WebXRInterfaceSessionSupportedSignalFn func(session_mode String, supported bool)

func (me *WebXRInterface) ConnectSessionSupported(subs SignalSubscribers, fn WebXRInterfaceSessionSupportedSignalFn) {
	sig := StringNameFromStr("session_supported")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSessionSupported(subs SignalSubscribers, fn WebXRInterfaceSessionSupportedSignalFn) {
	sig := StringNameFromStr("session_supported")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceSessionStartedSignalFn func()

func (me *WebXRInterface) ConnectSessionStarted(subs SignalSubscribers, fn WebXRInterfaceSessionStartedSignalFn) {
	sig := StringNameFromStr("session_started")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSessionStarted(subs SignalSubscribers, fn WebXRInterfaceSessionStartedSignalFn) {
	sig := StringNameFromStr("session_started")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceSessionEndedSignalFn func()

func (me *WebXRInterface) ConnectSessionEnded(subs SignalSubscribers, fn WebXRInterfaceSessionEndedSignalFn) {
	sig := StringNameFromStr("session_ended")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSessionEnded(subs SignalSubscribers, fn WebXRInterfaceSessionEndedSignalFn) {
	sig := StringNameFromStr("session_ended")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceSessionFailedSignalFn func(message String)

func (me *WebXRInterface) ConnectSessionFailed(subs SignalSubscribers, fn WebXRInterfaceSessionFailedSignalFn) {
	sig := StringNameFromStr("session_failed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSessionFailed(subs SignalSubscribers, fn WebXRInterfaceSessionFailedSignalFn) {
	sig := StringNameFromStr("session_failed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceSelectstartSignalFn func(input_source_id int)

func (me *WebXRInterface) ConnectSelectstart(subs SignalSubscribers, fn WebXRInterfaceSelectstartSignalFn) {
	sig := StringNameFromStr("selectstart")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSelectstart(subs SignalSubscribers, fn WebXRInterfaceSelectstartSignalFn) {
	sig := StringNameFromStr("selectstart")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceSelectSignalFn func(input_source_id int)

func (me *WebXRInterface) ConnectSelect(subs SignalSubscribers, fn WebXRInterfaceSelectSignalFn) {
	sig := StringNameFromStr("select")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSelect(subs SignalSubscribers, fn WebXRInterfaceSelectSignalFn) {
	sig := StringNameFromStr("select")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceSelectendSignalFn func(input_source_id int)

func (me *WebXRInterface) ConnectSelectend(subs SignalSubscribers, fn WebXRInterfaceSelectendSignalFn) {
	sig := StringNameFromStr("selectend")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSelectend(subs SignalSubscribers, fn WebXRInterfaceSelectendSignalFn) {
	sig := StringNameFromStr("selectend")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceSqueezestartSignalFn func(input_source_id int)

func (me *WebXRInterface) ConnectSqueezestart(subs SignalSubscribers, fn WebXRInterfaceSqueezestartSignalFn) {
	sig := StringNameFromStr("squeezestart")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSqueezestart(subs SignalSubscribers, fn WebXRInterfaceSqueezestartSignalFn) {
	sig := StringNameFromStr("squeezestart")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceSqueezeSignalFn func(input_source_id int)

func (me *WebXRInterface) ConnectSqueeze(subs SignalSubscribers, fn WebXRInterfaceSqueezeSignalFn) {
	sig := StringNameFromStr("squeeze")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSqueeze(subs SignalSubscribers, fn WebXRInterfaceSqueezeSignalFn) {
	sig := StringNameFromStr("squeeze")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceSqueezeendSignalFn func(input_source_id int)

func (me *WebXRInterface) ConnectSqueezeend(subs SignalSubscribers, fn WebXRInterfaceSqueezeendSignalFn) {
	sig := StringNameFromStr("squeezeend")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectSqueezeend(subs SignalSubscribers, fn WebXRInterfaceSqueezeendSignalFn) {
	sig := StringNameFromStr("squeezeend")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceVisibilityStateChangedSignalFn func()

func (me *WebXRInterface) ConnectVisibilityStateChanged(subs SignalSubscribers, fn WebXRInterfaceVisibilityStateChangedSignalFn) {
	sig := StringNameFromStr("visibility_state_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectVisibilityStateChanged(subs SignalSubscribers, fn WebXRInterfaceVisibilityStateChangedSignalFn) {
	sig := StringNameFromStr("visibility_state_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceReferenceSpaceResetSignalFn func()

func (me *WebXRInterface) ConnectReferenceSpaceReset(subs SignalSubscribers, fn WebXRInterfaceReferenceSpaceResetSignalFn) {
	sig := StringNameFromStr("reference_space_reset")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectReferenceSpaceReset(subs SignalSubscribers, fn WebXRInterfaceReferenceSpaceResetSignalFn) {
	sig := StringNameFromStr("reference_space_reset")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebXRInterfaceDisplayRefreshRateChangedSignalFn func()

func (me *WebXRInterface) ConnectDisplayRefreshRateChanged(subs SignalSubscribers, fn WebXRInterfaceDisplayRefreshRateChangedSignalFn) {
	sig := StringNameFromStr("display_refresh_rate_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebXRInterface) DisconnectDisplayRefreshRateChanged(subs SignalSubscribers, fn WebXRInterfaceDisplayRefreshRateChangedSignalFn) {
	sig := StringNameFromStr("display_refresh_rate_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
