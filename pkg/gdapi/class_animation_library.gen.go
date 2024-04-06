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

type ptrsForAnimationLibraryList struct {
	fnAddAnimation     gdc.MethodBindPtr
	fnRemoveAnimation  gdc.MethodBindPtr
	fnRenameAnimation  gdc.MethodBindPtr
	fnHasAnimation     gdc.MethodBindPtr
	fnGetAnimation     gdc.MethodBindPtr
	fnGetAnimationList gdc.MethodBindPtr
}

var ptrsForAnimationLibrary ptrsForAnimationLibraryList

func initAnimationLibraryPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationLibrary")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_animation")
		defer methodName.Destroy()
		ptrsForAnimationLibrary.fnAddAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1811855551))
	}
	{
		methodName := StringNameFromStr("remove_animation")
		defer methodName.Destroy()
		ptrsForAnimationLibrary.fnRemoveAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("rename_animation")
		defer methodName.Destroy()
		ptrsForAnimationLibrary.fnRenameAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("has_animation")
		defer methodName.Destroy()
		ptrsForAnimationLibrary.fnHasAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("get_animation")
		defer methodName.Destroy()
		ptrsForAnimationLibrary.fnGetAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2933122410))
	}
	{
		methodName := StringNameFromStr("get_animation_list")
		defer methodName.Destroy()
		ptrsForAnimationLibrary.fnGetAnimationList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}

}

type AnimationLibrary struct {
	Resource
}

func (me *AnimationLibrary) BaseClass() string {
	return "AnimationLibrary"
}

func NewAnimationLibrary() *AnimationLibrary {
	str := StringNameFromStr("AnimationLibrary") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationLibrary{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AnimationLibrary) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationLibrary) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationLibrary) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimationLibrary) AddAnimation(name StringName, animation Animation) Error {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), animation.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationLibrary.fnAddAnimation), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationLibrary) RemoveAnimation(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationLibrary.fnRemoveAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationLibrary) RenameAnimation(name StringName, newname StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), newname.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationLibrary.fnRenameAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationLibrary) HasAnimation(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationLibrary.fnHasAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationLibrary) GetAnimation(name StringName) Animation {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAnimation()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationLibrary.fnGetAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationLibrary) GetAnimationList() []StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationLibrary.fnGetAnimationList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[StringName](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Signals

type AnimationLibraryAnimationAddedSignalFn func(name StringName)

func (me *AnimationLibrary) ConnectAnimationAdded(subs SignalSubscribers, fn AnimationLibraryAnimationAddedSignalFn) {
	sig := StringNameFromStr("animation_added")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationLibrary) DisconnectAnimationAdded(subs SignalSubscribers, fn AnimationLibraryAnimationAddedSignalFn) {
	sig := StringNameFromStr("animation_added")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationLibraryAnimationRemovedSignalFn func(name StringName)

func (me *AnimationLibrary) ConnectAnimationRemoved(subs SignalSubscribers, fn AnimationLibraryAnimationRemovedSignalFn) {
	sig := StringNameFromStr("animation_removed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationLibrary) DisconnectAnimationRemoved(subs SignalSubscribers, fn AnimationLibraryAnimationRemovedSignalFn) {
	sig := StringNameFromStr("animation_removed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationLibraryAnimationRenamedSignalFn func(name StringName, to_name StringName)

func (me *AnimationLibrary) ConnectAnimationRenamed(subs SignalSubscribers, fn AnimationLibraryAnimationRenamedSignalFn) {
	sig := StringNameFromStr("animation_renamed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationLibrary) DisconnectAnimationRenamed(subs SignalSubscribers, fn AnimationLibraryAnimationRenamedSignalFn) {
	sig := StringNameFromStr("animation_renamed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationLibraryAnimationChangedSignalFn func(name StringName)

func (me *AnimationLibrary) ConnectAnimationChanged(subs SignalSubscribers, fn AnimationLibraryAnimationChangedSignalFn) {
	sig := StringNameFromStr("animation_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationLibrary) DisconnectAnimationChanged(subs SignalSubscribers, fn AnimationLibraryAnimationChangedSignalFn) {
	sig := StringNameFromStr("animation_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
