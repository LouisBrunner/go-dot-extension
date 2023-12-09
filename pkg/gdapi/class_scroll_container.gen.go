// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *ScrollContainer) GetHScroll()  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) SetVScroll(value int, )  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) GetVScroll()  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) SetHorizontalCustomStep(value float32, )  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) GetHorizontalCustomStep()  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) SetVerticalCustomStep(value float32, )  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) GetVerticalCustomStep()  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) SetHorizontalScrollMode(enable ScrollContainerScrollMode, )  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) GetHorizontalScrollMode()  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) SetVerticalScrollMode(enable ScrollContainerScrollMode, )  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) GetVerticalScrollMode()  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) SetDeadzone(deadzone int, )  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) GetDeadzone()  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) SetFollowFocus(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) IsFollowingFocus()  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) GetHScrollBar()  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) GetVScrollBar()  {
  panic("TODO: implement")
}

func  (me *ScrollContainer) EnsureControlVisible(control Control, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
