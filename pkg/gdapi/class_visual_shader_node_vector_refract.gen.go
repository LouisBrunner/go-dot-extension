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

type ptrsForVisualShaderNodeVectorRefractList struct {
}

var ptrsForVisualShaderNodeVectorRefract ptrsForVisualShaderNodeVectorRefractList

func initVisualShaderNodeVectorRefractPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeVectorRefract")
	defer className.Destroy()
}

type VisualShaderNodeVectorRefract struct {
	VisualShaderNodeVectorBase
}

func (me *VisualShaderNodeVectorRefract) BaseClass() string {
	return "VisualShaderNodeVectorRefract"
}

func NewVisualShaderNodeVectorRefract() *VisualShaderNodeVectorRefract {
	str := StringNameFromStr("VisualShaderNodeVectorRefract") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeVectorRefract{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeVectorRefract) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorRefract) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorRefract) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
