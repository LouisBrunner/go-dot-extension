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

type ptrsForPhysicalBoneSimulator3DList struct {
	fnIsSimulatingPhysics                   gdc.MethodBindPtr
	fnPhysicalBonesStopSimulation           gdc.MethodBindPtr
	fnPhysicalBonesStartSimulation          gdc.MethodBindPtr
	fnPhysicalBonesAddCollisionException    gdc.MethodBindPtr
	fnPhysicalBonesRemoveCollisionException gdc.MethodBindPtr
}

var ptrsForPhysicalBoneSimulator3D ptrsForPhysicalBoneSimulator3DList

func initPhysicalBoneSimulator3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicalBoneSimulator3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("is_simulating_physics")
		defer methodName.Destroy()
		ptrsForPhysicalBoneSimulator3D.fnIsSimulatingPhysics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("physical_bones_stop_simulation")
		defer methodName.Destroy()
		ptrsForPhysicalBoneSimulator3D.fnPhysicalBonesStopSimulation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("physical_bones_start_simulation")
		defer methodName.Destroy()
		ptrsForPhysicalBoneSimulator3D.fnPhysicalBonesStartSimulation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2787316981))
	}
	{
		methodName := StringNameFromStr("physical_bones_add_collision_exception")
		defer methodName.Destroy()
		ptrsForPhysicalBoneSimulator3D.fnPhysicalBonesAddCollisionException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("physical_bones_remove_collision_exception")
		defer methodName.Destroy()
		ptrsForPhysicalBoneSimulator3D.fnPhysicalBonesRemoveCollisionException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}

}

type PhysicalBoneSimulator3D struct {
	SkeletonModifier3D
}

func (me *PhysicalBoneSimulator3D) BaseClass() string {
	return "PhysicalBoneSimulator3D"
}

func NewPhysicalBoneSimulator3D() *PhysicalBoneSimulator3D {
	str := StringNameFromStr("PhysicalBoneSimulator3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicalBoneSimulator3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicalBoneSimulator3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicalBoneSimulator3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicalBoneSimulator3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicalBoneSimulator3D) IsSimulatingPhysics() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBoneSimulator3D.fnIsSimulatingPhysics), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBoneSimulator3D) PhysicalBonesStopSimulation() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBoneSimulator3D.fnPhysicalBonesStopSimulation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBoneSimulator3D) PhysicalBonesStartSimulation(bones []StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bones)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBoneSimulator3D.fnPhysicalBonesStartSimulation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBoneSimulator3D) PhysicalBonesAddCollisionException(exception RID) {
	cargs := []gdc.ConstTypePtr{exception.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBoneSimulator3D.fnPhysicalBonesAddCollisionException), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBoneSimulator3D) PhysicalBonesRemoveCollisionException(exception RID) {
	cargs := []gdc.ConstTypePtr{exception.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBoneSimulator3D.fnPhysicalBonesRemoveCollisionException), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
