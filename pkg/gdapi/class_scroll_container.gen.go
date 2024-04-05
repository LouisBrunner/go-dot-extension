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

type ptrsForScrollContainerList struct {
  fnSetHScroll gdc.MethodBindPtr
  fnGetHScroll gdc.MethodBindPtr
  fnSetVScroll gdc.MethodBindPtr
  fnGetVScroll gdc.MethodBindPtr
  fnSetHorizontalCustomStep gdc.MethodBindPtr
  fnGetHorizontalCustomStep gdc.MethodBindPtr
  fnSetVerticalCustomStep gdc.MethodBindPtr
  fnGetVerticalCustomStep gdc.MethodBindPtr
  fnSetHorizontalScrollMode gdc.MethodBindPtr
  fnGetHorizontalScrollMode gdc.MethodBindPtr
  fnSetVerticalScrollMode gdc.MethodBindPtr
  fnGetVerticalScrollMode gdc.MethodBindPtr
  fnSetDeadzone gdc.MethodBindPtr
  fnGetDeadzone gdc.MethodBindPtr
  fnSetFollowFocus gdc.MethodBindPtr
  fnIsFollowingFocus gdc.MethodBindPtr
  fnGetHScrollBar gdc.MethodBindPtr
  fnGetVScrollBar gdc.MethodBindPtr
  fnEnsureControlVisible gdc.MethodBindPtr
}

var ptrsForScrollContainer ptrsForScrollContainerList

func initScrollContainerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ScrollContainer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_h_scroll")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnSetHScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_h_scroll")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnGetHScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_v_scroll")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnSetVScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_v_scroll")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnGetVScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_horizontal_custom_step")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnSetHorizontalCustomStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_horizontal_custom_step")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnGetHorizontalCustomStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_vertical_custom_step")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnSetVerticalCustomStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_vertical_custom_step")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnGetVerticalCustomStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_horizontal_scroll_mode")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnSetHorizontalScrollMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2750506364))
  }
  {
    methodName := StringNameFromStr("get_horizontal_scroll_mode")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnGetHorizontalScrollMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3987985145))
  }
  {
    methodName := StringNameFromStr("set_vertical_scroll_mode")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnSetVerticalScrollMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2750506364))
  }
  {
    methodName := StringNameFromStr("get_vertical_scroll_mode")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnGetVerticalScrollMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3987985145))
  }
  {
    methodName := StringNameFromStr("set_deadzone")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnSetDeadzone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_deadzone")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnGetDeadzone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_follow_focus")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnSetFollowFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_following_focus")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnIsFollowingFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_h_scroll_bar")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnGetHScrollBar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4004517983))
  }
  {
    methodName := StringNameFromStr("get_v_scroll_bar")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnGetVScrollBar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2630340773))
  }
  {
    methodName := StringNameFromStr("ensure_control_visible")
    defer methodName.Destroy()
    ptrsForScrollContainer.fnEnsureControlVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
  }
}

type ScrollContainer struct {
  Container
}

func (me *ScrollContainer) BaseClass() string {
  return "ScrollContainer"
}

func NewScrollContainer() *ScrollContainer {
  str := StringNameFromStr("ScrollContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ScrollContainer{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *ScrollContainer) SetHScroll(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnSetHScroll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScrollContainer) GetHScroll() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnGetHScroll), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ScrollContainer) SetVScroll(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnSetVScroll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScrollContainer) GetVScroll() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnGetVScroll), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ScrollContainer) SetHorizontalCustomStep(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnSetHorizontalCustomStep), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScrollContainer) GetHorizontalCustomStep() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnGetHorizontalCustomStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ScrollContainer) SetVerticalCustomStep(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnSetVerticalCustomStep), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScrollContainer) GetVerticalCustomStep() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnGetVerticalCustomStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ScrollContainer) SetHorizontalScrollMode(enable ScrollContainerScrollMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnSetHorizontalScrollMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScrollContainer) GetHorizontalScrollMode() ScrollContainerScrollMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ScrollContainerScrollMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnGetHorizontalScrollMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ScrollContainer) SetVerticalScrollMode(enable ScrollContainerScrollMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnSetVerticalScrollMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScrollContainer) GetVerticalScrollMode() ScrollContainerScrollMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ScrollContainerScrollMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnGetVerticalScrollMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ScrollContainer) SetDeadzone(deadzone int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&deadzone) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnSetDeadzone), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScrollContainer) GetDeadzone() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnGetDeadzone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ScrollContainer) SetFollowFocus(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnSetFollowFocus), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScrollContainer) IsFollowingFocus() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnIsFollowingFocus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ScrollContainer) GetHScrollBar() HScrollBar {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewHScrollBar()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnGetHScrollBar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ScrollContainer) GetVScrollBar() VScrollBar {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVScrollBar()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnGetVScrollBar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ScrollContainer) EnsureControlVisible(control Control, )  {
  cargs := []gdc.ConstTypePtr{control.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollContainer.fnEnsureControlVisible), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ScrollContainerScrollStartedSignalFn func()

func (me *ScrollContainer) ConnectScrollStarted(subs SignalSubscribers, fn ScrollContainerScrollStartedSignalFn) {
  sig := StringNameFromStr("scroll_started")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScrollContainer) DisconnectScrollStarted(subs SignalSubscribers, fn ScrollContainerScrollStartedSignalFn) {
  sig := StringNameFromStr("scroll_started")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ScrollContainerScrollEndedSignalFn func()

func (me *ScrollContainer) ConnectScrollEnded(subs SignalSubscribers, fn ScrollContainerScrollEndedSignalFn) {
  sig := StringNameFromStr("scroll_ended")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScrollContainer) DisconnectScrollEnded(subs SignalSubscribers, fn ScrollContainerScrollEndedSignalFn) {
  sig := StringNameFromStr("scroll_ended")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
