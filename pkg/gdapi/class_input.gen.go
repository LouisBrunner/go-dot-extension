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



// Enums

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

func (me *Input) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Input) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Input) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Input) IsAnythingPressed()  {
  panic("TODO: implement")
}

func  (me *Input) IsKeyPressed(keycode Key, )  {
  panic("TODO: implement")
}

func  (me *Input) IsPhysicalKeyPressed(keycode Key, )  {
  panic("TODO: implement")
}

func  (me *Input) IsKeyLabelPressed(keycode Key, )  {
  panic("TODO: implement")
}

func  (me *Input) IsMouseButtonPressed(button MouseButton, )  {
  panic("TODO: implement")
}

func  (me *Input) IsJoyButtonPressed(device int, button JoyButton, )  {
  panic("TODO: implement")
}

func  (me *Input) IsActionPressed(action StringName, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *Input) IsActionJustPressed(action StringName, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *Input) IsActionJustReleased(action StringName, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *Input) GetActionStrength(action StringName, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *Input) GetActionRawStrength(action StringName, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *Input) GetAxis(negative_action StringName, positive_action StringName, )  {
  panic("TODO: implement")
}

func  (me *Input) GetVector(negative_x StringName, positive_x StringName, negative_y StringName, positive_y StringName, deadzone float32, )  {
  panic("TODO: implement")
}

func  (me *Input) AddJoyMapping(mapping String, update_existing bool, )  {
  panic("TODO: implement")
}

func  (me *Input) RemoveJoyMapping(guid String, )  {
  panic("TODO: implement")
}

func  (me *Input) IsJoyKnown(device int, )  {
  panic("TODO: implement")
}

func  (me *Input) GetJoyAxis(device int, axis JoyAxis, )  {
  panic("TODO: implement")
}

func  (me *Input) GetJoyName(device int, )  {
  panic("TODO: implement")
}

func  (me *Input) GetJoyGuid(device int, )  {
  panic("TODO: implement")
}

func  (me *Input) ShouldIgnoreDevice(vendor_id int, product_id int, )  {
  panic("TODO: implement")
}

func  (me *Input) GetConnectedJoypads()  {
  panic("TODO: implement")
}

func  (me *Input) GetJoyVibrationStrength(device int, )  {
  panic("TODO: implement")
}

func  (me *Input) GetJoyVibrationDuration(device int, )  {
  panic("TODO: implement")
}

func  (me *Input) StartJoyVibration(device int, weak_magnitude float32, strong_magnitude float32, duration float32, )  {
  panic("TODO: implement")
}

func  (me *Input) StopJoyVibration(device int, )  {
  panic("TODO: implement")
}

func  (me *Input) VibrateHandheld(duration_ms int, )  {
  panic("TODO: implement")
}

func  (me *Input) GetGravity()  {
  panic("TODO: implement")
}

func  (me *Input) GetAccelerometer()  {
  panic("TODO: implement")
}

func  (me *Input) GetMagnetometer()  {
  panic("TODO: implement")
}

func  (me *Input) GetGyroscope()  {
  panic("TODO: implement")
}

func  (me *Input) SetGravity(value Vector3, )  {
  panic("TODO: implement")
}

func  (me *Input) SetAccelerometer(value Vector3, )  {
  panic("TODO: implement")
}

func  (me *Input) SetMagnetometer(value Vector3, )  {
  panic("TODO: implement")
}

func  (me *Input) SetGyroscope(value Vector3, )  {
  panic("TODO: implement")
}

func  (me *Input) GetLastMouseVelocity()  {
  panic("TODO: implement")
}

func  (me *Input) GetMouseButtonMask()  {
  panic("TODO: implement")
}

func  (me *Input) SetMouseMode(mode InputMouseMode, )  {
  panic("TODO: implement")
}

func  (me *Input) GetMouseMode()  {
  panic("TODO: implement")
}

func  (me *Input) WarpMouse(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Input) ActionPress(action StringName, strength float32, )  {
  panic("TODO: implement")
}

func  (me *Input) ActionRelease(action StringName, )  {
  panic("TODO: implement")
}

func  (me *Input) SetDefaultCursorShape(shape InputCursorShape, )  {
  panic("TODO: implement")
}

func  (me *Input) GetCurrentCursorShape()  {
  panic("TODO: implement")
}

func  (me *Input) SetCustomMouseCursor(image Resource, shape InputCursorShape, hotspot Vector2, )  {
  panic("TODO: implement")
}

func  (me *Input) ParseInputEvent(event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *Input) SetUseAccumulatedInput(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Input) IsUsingAccumulatedInput()  {
  panic("TODO: implement")
}

func  (me *Input) FlushBufferedEvents()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
