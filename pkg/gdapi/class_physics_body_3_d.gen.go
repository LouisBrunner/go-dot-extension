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

type ptrsForPhysicsBody3DList struct {
	fnMoveAndCollide               gdc.MethodBindPtr
	fnTestMove                     gdc.MethodBindPtr
	fnSetAxisLock                  gdc.MethodBindPtr
	fnGetAxisLock                  gdc.MethodBindPtr
	fnGetCollisionExceptions       gdc.MethodBindPtr
	fnAddCollisionExceptionWith    gdc.MethodBindPtr
	fnRemoveCollisionExceptionWith gdc.MethodBindPtr
}

var ptrsForPhysicsBody3D ptrsForPhysicsBody3DList

func initPhysicsBody3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsBody3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("move_and_collide")
		defer methodName.Destroy()
		ptrsForPhysicsBody3D.fnMoveAndCollide = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3208792678))
	}
	{
		methodName := StringNameFromStr("test_move")
		defer methodName.Destroy()
		ptrsForPhysicsBody3D.fnTestMove = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2481691619))
	}
	{
		methodName := StringNameFromStr("set_axis_lock")
		defer methodName.Destroy()
		ptrsForPhysicsBody3D.fnSetAxisLock = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1787895195))
	}
	{
		methodName := StringNameFromStr("get_axis_lock")
		defer methodName.Destroy()
		ptrsForPhysicsBody3D.fnGetAxisLock = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2264617709))
	}
	{
		methodName := StringNameFromStr("get_collision_exceptions")
		defer methodName.Destroy()
		ptrsForPhysicsBody3D.fnGetCollisionExceptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("add_collision_exception_with")
		defer methodName.Destroy()
		ptrsForPhysicsBody3D.fnAddCollisionExceptionWith = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("remove_collision_exception_with")
		defer methodName.Destroy()
		ptrsForPhysicsBody3D.fnRemoveCollisionExceptionWith = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
}

type PhysicsBody3D struct {
	CollisionObject3D
}

func (me *PhysicsBody3D) BaseClass() string {
	return "PhysicsBody3D"
}

func NewPhysicsBody3D() *PhysicsBody3D {
	str := StringNameFromStr("PhysicsBody3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsBody3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsBody3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsBody3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsBody3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsBody3D) MoveAndCollide(motion Vector3, test_only bool, safe_margin float64, recovery_as_collision bool, max_collisions int64) KinematicCollision3D {
	cargs := []gdc.ConstTypePtr{motion.AsCTypePtr(), gdc.ConstTypePtr(&test_only), gdc.ConstTypePtr(&safe_margin), gdc.ConstTypePtr(&recovery_as_collision), gdc.ConstTypePtr(&max_collisions)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewKinematicCollision3D()
	pinner.Pin(&test_only)
	pinner.Pin(&safe_margin)
	pinner.Pin(&recovery_as_collision)
	pinner.Pin(&max_collisions)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody3D.fnMoveAndCollide), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsBody3D) TestMove(from Transform3D, motion Vector3, collision KinematicCollision3D, safe_margin float64, recovery_as_collision bool, max_collisions int64) bool {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), motion.AsCTypePtr(), collision.AsCTypePtr(), gdc.ConstTypePtr(&safe_margin), gdc.ConstTypePtr(&recovery_as_collision), gdc.ConstTypePtr(&max_collisions)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&safe_margin)
	pinner.Pin(&recovery_as_collision)
	pinner.Pin(&max_collisions)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody3D.fnTestMove), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsBody3D) SetAxisLock(axis PhysicsServer3DBodyAxis, lock bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&lock)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody3D.fnSetAxisLock), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsBody3D) GetAxisLock(axis PhysicsServer3DBodyAxis) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&axis)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody3D.fnGetAxisLock), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsBody3D) GetCollisionExceptions() []PhysicsBody3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody3D.fnGetCollisionExceptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PhysicsBody3D](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsBody3D) AddCollisionExceptionWith(body Node) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody3D.fnAddCollisionExceptionWith), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsBody3D) RemoveCollisionExceptionWith(body Node) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody3D.fnRemoveCollisionExceptionWith), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
