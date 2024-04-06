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

type ptrsForVisualShaderNodeSDFRaymarchList struct {
}

var ptrsForVisualShaderNodeSDFRaymarch ptrsForVisualShaderNodeSDFRaymarchList

func initVisualShaderNodeSDFRaymarchPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeSDFRaymarch")
	defer className.Destroy()

}

type VisualShaderNodeSDFRaymarch struct {
	VisualShaderNode
}

func (me *VisualShaderNodeSDFRaymarch) BaseClass() string {
	return "VisualShaderNodeSDFRaymarch"
}

func NewVisualShaderNodeSDFRaymarch() *VisualShaderNodeSDFRaymarch {
	str := StringNameFromStr("VisualShaderNodeSDFRaymarch") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeSDFRaymarch{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeSDFRaymarch) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSDFRaymarch) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSDFRaymarch) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
