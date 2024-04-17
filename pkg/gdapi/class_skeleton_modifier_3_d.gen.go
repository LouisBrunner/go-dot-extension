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

type ptrsForSkeletonModifier3DList struct {
	fnSetActive    gdc.MethodBindPtr
	fnIsActive     gdc.MethodBindPtr
	fnSetInfluence gdc.MethodBindPtr
	fnGetInfluence gdc.MethodBindPtr
}

var ptrsForSkeletonModifier3D ptrsForSkeletonModifier3DList

func initSkeletonModifier3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SkeletonModifier3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_active")
		defer methodName.Destroy()
		ptrsForSkeletonModifier3D.fnSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_active")
		defer methodName.Destroy()
		ptrsForSkeletonModifier3D.fnIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_influence")
		defer methodName.Destroy()
		ptrsForSkeletonModifier3D.fnSetInfluence = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_influence")
		defer methodName.Destroy()
		ptrsForSkeletonModifier3D.fnGetInfluence = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type SkeletonModifier3D struct {
	Node3D
}

func (me *SkeletonModifier3D) BaseClass() string {
	return "SkeletonModifier3D"
}

func NewSkeletonModifier3D() *SkeletonModifier3D {
	str := StringNameFromStr("SkeletonModifier3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SkeletonModifier3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SkeletonModifier3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SkeletonModifier3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SkeletonModifier3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SkeletonModifier3D) SetActive(active bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModifier3D.fnSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModifier3D) IsActive() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModifier3D.fnIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModifier3D) SetInfluence(influence float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&influence)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModifier3D.fnSetInfluence), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModifier3D) GetInfluence() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModifier3D.fnGetInfluence), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SkeletonModifier3DModificationProcessedSignalFn func()

func (me *SkeletonModifier3D) ConnectModificationProcessed(subs SignalSubscribers, fn SkeletonModifier3DModificationProcessedSignalFn) {
	sig := StringNameFromStr("modification_processed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *SkeletonModifier3D) DisconnectModificationProcessed(subs SignalSubscribers, fn SkeletonModifier3DModificationProcessedSignalFn) {
	sig := StringNameFromStr("modification_processed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
