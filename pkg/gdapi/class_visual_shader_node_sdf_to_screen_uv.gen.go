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

type ptrsForVisualShaderNodeSDFToScreenUVList struct {
}

var ptrsForVisualShaderNodeSDFToScreenUV ptrsForVisualShaderNodeSDFToScreenUVList

func initVisualShaderNodeSDFToScreenUVPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeSDFToScreenUV")
	defer className.Destroy()

}

type VisualShaderNodeSDFToScreenUV struct {
	VisualShaderNode
}

func (me *VisualShaderNodeSDFToScreenUV) BaseClass() string {
	return "VisualShaderNodeSDFToScreenUV"
}

func NewVisualShaderNodeSDFToScreenUV() *VisualShaderNodeSDFToScreenUV {
	str := StringNameFromStr("VisualShaderNodeSDFToScreenUV") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeSDFToScreenUV{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeSDFToScreenUV) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSDFToScreenUV) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSDFToScreenUV) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
