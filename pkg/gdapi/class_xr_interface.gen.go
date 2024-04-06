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

type ptrsForXRInterfaceList struct {
	fnGetName                           gdc.MethodBindPtr
	fnGetCapabilities                   gdc.MethodBindPtr
	fnIsPrimary                         gdc.MethodBindPtr
	fnSetPrimary                        gdc.MethodBindPtr
	fnIsInitialized                     gdc.MethodBindPtr
	fnInitialize                        gdc.MethodBindPtr
	fnUninitialize                      gdc.MethodBindPtr
	fnGetSystemInfo                     gdc.MethodBindPtr
	fnGetTrackingStatus                 gdc.MethodBindPtr
	fnGetRenderTargetSize               gdc.MethodBindPtr
	fnGetViewCount                      gdc.MethodBindPtr
	fnTriggerHapticPulse                gdc.MethodBindPtr
	fnSupportsPlayAreaMode              gdc.MethodBindPtr
	fnGetPlayAreaMode                   gdc.MethodBindPtr
	fnSetPlayAreaMode                   gdc.MethodBindPtr
	fnGetPlayArea                       gdc.MethodBindPtr
	fnGetAnchorDetectionIsEnabled       gdc.MethodBindPtr
	fnSetAnchorDetectionIsEnabled       gdc.MethodBindPtr
	fnGetCameraFeedId                   gdc.MethodBindPtr
	fnIsPassthroughSupported            gdc.MethodBindPtr
	fnIsPassthroughEnabled              gdc.MethodBindPtr
	fnStartPassthrough                  gdc.MethodBindPtr
	fnStopPassthrough                   gdc.MethodBindPtr
	fnGetTransformForView               gdc.MethodBindPtr
	fnGetProjectionForView              gdc.MethodBindPtr
	fnGetSupportedEnvironmentBlendModes gdc.MethodBindPtr
	fnSetEnvironmentBlendMode           gdc.MethodBindPtr
	fnGetEnvironmentBlendMode           gdc.MethodBindPtr
}

var ptrsForXRInterface ptrsForXRInterfaceList

func initXRInterfacePtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRInterface")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_name")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("get_capabilities")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetCapabilities = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("is_primary")
		defer methodName.Destroy()
		ptrsForXRInterface.fnIsPrimary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_primary")
		defer methodName.Destroy()
		ptrsForXRInterface.fnSetPrimary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_initialized")
		defer methodName.Destroy()
		ptrsForXRInterface.fnIsInitialized = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("initialize")
		defer methodName.Destroy()
		ptrsForXRInterface.fnInitialize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("uninitialize")
		defer methodName.Destroy()
		ptrsForXRInterface.fnUninitialize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_system_info")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetSystemInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2382534195))
	}
	{
		methodName := StringNameFromStr("get_tracking_status")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetTrackingStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 167423259))
	}
	{
		methodName := StringNameFromStr("get_render_target_size")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetRenderTargetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1497962370))
	}
	{
		methodName := StringNameFromStr("get_view_count")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetViewCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("trigger_haptic_pulse")
		defer methodName.Destroy()
		ptrsForXRInterface.fnTriggerHapticPulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3752640163))
	}
	{
		methodName := StringNameFromStr("supports_play_area_mode")
		defer methodName.Destroy()
		ptrsForXRInterface.fnSupportsPlayAreaMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3429955281))
	}
	{
		methodName := StringNameFromStr("get_play_area_mode")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetPlayAreaMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1615132885))
	}
	{
		methodName := StringNameFromStr("set_play_area_mode")
		defer methodName.Destroy()
		ptrsForXRInterface.fnSetPlayAreaMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3429955281))
	}
	{
		methodName := StringNameFromStr("get_play_area")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetPlayArea = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
	}
	{
		methodName := StringNameFromStr("get_anchor_detection_is_enabled")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetAnchorDetectionIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_anchor_detection_is_enabled")
		defer methodName.Destroy()
		ptrsForXRInterface.fnSetAnchorDetectionIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_camera_feed_id")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetCameraFeedId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("is_passthrough_supported")
		defer methodName.Destroy()
		ptrsForXRInterface.fnIsPassthroughSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("is_passthrough_enabled")
		defer methodName.Destroy()
		ptrsForXRInterface.fnIsPassthroughEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("start_passthrough")
		defer methodName.Destroy()
		ptrsForXRInterface.fnStartPassthrough = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("stop_passthrough")
		defer methodName.Destroy()
		ptrsForXRInterface.fnStopPassthrough = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_transform_for_view")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetTransformForView = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 518934792))
	}
	{
		methodName := StringNameFromStr("get_projection_for_view")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetProjectionForView = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3766090294))
	}
	{
		methodName := StringNameFromStr("get_supported_environment_blend_modes")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetSupportedEnvironmentBlendModes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("set_environment_blend_mode")
		defer methodName.Destroy()
		ptrsForXRInterface.fnSetEnvironmentBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 551152418))
	}
	{
		methodName := StringNameFromStr("get_environment_blend_mode")
		defer methodName.Destroy()
		ptrsForXRInterface.fnGetEnvironmentBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1984334071))
	}

}

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
	XRInterfaceCapabilitiesXrNone     XRInterfaceCapabilities = 0
	XRInterfaceCapabilitiesXrMono     XRInterfaceCapabilities = 1
	XRInterfaceCapabilitiesXrStereo   XRInterfaceCapabilities = 2
	XRInterfaceCapabilitiesXrQuad     XRInterfaceCapabilities = 4
	XRInterfaceCapabilitiesXrVr       XRInterfaceCapabilities = 8
	XRInterfaceCapabilitiesXrAr       XRInterfaceCapabilities = 16
	XRInterfaceCapabilitiesXrExternal XRInterfaceCapabilities = 32
)

