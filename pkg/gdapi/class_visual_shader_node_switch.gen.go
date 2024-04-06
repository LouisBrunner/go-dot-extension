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

type ptrsForVisualShaderNodeSwitchList struct {
	fnSetOpType gdc.MethodBindPtr
	fnGetOpType gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeSwitch ptrsForVisualShaderNodeSwitchList

func initVisualShaderNodeSwitchPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeSwitch")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_op_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeSwitch.fnSetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 510471861))
	}
	{
		methodName := StringNameFromStr("get_op_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeSwitch.fnGetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2517845071))
	}

}

type VisualShaderNodeSwitch struct {
	VisualShaderNode
}

func (me *VisualShaderNodeSwitch) BaseClass() string {
	return "VisualShaderNodeSwitch"
}

func NewVisualShaderNodeSwitch() *VisualShaderNodeSwitch {
	str := StringNameFromStr("VisualShaderNodeSwitch") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeSwitch{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeSwitchOpType int

const (
	VisualShaderNodeSwitchOpTypeOpTypeFloat     VisualShaderNodeSwitchOpType = 0
	VisualShaderNodeSwitchOpTypeOpTypeInt       VisualShaderNodeSwitchOpType = 1
	VisualShaderNodeSwitchOpTypeOpTypeUint      VisualShaderNodeSwitchOpType = 2
	VisualShaderNodeSwitchOpTypeOpTypeVector2D  VisualShaderNodeSwitchOpType = 3
	VisualShaderNodeSwitchOpTypeOpTypeVector3D  VisualShaderNodeSwitchOpType = 4
	VisualShaderNodeSwitchOpTypeOpTypeVector4D  VisualShaderNodeSwitchOpType = 5
	VisualShaderNodeSwitchOpTypeOpTypeBoolean   VisualShaderNodeSwitchOpType = 6
	VisualShaderNodeSwitchOpTypeOpTypeTransform VisualShaderNodeSwitchOpType = 7
	VisualShaderNodeSwitchOpTypeOpTypeMax       VisualShaderNodeSwitchOpType = 8
)

func (me *VisualShaderNodeSwitch) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSwitch) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSwitch) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeSwitch) SetOpType(type_ VisualShaderNodeSwitchOpType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeSwitch.fnSetOpType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeSwitch) GetOpType() VisualShaderNodeSwitchOpType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeSwitchOpType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeSwitch.fnGetOpType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
