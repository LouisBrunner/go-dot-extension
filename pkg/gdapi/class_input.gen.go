// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Input struct {
  obj gdc.ObjectPtr
}

func (me *Input) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Input) BaseClass() string {
  return "Input"
}

type InputMouseMode int
const (
  InputMouseModeMouseModeVisible InputMouseMode = 0
  InputMouseModeMouseModeHidden InputMouseMode = 1
  InputMouseModeMouseModeCaptured InputMouseMode = 2
  InputMouseModeMouseModeConfined InputMouseMode = 3
  InputMouseModeMouseModeConfinedHidden InputMouseMode = 4
)

type InputCursorShape int
const (
  InputCursorShapeCursorArrow InputCursorShape = 0
  InputCursorShapeCursorIbeam InputCursorShape = 1
  InputCursorShapeCursorPointingHand InputCursorShape = 2
  InputCursorShapeCursorCross InputCursorShape = 3
  InputCursorShapeCursorWait InputCursorShape = 4
  InputCursorShapeCursorBusy InputCursorShape = 5
  InputCursorShapeCursorDrag InputCursorShape = 6
  InputCursorShapeCursorCanDrop InputCursorShape = 7
  InputCursorShapeCursorForbidden InputCursorShape = 8
  InputCursorShapeCursorVsize InputCursorShape = 9
  InputCursorShapeCursorHsize InputCursorShape = 10
  InputCursorShapeCursorBdiagsize InputCursorShape = 11
  InputCursorShapeCursorFdiagsize InputCursorShape = 12
  InputCursorShapeCursorMove InputCursorShape = 13
  InputCursorShapeCursorVsplit InputCursorShape = 14
  InputCursorShapeCursorHsplit InputCursorShape = 15
  InputCursorShapeCursorHelp InputCursorShape = 16
)

func  (me *Input) IsAnythingPressed() { // TODO: return value
  // TODO: implement
}

func  (me *Input) IsKeyPressed(keycode Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) IsPhysicalKeyPressed(keycode Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) IsKeyLabelPressed(keycode Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) IsMouseButtonPressed(button MouseButton, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) IsJoyButtonPressed(device int, button JoyButton, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) IsActionPressed(action StringName, exact_match bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) IsActionJustPressed(action StringName, exact_match bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) IsActionJustReleased(action StringName, exact_match bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetActionStrength(action StringName, exact_match bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetActionRawStrength(action StringName, exact_match bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetAxis(negative_action StringName, positive_action StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetVector(negative_x StringName, positive_x StringName, negative_y StringName, positive_y StringName, deadzone float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) AddJoyMapping(mapping String, update_existing bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) RemoveJoyMapping(guid String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) IsJoyKnown(device int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetJoyAxis(device int, axis JoyAxis, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetJoyName(device int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetJoyGuid(device int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) ShouldIgnoreDevice(vendor_id int, product_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetConnectedJoypads() { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetJoyVibrationStrength(device int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetJoyVibrationDuration(device int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) StartJoyVibration(device int, weak_magnitude float32, strong_magnitude float32, duration float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) StopJoyVibration(device int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) VibrateHandheld(duration_ms int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetGravity() { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetAccelerometer() { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetMagnetometer() { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetGyroscope() { // TODO: return value
  // TODO: implement
}

func  (me *Input) SetGravity(value Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) SetAccelerometer(value Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) SetMagnetometer(value Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) SetGyroscope(value Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetLastMouseVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetMouseButtonMask() { // TODO: return value
  // TODO: implement
}

func  (me *Input) SetMouseMode(mode InputMouseMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetMouseMode() { // TODO: return value
  // TODO: implement
}

func  (me *Input) WarpMouse(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) ActionPress(action StringName, strength float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) ActionRelease(action StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) SetDefaultCursorShape(shape InputCursorShape, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) GetCurrentCursorShape() { // TODO: return value
  // TODO: implement
}

func  (me *Input) SetCustomMouseCursor(image Resource, shape InputCursorShape, hotspot Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) ParseInputEvent(event InputEvent, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) SetUseAccumulatedInput(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Input) IsUsingAccumulatedInput() { // TODO: return value
  // TODO: implement
}

func  (me *Input) FlushBufferedEvents() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
