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

type ptrsForVisualShaderNodeVaryingSetterList struct {
}

var ptrsForVisualShaderNodeVaryingSetter ptrsForVisualShaderNodeVaryingSetterList

func initVisualShaderNodeVaryingSetterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeVaryingSetter")
	defer className.Destroy()

}

type VisualShaderNodeVaryingSetter struct {
	VisualShaderNodeVarying
}

func (me *VisualShaderNodeVaryingSetter) BaseClass() string {
	return "VisualShaderNodeVaryingSetter"
}

func NewVisualShaderNodeVaryingSetter() *VisualShaderNodeVaryingSetter {
	str := StringNameFromStr("VisualShaderNodeVaryingSetter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeVaryingSetter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeVaryingSetter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVaryingSetter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVaryingSetter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
