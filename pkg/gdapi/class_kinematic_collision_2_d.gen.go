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

type ptrsForKinematicCollision2DList struct {
	fnGetPosition           gdc.MethodBindPtr
	fnGetNormal             gdc.MethodBindPtr
	fnGetTravel             gdc.MethodBindPtr
	fnGetRemainder          gdc.MethodBindPtr
	fnGetAngle              gdc.MethodBindPtr
	fnGetDepth              gdc.MethodBindPtr
	fnGetLocalShape         gdc.MethodBindPtr
	fnGetCollider           gdc.MethodBindPtr
	fnGetColliderId         gdc.MethodBindPtr
	fnGetColliderRid        gdc.MethodBindPtr
	fnGetColliderShape      gdc.MethodBindPtr
	fnGetColliderShapeIndex gdc.MethodBindPtr
	fnGetColliderVelocity   gdc.MethodBindPtr
}

var ptrsForKinematicCollision2D ptrsForKinematicCollision2DList

func initKinematicCollision2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("KinematicCollision2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_normal")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_travel")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetTravel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_remainder")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetRemainder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_angle")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841063350))
	}
	{
		methodName := StringNameFromStr("get_depth")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_local_shape")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetLocalShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1981248198))
	}
	{
		methodName := StringNameFromStr("get_collider")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetCollider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1981248198))
	}
	{
		methodName := StringNameFromStr("get_collider_id")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetColliderId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_collider_rid")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetColliderRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_collider_shape")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetColliderShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1981248198))
	}
	{
		methodName := StringNameFromStr("get_collider_shape_index")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetColliderShapeIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_collider_velocity")
		defer methodName.Destroy()
		ptrsForKinematicCollision2D.fnGetColliderVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
}

type KinematicCollision2D struct {
	RefCounted
}

func (me *KinematicCollision2D) BaseClass() string {
	return "KinematicCollision2D"
}

func NewKinematicCollision2D() *KinematicCollision2D {
	str := StringNameFromStr("KinematicCollision2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &KinematicCollision2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *KinematicCollision2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *KinematicCollision2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *KinematicCollision2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *KinematicCollision2D) GetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision2D) GetNormal() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision2D) GetTravel() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetTravel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision2D) GetRemainder() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetRemainder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision2D) GetAngle(up_direction Vector2) float64 {
	cargs := []gdc.ConstTypePtr{up_direction.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *KinematicCollision2D) GetDepth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *KinematicCollision2D) GetLocalShape() Object {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetLocalShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision2D) GetCollider() Object {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetCollider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision2D) GetColliderId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetColliderId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *KinematicCollision2D) GetColliderRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetColliderRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision2D) GetColliderShape() Object {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetColliderShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision2D) GetColliderShapeIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetColliderShapeIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *KinematicCollision2D) GetColliderVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision2D.fnGetColliderVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
