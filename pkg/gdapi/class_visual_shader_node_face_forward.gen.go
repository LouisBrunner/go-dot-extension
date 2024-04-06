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

type ptrsForVisualShaderNodeFaceForwardList struct {
}

var ptrsForVisualShaderNodeFaceForward ptrsForVisualShaderNodeFaceForwardList

func initVisualShaderNodeFaceForwardPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeFaceForward")
	defer className.Destroy()
}

type VisualShaderNodeFaceForward struct {
	VisualShaderNodeVectorBase
}

func (me *VisualShaderNodeFaceForward) BaseClass() string {
	return "VisualShaderNodeFaceForward"
}

func NewVisualShaderNodeFaceForward() *VisualShaderNodeFaceForward {
	str := StringNameFromStr("VisualShaderNodeFaceForward") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeFaceForward{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeFaceForward) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFaceForward) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFaceForward) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
