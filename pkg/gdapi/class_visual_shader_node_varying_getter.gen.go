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

type ptrsForVisualShaderNodeVaryingGetterList struct {
}

var ptrsForVisualShaderNodeVaryingGetter ptrsForVisualShaderNodeVaryingGetterList

func initVisualShaderNodeVaryingGetterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeVaryingGetter")
	defer className.Destroy()
}

type VisualShaderNodeVaryingGetter struct {
	VisualShaderNodeVarying
}

func (me *VisualShaderNodeVaryingGetter) BaseClass() string {
	return "VisualShaderNodeVaryingGetter"
}

func NewVisualShaderNodeVaryingGetter() *VisualShaderNodeVaryingGetter {
	str := StringNameFromStr("VisualShaderNodeVaryingGetter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeVaryingGetter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeVaryingGetter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVaryingGetter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVaryingGetter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
