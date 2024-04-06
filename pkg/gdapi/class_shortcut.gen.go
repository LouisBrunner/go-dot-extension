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

type ptrsForShortcutList struct {
	fnSetEvents     gdc.MethodBindPtr
	fnGetEvents     gdc.MethodBindPtr
	fnHasValidEvent gdc.MethodBindPtr
	fnMatchesEvent  gdc.MethodBindPtr
	fnGetAsText     gdc.MethodBindPtr
}

var ptrsForShortcut ptrsForShortcutList

func initShortcutPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Shortcut")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_events")
		defer methodName.Destroy()
		ptrsForShortcut.fnSetEvents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_events")
		defer methodName.Destroy()
		ptrsForShortcut.fnGetEvents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("has_valid_event")
		defer methodName.Destroy()
		ptrsForShortcut.fnHasValidEvent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("matches_event")
		defer methodName.Destroy()
		ptrsForShortcut.fnMatchesEvent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3738334489))
	}
	{
		methodName := StringNameFromStr("get_as_text")
		defer methodName.Destroy()
		ptrsForShortcut.fnGetAsText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}

}

type Shortcut struct {
	Resource
}

func (me *Shortcut) BaseClass() string {
	return "Shortcut"
}

func NewShortcut() *Shortcut {
	str := StringNameFromStr("Shortcut") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Shortcut{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Shortcut) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Shortcut) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Shortcut) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Shortcut) SetEvents(events Array) {
	cargs := []gdc.ConstTypePtr{events.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShortcut.fnSetEvents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Shortcut) GetEvents() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShortcut.fnGetEvents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Shortcut) HasValidEvent() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShortcut.fnHasValidEvent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Shortcut) MatchesEvent(event InputEvent) bool {
	cargs := []gdc.ConstTypePtr{event.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShortcut.fnMatchesEvent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Shortcut) GetAsText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShortcut.fnGetAsText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
