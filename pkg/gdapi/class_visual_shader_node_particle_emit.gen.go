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

type ptrsForVisualShaderNodeParticleEmitList struct {
	fnSetFlags gdc.MethodBindPtr
	fnGetFlags gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeParticleEmit ptrsForVisualShaderNodeParticleEmitList

func initVisualShaderNodeParticleEmitPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeParticleEmit")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_flags")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleEmit.fnSetFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3960756792))
	}
	{
		methodName := StringNameFromStr("get_flags")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleEmit.fnGetFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 171277835))
	}
}

type VisualShaderNodeParticleEmit struct {
	VisualShaderNode
}

func (me *VisualShaderNodeParticleEmit) BaseClass() string {
	return "VisualShaderNodeParticleEmit"
}

func NewVisualShaderNodeParticleEmit() *VisualShaderNodeParticleEmit {
	str := StringNameFromStr("VisualShaderNodeParticleEmit") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeParticleEmit{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeParticleEmitEmitFlags int

const (
	VisualShaderNodeParticleEmitEmitFlagsEmitFlagPosition VisualShaderNodeParticleEmitEmitFlags = 1
	VisualShaderNodeParticleEmitEmitFlagsEmitFlagRotScale VisualShaderNodeParticleEmitEmitFlags = 2
	VisualShaderNodeParticleEmitEmitFlagsEmitFlagVelocity VisualShaderNodeParticleEmitEmitFlags = 4
	VisualShaderNodeParticleEmitEmitFlagsEmitFlagColor    VisualShaderNodeParticleEmitEmitFlags = 8
	VisualShaderNodeParticleEmitEmitFlagsEmitFlagCustom   VisualShaderNodeParticleEmitEmitFlags = 16
)

func (me *VisualShaderNodeParticleEmit) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleEmit) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleEmit) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeParticleEmit) SetFlags(flags VisualShaderNodeParticleEmitEmitFlags) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleEmit.fnSetFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeParticleEmit) GetFlags() VisualShaderNodeParticleEmitEmitFlags {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeParticleEmitEmitFlags

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleEmit.fnGetFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
