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

type ptrsForVisualShaderNodeDeterminantList struct {
}

var ptrsForVisualShaderNodeDeterminant ptrsForVisualShaderNodeDeterminantList

func initVisualShaderNodeDeterminantPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeDeterminant")
	defer className.Destroy()

}

type VisualShaderNodeDeterminant struct {
	VisualShaderNode
}

func (me *VisualShaderNodeDeterminant) BaseClass() string {
	return "VisualShaderNodeDeterminant"
}

func NewVisualShaderNodeDeterminant() *VisualShaderNodeDeterminant {
	str := StringNameFromStr("VisualShaderNodeDeterminant") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeDeterminant{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeDeterminant) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeDeterminant) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeDeterminant) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
