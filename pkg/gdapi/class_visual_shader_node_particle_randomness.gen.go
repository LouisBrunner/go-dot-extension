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

type ptrsForVisualShaderNodeParticleRandomnessList struct {
	fnSetOpType gdc.MethodBindPtr
	fnGetOpType gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeParticleRandomness ptrsForVisualShaderNodeParticleRandomnessList

func initVisualShaderNodeParticleRandomnessPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeParticleRandomness")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_op_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleRandomness.fnSetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2060089061))
	}
	{
		methodName := StringNameFromStr("get_op_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleRandomness.fnGetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3597061078))
	}
}

type VisualShaderNodeParticleRandomness struct {
	VisualShaderNode
}

func (me *VisualShaderNodeParticleRandomness) BaseClass() string {
	return "VisualShaderNodeParticleRandomness"
}

func NewVisualShaderNodeParticleRandomness() *VisualShaderNodeParticleRandomness {
	str := StringNameFromStr("VisualShaderNodeParticleRandomness") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeParticleRandomness{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeParticleRandomnessOpType int

const (
	VisualShaderNodeParticleRandomnessOpTypeOpTypeScalar   VisualShaderNodeParticleRandomnessOpType = 0
	VisualShaderNodeParticleRandomnessOpTypeOpTypeVector2D VisualShaderNodeParticleRandomnessOpType = 1
	VisualShaderNodeParticleRandomnessOpTypeOpTypeVector3D VisualShaderNodeParticleRandomnessOpType = 2
	VisualShaderNodeParticleRandomnessOpTypeOpTypeVector4D VisualShaderNodeParticleRandomnessOpType = 3
	VisualShaderNodeParticleRandomnessOpTypeOpTypeMax      VisualShaderNodeParticleRandomnessOpType = 4
)

func (me *VisualShaderNodeParticleRandomness) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleRandomness) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleRandomness) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeParticleRandomness) SetOpType(type_ VisualShaderNodeParticleRandomnessOpType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleRandomness.fnSetOpType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeParticleRandomness) GetOpType() VisualShaderNodeParticleRandomnessOpType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeParticleRandomnessOpType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleRandomness.fnGetOpType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
