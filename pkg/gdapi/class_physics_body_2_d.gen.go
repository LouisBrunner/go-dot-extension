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

type ptrsForPhysicsBody2DList struct {
	fnMoveAndCollide               gdc.MethodBindPtr
	fnTestMove                     gdc.MethodBindPtr
	fnGetCollisionExceptions       gdc.MethodBindPtr
	fnAddCollisionExceptionWith    gdc.MethodBindPtr
	fnRemoveCollisionExceptionWith gdc.MethodBindPtr
}

var ptrsForPhysicsBody2D ptrsForPhysicsBody2DList

func initPhysicsBody2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsBody2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("move_and_collide")
		defer methodName.Destroy()
		ptrsForPhysicsBody2D.fnMoveAndCollide = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3681923724))
	}
	{
		methodName := StringNameFromStr("test_move")
		defer methodName.Destroy()
		ptrsForPhysicsBody2D.fnTestMove = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3324464701))
	}
	{
		methodName := StringNameFromStr("get_collision_exceptions")
		defer methodName.Destroy()
		ptrsForPhysicsBody2D.fnGetCollisionExceptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("add_collision_exception_with")
		defer methodName.Destroy()
		ptrsForPhysicsBody2D.fnAddCollisionExceptionWith = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("remove_collision_exception_with")
		defer methodName.Destroy()
		ptrsForPhysicsBody2D.fnRemoveCollisionExceptionWith = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}

}

type PhysicsBody2D struct {
	CollisionObject2D
}

func (me *PhysicsBody2D) BaseClass() string {
	return "PhysicsBody2D"
}

func NewPhysicsBody2D() *PhysicsBody2D {
	str := StringNameFromStr("PhysicsBody2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsBody2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsBody2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsBody2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsBody2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsBody2D) MoveAndCollide(motion Vector2, test_only bool, safe_margin float64, recovery_as_collision bool) KinematicCollision2D {
	cargs := []gdc.ConstTypePtr{motion.AsCTypePtr(), gdc.ConstTypePtr(&test_only), gdc.ConstTypePtr(&safe_margin), gdc.ConstTypePtr(&recovery_as_collision)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewKinematicCollision2D()
	pinner.Pin(&test_only)
	pinner.Pin(&safe_margin)
	pinner.Pin(&recovery_as_collision)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody2D.fnMoveAndCollide), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsBody2D) TestMove(from Transform2D, motion Vector2, collision KinematicCollision2D, safe_margin float64, recovery_as_collision bool) bool {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), motion.AsCTypePtr(), collision.AsCTypePtr(), gdc.ConstTypePtr(&safe_margin), gdc.ConstTypePtr(&recovery_as_collision)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&safe_margin)
	pinner.Pin(&recovery_as_collision)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody2D.fnTestMove), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsBody2D) GetCollisionExceptions() []PhysicsBody2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody2D.fnGetCollisionExceptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PhysicsBody2D](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsBody2D) AddCollisionExceptionWith(body Node) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody2D.fnAddCollisionExceptionWith), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsBody2D) RemoveCollisionExceptionWith(body Node) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsBody2D.fnRemoveCollisionExceptionWith), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
