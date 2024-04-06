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

type ptrsForEditorScriptPickerList struct {
	fnSetScriptOwner gdc.MethodBindPtr
	fnGetScriptOwner gdc.MethodBindPtr
}

var ptrsForEditorScriptPicker ptrsForEditorScriptPickerList

func initEditorScriptPickerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorScriptPicker")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_script_owner")
		defer methodName.Destroy()
		ptrsForEditorScriptPicker.fnSetScriptOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("get_script_owner")
		defer methodName.Destroy()
		ptrsForEditorScriptPicker.fnGetScriptOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
	}
}

type EditorScriptPicker struct {
	EditorResourcePicker
}

func (me *EditorScriptPicker) BaseClass() string {
	return "EditorScriptPicker"
}

func NewEditorScriptPicker() *EditorScriptPicker {
	str := StringNameFromStr("EditorScriptPicker") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorScriptPicker{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorScriptPicker) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorScriptPicker) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorScriptPicker) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorScriptPicker) SetScriptOwner(owner_node Node) {
	cargs := []gdc.ConstTypePtr{owner_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorScriptPicker.fnSetScriptOwner), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorScriptPicker) GetScriptOwner() Node {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorScriptPicker.fnGetScriptOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
