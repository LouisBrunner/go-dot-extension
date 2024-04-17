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

type ptrsForInputList struct {
	fnIsAnythingPressed          gdc.MethodBindPtr
	fnIsKeyPressed               gdc.MethodBindPtr
	fnIsPhysicalKeyPressed       gdc.MethodBindPtr
	fnIsKeyLabelPressed          gdc.MethodBindPtr
	fnIsMouseButtonPressed       gdc.MethodBindPtr
	fnIsJoyButtonPressed         gdc.MethodBindPtr
	fnIsActionPressed            gdc.MethodBindPtr
	fnIsActionJustPressed        gdc.MethodBindPtr
	fnIsActionJustReleased       gdc.MethodBindPtr
	fnGetActionStrength          gdc.MethodBindPtr
	fnGetActionRawStrength       gdc.MethodBindPtr
	fnGetAxis                    gdc.MethodBindPtr
	fnGetVector                  gdc.MethodBindPtr
	fnAddJoyMapping              gdc.MethodBindPtr
	fnRemoveJoyMapping           gdc.MethodBindPtr
	fnIsJoyKnown                 gdc.MethodBindPtr
	fnGetJoyAxis                 gdc.MethodBindPtr
	fnGetJoyName                 gdc.MethodBindPtr
	fnGetJoyGuid                 gdc.MethodBindPtr
	fnGetJoyInfo                 gdc.MethodBindPtr
	fnShouldIgnoreDevice         gdc.MethodBindPtr
	fnGetConnectedJoypads        gdc.MethodBindPtr
	fnGetJoyVibrationStrength    gdc.MethodBindPtr
	fnGetJoyVibrationDuration    gdc.MethodBindPtr
	fnStartJoyVibration          gdc.MethodBindPtr
	fnStopJoyVibration           gdc.MethodBindPtr
	fnVibrateHandheld            gdc.MethodBindPtr
	fnGetGravity                 gdc.MethodBindPtr
	fnGetAccelerometer           gdc.MethodBindPtr
	fnGetMagnetometer            gdc.MethodBindPtr
	fnGetGyroscope               gdc.MethodBindPtr
	fnSetGravity                 gdc.MethodBindPtr
	fnSetAccelerometer           gdc.MethodBindPtr
	fnSetMagnetometer            gdc.MethodBindPtr
	fnSetGyroscope               gdc.MethodBindPtr
	fnGetLastMouseVelocity       gdc.MethodBindPtr
	fnGetLastMouseScreenVelocity gdc.MethodBindPtr
	fnGetMouseButtonMask         gdc.MethodBindPtr
	fnSetMouseMode               gdc.MethodBindPtr
	fnGetMouseMode               gdc.MethodBindPtr
	fnWarpMouse                  gdc.MethodBindPtr
	fnActionPress                gdc.MethodBindPtr
	fnActionRelease              gdc.MethodBindPtr
	fnSetDefaultCursorShape      gdc.MethodBindPtr
	fnGetCurrentCursorShape      gdc.MethodBindPtr
	fnSetCustomMouseCursor       gdc.MethodBindPtr
	fnParseInputEvent            gdc.MethodBindPtr
	fnSetUseAccumulatedInput     gdc.MethodBindPtr
	fnIsUsingAccumulatedInput    gdc.MethodBindPtr
	fnFlushBufferedEvents        gdc.MethodBindPtr
	fnSetEmulateMouseFromTouch   gdc.MethodBindPtr
	fnIsEmulatingMouseFromTouch  gdc.MethodBindPtr
	fnSetEmulateTouchFromMouse   gdc.MethodBindPtr
	fnIsEmulatingTouchFromMouse  gdc.MethodBindPtr
}

var ptrsForInput ptrsForInputList

func initInputPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Input")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("is_anything_pressed")
		defer methodName.Destroy()
		ptrsForInput.fnIsAnythingPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_key_pressed")
		defer methodName.Destroy()
		ptrsForInput.fnIsKeyPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1938909964))
	}
	{
		methodName := StringNameFromStr("is_physical_key_pressed")
		defer methodName.Destroy()
		ptrsForInput.fnIsPhysicalKeyPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1938909964))
	}
	{
		methodName := StringNameFromStr("is_key_label_pressed")
		defer methodName.Destroy()
		ptrsForInput.fnIsKeyLabelPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1938909964))
	}
	{
		methodName := StringNameFromStr("is_mouse_button_pressed")
		defer methodName.Destroy()
		ptrsForInput.fnIsMouseButtonPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1821097125))
	}
	{
		methodName := StringNameFromStr("is_joy_button_pressed")
		defer methodName.Destroy()
		ptrsForInput.fnIsJoyButtonPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 787208542))
	}
	{
		methodName := StringNameFromStr("is_action_pressed")
		defer methodName.Destroy()
		ptrsForInput.fnIsActionPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1558498928))
	}
	{
		methodName := StringNameFromStr("is_action_just_pressed")
		defer methodName.Destroy()
		ptrsForInput.fnIsActionJustPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1558498928))
	}
	{
		methodName := StringNameFromStr("is_action_just_released")
		defer methodName.Destroy()
		ptrsForInput.fnIsActionJustReleased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1558498928))
	}
	{
		methodName := StringNameFromStr("get_action_strength")
		defer methodName.Destroy()
		ptrsForInput.fnGetActionStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 801543509))
	}
	{
		methodName := StringNameFromStr("get_action_raw_strength")
		defer methodName.Destroy()
		ptrsForInput.fnGetActionRawStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 801543509))
	}
	{
		methodName := StringNameFromStr("get_axis")
		defer methodName.Destroy()
		ptrsForInput.fnGetAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1958752504))
	}
	{
		methodName := StringNameFromStr("get_vector")
		defer methodName.Destroy()
		ptrsForInput.fnGetVector = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2479607902))
	}
	{
		methodName := StringNameFromStr("add_joy_mapping")
		defer methodName.Destroy()
		ptrsForInput.fnAddJoyMapping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1168363258))
	}
	{
		methodName := StringNameFromStr("remove_joy_mapping")
		defer methodName.Destroy()
		ptrsForInput.fnRemoveJoyMapping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("is_joy_known")
		defer methodName.Destroy()
		ptrsForInput.fnIsJoyKnown = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3067735520))
	}
	{
		methodName := StringNameFromStr("get_joy_axis")
		defer methodName.Destroy()
		ptrsForInput.fnGetJoyAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4063175957))
	}
	{
		methodName := StringNameFromStr("get_joy_name")
		defer methodName.Destroy()
		ptrsForInput.fnGetJoyName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 990163283))
	}
	{
		methodName := StringNameFromStr("get_joy_guid")
		defer methodName.Destroy()
		ptrsForInput.fnGetJoyGuid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_joy_info")
		defer methodName.Destroy()
		ptrsForInput.fnGetJoyInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3485342025))
	}
	{
		methodName := StringNameFromStr("should_ignore_device")
		defer methodName.Destroy()
		ptrsForInput.fnShouldIgnoreDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
	}
	{
		methodName := StringNameFromStr("get_connected_joypads")
		defer methodName.Destroy()
		ptrsForInput.fnGetConnectedJoypads = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("get_joy_vibration_strength")
		defer methodName.Destroy()
		ptrsForInput.fnGetJoyVibrationStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3114997196))
	}
	{
		methodName := StringNameFromStr("get_joy_vibration_duration")
		defer methodName.Destroy()
		ptrsForInput.fnGetJoyVibrationDuration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4025615559))
	}
	{
		methodName := StringNameFromStr("start_joy_vibration")
		defer methodName.Destroy()
		ptrsForInput.fnStartJoyVibration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2576575033))
	}
	{
		methodName := StringNameFromStr("stop_joy_vibration")
		defer methodName.Destroy()
		ptrsForInput.fnStopJoyVibration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("vibrate_handheld")
		defer methodName.Destroy()
		ptrsForInput.fnVibrateHandheld = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 955504365))
	}
	{
		methodName := StringNameFromStr("get_gravity")
		defer methodName.Destroy()
		ptrsForInput.fnGetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_accelerometer")
		defer methodName.Destroy()
		ptrsForInput.fnGetAccelerometer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_magnetometer")
		defer methodName.Destroy()
		ptrsForInput.fnGetMagnetometer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_gyroscope")
		defer methodName.Destroy()
		ptrsForInput.fnGetGyroscope = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_gravity")
		defer methodName.Destroy()
		ptrsForInput.fnSetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("set_accelerometer")
		defer methodName.Destroy()
		ptrsForInput.fnSetAccelerometer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("set_magnetometer")
		defer methodName.Destroy()
		ptrsForInput.fnSetMagnetometer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("set_gyroscope")
		defer methodName.Destroy()
		ptrsForInput.fnSetGyroscope = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_last_mouse_velocity")
		defer methodName.Destroy()
		ptrsForInput.fnGetLastMouseVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1497962370))
	}
	{
		methodName := StringNameFromStr("get_last_mouse_screen_velocity")
		defer methodName.Destroy()
		ptrsForInput.fnGetLastMouseScreenVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1497962370))
	}
	{
		methodName := StringNameFromStr("get_mouse_button_mask")
		defer methodName.Destroy()
		ptrsForInput.fnGetMouseButtonMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2512161324))
	}
	{
		methodName := StringNameFromStr("set_mouse_mode")
		defer methodName.Destroy()
		ptrsForInput.fnSetMouseMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2228490894))
	}
	{
		methodName := StringNameFromStr("get_mouse_mode")
		defer methodName.Destroy()
		ptrsForInput.fnGetMouseMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 965286182))
	}
	{
		methodName := StringNameFromStr("warp_mouse")
		defer methodName.Destroy()
		ptrsForInput.fnWarpMouse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("action_press")
		defer methodName.Destroy()
		ptrsForInput.fnActionPress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1713091165))
	}
	{
		methodName := StringNameFromStr("action_release")
		defer methodName.Destroy()
		ptrsForInput.fnActionRelease = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("set_default_cursor_shape")
		defer methodName.Destroy()
		ptrsForInput.fnSetDefaultCursorShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2124816902))
	}
	{
		methodName := StringNameFromStr("get_current_cursor_shape")
		defer methodName.Destroy()
		ptrsForInput.fnGetCurrentCursorShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3455658929))
	}
	{
		methodName := StringNameFromStr("set_custom_mouse_cursor")
		defer methodName.Destroy()
		ptrsForInput.fnSetCustomMouseCursor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 703945977))
	}
	{
		methodName := StringNameFromStr("parse_input_event")
		defer methodName.Destroy()
		ptrsForInput.fnParseInputEvent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3754044979))
	}
	{
		methodName := StringNameFromStr("set_use_accumulated_input")
		defer methodName.Destroy()
		ptrsForInput.fnSetUseAccumulatedInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_accumulated_input")
		defer methodName.Destroy()
		ptrsForInput.fnIsUsingAccumulatedInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("flush_buffered_events")
		defer methodName.Destroy()
		ptrsForInput.fnFlushBufferedEvents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_emulate_mouse_from_touch")
		defer methodName.Destroy()
		ptrsForInput.fnSetEmulateMouseFromTouch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_emulating_mouse_from_touch")
		defer methodName.Destroy()
		ptrsForInput.fnIsEmulatingMouseFromTouch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_emulate_touch_from_mouse")
		defer methodName.Destroy()
		ptrsForInput.fnSetEmulateTouchFromMouse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_emulating_touch_from_mouse")
		defer methodName.Destroy()
		ptrsForInput.fnIsEmulatingTouchFromMouse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

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
	InputMouseModeMouseModeVisible        InputMouseMode = 0
	InputMouseModeMouseModeHidden         InputMouseMode = 1
	InputMouseModeMouseModeCaptured       InputMouseMode = 2
	InputMouseModeMouseModeConfined       InputMouseMode = 3
	InputMouseModeMouseModeConfinedHidden InputMouseMode = 4
)

