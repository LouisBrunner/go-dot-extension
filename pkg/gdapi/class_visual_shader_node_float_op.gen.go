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

type ptrsForVisualShaderNodeFloatOpList struct {
	fnSetOperator gdc.MethodBindPtr
	fnGetOperator gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeFloatOp ptrsForVisualShaderNodeFloatOpList

func initVisualShaderNodeFloatOpPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeFloatOp")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_operator")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFloatOp.fnSetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2488468047))
	}
	{
		methodName := StringNameFromStr("get_operator")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFloatOp.fnGetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1867979390))
	}
}

type VisualShaderNodeFloatOp struct {
	VisualShaderNode
}

func (me *VisualShaderNodeFloatOp) BaseClass() string {
	return "VisualShaderNodeFloatOp"
}

func NewVisualShaderNodeFloatOp() *VisualShaderNodeFloatOp {
	str := StringNameFromStr("VisualShaderNodeFloatOp") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeFloatOp{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeFloatOpOperator int

const (
	VisualShaderNodeFloatOpOperatorOpAdd      VisualShaderNodeFloatOpOperator = 0
	VisualShaderNodeFloatOpOperatorOpSub      VisualShaderNodeFloatOpOperator = 1
	VisualShaderNodeFloatOpOperatorOpMul      VisualShaderNodeFloatOpOperator = 2
	VisualShaderNodeFloatOpOperatorOpDiv      VisualShaderNodeFloatOpOperator = 3
	VisualShaderNodeFloatOpOperatorOpMod      VisualShaderNodeFloatOpOperator = 4
	VisualShaderNodeFloatOpOperatorOpPow      VisualShaderNodeFloatOpOperator = 5
	VisualShaderNodeFloatOpOperatorOpMax      VisualShaderNodeFloatOpOperator = 6
	VisualShaderNodeFloatOpOperatorOpMin      VisualShaderNodeFloatOpOperator = 7
	VisualShaderNodeFloatOpOperatorOpAtan2    VisualShaderNodeFloatOpOperator = 8
	VisualShaderNodeFloatOpOperatorOpStep     VisualShaderNodeFloatOpOperator = 9
	VisualShaderNodeFloatOpOperatorOpEnumSize VisualShaderNodeFloatOpOperator = 10
)

func (me *VisualShaderNodeFloatOp) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFloatOp) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFloatOp) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeFloatOp) SetOperator(op VisualShaderNodeFloatOpOperator) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatOp.fnSetOperator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeFloatOp) GetOperator() VisualShaderNodeFloatOpOperator {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeFloatOpOperator

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatOp.fnGetOperator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
