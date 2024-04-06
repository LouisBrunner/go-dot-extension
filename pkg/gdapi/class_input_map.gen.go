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

type ptrsForInputMapList struct {
	fnHasAction               gdc.MethodBindPtr
	fnGetActions              gdc.MethodBindPtr
	fnAddAction               gdc.MethodBindPtr
	fnEraseAction             gdc.MethodBindPtr
	fnActionSetDeadzone       gdc.MethodBindPtr
	fnActionGetDeadzone       gdc.MethodBindPtr
	fnActionAddEvent          gdc.MethodBindPtr
	fnActionHasEvent          gdc.MethodBindPtr
	fnActionEraseEvent        gdc.MethodBindPtr
	fnActionEraseEvents       gdc.MethodBindPtr
	fnActionGetEvents         gdc.MethodBindPtr
	fnEventIsAction           gdc.MethodBindPtr
	fnLoadFromProjectSettings gdc.MethodBindPtr
}

var ptrsForInputMap ptrsForInputMapList

func initInputMapPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputMap")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("has_action")
		defer methodName.Destroy()
		ptrsForInputMap.fnHasAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("get_actions")
		defer methodName.Destroy()
		ptrsForInputMap.fnGetActions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("add_action")
		defer methodName.Destroy()
		ptrsForInputMap.fnAddAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4100757082))
	}
	{
		methodName := StringNameFromStr("erase_action")
		defer methodName.Destroy()
		ptrsForInputMap.fnEraseAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("action_set_deadzone")
		defer methodName.Destroy()
		ptrsForInputMap.fnActionSetDeadzone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4135858297))
	}
	{
		methodName := StringNameFromStr("action_get_deadzone")
		defer methodName.Destroy()
		ptrsForInputMap.fnActionGetDeadzone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1391627649))
	}
	{
		methodName := StringNameFromStr("action_add_event")
		defer methodName.Destroy()
		ptrsForInputMap.fnActionAddEvent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 518302593))
	}
	{
		methodName := StringNameFromStr("action_has_event")
		defer methodName.Destroy()
		ptrsForInputMap.fnActionHasEvent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1185871985))
	}
	{
		methodName := StringNameFromStr("action_erase_event")
		defer methodName.Destroy()
		ptrsForInputMap.fnActionEraseEvent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 518302593))
	}
	{
		methodName := StringNameFromStr("action_erase_events")
		defer methodName.Destroy()
		ptrsForInputMap.fnActionEraseEvents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("action_get_events")
		defer methodName.Destroy()
		ptrsForInputMap.fnActionGetEvents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 689397652))
	}
	{
		methodName := StringNameFromStr("event_is_action")
		defer methodName.Destroy()
		ptrsForInputMap.fnEventIsAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3193353650))
	}
	{
		methodName := StringNameFromStr("load_from_project_settings")
		defer methodName.Destroy()
		ptrsForInputMap.fnLoadFromProjectSettings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type InputMap struct {
	Object
}

func (me *InputMap) BaseClass() string {
	return "InputMap"
}

func NewInputMap() *InputMap {
	str := StringNameFromStr("InputMap") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputMap{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputMap) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputMap) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputMap) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputMap) HasAction(action StringName) bool {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnHasAction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputMap) GetActions() []StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnGetActions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[StringName](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *InputMap) AddAction(action StringName, deadzone float64) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&deadzone)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnAddAction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputMap) EraseAction(action StringName) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnEraseAction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputMap) ActionSetDeadzone(action StringName, deadzone float64) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&deadzone)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnActionSetDeadzone), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputMap) ActionGetDeadzone(action StringName) float64 {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnActionGetDeadzone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputMap) ActionAddEvent(action StringName, event InputEvent) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), event.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnActionAddEvent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputMap) ActionHasEvent(action StringName, event InputEvent) bool {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), event.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnActionHasEvent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputMap) ActionEraseEvent(action StringName, event InputEvent) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), event.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnActionEraseEvent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputMap) ActionEraseEvents(action StringName) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnActionEraseEvents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputMap) ActionGetEvents(action StringName) []InputEvent {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnActionGetEvents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[InputEvent](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *InputMap) EventIsAction(event InputEvent, action StringName, exact_match bool) bool {
	cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnEventIsAction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputMap) LoadFromProjectSettings() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputMap.fnLoadFromProjectSettings), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