type XRInterfaceTrackingStatus int

const (
	XRInterfaceTrackingStatusXrNormalTracking       XRInterfaceTrackingStatus = 0
	XRInterfaceTrackingStatusXrExcessiveMotion      XRInterfaceTrackingStatus = 1
	XRInterfaceTrackingStatusXrInsufficientFeatures XRInterfaceTrackingStatus = 2
	XRInterfaceTrackingStatusXrUnknownTracking      XRInterfaceTrackingStatus = 3
	XRInterfaceTrackingStatusXrNotTracking          XRInterfaceTrackingStatus = 4
)

type XRInterfacePlayAreaMode int

const (
	XRInterfacePlayAreaModeXrPlayAreaUnknown   XRInterfacePlayAreaMode = 0
	XRInterfacePlayAreaModeXrPlayArea3Dof      XRInterfacePlayAreaMode = 1
	XRInterfacePlayAreaModeXrPlayAreaSitting   XRInterfacePlayAreaMode = 2
	XRInterfacePlayAreaModeXrPlayAreaRoomscale XRInterfacePlayAreaMode = 3
	XRInterfacePlayAreaModeXrPlayAreaStage     XRInterfacePlayAreaMode = 4
)

type XRInterfaceEnvironmentBlendMode int

const (
	XRInterfaceEnvironmentBlendModeXrEnvBlendModeOpaque     XRInterfaceEnvironmentBlendMode = 0
	XRInterfaceEnvironmentBlendModeXrEnvBlendModeAdditive   XRInterfaceEnvironmentBlendMode = 1
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

func (me *XRInterface) GetName() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRInterface) GetCapabilities() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetCapabilities), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) IsPrimary() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnIsPrimary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) SetPrimary(primary bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primary)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnSetPrimary), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRInterface) IsInitialized() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnIsInitialized), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) Initialize() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnInitialize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) Uninitialize() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnUninitialize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRInterface) GetSystemInfo() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetSystemInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRInterface) GetTrackingStatus() XRInterfaceTrackingStatus {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRInterfaceTrackingStatus

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetTrackingStatus), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRInterface) GetRenderTargetSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetRenderTargetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRInterface) GetViewCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetViewCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) TriggerHapticPulse(action_name String, tracker_name StringName, frequency float64, amplitude float64, duration_sec float64, delay_sec float64) {
	cargs := []gdc.ConstTypePtr{action_name.AsCTypePtr(), tracker_name.AsCTypePtr(), gdc.ConstTypePtr(&frequency), gdc.ConstTypePtr(&amplitude), gdc.ConstTypePtr(&duration_sec), gdc.ConstTypePtr(&delay_sec)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnTriggerHapticPulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRInterface) SupportsPlayAreaMode(mode XRInterfacePlayAreaMode) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&mode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnSupportsPlayAreaMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) GetPlayAreaMode() XRInterfacePlayAreaMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRInterfacePlayAreaMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetPlayAreaMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRInterface) SetPlayAreaMode(mode XRInterfacePlayAreaMode) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&mode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnSetPlayAreaMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) GetPlayArea() PackedVector3Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector3Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetPlayArea), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRInterface) GetAnchorDetectionIsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetAnchorDetectionIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) SetAnchorDetectionIsEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnSetAnchorDetectionIsEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRInterface) GetCameraFeedId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetCameraFeedId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) IsPassthroughSupported() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnIsPassthroughSupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) IsPassthroughEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnIsPassthroughEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) StartPassthrough() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnStartPassthrough), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) StopPassthrough() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnStopPassthrough), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRInterface) GetTransformForView(view int64, cam_transform Transform3D) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&view), cam_transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&view)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetTransformForView), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRInterface) GetProjectionForView(view int64, aspect float64, near float64, far float64) Projection {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&view), gdc.ConstTypePtr(&aspect), gdc.ConstTypePtr(&near), gdc.ConstTypePtr(&far)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewProjection()
	pinner.Pin(&view)
	pinner.Pin(&aspect)
	pinner.Pin(&near)
	pinner.Pin(&far)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetProjectionForView), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRInterface) GetSupportedEnvironmentBlendModes() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetSupportedEnvironmentBlendModes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRInterface) SetEnvironmentBlendMode(mode XRInterfaceEnvironmentBlendMode) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&mode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnSetEnvironmentBlendMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRInterface) GetEnvironmentBlendMode() XRInterfaceEnvironmentBlendMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRInterfaceEnvironmentBlendMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterface.fnGetEnvironmentBlendMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type XRInterfacePlayAreaChangedSignalFn func(mode int)

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
