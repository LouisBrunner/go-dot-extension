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

type ptrsForCamera2DList struct {
	fnSetOffset                   gdc.MethodBindPtr
	fnGetOffset                   gdc.MethodBindPtr
	fnSetAnchorMode               gdc.MethodBindPtr
	fnGetAnchorMode               gdc.MethodBindPtr
	fnSetIgnoreRotation           gdc.MethodBindPtr
	fnIsIgnoringRotation          gdc.MethodBindPtr
	fnSetProcessCallback          gdc.MethodBindPtr
	fnGetProcessCallback          gdc.MethodBindPtr
	fnSetEnabled                  gdc.MethodBindPtr
	fnIsEnabled                   gdc.MethodBindPtr
	fnMakeCurrent                 gdc.MethodBindPtr
	fnIsCurrent                   gdc.MethodBindPtr
	fnSetLimit                    gdc.MethodBindPtr
	fnGetLimit                    gdc.MethodBindPtr
	fnSetLimitSmoothingEnabled    gdc.MethodBindPtr
	fnIsLimitSmoothingEnabled     gdc.MethodBindPtr
	fnSetDragVerticalEnabled      gdc.MethodBindPtr
	fnIsDragVerticalEnabled       gdc.MethodBindPtr
	fnSetDragHorizontalEnabled    gdc.MethodBindPtr
	fnIsDragHorizontalEnabled     gdc.MethodBindPtr
	fnSetDragVerticalOffset       gdc.MethodBindPtr
	fnGetDragVerticalOffset       gdc.MethodBindPtr
	fnSetDragHorizontalOffset     gdc.MethodBindPtr
	fnGetDragHorizontalOffset     gdc.MethodBindPtr
	fnSetDragMargin               gdc.MethodBindPtr
	fnGetDragMargin               gdc.MethodBindPtr
	fnGetTargetPosition           gdc.MethodBindPtr
	fnGetScreenCenterPosition     gdc.MethodBindPtr
	fnSetZoom                     gdc.MethodBindPtr
	fnGetZoom                     gdc.MethodBindPtr
	fnSetCustomViewport           gdc.MethodBindPtr
	fnGetCustomViewport           gdc.MethodBindPtr
	fnSetPositionSmoothingSpeed   gdc.MethodBindPtr
	fnGetPositionSmoothingSpeed   gdc.MethodBindPtr
	fnSetPositionSmoothingEnabled gdc.MethodBindPtr
	fnIsPositionSmoothingEnabled  gdc.MethodBindPtr
	fnSetRotationSmoothingEnabled gdc.MethodBindPtr
	fnIsRotationSmoothingEnabled  gdc.MethodBindPtr
	fnSetRotationSmoothingSpeed   gdc.MethodBindPtr
	fnGetRotationSmoothingSpeed   gdc.MethodBindPtr
	fnForceUpdateScroll           gdc.MethodBindPtr
	fnResetSmoothing              gdc.MethodBindPtr
	fnAlign                       gdc.MethodBindPtr
	fnSetScreenDrawingEnabled     gdc.MethodBindPtr
	fnIsScreenDrawingEnabled      gdc.MethodBindPtr
	fnSetLimitDrawingEnabled      gdc.MethodBindPtr
	fnIsLimitDrawingEnabled       gdc.MethodBindPtr
	fnSetMarginDrawingEnabled     gdc.MethodBindPtr
	fnIsMarginDrawingEnabled      gdc.MethodBindPtr
}

var ptrsForCamera2D ptrsForCamera2DList

func initCamera2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Camera2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_offset")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_offset")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_anchor_mode")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetAnchorMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2050398218))
	}
	{
		methodName := StringNameFromStr("get_anchor_mode")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetAnchorMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 155978067))
	}
	{
		methodName := StringNameFromStr("set_ignore_rotation")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetIgnoreRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_ignoring_rotation")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsIgnoringRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_process_callback")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetProcessCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4201947462))
	}
	{
		methodName := StringNameFromStr("get_process_callback")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetProcessCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2325344499))
	}
	{
		methodName := StringNameFromStr("set_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("make_current")
		defer methodName.Destroy()
		ptrsForCamera2D.fnMakeCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_current")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_limit")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetLimit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 437707142))
	}
	{
		methodName := StringNameFromStr("get_limit")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetLimit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1983885014))
	}
	{
		methodName := StringNameFromStr("set_limit_smoothing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetLimitSmoothingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_limit_smoothing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsLimitSmoothingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_drag_vertical_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetDragVerticalEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_drag_vertical_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsDragVerticalEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_drag_horizontal_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetDragHorizontalEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_drag_horizontal_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsDragHorizontalEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_drag_vertical_offset")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetDragVerticalOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_drag_vertical_offset")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetDragVerticalOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_drag_horizontal_offset")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetDragHorizontalOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_drag_horizontal_offset")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetDragHorizontalOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_drag_margin")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetDragMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4290182280))
	}
	{
		methodName := StringNameFromStr("get_drag_margin")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetDragMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2869120046))
	}
	{
		methodName := StringNameFromStr("get_target_position")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_screen_center_position")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetScreenCenterPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_zoom")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetZoom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_zoom")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetZoom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_custom_viewport")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetCustomViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("get_custom_viewport")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetCustomViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
	}
	{
		methodName := StringNameFromStr("set_position_smoothing_speed")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetPositionSmoothingSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_position_smoothing_speed")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetPositionSmoothingSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_position_smoothing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetPositionSmoothingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_position_smoothing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsPositionSmoothingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_rotation_smoothing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetRotationSmoothingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_rotation_smoothing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsRotationSmoothingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_rotation_smoothing_speed")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetRotationSmoothingSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_rotation_smoothing_speed")
		defer methodName.Destroy()
		ptrsForCamera2D.fnGetRotationSmoothingSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("force_update_scroll")
		defer methodName.Destroy()
		ptrsForCamera2D.fnForceUpdateScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("reset_smoothing")
		defer methodName.Destroy()
		ptrsForCamera2D.fnResetSmoothing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("align")
		defer methodName.Destroy()
		ptrsForCamera2D.fnAlign = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_screen_drawing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetScreenDrawingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_screen_drawing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsScreenDrawingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_limit_drawing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetLimitDrawingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_limit_drawing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsLimitDrawingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_margin_drawing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnSetMarginDrawingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_margin_drawing_enabled")
		defer methodName.Destroy()
		ptrsForCamera2D.fnIsMarginDrawingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type Camera2D struct {
	Node2D
}

