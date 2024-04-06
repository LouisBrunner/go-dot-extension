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

type ptrsForVisualShaderNodeExpressionList struct {
	fnSetExpression gdc.MethodBindPtr
	fnGetExpression gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeExpression ptrsForVisualShaderNodeExpressionList

func initVisualShaderNodeExpressionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeExpression")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_expression")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeExpression.fnSetExpression = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_expression")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeExpression.fnGetExpression = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
}

type VisualShaderNodeExpression struct {
	VisualShaderNodeGroupBase
}

func (me *VisualShaderNodeExpression) BaseClass() string {
	return "VisualShaderNodeExpression"
}

func NewVisualShaderNodeExpression() *VisualShaderNodeExpression {
	str := StringNameFromStr("VisualShaderNodeExpression") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeExpression{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeExpression) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeExpression) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeExpression) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeExpression) SetExpression(expression String) {
	cargs := []gdc.ConstTypePtr{expression.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeExpression.fnSetExpression), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeExpression) GetExpression() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeExpression.fnGetExpression), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
