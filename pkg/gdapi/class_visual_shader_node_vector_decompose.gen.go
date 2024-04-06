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

type ptrsForVisualShaderNodeVectorDecomposeList struct {
}

var ptrsForVisualShaderNodeVectorDecompose ptrsForVisualShaderNodeVectorDecomposeList

func initVisualShaderNodeVectorDecomposePtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeVectorDecompose")
	defer className.Destroy()

}

type VisualShaderNodeVectorDecompose struct {
	VisualShaderNodeVectorBase
}

func (me *VisualShaderNodeVectorDecompose) BaseClass() string {
	return "VisualShaderNodeVectorDecompose"
}

func NewVisualShaderNodeVectorDecompose() *VisualShaderNodeVectorDecompose {
	str := StringNameFromStr("VisualShaderNodeVectorDecompose") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeVectorDecompose{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeVectorDecompose) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorDecompose) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorDecompose) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
