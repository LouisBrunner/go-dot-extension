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

type ptrsForOpenXRAPIExtensionList struct {
	fnGetInstance                              gdc.MethodBindPtr
	fnGetSystemId                              gdc.MethodBindPtr
	fnGetSession                               gdc.MethodBindPtr
	fnTransformFromPose                        gdc.MethodBindPtr
	fnXrResult                                 gdc.MethodBindPtr
	fnOpenxrIsEnabled                          gdc.MethodBindPtr
	fnGetInstanceProcAddr                      gdc.MethodBindPtr
	fnGetErrorString                           gdc.MethodBindPtr
	fnGetSwapchainFormatName                   gdc.MethodBindPtr
	fnIsInitialized                            gdc.MethodBindPtr
	fnIsRunning                                gdc.MethodBindPtr
	fnGetPlaySpace                             gdc.MethodBindPtr
	fnGetNextFrameTime                         gdc.MethodBindPtr
	fnCanRender                                gdc.MethodBindPtr
	fnGetHandTracker                           gdc.MethodBindPtr
	fnRegisterCompositionLayerProvider         gdc.MethodBindPtr
	fnUnregisterCompositionLayerProvider       gdc.MethodBindPtr
	fnSetEmulateEnvironmentBlendModeAlphaBlend gdc.MethodBindPtr
	fnIsEnvironmentBlendModeAlphaSupported     gdc.MethodBindPtr
}

var ptrsForOpenXRAPIExtension ptrsForOpenXRAPIExtensionList

func initOpenXRAPIExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRAPIExtension")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_instance")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnGetInstance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_system_id")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnGetSystemId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_session")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnGetSession = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("transform_from_pose")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnTransformFromPose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3255299855))
	}
	{
		methodName := StringNameFromStr("xr_result")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnXrResult = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3886436197))
	}
	{
		methodName := StringNameFromStr("openxr_is_enabled")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnOpenxrIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2703660260))
	}
	{
		methodName := StringNameFromStr("get_instance_proc_addr")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnGetInstanceProcAddr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1597066294))
	}
	{
		methodName := StringNameFromStr("get_error_string")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnGetErrorString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 990163283))
	}
	{
		methodName := StringNameFromStr("get_swapchain_format_name")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnGetSwapchainFormatName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 990163283))
	}
	{
		methodName := StringNameFromStr("is_initialized")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnIsInitialized = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("is_running")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnIsRunning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("get_play_space")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnGetPlaySpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_next_frame_time")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnGetNextFrameTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("can_render")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnCanRender = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("get_hand_tracker")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnGetHandTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
	}
	{
		methodName := StringNameFromStr("register_composition_layer_provider")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnRegisterCompositionLayerProvider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1997997368))
	}
	{
		methodName := StringNameFromStr("unregister_composition_layer_provider")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnUnregisterCompositionLayerProvider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1997997368))
	}
	{
		methodName := StringNameFromStr("set_emulate_environment_blend_mode_alpha_blend")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnSetEmulateEnvironmentBlendModeAlphaBlend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_environment_blend_mode_alpha_supported")
		defer methodName.Destroy()
		ptrsForOpenXRAPIExtension.fnIsEnvironmentBlendModeAlphaSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1579290861))
	}

}

type OpenXRAPIExtension struct {
	RefCounted
}

func (me *OpenXRAPIExtension) BaseClass() string {
	return "OpenXRAPIExtension"
}

func NewOpenXRAPIExtension() *OpenXRAPIExtension {
	str := StringNameFromStr("OpenXRAPIExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRAPIExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type OpenXRAPIExtensionOpenXRAlphaBlendModeSupport int

const (
	OpenXRAPIExtensionOpenXRAlphaBlendModeSupportOpenxrAlphaBlendModeSupportNone      OpenXRAPIExtensionOpenXRAlphaBlendModeSupport = 0
	OpenXRAPIExtensionOpenXRAlphaBlendModeSupportOpenxrAlphaBlendModeSupportReal      OpenXRAPIExtensionOpenXRAlphaBlendModeSupport = 1
	OpenXRAPIExtensionOpenXRAlphaBlendModeSupportOpenxrAlphaBlendModeSupportEmulating OpenXRAPIExtensionOpenXRAlphaBlendModeSupport = 2
)

func (me *OpenXRAPIExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRAPIExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRAPIExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRAPIExtension) GetInstance() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnGetInstance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) GetSystemId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnGetSystemId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) GetSession() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnGetSession), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) TransformFromPose(pose unsafe.Pointer) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pose)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&pose)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnTransformFromPose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRAPIExtension) XrResult(result int64, format String, args Array) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&result), format.AsCTypePtr(), args.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&result)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnXrResult), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func OpenXRAPIExtensionOpenxrIsEnabled(check_run_in_editor bool) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&check_run_in_editor)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&check_run_in_editor)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnOpenxrIsEnabled), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) GetInstanceProcAddr(name String) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnGetInstanceProcAddr), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) GetErrorString(result int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&result)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&result)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnGetErrorString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRAPIExtension) GetSwapchainFormatName(swapchain_format int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&swapchain_format)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&swapchain_format)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnGetSwapchainFormatName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRAPIExtension) IsInitialized() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnIsInitialized), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) IsRunning() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnIsRunning), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) GetPlaySpace() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnGetPlaySpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) GetNextFrameTime() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnGetNextFrameTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) CanRender() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnCanRender), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) GetHandTracker(hand_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&hand_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnGetHandTracker), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRAPIExtension) RegisterCompositionLayerProvider(extension OpenXRExtensionWrapperExtension) {
	cargs := []gdc.ConstTypePtr{extension.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnRegisterCompositionLayerProvider), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRAPIExtension) UnregisterCompositionLayerProvider(extension OpenXRExtensionWrapperExtension) {
	cargs := []gdc.ConstTypePtr{extension.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnUnregisterCompositionLayerProvider), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRAPIExtension) SetEmulateEnvironmentBlendModeAlphaBlend(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnSetEmulateEnvironmentBlendModeAlphaBlend), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRAPIExtension) IsEnvironmentBlendModeAlphaSupported() OpenXRAPIExtensionOpenXRAlphaBlendModeSupport {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret OpenXRAPIExtensionOpenXRAlphaBlendModeSupport

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAPIExtension.fnIsEnvironmentBlendModeAlphaSupported), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals
