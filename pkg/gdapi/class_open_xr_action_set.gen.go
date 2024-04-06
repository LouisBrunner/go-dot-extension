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

type ptrsForOpenXRActionSetList struct {
	fnSetLocalizedName gdc.MethodBindPtr
	fnGetLocalizedName gdc.MethodBindPtr
	fnSetPriority      gdc.MethodBindPtr
	fnGetPriority      gdc.MethodBindPtr
	fnGetActionCount   gdc.MethodBindPtr
	fnSetActions       gdc.MethodBindPtr
	fnGetActions       gdc.MethodBindPtr
	fnAddAction        gdc.MethodBindPtr
	fnRemoveAction     gdc.MethodBindPtr
}

var ptrsForOpenXRActionSet ptrsForOpenXRActionSetList

func initOpenXRActionSetPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRActionSet")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_localized_name")
		defer methodName.Destroy()
		ptrsForOpenXRActionSet.fnSetLocalizedName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_localized_name")
		defer methodName.Destroy()
		ptrsForOpenXRActionSet.fnGetLocalizedName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_priority")
		defer methodName.Destroy()
		ptrsForOpenXRActionSet.fnSetPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_priority")
		defer methodName.Destroy()
		ptrsForOpenXRActionSet.fnGetPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_action_count")
		defer methodName.Destroy()
		ptrsForOpenXRActionSet.fnGetActionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_actions")
		defer methodName.Destroy()
		ptrsForOpenXRActionSet.fnSetActions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_actions")
		defer methodName.Destroy()
		ptrsForOpenXRActionSet.fnGetActions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("add_action")
		defer methodName.Destroy()
		ptrsForOpenXRActionSet.fnAddAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 349361333))
	}
	{
		methodName := StringNameFromStr("remove_action")
		defer methodName.Destroy()
		ptrsForOpenXRActionSet.fnRemoveAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 349361333))
	}

}

type OpenXRActionSet struct {
	Resource
}

func (me *OpenXRActionSet) BaseClass() string {
	return "OpenXRActionSet"
}

func NewOpenXRActionSet() *OpenXRActionSet {
	str := StringNameFromStr("OpenXRActionSet") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRActionSet{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OpenXRActionSet) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRActionSet) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRActionSet) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRActionSet) SetLocalizedName(localized_name String) {
	cargs := []gdc.ConstTypePtr{localized_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionSet.fnSetLocalizedName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRActionSet) GetLocalizedName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionSet.fnGetLocalizedName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRActionSet) SetPriority(priority int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionSet.fnSetPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRActionSet) GetPriority() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionSet.fnGetPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRActionSet) GetActionCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionSet.fnGetActionCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRActionSet) SetActions(actions Array) {
	cargs := []gdc.ConstTypePtr{actions.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionSet.fnSetActions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRActionSet) GetActions() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionSet.fnGetActions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRActionSet) AddAction(action OpenXRAction) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionSet.fnAddAction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRActionSet) RemoveAction(action OpenXRAction) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionSet.fnRemoveAction), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
