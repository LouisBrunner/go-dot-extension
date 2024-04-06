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

type ptrsForVisualShaderNodeVectorBaseList struct {
	fnSetOpType gdc.MethodBindPtr
	fnGetOpType gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeVectorBase ptrsForVisualShaderNodeVectorBaseList

func initVisualShaderNodeVectorBasePtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeVectorBase")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_op_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVectorBase.fnSetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1692596998))
	}
	{
		methodName := StringNameFromStr("get_op_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVectorBase.fnGetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2568738462))
	}

}

type VisualShaderNodeVectorBase struct {
	VisualShaderNode
}

func (me *VisualShaderNodeVectorBase) BaseClass() string {
	return "VisualShaderNodeVectorBase"
}

func NewVisualShaderNodeVectorBase() *VisualShaderNodeVectorBase {
	str := StringNameFromStr("VisualShaderNodeVectorBase") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeVectorBase{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeVectorBaseOpType int

const (
	VisualShaderNodeVectorBaseOpTypeOpTypeVector2D VisualShaderNodeVectorBaseOpType = 0
	VisualShaderNodeVectorBaseOpTypeOpTypeVector3D VisualShaderNodeVectorBaseOpType = 1
	VisualShaderNodeVectorBaseOpTypeOpTypeVector4D VisualShaderNodeVectorBaseOpType = 2
	VisualShaderNodeVectorBaseOpTypeOpTypeMax      VisualShaderNodeVectorBaseOpType = 3
)

func (me *VisualShaderNodeVectorBase) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorBase) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorBase) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeVectorBase) SetOpType(type_ VisualShaderNodeVectorBaseOpType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVectorBase.fnSetOpType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeVectorBase) GetOpType() VisualShaderNodeVectorBaseOpType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeVectorBaseOpType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVectorBase.fnGetOpType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
