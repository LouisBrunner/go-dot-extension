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

type ptrsForCollisionShape3DList struct {
	fnResourceChanged        gdc.MethodBindPtr
	fnSetShape               gdc.MethodBindPtr
	fnGetShape               gdc.MethodBindPtr
	fnSetDisabled            gdc.MethodBindPtr
	fnIsDisabled             gdc.MethodBindPtr
	fnMakeConvexFromSiblings gdc.MethodBindPtr
}

var ptrsForCollisionShape3D ptrsForCollisionShape3DList

func initCollisionShape3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CollisionShape3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("resource_changed")
		defer methodName.Destroy()
		ptrsForCollisionShape3D.fnResourceChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 968641751))
	}
	{
		methodName := StringNameFromStr("set_shape")
		defer methodName.Destroy()
		ptrsForCollisionShape3D.fnSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1549710052))
	}
	{
		methodName := StringNameFromStr("get_shape")
		defer methodName.Destroy()
		ptrsForCollisionShape3D.fnGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3214262478))
	}
	{
		methodName := StringNameFromStr("set_disabled")
		defer methodName.Destroy()
		ptrsForCollisionShape3D.fnSetDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_disabled")
		defer methodName.Destroy()
		ptrsForCollisionShape3D.fnIsDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("make_convex_from_siblings")
		defer methodName.Destroy()
		ptrsForCollisionShape3D.fnMakeConvexFromSiblings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type CollisionShape3D struct {
	Node3D
}

func (me *CollisionShape3D) BaseClass() string {
	return "CollisionShape3D"
}

func NewCollisionShape3D() *CollisionShape3D {
	str := StringNameFromStr("CollisionShape3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CollisionShape3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CollisionShape3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CollisionShape3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CollisionShape3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CollisionShape3D) ResourceChanged(resource Resource) {
	cargs := []gdc.ConstTypePtr{resource.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape3D.fnResourceChanged), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionShape3D) SetShape(shape Shape3D) {
	cargs := []gdc.ConstTypePtr{shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape3D.fnSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionShape3D) GetShape() Shape3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewShape3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape3D.fnGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionShape3D) SetDisabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape3D.fnSetDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionShape3D) IsDisabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape3D.fnIsDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionShape3D) MakeConvexFromSiblings() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape3D.fnMakeConvexFromSiblings), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
