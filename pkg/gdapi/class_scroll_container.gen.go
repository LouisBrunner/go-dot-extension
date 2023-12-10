// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ScrollContainer struct {
  obj gdc.ObjectPtr
}

func (me *ScrollContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ScrollContainer) BaseClass() string {
  return "ScrollContainer"
}



// Enums

type ScrollContainerScrollMode int
const (
  ScrollContainerScrollModeScrollModeDisabled ScrollContainerScrollMode = 0
  ScrollContainerScrollModeScrollModeAuto ScrollContainerScrollMode = 1
  ScrollContainerScrollModeScrollModeShowAlways ScrollContainerScrollMode = 2
  ScrollContainerScrollModeScrollModeShowNever ScrollContainerScrollMode = 3
)

func (me *ScrollContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScrollContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScrollContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ScrollContainer) SetHScroll(value int, )  {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_scroll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScrollContainer) GetHScroll() int {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_scroll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScrollContainer) SetVScroll(value int, )  {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_scroll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScrollContainer) GetVScroll() int {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_scroll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScrollContainer) SetHorizontalCustomStep(value float32, )  {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_horizontal_custom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScrollContainer) GetHorizontalCustomStep() float32 {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_horizontal_custom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScrollContainer) SetVerticalCustomStep(value float32, )  {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical_custom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScrollContainer) GetVerticalCustomStep() float32 {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertical_custom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScrollContainer) SetHorizontalScrollMode(enable ScrollContainerScrollMode, )  {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_horizontal_scroll_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2750506364) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScrollContainer) GetHorizontalScrollMode() ScrollContainerScrollMode {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_horizontal_scroll_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3987985145) // FIXME: should cache?
  var ret ScrollContainerScrollMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScrollContainer) SetVerticalScrollMode(enable ScrollContainerScrollMode, )  {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical_scroll_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2750506364) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScrollContainer) GetVerticalScrollMode() ScrollContainerScrollMode {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertical_scroll_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3987985145) // FIXME: should cache?
  var ret ScrollContainerScrollMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScrollContainer) SetDeadzone(deadzone int, )  {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_deadzone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&deadzone), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScrollContainer) GetDeadzone() int {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_deadzone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScrollContainer) SetFollowFocus(enabled bool, )  {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_follow_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScrollContainer) IsFollowingFocus() bool {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_following_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScrollContainer) GetHScrollBar() HScrollBar {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_scroll_bar")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4004517983) // FIXME: should cache?
  var ret HScrollBar
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScrollContainer) GetVScrollBar() VScrollBar {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_scroll_bar")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2630340773) // FIXME: should cache?
  var ret VScrollBar
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScrollContainer) EnsureControlVisible(control Control, )  {
  classNameV := StringNameFromStr("ScrollContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("ensure_control_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(control.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *ScrollContainer) GetPropFollowFocus() bool {
  panic("TODO: implement")
}

func (me *ScrollContainer) SetPropFollowFocus(value bool) {
  panic("TODO: implement")
}

func (me *ScrollContainer) GetPropScrollHorizontal() int {
  panic("TODO: implement")
}

func (me *ScrollContainer) SetPropScrollHorizontal(value int) {
  panic("TODO: implement")
}

func (me *ScrollContainer) GetPropScrollVertical() int {
  panic("TODO: implement")
}

func (me *ScrollContainer) SetPropScrollVertical(value int) {
  panic("TODO: implement")
}

func (me *ScrollContainer) GetPropScrollHorizontalCustomStep() float32 {
  panic("TODO: implement")
}

func (me *ScrollContainer) SetPropScrollHorizontalCustomStep(value float32) {
  panic("TODO: implement")
}

func (me *ScrollContainer) GetPropScrollVerticalCustomStep() float32 {
  panic("TODO: implement")
}

func (me *ScrollContainer) SetPropScrollVerticalCustomStep(value float32) {
  panic("TODO: implement")
}

func (me *ScrollContainer) GetPropHorizontalScrollMode() int {
  panic("TODO: implement")
}

func (me *ScrollContainer) SetPropHorizontalScrollMode(value int) {
  panic("TODO: implement")
}

func (me *ScrollContainer) GetPropVerticalScrollMode() int {
  panic("TODO: implement")
}

func (me *ScrollContainer) SetPropVerticalScrollMode(value int) {
  panic("TODO: implement")
}

func (me *ScrollContainer) GetPropScrollDeadzone() int {
  panic("TODO: implement")
}

func (me *ScrollContainer) SetPropScrollDeadzone(value int) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API