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

type ptrsForAnimationNodeBlend2List struct {
}

var ptrsForAnimationNodeBlend2 ptrsForAnimationNodeBlend2List

func initAnimationNodeBlend2Ptrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationNodeBlend2")
	defer className.Destroy()

}

type AnimationNodeBlend2 struct {
	AnimationNodeSync
}

func (me *AnimationNodeBlend2) BaseClass() string {
	return "AnimationNodeBlend2"
}

func NewAnimationNodeBlend2() *AnimationNodeBlend2 {
	str := StringNameFromStr("AnimationNodeBlend2") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationNodeBlend2{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AnimationNodeBlend2) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationNodeBlend2) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeBlend2) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
