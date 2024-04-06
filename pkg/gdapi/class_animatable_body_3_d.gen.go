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

type ptrsForAnimatableBody3DList struct {
	fnSetSyncToPhysics       gdc.MethodBindPtr
	fnIsSyncToPhysicsEnabled gdc.MethodBindPtr
}

var ptrsForAnimatableBody3D ptrsForAnimatableBody3DList

func initAnimatableBody3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimatableBody3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_sync_to_physics")
		defer methodName.Destroy()
		ptrsForAnimatableBody3D.fnSetSyncToPhysics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_sync_to_physics_enabled")
		defer methodName.Destroy()
		ptrsForAnimatableBody3D.fnIsSyncToPhysicsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type AnimatableBody3D struct {
	StaticBody3D
}

func (me *AnimatableBody3D) BaseClass() string {
	return "AnimatableBody3D"
}

func NewAnimatableBody3D() *AnimatableBody3D {
	str := StringNameFromStr("AnimatableBody3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimatableBody3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AnimatableBody3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimatableBody3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimatableBody3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimatableBody3D) SetSyncToPhysics(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatableBody3D.fnSetSyncToPhysics), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatableBody3D) IsSyncToPhysicsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatableBody3D.fnIsSyncToPhysicsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
