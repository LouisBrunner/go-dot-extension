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

type ptrsForJoint3DList struct {
	fnSetNodeA                     gdc.MethodBindPtr
	fnGetNodeA                     gdc.MethodBindPtr
	fnSetNodeB                     gdc.MethodBindPtr
	fnGetNodeB                     gdc.MethodBindPtr
	fnSetSolverPriority            gdc.MethodBindPtr
	fnGetSolverPriority            gdc.MethodBindPtr
	fnSetExcludeNodesFromCollision gdc.MethodBindPtr
	fnGetExcludeNodesFromCollision gdc.MethodBindPtr
	fnGetRid                       gdc.MethodBindPtr
}

var ptrsForJoint3D ptrsForJoint3DList

func initJoint3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Joint3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_node_a")
		defer methodName.Destroy()
		ptrsForJoint3D.fnSetNodeA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_node_a")
		defer methodName.Destroy()
		ptrsForJoint3D.fnGetNodeA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_node_b")
		defer methodName.Destroy()
		ptrsForJoint3D.fnSetNodeB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_node_b")
		defer methodName.Destroy()
		ptrsForJoint3D.fnGetNodeB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_solver_priority")
		defer methodName.Destroy()
		ptrsForJoint3D.fnSetSolverPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_solver_priority")
		defer methodName.Destroy()
		ptrsForJoint3D.fnGetSolverPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_exclude_nodes_from_collision")
		defer methodName.Destroy()
		ptrsForJoint3D.fnSetExcludeNodesFromCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_exclude_nodes_from_collision")
		defer methodName.Destroy()
		ptrsForJoint3D.fnGetExcludeNodesFromCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForJoint3D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
}

type Joint3D struct {
	Node3D
}

func (me *Joint3D) BaseClass() string {
	return "Joint3D"
}

func NewJoint3D() *Joint3D {
	str := StringNameFromStr("Joint3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Joint3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Joint3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Joint3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Joint3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Joint3D) SetNodeA(node NodePath) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint3D.fnSetNodeA), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Joint3D) GetNodeA() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint3D.fnGetNodeA), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Joint3D) SetNodeB(node NodePath) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint3D.fnSetNodeB), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Joint3D) GetNodeB() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint3D.fnGetNodeB), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Joint3D) SetSolverPriority(priority int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint3D.fnSetSolverPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Joint3D) GetSolverPriority() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint3D.fnGetSolverPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Joint3D) SetExcludeNodesFromCollision(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint3D.fnSetExcludeNodesFromCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Joint3D) GetExcludeNodesFromCollision() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint3D.fnGetExcludeNodesFromCollision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Joint3D) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint3D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
