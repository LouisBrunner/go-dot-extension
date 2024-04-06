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

type ptrsForVisualShaderNodeBillboardList struct {
	fnSetBillboardType    gdc.MethodBindPtr
	fnGetBillboardType    gdc.MethodBindPtr
	fnSetKeepScaleEnabled gdc.MethodBindPtr
	fnIsKeepScaleEnabled  gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeBillboard ptrsForVisualShaderNodeBillboardList

func initVisualShaderNodeBillboardPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeBillboard")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_billboard_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeBillboard.fnSetBillboardType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1227463289))
	}
	{
		methodName := StringNameFromStr("get_billboard_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeBillboard.fnGetBillboardType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3724188517))
	}
	{
		methodName := StringNameFromStr("set_keep_scale_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeBillboard.fnSetKeepScaleEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_keep_scale_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeBillboard.fnIsKeepScaleEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type VisualShaderNodeBillboard struct {
	VisualShaderNode
}

func (me *VisualShaderNodeBillboard) BaseClass() string {
	return "VisualShaderNodeBillboard"
}

func NewVisualShaderNodeBillboard() *VisualShaderNodeBillboard {
	str := StringNameFromStr("VisualShaderNodeBillboard") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeBillboard{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeBillboardBillboardType int

const (
	VisualShaderNodeBillboardBillboardTypeBillboardTypeDisabled  VisualShaderNodeBillboardBillboardType = 0
	VisualShaderNodeBillboardBillboardTypeBillboardTypeEnabled   VisualShaderNodeBillboardBillboardType = 1
	VisualShaderNodeBillboardBillboardTypeBillboardTypeFixedY    VisualShaderNodeBillboardBillboardType = 2
	VisualShaderNodeBillboardBillboardTypeBillboardTypeParticles VisualShaderNodeBillboardBillboardType = 3
	VisualShaderNodeBillboardBillboardTypeBillboardTypeMax       VisualShaderNodeBillboardBillboardType = 4
)

func (me *VisualShaderNodeBillboard) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeBillboard) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeBillboard) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeBillboard) SetBillboardType(billboard_type VisualShaderNodeBillboardBillboardType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&billboard_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeBillboard.fnSetBillboardType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeBillboard) GetBillboardType() VisualShaderNodeBillboardBillboardType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeBillboardBillboardType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeBillboard.fnGetBillboardType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisualShaderNodeBillboard) SetKeepScaleEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeBillboard.fnSetKeepScaleEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeBillboard) IsKeepScaleEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeBillboard.fnIsKeepScaleEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
