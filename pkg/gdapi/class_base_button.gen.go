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

func  (me *BaseButton) XPressed() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) XToggled(button_pressed bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetPressed(pressed bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) IsPressed() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetPressedNoSignal(pressed bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) IsHovered() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetToggleMode(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) IsToggleMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetShortcutInTooltip(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) IsShortcutInTooltipEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetDisabled(disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) IsDisabled() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetActionMode(mode BaseButtonActionMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) GetActionMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetButtonMask(mask MouseButtonMask, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) GetButtonMask() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) GetDrawMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetKeepPressedOutside(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) IsKeepPressedOutside() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetShortcutFeedback(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) IsShortcutFeedback() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetShortcut(shortcut Shortcut, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) GetShortcut() { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) SetButtonGroup(button_group ButtonGroup, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseButton) GetButtonGroup() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
