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

type ptrsForVisualShaderNodeTextureSDFList struct {
}

var ptrsForVisualShaderNodeTextureSDF ptrsForVisualShaderNodeTextureSDFList

func initVisualShaderNodeTextureSDFPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTextureSDF")
	defer className.Destroy()
}

type VisualShaderNodeTextureSDF struct {
	VisualShaderNode
}

func (me *VisualShaderNodeTextureSDF) BaseClass() string {
	return "VisualShaderNodeTextureSDF"
}

func NewVisualShaderNodeTextureSDF() *VisualShaderNodeTextureSDF {
	str := StringNameFromStr("VisualShaderNodeTextureSDF") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTextureSDF{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeTextureSDF) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTextureSDF) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTextureSDF) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
