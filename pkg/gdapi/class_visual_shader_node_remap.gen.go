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

type ptrsForVisualShaderNodeRemapList struct {
}

var ptrsForVisualShaderNodeRemap ptrsForVisualShaderNodeRemapList

func initVisualShaderNodeRemapPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeRemap")
	defer className.Destroy()
}

type VisualShaderNodeRemap struct {
	VisualShaderNode
}

func (me *VisualShaderNodeRemap) BaseClass() string {
	return "VisualShaderNodeRemap"
}

func NewVisualShaderNodeRemap() *VisualShaderNodeRemap {
	str := StringNameFromStr("VisualShaderNodeRemap") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeRemap{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeRemap) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeRemap) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeRemap) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
