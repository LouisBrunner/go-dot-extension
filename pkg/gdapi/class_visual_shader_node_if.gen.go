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

type ptrsForVisualShaderNodeIfList struct {
}

var ptrsForVisualShaderNodeIf ptrsForVisualShaderNodeIfList

func initVisualShaderNodeIfPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeIf")
	defer className.Destroy()
}

type VisualShaderNodeIf struct {
	VisualShaderNode
}

func (me *VisualShaderNodeIf) BaseClass() string {
	return "VisualShaderNodeIf"
}

func NewVisualShaderNodeIf() *VisualShaderNodeIf {
	str := StringNameFromStr("VisualShaderNodeIf") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeIf{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeIf) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIf) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIf) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
