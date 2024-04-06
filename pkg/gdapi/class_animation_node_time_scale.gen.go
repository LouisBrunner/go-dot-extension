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

type ptrsForAnimationNodeTimeScaleList struct {
}

var ptrsForAnimationNodeTimeScale ptrsForAnimationNodeTimeScaleList

func initAnimationNodeTimeScalePtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationNodeTimeScale")
	defer className.Destroy()
}

type AnimationNodeTimeScale struct {
	AnimationNode
}

func (me *AnimationNodeTimeScale) BaseClass() string {
	return "AnimationNodeTimeScale"
}

func NewAnimationNodeTimeScale() *AnimationNodeTimeScale {
	str := StringNameFromStr("AnimationNodeTimeScale") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationNodeTimeScale{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AnimationNodeTimeScale) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationNodeTimeScale) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeTimeScale) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
