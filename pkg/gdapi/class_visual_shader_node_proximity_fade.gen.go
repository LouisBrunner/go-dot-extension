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

type ptrsForVisualShaderNodeProximityFadeList struct {
}

var ptrsForVisualShaderNodeProximityFade ptrsForVisualShaderNodeProximityFadeList

func initVisualShaderNodeProximityFadePtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeProximityFade")
	defer className.Destroy()
}

type VisualShaderNodeProximityFade struct {
	VisualShaderNode
}

func (me *VisualShaderNodeProximityFade) BaseClass() string {
	return "VisualShaderNodeProximityFade"
}

func NewVisualShaderNodeProximityFade() *VisualShaderNodeProximityFade {
	str := StringNameFromStr("VisualShaderNodeProximityFade") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeProximityFade{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeProximityFade) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeProximityFade) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeProximityFade) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
