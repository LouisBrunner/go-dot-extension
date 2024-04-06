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

type ptrsForOpenXRActionMapList struct {
	fnSetActionSets              gdc.MethodBindPtr
	fnGetActionSets              gdc.MethodBindPtr
	fnGetActionSetCount          gdc.MethodBindPtr
	fnFindActionSet              gdc.MethodBindPtr
	fnGetActionSet               gdc.MethodBindPtr
	fnAddActionSet               gdc.MethodBindPtr
	fnRemoveActionSet            gdc.MethodBindPtr
	fnSetInteractionProfiles     gdc.MethodBindPtr
	fnGetInteractionProfiles     gdc.MethodBindPtr
	fnGetInteractionProfileCount gdc.MethodBindPtr
	fnFindInteractionProfile     gdc.MethodBindPtr
	fnGetInteractionProfile      gdc.MethodBindPtr
	fnAddInteractionProfile      gdc.MethodBindPtr
	fnRemoveInteractionProfile   gdc.MethodBindPtr
	fnCreateDefaultActionSets    gdc.MethodBindPtr
}

var ptrsForOpenXRActionMap ptrsForOpenXRActionMapList

func initOpenXRActionMapPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRActionMap")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_action_sets")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnSetActionSets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_action_sets")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnGetActionSets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_action_set_count")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnGetActionSetCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("find_action_set")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnFindActionSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1888809267))
	}
	{
		methodName := StringNameFromStr("get_action_set")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnGetActionSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1789580336))
	}
	{
		methodName := StringNameFromStr("add_action_set")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnAddActionSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2093310581))
	}
	{
		methodName := StringNameFromStr("remove_action_set")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnRemoveActionSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2093310581))
	}
	{
		methodName := StringNameFromStr("set_interaction_profiles")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnSetInteractionProfiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_interaction_profiles")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnGetInteractionProfiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_interaction_profile_count")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnGetInteractionProfileCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("find_interaction_profile")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnFindInteractionProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3095875538))
	}
	{
		methodName := StringNameFromStr("get_interaction_profile")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnGetInteractionProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2546151210))
	}
	{
		methodName := StringNameFromStr("add_interaction_profile")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnAddInteractionProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2697953512))
	}
	{
		methodName := StringNameFromStr("remove_interaction_profile")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnRemoveInteractionProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2697953512))
	}
	{
		methodName := StringNameFromStr("create_default_action_sets")
		defer methodName.Destroy()
		ptrsForOpenXRActionMap.fnCreateDefaultActionSets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type OpenXRActionMap struct {
	Resource
}

func (me *OpenXRActionMap) BaseClass() string {
	return "OpenXRActionMap"
}

func NewOpenXRActionMap() *OpenXRActionMap {
	str := StringNameFromStr("OpenXRActionMap") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRActionMap{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OpenXRActionMap) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRActionMap) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRActionMap) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRActionMap) SetActionSets(action_sets Array) {
	cargs := []gdc.ConstTypePtr{action_sets.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnSetActionSets), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRActionMap) GetActionSets() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnGetActionSets), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRActionMap) GetActionSetCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnGetActionSetCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRActionMap) FindActionSet(name String) OpenXRActionSet {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewOpenXRActionSet()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnFindActionSet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRActionMap) GetActionSet(idx int64) OpenXRActionSet {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewOpenXRActionSet()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnGetActionSet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRActionMap) AddActionSet(action_set OpenXRActionSet) {
	cargs := []gdc.ConstTypePtr{action_set.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnAddActionSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRActionMap) RemoveActionSet(action_set OpenXRActionSet) {
	cargs := []gdc.ConstTypePtr{action_set.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnRemoveActionSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRActionMap) SetInteractionProfiles(interaction_profiles Array) {
	cargs := []gdc.ConstTypePtr{interaction_profiles.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnSetInteractionProfiles), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRActionMap) GetInteractionProfiles() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnGetInteractionProfiles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRActionMap) GetInteractionProfileCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnGetInteractionProfileCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRActionMap) FindInteractionProfile(name String) OpenXRInteractionProfile {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewOpenXRInteractionProfile()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnFindInteractionProfile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRActionMap) GetInteractionProfile(idx int64) OpenXRInteractionProfile {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewOpenXRInteractionProfile()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnGetInteractionProfile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRActionMap) AddInteractionProfile(interaction_profile OpenXRInteractionProfile) {
	cargs := []gdc.ConstTypePtr{interaction_profile.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnAddInteractionProfile), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRActionMap) RemoveInteractionProfile(interaction_profile OpenXRInteractionProfile) {
	cargs := []gdc.ConstTypePtr{interaction_profile.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnRemoveInteractionProfile), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRActionMap) CreateDefaultActionSets() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRActionMap.fnCreateDefaultActionSets), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
