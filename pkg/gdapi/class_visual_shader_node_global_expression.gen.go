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

type ptrsForVisualShaderNodeGlobalExpressionList struct {
}

var ptrsForVisualShaderNodeGlobalExpression ptrsForVisualShaderNodeGlobalExpressionList

func initVisualShaderNodeGlobalExpressionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeGlobalExpression")
	defer className.Destroy()
}

type VisualShaderNodeGlobalExpression struct {
	VisualShaderNodeExpression
}

func (me *VisualShaderNodeGlobalExpression) BaseClass() string {
	return "VisualShaderNodeGlobalExpression"
}

func NewVisualShaderNodeGlobalExpression() *VisualShaderNodeGlobalExpression {
	str := StringNameFromStr("VisualShaderNodeGlobalExpression") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeGlobalExpression{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeGlobalExpression) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeGlobalExpression) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeGlobalExpression) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
