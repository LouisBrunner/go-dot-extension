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

type ptrsForVisualShaderNodeDotProductList struct {
}

var ptrsForVisualShaderNodeDotProduct ptrsForVisualShaderNodeDotProductList

func initVisualShaderNodeDotProductPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeDotProduct")
	defer className.Destroy()
}

type VisualShaderNodeDotProduct struct {
	VisualShaderNode
}

func (me *VisualShaderNodeDotProduct) BaseClass() string {
	return "VisualShaderNodeDotProduct"
}

func NewVisualShaderNodeDotProduct() *VisualShaderNodeDotProduct {
	str := StringNameFromStr("VisualShaderNodeDotProduct") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeDotProduct{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeDotProduct) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeDotProduct) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeDotProduct) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
