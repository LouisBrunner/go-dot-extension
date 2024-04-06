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

type ptrsForAnimationRootNodeList struct {
}

var ptrsForAnimationRootNode ptrsForAnimationRootNodeList

func initAnimationRootNodePtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationRootNode")
	defer className.Destroy()

}

type AnimationRootNode struct {
	AnimationNode
}

func (me *AnimationRootNode) BaseClass() string {
	return "AnimationRootNode"
}

func NewAnimationRootNode() *AnimationRootNode {
	str := StringNameFromStr("AnimationRootNode") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationRootNode{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AnimationRootNode) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationRootNode) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationRootNode) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
