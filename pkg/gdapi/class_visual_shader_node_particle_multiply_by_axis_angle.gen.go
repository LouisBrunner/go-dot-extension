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

type ptrsForVisualShaderNodeParticleMultiplyByAxisAngleList struct {
	fnSetDegreesMode gdc.MethodBindPtr
	fnIsDegreesMode  gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeParticleMultiplyByAxisAngle ptrsForVisualShaderNodeParticleMultiplyByAxisAngleList

func initVisualShaderNodeParticleMultiplyByAxisAnglePtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeParticleMultiplyByAxisAngle")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_degrees_mode")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleMultiplyByAxisAngle.fnSetDegreesMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_degrees_mode")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleMultiplyByAxisAngle.fnIsDegreesMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type VisualShaderNodeParticleMultiplyByAxisAngle struct {
	VisualShaderNode
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) BaseClass() string {
	return "VisualShaderNodeParticleMultiplyByAxisAngle"
}

func NewVisualShaderNodeParticleMultiplyByAxisAngle() *VisualShaderNodeParticleMultiplyByAxisAngle {
	str := StringNameFromStr("VisualShaderNodeParticleMultiplyByAxisAngle") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeParticleMultiplyByAxisAngle{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) SetDegreesMode(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleMultiplyByAxisAngle.fnSetDegreesMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) IsDegreesMode() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleMultiplyByAxisAngle.fnIsDegreesMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
