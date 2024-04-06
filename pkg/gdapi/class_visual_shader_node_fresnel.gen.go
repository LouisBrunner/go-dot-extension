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

type ptrsForVisualShaderNodeFresnelList struct {
}

var ptrsForVisualShaderNodeFresnel ptrsForVisualShaderNodeFresnelList

func initVisualShaderNodeFresnelPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeFresnel")
	defer className.Destroy()

}

type VisualShaderNodeFresnel struct {
	VisualShaderNode
}

func (me *VisualShaderNodeFresnel) BaseClass() string {
	return "VisualShaderNodeFresnel"
}

func NewVisualShaderNodeFresnel() *VisualShaderNodeFresnel {
	str := StringNameFromStr("VisualShaderNodeFresnel") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeFresnel{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeFresnel) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFresnel) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFresnel) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
