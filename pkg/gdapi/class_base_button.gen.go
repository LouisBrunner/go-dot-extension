// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BaseButton struct {
  obj gdc.ObjectPtr
}

func (me *BaseButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BaseButton) BaseClass() string {
  return "BaseButton"
}



// Enums

type BaseButtonDrawMode int
const (
  BaseButtonDrawModeDrawNormal BaseButtonDrawMode = 0
  BaseButtonDrawModeDrawPressed BaseButtonDrawMode = 1
  BaseButtonDrawModeDrawHover BaseButtonDrawMode = 2
  BaseButtonDrawModeDrawDisabled BaseButtonDrawMode = 3
  BaseButtonDrawModeDrawHoverPressed BaseButtonDrawMode = 4
)

type BaseButtonActionMode int
const (
  BaseButtonActionModeActionModeButtonPress BaseButtonActionMode = 0
  BaseButtonActionModeActionModeButtonRelease BaseButtonActionMode = 1
)

func (me *BaseButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BaseButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *BaseButton) XPressed()  {
  panic("TODO: implement")
}

func  (me *BaseButton) XToggled(button_pressed bool, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetPressed(pressed bool, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) IsPressed()  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetPressedNoSignal(pressed bool, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) IsHovered()  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetToggleMode(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) IsToggleMode()  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetShortcutInTooltip(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) IsShortcutInTooltipEnabled()  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetDisabled(disabled bool, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) IsDisabled()  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetActionMode(mode BaseButtonActionMode, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) GetActionMode()  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetButtonMask(mask MouseButtonMask, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) GetButtonMask()  {
  panic("TODO: implement")
}

func  (me *BaseButton) GetDrawMode()  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetKeepPressedOutside(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) IsKeepPressedOutside()  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetShortcutFeedback(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) IsShortcutFeedback()  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetShortcut(shortcut Shortcut, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) GetShortcut()  {
  panic("TODO: implement")
}

func  (me *BaseButton) SetButtonGroup(button_group ButtonGroup, )  {
  panic("TODO: implement")
}

func  (me *BaseButton) GetButtonGroup()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
