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

type ptrsForEditorResourcePickerList struct {
	fnXSetCreateOptions   gdc.MethodBindPtr
	fnXHandleMenuSelected gdc.MethodBindPtr
	fnSetBaseType         gdc.MethodBindPtr
	fnGetBaseType         gdc.MethodBindPtr
	fnGetAllowedTypes     gdc.MethodBindPtr
	fnSetEditedResource   gdc.MethodBindPtr
	fnGetEditedResource   gdc.MethodBindPtr
	fnSetToggleMode       gdc.MethodBindPtr
	fnIsToggleMode        gdc.MethodBindPtr
	fnSetTogglePressed    gdc.MethodBindPtr
	fnSetEditable         gdc.MethodBindPtr
	fnIsEditable          gdc.MethodBindPtr
}

var ptrsForEditorResourcePicker ptrsForEditorResourcePickerList

func initEditorResourcePickerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorResourcePicker")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_base_type")
		defer methodName.Destroy()
		ptrsForEditorResourcePicker.fnSetBaseType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_base_type")
		defer methodName.Destroy()
		ptrsForEditorResourcePicker.fnGetBaseType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_allowed_types")
		defer methodName.Destroy()
		ptrsForEditorResourcePicker.fnGetAllowedTypes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_edited_resource")
		defer methodName.Destroy()
		ptrsForEditorResourcePicker.fnSetEditedResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 968641751))
	}
	{
		methodName := StringNameFromStr("get_edited_resource")
		defer methodName.Destroy()
		ptrsForEditorResourcePicker.fnGetEditedResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2674603643))
	}
	{
		methodName := StringNameFromStr("set_toggle_mode")
		defer methodName.Destroy()
		ptrsForEditorResourcePicker.fnSetToggleMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_toggle_mode")
		defer methodName.Destroy()
		ptrsForEditorResourcePicker.fnIsToggleMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_toggle_pressed")
		defer methodName.Destroy()
		ptrsForEditorResourcePicker.fnSetTogglePressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_editable")
		defer methodName.Destroy()
		ptrsForEditorResourcePicker.fnSetEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_editable")
		defer methodName.Destroy()
		ptrsForEditorResourcePicker.fnIsEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type EditorResourcePicker struct {
	HBoxContainer
}

func (me *EditorResourcePicker) BaseClass() string {
	return "EditorResourcePicker"
}

func NewEditorResourcePicker() *EditorResourcePicker {
	str := StringNameFromStr("EditorResourcePicker") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorResourcePicker{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorResourcePicker) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorResourcePicker) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorResourcePicker) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorResourcePicker) SetBaseType(base_type String) {
	cargs := []gdc.ConstTypePtr{base_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePicker.fnSetBaseType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorResourcePicker) GetBaseType() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePicker.fnGetBaseType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorResourcePicker) GetAllowedTypes() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePicker.fnGetAllowedTypes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorResourcePicker) SetEditedResource(resource Resource) {
	cargs := []gdc.ConstTypePtr{resource.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePicker.fnSetEditedResource), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorResourcePicker) GetEditedResource() Resource {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewResource()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePicker.fnGetEditedResource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorResourcePicker) SetToggleMode(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePicker.fnSetToggleMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorResourcePicker) IsToggleMode() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePicker.fnIsToggleMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorResourcePicker) SetTogglePressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePicker.fnSetTogglePressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorResourcePicker) SetEditable(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePicker.fnSetEditable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorResourcePicker) IsEditable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePicker.fnIsEditable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type EditorResourcePickerResourceSelectedSignalFn func(resource Resource, inspect bool)

func (me *EditorResourcePicker) ConnectResourceSelected(subs SignalSubscribers, fn EditorResourcePickerResourceSelectedSignalFn) {
	sig := StringNameFromStr("resource_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorResourcePicker) DisconnectResourceSelected(subs SignalSubscribers, fn EditorResourcePickerResourceSelectedSignalFn) {
	sig := StringNameFromStr("resource_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type EditorResourcePickerResourceChangedSignalFn func(resource Resource)

func (me *EditorResourcePicker) ConnectResourceChanged(subs SignalSubscribers, fn EditorResourcePickerResourceChangedSignalFn) {
	sig := StringNameFromStr("resource_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorResourcePicker) DisconnectResourceChanged(subs SignalSubscribers, fn EditorResourcePickerResourceChangedSignalFn) {
	sig := StringNameFromStr("resource_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
