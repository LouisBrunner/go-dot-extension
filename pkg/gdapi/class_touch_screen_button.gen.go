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

type ptrsForTouchScreenButtonList struct {
	fnSetTextureNormal     gdc.MethodBindPtr
	fnGetTextureNormal     gdc.MethodBindPtr
	fnSetTexturePressed    gdc.MethodBindPtr
	fnGetTexturePressed    gdc.MethodBindPtr
	fnSetBitmask           gdc.MethodBindPtr
	fnGetBitmask           gdc.MethodBindPtr
	fnSetShape             gdc.MethodBindPtr
	fnGetShape             gdc.MethodBindPtr
	fnSetShapeCentered     gdc.MethodBindPtr
	fnIsShapeCentered      gdc.MethodBindPtr
	fnSetShapeVisible      gdc.MethodBindPtr
	fnIsShapeVisible       gdc.MethodBindPtr
	fnSetAction            gdc.MethodBindPtr
	fnGetAction            gdc.MethodBindPtr
	fnSetVisibilityMode    gdc.MethodBindPtr
	fnGetVisibilityMode    gdc.MethodBindPtr
	fnSetPassbyPress       gdc.MethodBindPtr
	fnIsPassbyPressEnabled gdc.MethodBindPtr
	fnIsPressed            gdc.MethodBindPtr
}

var ptrsForTouchScreenButton ptrsForTouchScreenButtonList

func initTouchScreenButtonPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TouchScreenButton")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture_normal")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnSetTextureNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture_normal")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnGetTextureNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_texture_pressed")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnSetTexturePressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture_pressed")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnGetTexturePressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_bitmask")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnSetBitmask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 698588216))
	}
	{
		methodName := StringNameFromStr("get_bitmask")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnGetBitmask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2459671998))
	}
	{
		methodName := StringNameFromStr("set_shape")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 771364740))
	}
	{
		methodName := StringNameFromStr("get_shape")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 522005891))
	}
	{
		methodName := StringNameFromStr("set_shape_centered")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnSetShapeCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_shape_centered")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnIsShapeCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_shape_visible")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnSetShapeVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_shape_visible")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnIsShapeVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_action")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnSetAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_action")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnGetAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_visibility_mode")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnSetVisibilityMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3031128463))
	}
	{
		methodName := StringNameFromStr("get_visibility_mode")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnGetVisibilityMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2558996468))
	}
	{
		methodName := StringNameFromStr("set_passby_press")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnSetPassbyPress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_passby_press_enabled")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnIsPassbyPressEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_pressed")
		defer methodName.Destroy()
		ptrsForTouchScreenButton.fnIsPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type TouchScreenButton struct {
	Node2D
}

func (me *TouchScreenButton) BaseClass() string {
	return "TouchScreenButton"
}

func NewTouchScreenButton() *TouchScreenButton {
	str := StringNameFromStr("TouchScreenButton") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TouchScreenButton{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TouchScreenButtonVisibilityMode int

const (
	TouchScreenButtonVisibilityModeVisibilityAlways          TouchScreenButtonVisibilityMode = 0
	TouchScreenButtonVisibilityModeVisibilityTouchscreenOnly TouchScreenButtonVisibilityMode = 1
)

func (me *TouchScreenButton) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TouchScreenButton) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TouchScreenButton) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TouchScreenButton) SetTextureNormal(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnSetTextureNormal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TouchScreenButton) GetTextureNormal() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnGetTextureNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TouchScreenButton) SetTexturePressed(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnSetTexturePressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TouchScreenButton) GetTexturePressed() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnGetTexturePressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TouchScreenButton) SetBitmask(bitmask BitMap) {
	cargs := []gdc.ConstTypePtr{bitmask.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnSetBitmask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TouchScreenButton) GetBitmask() BitMap {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBitMap()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnGetBitmask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TouchScreenButton) SetShape(shape Shape2D) {
	cargs := []gdc.ConstTypePtr{shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TouchScreenButton) GetShape() Shape2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewShape2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TouchScreenButton) SetShapeCentered(bool bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bool)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnSetShapeCentered), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TouchScreenButton) IsShapeCentered() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnIsShapeCentered), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TouchScreenButton) SetShapeVisible(bool bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bool)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnSetShapeVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TouchScreenButton) IsShapeVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnIsShapeVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TouchScreenButton) SetAction(action String) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnSetAction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TouchScreenButton) GetAction() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnGetAction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TouchScreenButton) SetVisibilityMode(mode TouchScreenButtonVisibilityMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnSetVisibilityMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TouchScreenButton) GetVisibilityMode() TouchScreenButtonVisibilityMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TouchScreenButtonVisibilityMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnGetVisibilityMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TouchScreenButton) SetPassbyPress(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnSetPassbyPress), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TouchScreenButton) IsPassbyPressEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnIsPassbyPressEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TouchScreenButton) IsPressed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTouchScreenButton.fnIsPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TouchScreenButtonPressedSignalFn func()

func (me *TouchScreenButton) ConnectPressed(subs SignalSubscribers, fn TouchScreenButtonPressedSignalFn) {
	sig := StringNameFromStr("pressed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TouchScreenButton) DisconnectPressed(subs SignalSubscribers, fn TouchScreenButtonPressedSignalFn) {
	sig := StringNameFromStr("pressed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TouchScreenButtonReleasedSignalFn func()

func (me *TouchScreenButton) ConnectReleased(subs SignalSubscribers, fn TouchScreenButtonReleasedSignalFn) {
	sig := StringNameFromStr("released")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TouchScreenButton) DisconnectReleased(subs SignalSubscribers, fn TouchScreenButtonReleasedSignalFn) {
	sig := StringNameFromStr("released")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
