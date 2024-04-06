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

type ptrsForJoint2DList struct {
	fnSetNodeA                     gdc.MethodBindPtr
	fnGetNodeA                     gdc.MethodBindPtr
	fnSetNodeB                     gdc.MethodBindPtr
	fnGetNodeB                     gdc.MethodBindPtr
	fnSetBias                      gdc.MethodBindPtr
	fnGetBias                      gdc.MethodBindPtr
	fnSetExcludeNodesFromCollision gdc.MethodBindPtr
	fnGetExcludeNodesFromCollision gdc.MethodBindPtr
	fnGetRid                       gdc.MethodBindPtr
}

var ptrsForJoint2D ptrsForJoint2DList

func initJoint2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Joint2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_node_a")
		defer methodName.Destroy()
		ptrsForJoint2D.fnSetNodeA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_node_a")
		defer methodName.Destroy()
		ptrsForJoint2D.fnGetNodeA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_node_b")
		defer methodName.Destroy()
		ptrsForJoint2D.fnSetNodeB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_node_b")
		defer methodName.Destroy()
		ptrsForJoint2D.fnGetNodeB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_bias")
		defer methodName.Destroy()
		ptrsForJoint2D.fnSetBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_bias")
		defer methodName.Destroy()
		ptrsForJoint2D.fnGetBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_exclude_nodes_from_collision")
		defer methodName.Destroy()
		ptrsForJoint2D.fnSetExcludeNodesFromCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_exclude_nodes_from_collision")
		defer methodName.Destroy()
		ptrsForJoint2D.fnGetExcludeNodesFromCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForJoint2D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}

}

type Joint2D struct {
	Node2D
}

func (me *Joint2D) BaseClass() string {
	return "Joint2D"
}

func NewJoint2D() *Joint2D {
	str := StringNameFromStr("Joint2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Joint2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Joint2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Joint2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Joint2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Joint2D) SetNodeA(node NodePath) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint2D.fnSetNodeA), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Joint2D) GetNodeA() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint2D.fnGetNodeA), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Joint2D) SetNodeB(node NodePath) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint2D.fnSetNodeB), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Joint2D) GetNodeB() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint2D.fnGetNodeB), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Joint2D) SetBias(bias float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint2D.fnSetBias), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Joint2D) GetBias() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint2D.fnGetBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Joint2D) SetExcludeNodesFromCollision(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint2D.fnSetExcludeNodesFromCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Joint2D) GetExcludeNodesFromCollision() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint2D.fnGetExcludeNodesFromCollision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Joint2D) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJoint2D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
