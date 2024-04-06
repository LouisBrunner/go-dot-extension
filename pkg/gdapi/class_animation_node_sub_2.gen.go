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

type ptrsForAnimationNodeSub2List struct {
}

var ptrsForAnimationNodeSub2 ptrsForAnimationNodeSub2List

func initAnimationNodeSub2Ptrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationNodeSub2")
	defer className.Destroy()

}

type AnimationNodeSub2 struct {
	AnimationNodeSync
}

func (me *AnimationNodeSub2) BaseClass() string {
	return "AnimationNodeSub2"
}

func NewAnimationNodeSub2() *AnimationNodeSub2 {
	str := StringNameFromStr("AnimationNodeSub2") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationNodeSub2{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AnimationNodeSub2) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationNodeSub2) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeSub2) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