func (me *Camera2D) BaseClass() string {
	return "Camera2D"
}

func NewCamera2D() *Camera2D {
	str := StringNameFromStr("Camera2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Camera2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type Camera2DAnchorMode int

const (
	Camera2DAnchorModeAnchorModeFixedTopLeft Camera2DAnchorMode = 0
	Camera2DAnchorModeAnchorModeDragCenter   Camera2DAnchorMode = 1
)

type Camera2DCamera2DProcessCallback int

const (
	Camera2DCamera2DProcessCallbackCamera2DProcessPhysics Camera2DCamera2DProcessCallback = 0
	Camera2DCamera2DProcessCallbackCamera2DProcessIdle    Camera2DCamera2DProcessCallback = 1
)

func (me *Camera2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Camera2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Camera2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Camera2D) SetOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera2D) SetAnchorMode(anchor_mode Camera2DAnchorMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&anchor_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetAnchorMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetAnchorMode() Camera2DAnchorMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Camera2DAnchorMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetAnchorMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Camera2D) SetIgnoreRotation(ignore bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ignore)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetIgnoreRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsIgnoringRotation() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsIgnoringRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetProcessCallback(mode Camera2DCamera2DProcessCallback) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetProcessCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetProcessCallback() Camera2DCamera2DProcessCallback {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Camera2DCamera2DProcessCallback

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetProcessCallback), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Camera2D) SetEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) MakeCurrent() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnMakeCurrent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsCurrent() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsCurrent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetLimit(margin Side, limit int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&limit)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetLimit), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetLimit(margin Side) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&margin)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetLimit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetLimitSmoothingEnabled(limit_smoothing_enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit_smoothing_enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetLimitSmoothingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsLimitSmoothingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsLimitSmoothingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetDragVerticalEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetDragVerticalEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsDragVerticalEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsDragVerticalEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetDragHorizontalEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetDragHorizontalEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsDragHorizontalEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsDragHorizontalEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetDragVerticalOffset(offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetDragVerticalOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetDragVerticalOffset() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetDragVerticalOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetDragHorizontalOffset(offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetDragHorizontalOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetDragHorizontalOffset() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetDragHorizontalOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetDragMargin(margin Side, drag_margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&drag_margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetDragMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetDragMargin(margin Side) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&margin)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetDragMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) GetTargetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetTargetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera2D) GetScreenCenterPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetScreenCenterPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera2D) SetZoom(zoom Vector2) {
	cargs := []gdc.ConstTypePtr{zoom.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetZoom), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetZoom() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetZoom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera2D) SetCustomViewport(viewport Node) {
	cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetCustomViewport), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetCustomViewport() Node {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetCustomViewport), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera2D) SetPositionSmoothingSpeed(position_smoothing_speed float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position_smoothing_speed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetPositionSmoothingSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetPositionSmoothingSpeed() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetPositionSmoothingSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetPositionSmoothingEnabled(position_smoothing_speed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position_smoothing_speed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetPositionSmoothingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsPositionSmoothingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsPositionSmoothingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetRotationSmoothingEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetRotationSmoothingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsRotationSmoothingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsRotationSmoothingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetRotationSmoothingSpeed(speed float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetRotationSmoothingSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) GetRotationSmoothingSpeed() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnGetRotationSmoothingSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) ForceUpdateScroll() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnForceUpdateScroll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) ResetSmoothing() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnResetSmoothing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) Align() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnAlign), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) SetScreenDrawingEnabled(screen_drawing_enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen_drawing_enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetScreenDrawingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsScreenDrawingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsScreenDrawingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetLimitDrawingEnabled(limit_drawing_enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit_drawing_enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetLimitDrawingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsLimitDrawingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsLimitDrawingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera2D) SetMarginDrawingEnabled(margin_drawing_enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin_drawing_enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnSetMarginDrawingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera2D) IsMarginDrawingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera2D.fnIsMarginDrawingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
