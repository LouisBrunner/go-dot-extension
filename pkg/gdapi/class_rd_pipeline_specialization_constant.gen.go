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

type ptrsForRDPipelineSpecializationConstantList struct {
	fnSetValue      gdc.MethodBindPtr
	fnGetValue      gdc.MethodBindPtr
	fnSetConstantId gdc.MethodBindPtr
	fnGetConstantId gdc.MethodBindPtr
}

var ptrsForRDPipelineSpecializationConstant ptrsForRDPipelineSpecializationConstantList

func initRDPipelineSpecializationConstantPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDPipelineSpecializationConstant")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_value")
		defer methodName.Destroy()
		ptrsForRDPipelineSpecializationConstant.fnSetValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1114965689))
	}
	{
		methodName := StringNameFromStr("get_value")
		defer methodName.Destroy()
		ptrsForRDPipelineSpecializationConstant.fnGetValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1214101251))
	}
	{
		methodName := StringNameFromStr("set_constant_id")
		defer methodName.Destroy()
		ptrsForRDPipelineSpecializationConstant.fnSetConstantId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_constant_id")
		defer methodName.Destroy()
		ptrsForRDPipelineSpecializationConstant.fnGetConstantId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type RDPipelineSpecializationConstant struct {
	RefCounted
}

func (me *RDPipelineSpecializationConstant) BaseClass() string {
	return "RDPipelineSpecializationConstant"
}

func NewRDPipelineSpecializationConstant() *RDPipelineSpecializationConstant {
	str := StringNameFromStr("RDPipelineSpecializationConstant") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDPipelineSpecializationConstant{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDPipelineSpecializationConstant) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDPipelineSpecializationConstant) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDPipelineSpecializationConstant) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDPipelineSpecializationConstant) SetValue(value Variant) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineSpecializationConstant.fnSetValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineSpecializationConstant) GetValue() Variant {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineSpecializationConstant.fnGetValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RDPipelineSpecializationConstant) SetConstantId(constant_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&constant_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineSpecializationConstant.fnSetConstantId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineSpecializationConstant) GetConstantId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineSpecializationConstant.fnGetConstantId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
