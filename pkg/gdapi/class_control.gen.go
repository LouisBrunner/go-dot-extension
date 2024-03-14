// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Control struct {
  CanvasItem
}

func (me *Control) BaseClass() string {
  return "Control"
}



// Constants

var (
  ControlNotificationResized = "40" // TODO: construct correctly
  ControlNotificationMouseEnter = "41" // TODO: construct correctly
  ControlNotificationMouseExit = "42" // TODO: construct correctly
  ControlNotificationMouseEnterSelf = "60" // TODO: construct correctly
  ControlNotificationMouseExitSelf = "61" // TODO: construct correctly
  ControlNotificationFocusEnter = "43" // TODO: construct correctly
  ControlNotificationFocusExit = "44" // TODO: construct correctly
  ControlNotificationThemeChanged = "45" // TODO: construct correctly
  ControlNotificationScrollBegin = "47" // TODO: construct correctly
  ControlNotificationScrollEnd = "48" // TODO: construct correctly
  ControlNotificationLayoutDirectionChanged = "49" // TODO: construct correctly
)

// Enums

type ControlFocusMode int
const (
  ControlFocusModeFocusNone ControlFocusMode = 0
  ControlFocusModeFocusClick ControlFocusMode = 1
  ControlFocusModeFocusAll ControlFocusMode = 2
)

type ControlCursorShape int
const (
  ControlCursorShapeCursorArrow ControlCursorShape = 0
  ControlCursorShapeCursorIbeam ControlCursorShape = 1
  ControlCursorShapeCursorPointingHand ControlCursorShape = 2
  ControlCursorShapeCursorCross ControlCursorShape = 3
  ControlCursorShapeCursorWait ControlCursorShape = 4
  ControlCursorShapeCursorBusy ControlCursorShape = 5
  ControlCursorShapeCursorDrag ControlCursorShape = 6
  ControlCursorShapeCursorCanDrop ControlCursorShape = 7
  ControlCursorShapeCursorForbidden ControlCursorShape = 8
  ControlCursorShapeCursorVsize ControlCursorShape = 9
  ControlCursorShapeCursorHsize ControlCursorShape = 10
  ControlCursorShapeCursorBdiagsize ControlCursorShape = 11
  ControlCursorShapeCursorFdiagsize ControlCursorShape = 12
  ControlCursorShapeCursorMove ControlCursorShape = 13
  ControlCursorShapeCursorVsplit ControlCursorShape = 14
  ControlCursorShapeCursorHsplit ControlCursorShape = 15
  ControlCursorShapeCursorHelp ControlCursorShape = 16
)

type ControlLayoutPreset int
const (
  ControlLayoutPresetPresetTopLeft ControlLayoutPreset = 0
  ControlLayoutPresetPresetTopRight ControlLayoutPreset = 1
  ControlLayoutPresetPresetBottomLeft ControlLayoutPreset = 2
  ControlLayoutPresetPresetBottomRight ControlLayoutPreset = 3
  ControlLayoutPresetPresetCenterLeft ControlLayoutPreset = 4
  ControlLayoutPresetPresetCenterTop ControlLayoutPreset = 5
  ControlLayoutPresetPresetCenterRight ControlLayoutPreset = 6
  ControlLayoutPresetPresetCenterBottom ControlLayoutPreset = 7
  ControlLayoutPresetPresetCenter ControlLayoutPreset = 8
  ControlLayoutPresetPresetLeftWide ControlLayoutPreset = 9
  ControlLayoutPresetPresetTopWide ControlLayoutPreset = 10
  ControlLayoutPresetPresetRightWide ControlLayoutPreset = 11
  ControlLayoutPresetPresetBottomWide ControlLayoutPreset = 12
  ControlLayoutPresetPresetVcenterWide ControlLayoutPreset = 13
  ControlLayoutPresetPresetHcenterWide ControlLayoutPreset = 14
  ControlLayoutPresetPresetFullRect ControlLayoutPreset = 15
)

