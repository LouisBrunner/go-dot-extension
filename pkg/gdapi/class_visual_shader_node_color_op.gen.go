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

type ptrsForVisualShaderNodeColorOpList struct {
	fnSetOperator gdc.MethodBindPtr
	fnGetOperator gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeColorOp ptrsForVisualShaderNodeColorOpList

func initVisualShaderNodeColorOpPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeColorOp")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_operator")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeColorOp.fnSetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4260370673))
	}
	{
		methodName := StringNameFromStr("get_operator")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeColorOp.fnGetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1950956529))
	}

}

type VisualShaderNodeColorOp struct {
	VisualShaderNode
}

func (me *VisualShaderNodeColorOp) BaseClass() string {
	return "VisualShaderNodeColorOp"
}

func NewVisualShaderNodeColorOp() *VisualShaderNodeColorOp {
	str := StringNameFromStr("VisualShaderNodeColorOp") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeColorOp{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeColorOpOperator int

const (
	VisualShaderNodeColorOpOperatorOpScreen     VisualShaderNodeColorOpOperator = 0
	VisualShaderNodeColorOpOperatorOpDifference VisualShaderNodeColorOpOperator = 1
	VisualShaderNodeColorOpOperatorOpDarken     VisualShaderNodeColorOpOperator = 2
	VisualShaderNodeColorOpOperatorOpLighten    VisualShaderNodeColorOpOperator = 3
	VisualShaderNodeColorOpOperatorOpOverlay    VisualShaderNodeColorOpOperator = 4
	VisualShaderNodeColorOpOperatorOpDodge      VisualShaderNodeColorOpOperator = 5
	VisualShaderNodeColorOpOperatorOpBurn       VisualShaderNodeColorOpOperator = 6
	VisualShaderNodeColorOpOperatorOpSoftLight  VisualShaderNodeColorOpOperator = 7
	VisualShaderNodeColorOpOperatorOpHardLight  VisualShaderNodeColorOpOperator = 8
	VisualShaderNodeColorOpOperatorOpMax        VisualShaderNodeColorOpOperator = 9
)

func (me *VisualShaderNodeColorOp) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeColorOp) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeColorOp) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeColorOp) SetOperator(op VisualShaderNodeColorOpOperator) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeColorOp.fnSetOperator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeColorOp) GetOperator() VisualShaderNodeColorOpOperator {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeColorOpOperator

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeColorOp.fnGetOperator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
