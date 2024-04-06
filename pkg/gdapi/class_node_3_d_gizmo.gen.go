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

type ptrsForNode3DGizmoList struct {
}

var ptrsForNode3DGizmo ptrsForNode3DGizmoList

func initNode3DGizmoPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Node3DGizmo")
	defer className.Destroy()
}

type Node3DGizmo struct {
	RefCounted
}

func (me *Node3DGizmo) BaseClass() string {
	return "Node3DGizmo"
}

func NewNode3DGizmo() *Node3DGizmo {
	str := StringNameFromStr("Node3DGizmo") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Node3DGizmo{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Node3DGizmo) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Node3DGizmo) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Node3DGizmo) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