type ControlLayoutPresetMode int
const (
  ControlLayoutPresetModePresetModeMinsize ControlLayoutPresetMode = 0
  ControlLayoutPresetModePresetModeKeepWidth ControlLayoutPresetMode = 1
  ControlLayoutPresetModePresetModeKeepHeight ControlLayoutPresetMode = 2
  ControlLayoutPresetModePresetModeKeepSize ControlLayoutPresetMode = 3
)

type ControlSizeFlags int
const (
  ControlSizeFlagsSizeShrinkBegin ControlSizeFlags = 0
  ControlSizeFlagsSizeFill ControlSizeFlags = 1
  ControlSizeFlagsSizeExpand ControlSizeFlags = 2
  ControlSizeFlagsSizeExpandFill ControlSizeFlags = 3
  ControlSizeFlagsSizeShrinkCenter ControlSizeFlags = 4
  ControlSizeFlagsSizeShrinkEnd ControlSizeFlags = 8
)

type ControlMouseFilter int
const (
  ControlMouseFilterMouseFilterStop ControlMouseFilter = 0
  ControlMouseFilterMouseFilterPass ControlMouseFilter = 1
  ControlMouseFilterMouseFilterIgnore ControlMouseFilter = 2
)

type ControlGrowDirection int
const (
  ControlGrowDirectionGrowDirectionBegin ControlGrowDirection = 0
  ControlGrowDirectionGrowDirectionEnd ControlGrowDirection = 1
  ControlGrowDirectionGrowDirectionBoth ControlGrowDirection = 2
)

type ControlAnchor int
const (
  ControlAnchorAnchorBegin ControlAnchor = 0
  ControlAnchorAnchorEnd ControlAnchor = 1
)

type ControlLayoutDirection int
const (
  ControlLayoutDirectionLayoutDirectionInherited ControlLayoutDirection = 0
  ControlLayoutDirectionLayoutDirectionLocale ControlLayoutDirection = 1
  ControlLayoutDirectionLayoutDirectionLtr ControlLayoutDirection = 2
  ControlLayoutDirectionLayoutDirectionRtl ControlLayoutDirection = 3
)

type ControlTextDirection int
const (
  ControlTextDirectionTextDirectionInherited ControlTextDirection = 3
  ControlTextDirectionTextDirectionAuto ControlTextDirection = 0
  ControlTextDirectionTextDirectionLtr ControlTextDirection = 1
  ControlTextDirectionTextDirectionRtl ControlTextDirection = 2
)

