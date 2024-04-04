// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type Input struct {
  Object
}

func (me *Input) BaseClass() string {
  return "Input"
}

func NewInput() *Input {
  str := StringNameFromStr("Input") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Input{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *Input) IsAnythingPressed() bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_anything_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) IsKeyPressed(keycode Key, ) bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_key_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1938909964) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&keycode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) IsPhysicalKeyPressed(keycode Key, ) bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_physical_key_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1938909964) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&keycode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) IsKeyLabelPressed(keycode Key, ) bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_key_label_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1938909964) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&keycode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) IsMouseButtonPressed(button MouseButton, ) bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_mouse_button_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1821097125) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&button) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&button)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) IsJoyButtonPressed(device int64, button JoyButton, ) bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_joy_button_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 787208542) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device) , gdc.ConstTypePtr(&button) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&device)
  pinner.Pin(&button)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) IsActionPressed(action StringName, exact_match bool, ) bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_action_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1558498928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&exact_match)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) IsActionJustPressed(action StringName, exact_match bool, ) bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_action_just_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1558498928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&exact_match)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) IsActionJustReleased(action StringName, exact_match bool, ) bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_action_just_released")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1558498928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&exact_match)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) GetActionStrength(action StringName, exact_match bool, ) float64 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 801543509) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&exact_match)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) GetActionRawStrength(action StringName, exact_match bool, ) float64 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_raw_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 801543509) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&exact_match)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) GetAxis(negative_action StringName, positive_action StringName, ) float64 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1958752504) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{negative_action.AsCTypePtr(), positive_action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) GetVector(negative_x StringName, positive_x StringName, negative_y StringName, positive_y StringName, deadzone float64, ) Vector2 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vector")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2479607902) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{negative_x.AsCTypePtr(), positive_x.AsCTypePtr(), negative_y.AsCTypePtr(), positive_y.AsCTypePtr(), gdc.ConstTypePtr(&deadzone) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&deadzone)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Input) AddJoyMapping(mapping String, update_existing bool, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_joy_mapping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1168363258) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{mapping.AsCTypePtr(), gdc.ConstTypePtr(&update_existing) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) RemoveJoyMapping(guid String, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_joy_mapping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{guid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) IsJoyKnown(device int64, ) bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_joy_known")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&device)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) GetJoyAxis(device int64, axis JoyAxis, ) float64 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joy_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4063175957) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device) , gdc.ConstTypePtr(&axis) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&device)
  pinner.Pin(&axis)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) GetJoyName(device int64, ) String {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joy_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 990163283) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&device)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Input) GetJoyGuid(device int64, ) String {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joy_guid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&device)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Input) GetJoyInfo(device int64, ) Dictionary {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joy_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3485342025) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&device)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Input) ShouldIgnoreDevice(vendor_id int64, product_id int64, ) bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("should_ignore_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vendor_id) , gdc.ConstTypePtr(&product_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&vendor_id)
  pinner.Pin(&product_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) GetConnectedJoypads() []int {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connected_joypads")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[int](ret)
}

func  (me *Input) GetJoyVibrationStrength(device int64, ) Vector2 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joy_vibration_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3114997196) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&device)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Input) GetJoyVibrationDuration(device int64, ) float64 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joy_vibration_duration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4025615559) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&device)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) StartJoyVibration(device int64, weak_magnitude float64, strong_magnitude float64, duration float64, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start_joy_vibration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2576575033) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device) , gdc.ConstTypePtr(&weak_magnitude) , gdc.ConstTypePtr(&strong_magnitude) , gdc.ConstTypePtr(&duration) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) StopJoyVibration(device int64, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop_joy_vibration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) VibrateHandheld(duration_ms int64, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("vibrate_handheld")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 955504365) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&duration_ms) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) GetGravity() Vector3 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Input) GetAccelerometer() Vector3 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_accelerometer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Input) GetMagnetometer() Vector3 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_magnetometer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Input) GetGyroscope() Vector3 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gyroscope")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Input) SetGravity(value Vector3, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) SetAccelerometer(value Vector3, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_accelerometer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) SetMagnetometer(value Vector3, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_magnetometer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) SetGyroscope(value Vector3, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gyroscope")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) GetLastMouseVelocity() Vector2 {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_mouse_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1497962370) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Input) GetMouseButtonMask() MouseButtonMask {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mouse_button_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2512161324) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MouseButtonMask

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Input) SetMouseMode(mode InputMouseMode, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mouse_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2228490894) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) GetMouseMode() InputMouseMode {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mouse_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 965286182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret InputMouseMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Input) WarpMouse(position Vector2, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("warp_mouse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) ActionPress(action StringName, strength float64, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_press")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1713091165) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) ActionRelease(action StringName, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_release")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) SetDefaultCursorShape(shape InputCursorShape, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_cursor_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2124816902) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) GetCurrentCursorShape() InputCursorShape {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_cursor_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3455658929) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret InputCursorShape

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Input) SetCustomMouseCursor(image Resource, shape InputCursorShape, hotspot Vector2, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_mouse_cursor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 703945977) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), gdc.ConstTypePtr(&shape) , hotspot.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) ParseInputEvent(event InputEvent, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse_input_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3754044979) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) SetUseAccumulatedInput(enable bool, )  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_accumulated_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Input) IsUsingAccumulatedInput() bool {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_accumulated_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Input) FlushBufferedEvents()  {
  classNameV := StringNameFromStr("Input")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("flush_buffered_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type InputJoyConnectionChangedSignalFn func(device int, connected bool, )

func (me *Input) ConnectJoyConnectionChanged(subs SignalSubscribers, fn InputJoyConnectionChangedSignalFn) {
  sig := StringNameFromStr("joy_connection_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Input) DisconnectJoyConnectionChanged(subs SignalSubscribers, fn InputJoyConnectionChangedSignalFn) {
  sig := StringNameFromStr("joy_connection_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
