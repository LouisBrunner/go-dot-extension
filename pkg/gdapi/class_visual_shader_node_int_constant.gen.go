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

type ptrsForVisualShaderNodeIntConstantList struct {
	fnSetConstant gdc.MethodBindPtr
	fnGetConstant gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeIntConstant ptrsForVisualShaderNodeIntConstantList

func initVisualShaderNodeIntConstantPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeIntConstant")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_constant")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntConstant.fnSetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_constant")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntConstant.fnGetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type VisualShaderNodeIntConstant struct {
	VisualShaderNodeConstant
}

func (me *VisualShaderNodeIntConstant) BaseClass() string {
	return "VisualShaderNodeIntConstant"
}

func NewVisualShaderNodeIntConstant() *VisualShaderNodeIntConstant {
	str := StringNameFromStr("VisualShaderNodeIntConstant") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeIntConstant{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeIntConstant) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIntConstant) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIntConstant) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeIntConstant) SetConstant(constant int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&constant)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntConstant.fnSetConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeIntConstant) GetConstant() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntConstant.fnGetConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
