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

type ptrsForVisualShaderNodeParticleBoxEmitterList struct {
}

var ptrsForVisualShaderNodeParticleBoxEmitter ptrsForVisualShaderNodeParticleBoxEmitterList

func initVisualShaderNodeParticleBoxEmitterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeParticleBoxEmitter")
	defer className.Destroy()
}

type VisualShaderNodeParticleBoxEmitter struct {
	VisualShaderNodeParticleEmitter
}

func (me *VisualShaderNodeParticleBoxEmitter) BaseClass() string {
	return "VisualShaderNodeParticleBoxEmitter"
}

func NewVisualShaderNodeParticleBoxEmitter() *VisualShaderNodeParticleBoxEmitter {
	str := StringNameFromStr("VisualShaderNodeParticleBoxEmitter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeParticleBoxEmitter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeParticleBoxEmitter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleBoxEmitter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleBoxEmitter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
