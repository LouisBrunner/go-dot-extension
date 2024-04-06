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

type ptrsForVisualShaderNodeColorConstantList struct {
	fnSetConstant gdc.MethodBindPtr
	fnGetConstant gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeColorConstant ptrsForVisualShaderNodeColorConstantList

func initVisualShaderNodeColorConstantPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeColorConstant")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_constant")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeColorConstant.fnSetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_constant")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeColorConstant.fnGetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}

}

type VisualShaderNodeColorConstant struct {
	VisualShaderNodeConstant
}

func (me *VisualShaderNodeColorConstant) BaseClass() string {
	return "VisualShaderNodeColorConstant"
}

func NewVisualShaderNodeColorConstant() *VisualShaderNodeColorConstant {
	str := StringNameFromStr("VisualShaderNodeColorConstant") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeColorConstant{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeColorConstant) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeColorConstant) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeColorConstant) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeColorConstant) SetConstant(constant Color) {
	cargs := []gdc.ConstTypePtr{constant.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeColorConstant.fnSetConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeColorConstant) GetConstant() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeColorConstant.fnGetConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
