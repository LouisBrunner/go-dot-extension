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

type ptrsForVisualShaderNodeClampList struct {
	fnSetOpType gdc.MethodBindPtr
	fnGetOpType gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeClamp ptrsForVisualShaderNodeClampList

func initVisualShaderNodeClampPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeClamp")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_op_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeClamp.fnSetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 405010749))
	}
	{
		methodName := StringNameFromStr("get_op_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeClamp.fnGetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 233276050))
	}

}

type VisualShaderNodeClamp struct {
	VisualShaderNode
}

func (me *VisualShaderNodeClamp) BaseClass() string {
	return "VisualShaderNodeClamp"
}

func NewVisualShaderNodeClamp() *VisualShaderNodeClamp {
	str := StringNameFromStr("VisualShaderNodeClamp") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeClamp{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeClampOpType int

const (
	VisualShaderNodeClampOpTypeOpTypeFloat    VisualShaderNodeClampOpType = 0
	VisualShaderNodeClampOpTypeOpTypeInt      VisualShaderNodeClampOpType = 1
	VisualShaderNodeClampOpTypeOpTypeUint     VisualShaderNodeClampOpType = 2
	VisualShaderNodeClampOpTypeOpTypeVector2D VisualShaderNodeClampOpType = 3
	VisualShaderNodeClampOpTypeOpTypeVector3D VisualShaderNodeClampOpType = 4
	VisualShaderNodeClampOpTypeOpTypeVector4D VisualShaderNodeClampOpType = 5
	VisualShaderNodeClampOpTypeOpTypeMax      VisualShaderNodeClampOpType = 6
)

func (me *VisualShaderNodeClamp) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeClamp) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeClamp) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeClamp) SetOpType(op_type VisualShaderNodeClampOpType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeClamp.fnSetOpType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeClamp) GetOpType() VisualShaderNodeClampOpType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeClampOpType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeClamp.fnGetOpType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
