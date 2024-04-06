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

type ptrsForVisualShaderNodeRotationByAxisList struct {
}

var ptrsForVisualShaderNodeRotationByAxis ptrsForVisualShaderNodeRotationByAxisList

func initVisualShaderNodeRotationByAxisPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeRotationByAxis")
	defer className.Destroy()

}

type VisualShaderNodeRotationByAxis struct {
	VisualShaderNode
}

func (me *VisualShaderNodeRotationByAxis) BaseClass() string {
	return "VisualShaderNodeRotationByAxis"
}

func NewVisualShaderNodeRotationByAxis() *VisualShaderNodeRotationByAxis {
	str := StringNameFromStr("VisualShaderNodeRotationByAxis") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeRotationByAxis{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeRotationByAxis) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeRotationByAxis) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeRotationByAxis) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
