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

type ptrsForStaticBody3DList struct {
	fnSetConstantLinearVelocity  gdc.MethodBindPtr
	fnSetConstantAngularVelocity gdc.MethodBindPtr
	fnGetConstantLinearVelocity  gdc.MethodBindPtr
	fnGetConstantAngularVelocity gdc.MethodBindPtr
	fnSetPhysicsMaterialOverride gdc.MethodBindPtr
	fnGetPhysicsMaterialOverride gdc.MethodBindPtr
}

var ptrsForStaticBody3D ptrsForStaticBody3DList

func initStaticBody3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("StaticBody3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_constant_linear_velocity")
		defer methodName.Destroy()
		ptrsForStaticBody3D.fnSetConstantLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("set_constant_angular_velocity")
		defer methodName.Destroy()
		ptrsForStaticBody3D.fnSetConstantAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_constant_linear_velocity")
		defer methodName.Destroy()
		ptrsForStaticBody3D.fnGetConstantLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_constant_angular_velocity")
		defer methodName.Destroy()
		ptrsForStaticBody3D.fnGetConstantAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_physics_material_override")
		defer methodName.Destroy()
		ptrsForStaticBody3D.fnSetPhysicsMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1784508650))
	}
	{
		methodName := StringNameFromStr("get_physics_material_override")
		defer methodName.Destroy()
		ptrsForStaticBody3D.fnGetPhysicsMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2521850424))
	}
}

type StaticBody3D struct {
	PhysicsBody3D
}

func (me *StaticBody3D) BaseClass() string {
	return "StaticBody3D"
}

func NewStaticBody3D() *StaticBody3D {
	str := StringNameFromStr("StaticBody3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &StaticBody3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *StaticBody3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *StaticBody3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *StaticBody3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *StaticBody3D) SetConstantLinearVelocity(vel Vector3) {
	cargs := []gdc.ConstTypePtr{vel.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody3D.fnSetConstantLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StaticBody3D) SetConstantAngularVelocity(vel Vector3) {
	cargs := []gdc.ConstTypePtr{vel.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody3D.fnSetConstantAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StaticBody3D) GetConstantLinearVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody3D.fnGetConstantLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StaticBody3D) GetConstantAngularVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody3D.fnGetConstantAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StaticBody3D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial) {
	cargs := []gdc.ConstTypePtr{physics_material_override.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody3D.fnSetPhysicsMaterialOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StaticBody3D) GetPhysicsMaterialOverride() PhysicsMaterial {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody3D.fnGetPhysicsMaterialOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
