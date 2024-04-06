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

type ptrsForVisualShaderNodeLinearSceneDepthList struct {
}

var ptrsForVisualShaderNodeLinearSceneDepth ptrsForVisualShaderNodeLinearSceneDepthList

func initVisualShaderNodeLinearSceneDepthPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeLinearSceneDepth")
	defer className.Destroy()

}

type VisualShaderNodeLinearSceneDepth struct {
	VisualShaderNode
}

func (me *VisualShaderNodeLinearSceneDepth) BaseClass() string {
	return "VisualShaderNodeLinearSceneDepth"
}

func NewVisualShaderNodeLinearSceneDepth() *VisualShaderNodeLinearSceneDepth {
	str := StringNameFromStr("VisualShaderNodeLinearSceneDepth") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeLinearSceneDepth{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeLinearSceneDepth) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeLinearSceneDepth) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeLinearSceneDepth) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