type InputCursorShape int

const (
	InputCursorShapeCursorArrow        InputCursorShape = 0
	InputCursorShapeCursorIbeam        InputCursorShape = 1
	InputCursorShapeCursorPointingHand InputCursorShape = 2
	InputCursorShapeCursorCross        InputCursorShape = 3
	InputCursorShapeCursorWait         InputCursorShape = 4
	InputCursorShapeCursorBusy         InputCursorShape = 5
	InputCursorShapeCursorDrag         InputCursorShape = 6
	InputCursorShapeCursorCanDrop      InputCursorShape = 7
	InputCursorShapeCursorForbidden    InputCursorShape = 8
	InputCursorShapeCursorVsize        InputCursorShape = 9
	InputCursorShapeCursorHsize        InputCursorShape = 10
	InputCursorShapeCursorBdiagsize    InputCursorShape = 11
	InputCursorShapeCursorFdiagsize    InputCursorShape = 12
	InputCursorShapeCursorMove         InputCursorShape = 13
	InputCursorShapeCursorVsplit       InputCursorShape = 14
	InputCursorShapeCursorHsplit       InputCursorShape = 15
	InputCursorShapeCursorHelp         InputCursorShape = 16
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

func (me *Input) IsAnythingPressed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsAnythingPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) IsKeyPressed(keycode Key) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&keycode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsKeyPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) IsPhysicalKeyPressed(keycode Key) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&keycode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsPhysicalKeyPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) IsKeyLabelPressed(keycode Key) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&keycode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsKeyLabelPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) IsMouseButtonPressed(button MouseButton) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&button)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&button)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsMouseButtonPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) IsJoyButtonPressed(device int64, button JoyButton) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device), gdc.ConstTypePtr(&button)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&device)
	pinner.Pin(&button)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsJoyButtonPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) IsActionPressed(action StringName, exact_match bool) bool {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsActionPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) IsActionJustPressed(action StringName, exact_match bool) bool {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsActionJustPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) IsActionJustReleased(action StringName, exact_match bool) bool {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsActionJustReleased), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) GetActionStrength(action StringName, exact_match bool) float64 {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetActionStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) GetActionRawStrength(action StringName, exact_match bool) float64 {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetActionRawStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) GetAxis(negative_action StringName, positive_action StringName) float64 {
	cargs := []gdc.ConstTypePtr{negative_action.AsCTypePtr(), positive_action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetAxis), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) GetVector(negative_x StringName, positive_x StringName, negative_y StringName, positive_y StringName, deadzone float64) Vector2 {
	cargs := []gdc.ConstTypePtr{negative_x.AsCTypePtr(), positive_x.AsCTypePtr(), negative_y.AsCTypePtr(), positive_y.AsCTypePtr(), gdc.ConstTypePtr(&deadzone)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&deadzone)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetVector), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) AddJoyMapping(mapping String, update_existing bool) {
	cargs := []gdc.ConstTypePtr{mapping.AsCTypePtr(), gdc.ConstTypePtr(&update_existing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnAddJoyMapping), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) RemoveJoyMapping(guid String) {
	cargs := []gdc.ConstTypePtr{guid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnRemoveJoyMapping), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) IsJoyKnown(device int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&device)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsJoyKnown), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) GetJoyAxis(device int64, axis JoyAxis) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device), gdc.ConstTypePtr(&axis)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&device)
	pinner.Pin(&axis)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetJoyAxis), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) GetJoyName(device int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&device)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetJoyName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) GetJoyGuid(device int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&device)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetJoyGuid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) GetJoyInfo(device int64) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&device)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetJoyInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) ShouldIgnoreDevice(vendor_id int64, product_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vendor_id), gdc.ConstTypePtr(&product_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&vendor_id)
	pinner.Pin(&product_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnShouldIgnoreDevice), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) GetConnectedJoypads() []int {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetConnectedJoypads), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[int](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Input) GetJoyVibrationStrength(device int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&device)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetJoyVibrationStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) GetJoyVibrationDuration(device int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&device)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetJoyVibrationDuration), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) StartJoyVibration(device int64, weak_magnitude float64, strong_magnitude float64, duration float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device), gdc.ConstTypePtr(&weak_magnitude), gdc.ConstTypePtr(&strong_magnitude), gdc.ConstTypePtr(&duration)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnStartJoyVibration), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) StopJoyVibration(device int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnStopJoyVibration), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) VibrateHandheld(duration_ms int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&duration_ms)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnVibrateHandheld), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) GetGravity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) GetAccelerometer() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetAccelerometer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) GetMagnetometer() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetMagnetometer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) GetGyroscope() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetGyroscope), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) SetGravity(value Vector3) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnSetGravity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) SetAccelerometer(value Vector3) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnSetAccelerometer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) SetMagnetometer(value Vector3) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnSetMagnetometer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) SetGyroscope(value Vector3) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnSetGyroscope), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) GetLastMouseVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetLastMouseVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) GetLastMouseScreenVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetLastMouseScreenVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Input) GetMouseButtonMask() MouseButtonMask {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret MouseButtonMask

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetMouseButtonMask), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Input) SetMouseMode(mode InputMouseMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnSetMouseMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) GetMouseMode() InputMouseMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret InputMouseMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetMouseMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Input) WarpMouse(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnWarpMouse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) ActionPress(action StringName, strength float64) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnActionPress), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) ActionRelease(action StringName) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnActionRelease), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) SetDefaultCursorShape(shape InputCursorShape) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnSetDefaultCursorShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) GetCurrentCursorShape() InputCursorShape {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret InputCursorShape

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnGetCurrentCursorShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Input) SetCustomMouseCursor(image Resource, shape InputCursorShape, hotspot Vector2) {
	cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), gdc.ConstTypePtr(&shape), hotspot.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnSetCustomMouseCursor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) ParseInputEvent(event InputEvent) {
	cargs := []gdc.ConstTypePtr{event.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnParseInputEvent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) SetUseAccumulatedInput(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnSetUseAccumulatedInput), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) IsUsingAccumulatedInput() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsUsingAccumulatedInput), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) FlushBufferedEvents() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnFlushBufferedEvents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) SetEmulateMouseFromTouch(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnSetEmulateMouseFromTouch), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) IsEmulatingMouseFromTouch() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsEmulatingMouseFromTouch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Input) SetEmulateTouchFromMouse(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnSetEmulateTouchFromMouse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Input) IsEmulatingTouchFromMouse() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInput.fnIsEmulatingTouchFromMouse), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type InputJoyConnectionChangedSignalFn func(device int, connected bool)

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
