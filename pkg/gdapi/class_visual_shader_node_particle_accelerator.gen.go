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

type ptrsForVisualShaderNodeParticleAcceleratorList struct {
	fnSetMode gdc.MethodBindPtr
	fnGetMode gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeParticleAccelerator ptrsForVisualShaderNodeParticleAcceleratorList

func initVisualShaderNodeParticleAcceleratorPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeParticleAccelerator")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_mode")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleAccelerator.fnSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457585749))
	}
	{
		methodName := StringNameFromStr("get_mode")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleAccelerator.fnGetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2660365633))
	}
}

type VisualShaderNodeParticleAccelerator struct {
	VisualShaderNode
}

func (me *VisualShaderNodeParticleAccelerator) BaseClass() string {
	return "VisualShaderNodeParticleAccelerator"
}

func NewVisualShaderNodeParticleAccelerator() *VisualShaderNodeParticleAccelerator {
	str := StringNameFromStr("VisualShaderNodeParticleAccelerator") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeParticleAccelerator{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeParticleAcceleratorMode int

const (
	VisualShaderNodeParticleAcceleratorModeModeLinear     VisualShaderNodeParticleAcceleratorMode = 0
	VisualShaderNodeParticleAcceleratorModeModeRadial     VisualShaderNodeParticleAcceleratorMode = 1
	VisualShaderNodeParticleAcceleratorModeModeTangential VisualShaderNodeParticleAcceleratorMode = 2
	VisualShaderNodeParticleAcceleratorModeModeMax        VisualShaderNodeParticleAcceleratorMode = 3
)

func (me *VisualShaderNodeParticleAccelerator) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleAccelerator) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleAccelerator) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeParticleAccelerator) SetMode(mode VisualShaderNodeParticleAcceleratorMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleAccelerator.fnSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeParticleAccelerator) GetMode() VisualShaderNodeParticleAcceleratorMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeParticleAcceleratorMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleAccelerator.fnGetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
