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

type ptrsForParallaxLayerList struct {
	fnSetMotionScale  gdc.MethodBindPtr
	fnGetMotionScale  gdc.MethodBindPtr
	fnSetMotionOffset gdc.MethodBindPtr
	fnGetMotionOffset gdc.MethodBindPtr
	fnSetMirroring    gdc.MethodBindPtr
	fnGetMirroring    gdc.MethodBindPtr
}

var ptrsForParallaxLayer ptrsForParallaxLayerList

func initParallaxLayerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ParallaxLayer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_motion_scale")
		defer methodName.Destroy()
		ptrsForParallaxLayer.fnSetMotionScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_motion_scale")
		defer methodName.Destroy()
		ptrsForParallaxLayer.fnGetMotionScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_motion_offset")
		defer methodName.Destroy()
		ptrsForParallaxLayer.fnSetMotionOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_motion_offset")
		defer methodName.Destroy()
		ptrsForParallaxLayer.fnGetMotionOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_mirroring")
		defer methodName.Destroy()
		ptrsForParallaxLayer.fnSetMirroring = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_mirroring")
		defer methodName.Destroy()
		ptrsForParallaxLayer.fnGetMirroring = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
}

type ParallaxLayer struct {
	Node2D
}

func (me *ParallaxLayer) BaseClass() string {
	return "ParallaxLayer"
}

func NewParallaxLayer() *ParallaxLayer {
	str := StringNameFromStr("ParallaxLayer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ParallaxLayer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ParallaxLayer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ParallaxLayer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ParallaxLayer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ParallaxLayer) SetMotionScale(scale Vector2) {
	cargs := []gdc.ConstTypePtr{scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxLayer.fnSetMotionScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParallaxLayer) GetMotionScale() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxLayer.fnGetMotionScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParallaxLayer) SetMotionOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxLayer.fnSetMotionOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParallaxLayer) GetMotionOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxLayer.fnGetMotionOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParallaxLayer) SetMirroring(mirror Vector2) {
	cargs := []gdc.ConstTypePtr{mirror.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxLayer.fnSetMirroring), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParallaxLayer) GetMirroring() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxLayer.fnGetMirroring), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
