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

type ptrsForOpenXRActionList struct {
	fnSetLocalizedName gdc.MethodBindPtr
	fnGetLocalizedName gdc.MethodBindPtr
	fnSetActionType    gdc.MethodBindPtr
	fnGetActionType    gdc.MethodBindPtr
	fnSetToplevelPaths gdc.MethodBindPtr
	fnGetToplevelPaths gdc.MethodBindPtr
}

var ptrsForOpenXRAction ptrsForOpenXRActionList

func initOpenXRActionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRAction")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_localized_name")
		defer methodName.Destroy()
		ptrsForOpenXRAction.fnSetLocalizedName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_localized_name")
		defer methodName.Destroy()
		ptrsForOpenXRAction.fnGetLocalizedName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_action_type")
		defer methodName.Destroy()
		ptrsForOpenXRAction.fnSetActionType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1675238366))
	}
	{
		methodName := StringNameFromStr("get_action_type")
		defer methodName.Destroy()
		ptrsForOpenXRAction.fnGetActionType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536542431))
	}
	{
		methodName := StringNameFromStr("set_toplevel_paths")
		defer methodName.Destroy()
		ptrsForOpenXRAction.fnSetToplevelPaths = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
	}
	{
		methodName := StringNameFromStr("get_toplevel_paths")
		defer methodName.Destroy()
		ptrsForOpenXRAction.fnGetToplevelPaths = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
}

type OpenXRAction struct {
	Resource
}

func (me *OpenXRAction) BaseClass() string {
	return "OpenXRAction"
}

func NewOpenXRAction() *OpenXRAction {
	str := StringNameFromStr("OpenXRAction") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRAction{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type OpenXRActionActionType int

const (
	OpenXRActionActionTypeOpenxrActionBool    OpenXRActionActionType = 0
	OpenXRActionActionTypeOpenxrActionFloat   OpenXRActionActionType = 1
	OpenXRActionActionTypeOpenxrActionVector2 OpenXRActionActionType = 2
	OpenXRActionActionTypeOpenxrActionPose    OpenXRActionActionType = 3
)

func (me *OpenXRAction) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRAction) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRAction) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRAction) SetLocalizedName(localized_name String) {
	cargs := []gdc.ConstTypePtr{localized_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAction.fnSetLocalizedName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRAction) GetLocalizedName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAction.fnGetLocalizedName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRAction) SetActionType(action_type OpenXRActionActionType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&action_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAction.fnSetActionType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRAction) GetActionType() OpenXRActionActionType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret OpenXRActionActionType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAction.fnGetActionType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *OpenXRAction) SetToplevelPaths(toplevel_paths PackedStringArray) {
	cargs := []gdc.ConstTypePtr{toplevel_paths.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAction.fnSetToplevelPaths), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRAction) GetToplevelPaths() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRAction.fnGetToplevelPaths), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
