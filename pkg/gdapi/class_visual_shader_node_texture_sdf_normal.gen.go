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

type ptrsForVisualShaderNodeTextureSDFNormalList struct {
}

var ptrsForVisualShaderNodeTextureSDFNormal ptrsForVisualShaderNodeTextureSDFNormalList

func initVisualShaderNodeTextureSDFNormalPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTextureSDFNormal")
	defer className.Destroy()
}

type VisualShaderNodeTextureSDFNormal struct {
	VisualShaderNode
}

func (me *VisualShaderNodeTextureSDFNormal) BaseClass() string {
	return "VisualShaderNodeTextureSDFNormal"
}

func NewVisualShaderNodeTextureSDFNormal() *VisualShaderNodeTextureSDFNormal {
	str := StringNameFromStr("VisualShaderNodeTextureSDFNormal") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTextureSDFNormal{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeTextureSDFNormal) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTextureSDFNormal) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTextureSDFNormal) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
