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

type ptrsForAnimatableBody2DList struct {
	fnSetSyncToPhysics       gdc.MethodBindPtr
	fnIsSyncToPhysicsEnabled gdc.MethodBindPtr
}

var ptrsForAnimatableBody2D ptrsForAnimatableBody2DList

func initAnimatableBody2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimatableBody2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_sync_to_physics")
		defer methodName.Destroy()
		ptrsForAnimatableBody2D.fnSetSyncToPhysics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_sync_to_physics_enabled")
		defer methodName.Destroy()
		ptrsForAnimatableBody2D.fnIsSyncToPhysicsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type AnimatableBody2D struct {
	StaticBody2D
}

func (me *AnimatableBody2D) BaseClass() string {
	return "AnimatableBody2D"
}

func NewAnimatableBody2D() *AnimatableBody2D {
	str := StringNameFromStr("AnimatableBody2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimatableBody2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AnimatableBody2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimatableBody2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimatableBody2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimatableBody2D) SetSyncToPhysics(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatableBody2D.fnSetSyncToPhysics), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatableBody2D) IsSyncToPhysicsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatableBody2D.fnIsSyncToPhysicsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
