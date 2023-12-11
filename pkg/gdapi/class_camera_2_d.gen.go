// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Camera2D struct {
  obj gdc.ObjectPtr
}

func (me *Camera2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Camera2D) BaseClass() string {
  return "Camera2D"
}



// Enums

type Camera2DAnchorMode int
const (
  Camera2DAnchorModeAnchorModeFixedTopLeft Camera2DAnchorMode = 0
  Camera2DAnchorModeAnchorModeDragCenter Camera2DAnchorMode = 1
)

type Camera2DCamera2DProcessCallback int
const (
  Camera2DCamera2DProcessCallbackCamera2DProcessPhysics Camera2DCamera2DProcessCallback = 0
  Camera2DCamera2DProcessCallbackCamera2DProcessIdle Camera2DCamera2DProcessCallback = 1
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

func  (me *Camera2D) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetOffset() Vector2 {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetAnchorMode(anchor_mode Camera2DAnchorMode, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_anchor_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2050398218) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&anchor_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetAnchorMode() Camera2DAnchorMode {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_anchor_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 155978067) // FIXME: should cache?
  var ret Camera2DAnchorMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetIgnoreRotation(ignore bool, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ignore_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ignore), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsIgnoringRotation() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ignoring_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetProcessCallback(mode Camera2DCamera2DProcessCallback, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4201947462) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetProcessCallback() Camera2DCamera2DProcessCallback {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2325344499) // FIXME: should cache?
  var ret Camera2DCamera2DProcessCallback
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsEnabled() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) MakeCurrent()  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsCurrent() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetLimit(margin Side, limit int, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 437707142) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&limit), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetLimit(margin Side, ) int {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1983885014) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetLimitSmoothingEnabled(limit_smoothing_enabled bool, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_limit_smoothing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit_smoothing_enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsLimitSmoothingEnabled() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_limit_smoothing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetDragVerticalEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_vertical_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsDragVerticalEnabled() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drag_vertical_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetDragHorizontalEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_horizontal_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsDragHorizontalEnabled() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drag_horizontal_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetDragVerticalOffset(offset float32, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_vertical_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetDragVerticalOffset() float32 {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drag_vertical_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetDragHorizontalOffset(offset float32, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_horizontal_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetDragHorizontalOffset() float32 {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drag_horizontal_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetDragMargin(margin Side, drag_margin float32, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4290182280) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&drag_margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetDragMargin(margin Side, ) float32 {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drag_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2869120046) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) GetTargetPosition() Vector2 {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) GetScreenCenterPosition() Vector2 {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_screen_center_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetZoom(zoom Vector2, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_zoom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(zoom.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetZoom() Vector2 {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zoom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetCustomViewport(viewport Node, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_viewport")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(viewport.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetCustomViewport() Node {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_viewport")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetPositionSmoothingSpeed(position_smoothing_speed float32, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position_smoothing_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position_smoothing_speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetPositionSmoothingSpeed() float32 {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position_smoothing_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetPositionSmoothingEnabled(position_smoothing_speed bool, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position_smoothing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position_smoothing_speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsPositionSmoothingEnabled() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_position_smoothing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetRotationSmoothingEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation_smoothing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsRotationSmoothingEnabled() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_rotation_smoothing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetRotationSmoothingSpeed(speed float32, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation_smoothing_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) GetRotationSmoothingSpeed() float32 {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation_smoothing_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) ForceUpdateScroll()  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_update_scroll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) ResetSmoothing()  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reset_smoothing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) Align()  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("align")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) SetScreenDrawingEnabled(screen_drawing_enabled bool, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_screen_drawing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen_drawing_enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsScreenDrawingEnabled() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_screen_drawing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetLimitDrawingEnabled(limit_drawing_enabled bool, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_limit_drawing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit_drawing_enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsLimitDrawingEnabled() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_limit_drawing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera2D) SetMarginDrawingEnabled(margin_drawing_enabled bool, )  {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margin_drawing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin_drawing_enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera2D) IsMarginDrawingEnabled() bool {
  classNameV := StringNameFromStr("Camera2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_margin_drawing_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
