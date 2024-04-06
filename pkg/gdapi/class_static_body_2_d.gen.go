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

type ptrsForStaticBody2DList struct {
	fnSetConstantLinearVelocity  gdc.MethodBindPtr
	fnSetConstantAngularVelocity gdc.MethodBindPtr
	fnGetConstantLinearVelocity  gdc.MethodBindPtr
	fnGetConstantAngularVelocity gdc.MethodBindPtr
	fnSetPhysicsMaterialOverride gdc.MethodBindPtr
	fnGetPhysicsMaterialOverride gdc.MethodBindPtr
}

var ptrsForStaticBody2D ptrsForStaticBody2DList

func initStaticBody2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("StaticBody2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_constant_linear_velocity")
		defer methodName.Destroy()
		ptrsForStaticBody2D.fnSetConstantLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_constant_angular_velocity")
		defer methodName.Destroy()
		ptrsForStaticBody2D.fnSetConstantAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_constant_linear_velocity")
		defer methodName.Destroy()
		ptrsForStaticBody2D.fnGetConstantLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_constant_angular_velocity")
		defer methodName.Destroy()
		ptrsForStaticBody2D.fnGetConstantAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_physics_material_override")
		defer methodName.Destroy()
		ptrsForStaticBody2D.fnSetPhysicsMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1784508650))
	}
	{
		methodName := StringNameFromStr("get_physics_material_override")
		defer methodName.Destroy()
		ptrsForStaticBody2D.fnGetPhysicsMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2521850424))
	}

}

type StaticBody2D struct {
	PhysicsBody2D
}

func (me *StaticBody2D) BaseClass() string {
	return "StaticBody2D"
}

func NewStaticBody2D() *StaticBody2D {
	str := StringNameFromStr("StaticBody2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &StaticBody2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *StaticBody2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *StaticBody2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *StaticBody2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *StaticBody2D) SetConstantLinearVelocity(vel Vector2) {
	cargs := []gdc.ConstTypePtr{vel.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody2D.fnSetConstantLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StaticBody2D) SetConstantAngularVelocity(vel float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody2D.fnSetConstantAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StaticBody2D) GetConstantLinearVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody2D.fnGetConstantLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StaticBody2D) GetConstantAngularVelocity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody2D.fnGetConstantAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *StaticBody2D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial) {
	cargs := []gdc.ConstTypePtr{physics_material_override.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody2D.fnSetPhysicsMaterialOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StaticBody2D) GetPhysicsMaterialOverride() PhysicsMaterial {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStaticBody2D.fnGetPhysicsMaterialOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
