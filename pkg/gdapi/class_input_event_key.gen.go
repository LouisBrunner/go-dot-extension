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

type ptrsForInputEventKeyList struct {
	fnSetPressed                      gdc.MethodBindPtr
	fnSetKeycode                      gdc.MethodBindPtr
	fnGetKeycode                      gdc.MethodBindPtr
	fnSetPhysicalKeycode              gdc.MethodBindPtr
	fnGetPhysicalKeycode              gdc.MethodBindPtr
	fnSetKeyLabel                     gdc.MethodBindPtr
	fnGetKeyLabel                     gdc.MethodBindPtr
	fnSetUnicode                      gdc.MethodBindPtr
	fnGetUnicode                      gdc.MethodBindPtr
	fnSetEcho                         gdc.MethodBindPtr
	fnGetKeycodeWithModifiers         gdc.MethodBindPtr
	fnGetPhysicalKeycodeWithModifiers gdc.MethodBindPtr
	fnGetKeyLabelWithModifiers        gdc.MethodBindPtr
	fnAsTextKeycode                   gdc.MethodBindPtr
	fnAsTextPhysicalKeycode           gdc.MethodBindPtr
	fnAsTextKeyLabel                  gdc.MethodBindPtr
}

var ptrsForInputEventKey ptrsForInputEventKeyList

func initInputEventKeyPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventKey")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_pressed")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnSetPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_keycode")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnSetKeycode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 888074362))
	}
	{
		methodName := StringNameFromStr("get_keycode")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnGetKeycode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1585896689))
	}
	{
		methodName := StringNameFromStr("set_physical_keycode")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnSetPhysicalKeycode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 888074362))
	}
	{
		methodName := StringNameFromStr("get_physical_keycode")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnGetPhysicalKeycode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1585896689))
	}
	{
		methodName := StringNameFromStr("set_key_label")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnSetKeyLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 888074362))
	}
	{
		methodName := StringNameFromStr("get_key_label")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnGetKeyLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1585896689))
	}
	{
		methodName := StringNameFromStr("set_unicode")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnSetUnicode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_unicode")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnGetUnicode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_echo")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnSetEcho = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_keycode_with_modifiers")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnGetKeycodeWithModifiers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1585896689))
	}
	{
		methodName := StringNameFromStr("get_physical_keycode_with_modifiers")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnGetPhysicalKeycodeWithModifiers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1585896689))
	}
	{
		methodName := StringNameFromStr("get_key_label_with_modifiers")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnGetKeyLabelWithModifiers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1585896689))
	}
	{
		methodName := StringNameFromStr("as_text_keycode")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnAsTextKeycode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("as_text_physical_keycode")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnAsTextPhysicalKeycode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("as_text_key_label")
		defer methodName.Destroy()
		ptrsForInputEventKey.fnAsTextKeyLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}

}

type InputEventKey struct {
	InputEventWithModifiers
}

func (me *InputEventKey) BaseClass() string {
	return "InputEventKey"
}

func NewInputEventKey() *InputEventKey {
	str := StringNameFromStr("InputEventKey") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventKey{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventKey) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventKey) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventKey) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventKey) SetPressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnSetPressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventKey) SetKeycode(keycode Key) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnSetKeycode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventKey) GetKeycode() Key {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnGetKeycode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *InputEventKey) SetPhysicalKeycode(physical_keycode Key) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&physical_keycode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnSetPhysicalKeycode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventKey) GetPhysicalKeycode() Key {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnGetPhysicalKeycode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *InputEventKey) SetKeyLabel(key_label Key) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&key_label)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnSetKeyLabel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventKey) GetKeyLabel() Key {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnGetKeyLabel), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *InputEventKey) SetUnicode(unicode int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unicode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnSetUnicode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventKey) GetUnicode() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnGetUnicode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventKey) SetEcho(echo bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&echo)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnSetEcho), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventKey) GetKeycodeWithModifiers() Key {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnGetKeycodeWithModifiers), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *InputEventKey) GetPhysicalKeycodeWithModifiers() Key {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnGetPhysicalKeycodeWithModifiers), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *InputEventKey) GetKeyLabelWithModifiers() Key {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnGetKeyLabelWithModifiers), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *InputEventKey) AsTextKeycode() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnAsTextKeycode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventKey) AsTextPhysicalKeycode() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnAsTextPhysicalKeycode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventKey) AsTextKeyLabel() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventKey.fnAsTextKeyLabel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
