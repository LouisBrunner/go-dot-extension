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

type ptrsForVisualShaderNodeTransformOpList struct {
	fnSetOperator gdc.MethodBindPtr
	fnGetOperator gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeTransformOp ptrsForVisualShaderNodeTransformOpList

func initVisualShaderNodeTransformOpPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTransformOp")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_operator")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTransformOp.fnSetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2287310733))
	}
	{
		methodName := StringNameFromStr("get_operator")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTransformOp.fnGetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1238663601))
	}
}

type VisualShaderNodeTransformOp struct {
	VisualShaderNode
}

func (me *VisualShaderNodeTransformOp) BaseClass() string {
	return "VisualShaderNodeTransformOp"
}

func NewVisualShaderNodeTransformOp() *VisualShaderNodeTransformOp {
	str := StringNameFromStr("VisualShaderNodeTransformOp") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTransformOp{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeTransformOpOperator int

const (
	VisualShaderNodeTransformOpOperatorOpAxb     VisualShaderNodeTransformOpOperator = 0
	VisualShaderNodeTransformOpOperatorOpBxa     VisualShaderNodeTransformOpOperator = 1
	VisualShaderNodeTransformOpOperatorOpAxbComp VisualShaderNodeTransformOpOperator = 2
	VisualShaderNodeTransformOpOperatorOpBxaComp VisualShaderNodeTransformOpOperator = 3
	VisualShaderNodeTransformOpOperatorOpAdd     VisualShaderNodeTransformOpOperator = 4
	VisualShaderNodeTransformOpOperatorOpAMinusB VisualShaderNodeTransformOpOperator = 5
	VisualShaderNodeTransformOpOperatorOpBMinusA VisualShaderNodeTransformOpOperator = 6
	VisualShaderNodeTransformOpOperatorOpADivB   VisualShaderNodeTransformOpOperator = 7
	VisualShaderNodeTransformOpOperatorOpBDivA   VisualShaderNodeTransformOpOperator = 8
	VisualShaderNodeTransformOpOperatorOpMax     VisualShaderNodeTransformOpOperator = 9
)

func (me *VisualShaderNodeTransformOp) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTransformOp) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformOp) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeTransformOp) SetOperator(op VisualShaderNodeTransformOpOperator) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTransformOp.fnSetOperator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTransformOp) GetOperator() VisualShaderNodeTransformOpOperator {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeTransformOpOperator

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTransformOp.fnGetOperator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