func (me *Control) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Control) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Control) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Control) AcceptEvent()  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("accept_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetMinimumSize() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minimum_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetCombinedMinimumSize() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_combined_minimum_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetAnchorsPreset(preset ControlLayoutPreset, keep_offsets bool, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_anchors_preset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 509135270) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&preset), gdc.ConstTypePtr(&keep_offsets), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetOffsetsPreset(preset ControlLayoutPreset, resize_mode ControlLayoutPresetMode, margin int, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offsets_preset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3724524307) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&preset), gdc.ConstTypePtr(&resize_mode), gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetAnchorsAndOffsetsPreset(preset ControlLayoutPreset, resize_mode ControlLayoutPresetMode, margin int, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_anchors_and_offsets_preset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3724524307) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&preset), gdc.ConstTypePtr(&resize_mode), gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetAnchor(side Side, anchor float32, keep_offset bool, push_opposite_anchor bool, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_anchor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2302782885) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), gdc.ConstTypePtr(&anchor), gdc.ConstTypePtr(&keep_offset), gdc.ConstTypePtr(&push_opposite_anchor), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetAnchor(side Side, ) float32 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_anchor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2869120046) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetOffset(side Side, offset float32, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4290182280) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetOffset(offset Side, ) float32 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2869120046) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetAnchorAndOffset(side Side, anchor float32, offset float32, push_opposite_anchor bool, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_anchor_and_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4031722181) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), gdc.ConstTypePtr(&anchor), gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&push_opposite_anchor), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetBegin(position Vector2, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetEnd(position Vector2, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetPosition(position Vector2, keep_offsets bool, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2436320129) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&keep_offsets), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetSize(size Vector2, keep_offsets bool, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2436320129) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&keep_offsets), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) ResetSize()  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reset_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetCustomMinimumSize(size Vector2, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_minimum_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetGlobalPosition(position Vector2, keep_offsets bool, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2436320129) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&keep_offsets), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetRotation(radians float32, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetRotationDegrees(degrees float32, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetScale(scale Vector2, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetPivotOffset(pivot_offset Vector2, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pivot_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pivot_offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetBegin() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetEnd() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetPosition() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetSize() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetRotation() float32 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetRotationDegrees() float32 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetScale() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetPivotOffset() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pivot_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetCustomMinimumSize() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_minimum_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetParentAreaSize() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parent_area_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetGlobalPosition() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetScreenPosition() Vector2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_screen_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetRect() Rect2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetGlobalRect() Rect2 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetFocusMode(mode ControlFocusMode, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_focus_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3232914922) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetFocusMode() ControlFocusMode {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_focus_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2132829277) // FIXME: should cache?
  var ret ControlFocusMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasFocus() bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GrabFocus()  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("grab_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) ReleaseFocus()  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("release_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) FindPrevValidFocus() Control {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_prev_valid_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2783021301) // FIXME: should cache?
  var ret Control
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) FindNextValidFocus() Control {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_next_valid_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2783021301) // FIXME: should cache?
  var ret Control
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) FindValidFocusNeighbor(side Side, ) Control {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_valid_focus_neighbor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1543910170) // FIXME: should cache?
  var ret Control
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetHSizeFlags(flags ControlSizeFlags, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_size_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 394851643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetHSizeFlags() ControlSizeFlags {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_size_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3781367401) // FIXME: should cache?
  var ret ControlSizeFlags
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetStretchRatio(ratio float32, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stretch_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetStretchRatio() float32 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stretch_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetVSizeFlags(flags ControlSizeFlags, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_size_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 394851643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetVSizeFlags() ControlSizeFlags {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_size_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3781367401) // FIXME: should cache?
  var ret ControlSizeFlags
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetTheme(theme Theme, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_theme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2326690814) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(theme.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetTheme() Theme {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3846893731) // FIXME: should cache?
  var ret Theme
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetThemeTypeVariation(theme_type StringName, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_theme_type_variation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetThemeTypeVariation() StringName {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_type_variation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) BeginBulkThemeOverride()  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("begin_bulk_theme_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) EndBulkThemeOverride()  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("end_bulk_theme_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) AddThemeIconOverride(name StringName, texture Texture2D, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_icon_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1373065600) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) AddThemeStyleboxOverride(name StringName, stylebox StyleBox, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_stylebox_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4188838905) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(stylebox.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) AddThemeFontOverride(name StringName, font Font, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_font_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3518018674) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(font.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) AddThemeFontSizeOverride(name StringName, font_size int, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_font_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2415702435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) AddThemeColorOverride(name StringName, color Color, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_color_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4260178595) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) AddThemeConstantOverride(name StringName, constant int, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_constant_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2415702435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&constant), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) RemoveThemeIconOverride(name StringName, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_icon_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) RemoveThemeStyleboxOverride(name StringName, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_stylebox_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) RemoveThemeFontOverride(name StringName, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_font_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) RemoveThemeFontSizeOverride(name StringName, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_font_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) RemoveThemeColorOverride(name StringName, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_color_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) RemoveThemeConstantOverride(name StringName, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_constant_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetThemeIcon(name StringName, theme_type StringName, ) Texture2D {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3163973443) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetThemeStylebox(name StringName, theme_type StringName, ) StyleBox {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 604739069) // FIXME: should cache?
  var ret StyleBox
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetThemeFont(name StringName, theme_type StringName, ) Font {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2826986490) // FIXME: should cache?
  var ret Font
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetThemeFontSize(name StringName, theme_type StringName, ) int {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1327056374) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetThemeColor(name StringName, theme_type StringName, ) Color {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2798751242) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetThemeConstant(name StringName, theme_type StringName, ) int {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1327056374) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeIconOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_icon_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeStyleboxOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_stylebox_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeFontOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_font_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeFontSizeOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_font_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeColorOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_color_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeConstantOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_constant_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeIcon(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866386512) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeStylebox(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866386512) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeFont(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866386512) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeFontSize(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866386512) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeColor(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866386512) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) HasThemeConstant(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866386512) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetThemeDefaultBaseScale() float32 {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_default_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetThemeDefaultFont() Font {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_default_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229501585) // FIXME: should cache?
  var ret Font
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetThemeDefaultFontSize() int {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_default_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetParentControl() Control {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parent_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2783021301) // FIXME: should cache?
  var ret Control
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetHGrowDirection(direction ControlGrowDirection, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_grow_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2022385301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetHGrowDirection() ControlGrowDirection {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_grow_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635610155) // FIXME: should cache?
  var ret ControlGrowDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetVGrowDirection(direction ControlGrowDirection, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_grow_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2022385301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetVGrowDirection() ControlGrowDirection {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_grow_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635610155) // FIXME: should cache?
  var ret ControlGrowDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetTooltipText(hint String, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tooltip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(hint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetTooltipText() String {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tooltip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetTooltip(at_position Vector2, ) String {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2895288280) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(at_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetDefaultCursorShape(shape ControlCursorShape, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_cursor_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 217062046) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetDefaultCursorShape() ControlCursorShape {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_cursor_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2359535750) // FIXME: should cache?
  var ret ControlCursorShape
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GetCursorShape(position Vector2, ) ControlCursorShape {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cursor_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1395773853) // FIXME: should cache?
  var ret ControlCursorShape
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetFocusNeighbor(side Side, neighbor NodePath, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_focus_neighbor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2024461774) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), gdc.ConstTypePtr(neighbor.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetFocusNeighbor(side Side, ) NodePath {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_focus_neighbor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757935761) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetFocusNext(next NodePath, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_focus_next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(next.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetFocusNext() NodePath {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_focus_next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetFocusPrevious(previous NodePath, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_focus_previous")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(previous.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetFocusPrevious() NodePath {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_focus_previous")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) ForceDrag(data Variant, preview Control, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_drag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3191844692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), gdc.ConstTypePtr(preview.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetMouseFilter(filter ControlMouseFilter, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mouse_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3891156122) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetMouseFilter() ControlMouseFilter {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mouse_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1572545674) // FIXME: should cache?
  var ret ControlMouseFilter
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetForcePassScrollEvents(force_pass_scroll_events bool, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_force_pass_scroll_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force_pass_scroll_events), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) IsForcePassScrollEvents() bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_force_pass_scroll_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetClipContents(enable bool, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clip_contents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) IsClippingContents() bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_clipping_contents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) GrabClickFocus()  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("grab_click_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetDragForwarding(drag_func Callable, can_drop_func Callable, drop_func Callable, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_forwarding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1076571380) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(drag_func.AsCTypePtr()), gdc.ConstTypePtr(can_drop_func.AsCTypePtr()), gdc.ConstTypePtr(drop_func.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetDragPreview(control Control, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_preview")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(control.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) IsDragSuccessful() bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drag_successful")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) WarpMouse(position Vector2, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("warp_mouse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetShortcutContext(node Node, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut_context")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetShortcutContext() Node {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shortcut_context")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) UpdateMinimumSize()  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_minimum_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) SetLayoutDirection(direction ControlLayoutDirection, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layout_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3310692370) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) GetLayoutDirection() ControlLayoutDirection {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layout_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1546772008) // FIXME: should cache?
  var ret ControlLayoutDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) IsLayoutRtl() bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_layout_rtl")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetAutoTranslate(enable bool, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_translate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) IsAutoTranslating() bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_auto_translating")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Control) SetLocalizeNumeralSystem(enable bool, )  {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_localize_numeral_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Control) IsLocalizingNumeralSystem() bool {
  classNameV := StringNameFromStr("Control")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_localizing_numeral_system")
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

type ControlResizedSignalFn func()

func (me *Control) ConnectResized(subs SignalSubscribers, fn ControlResizedSignalFn) {
  sig := StringNameFromStr("resized")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Control) DisconnectResized(subs SignalSubscribers, fn ControlResizedSignalFn) {
  sig := StringNameFromStr("resized")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ControlGuiInputSignalFn func(event InputEvent, )

func (me *Control) ConnectGuiInput(subs SignalSubscribers, fn ControlGuiInputSignalFn) {
  sig := StringNameFromStr("gui_input")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Control) DisconnectGuiInput(subs SignalSubscribers, fn ControlGuiInputSignalFn) {
  sig := StringNameFromStr("gui_input")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ControlMouseEnteredSignalFn func()

func (me *Control) ConnectMouseEntered(subs SignalSubscribers, fn ControlMouseEnteredSignalFn) {
  sig := StringNameFromStr("mouse_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Control) DisconnectMouseEntered(subs SignalSubscribers, fn ControlMouseEnteredSignalFn) {
  sig := StringNameFromStr("mouse_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ControlMouseExitedSignalFn func()

func (me *Control) ConnectMouseExited(subs SignalSubscribers, fn ControlMouseExitedSignalFn) {
  sig := StringNameFromStr("mouse_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Control) DisconnectMouseExited(subs SignalSubscribers, fn ControlMouseExitedSignalFn) {
  sig := StringNameFromStr("mouse_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ControlFocusEnteredSignalFn func()

func (me *Control) ConnectFocusEntered(subs SignalSubscribers, fn ControlFocusEnteredSignalFn) {
  sig := StringNameFromStr("focus_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Control) DisconnectFocusEntered(subs SignalSubscribers, fn ControlFocusEnteredSignalFn) {
  sig := StringNameFromStr("focus_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ControlFocusExitedSignalFn func()

func (me *Control) ConnectFocusExited(subs SignalSubscribers, fn ControlFocusExitedSignalFn) {
  sig := StringNameFromStr("focus_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Control) DisconnectFocusExited(subs SignalSubscribers, fn ControlFocusExitedSignalFn) {
  sig := StringNameFromStr("focus_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ControlSizeFlagsChangedSignalFn func()

func (me *Control) ConnectSizeFlagsChanged(subs SignalSubscribers, fn ControlSizeFlagsChangedSignalFn) {
  sig := StringNameFromStr("size_flags_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Control) DisconnectSizeFlagsChanged(subs SignalSubscribers, fn ControlSizeFlagsChangedSignalFn) {
  sig := StringNameFromStr("size_flags_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ControlMinimumSizeChangedSignalFn func()

func (me *Control) ConnectMinimumSizeChanged(subs SignalSubscribers, fn ControlMinimumSizeChangedSignalFn) {
  sig := StringNameFromStr("minimum_size_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Control) DisconnectMinimumSizeChanged(subs SignalSubscribers, fn ControlMinimumSizeChangedSignalFn) {
  sig := StringNameFromStr("minimum_size_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ControlThemeChangedSignalFn func()

func (me *Control) ConnectThemeChanged(subs SignalSubscribers, fn ControlThemeChangedSignalFn) {
  sig := StringNameFromStr("theme_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Control) DisconnectThemeChanged(subs SignalSubscribers, fn ControlThemeChangedSignalFn) {
  sig := StringNameFromStr("theme_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
