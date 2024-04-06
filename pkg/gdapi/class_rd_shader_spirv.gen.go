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

type ptrsForRDShaderSPIRVList struct {
	fnSetStageBytecode     gdc.MethodBindPtr
	fnGetStageBytecode     gdc.MethodBindPtr
	fnSetStageCompileError gdc.MethodBindPtr
	fnGetStageCompileError gdc.MethodBindPtr
}

var ptrsForRDShaderSPIRV ptrsForRDShaderSPIRVList

func initRDShaderSPIRVPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDShaderSPIRV")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_stage_bytecode")
		defer methodName.Destroy()
		ptrsForRDShaderSPIRV.fnSetStageBytecode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3514097977))
	}
	{
		methodName := StringNameFromStr("get_stage_bytecode")
		defer methodName.Destroy()
		ptrsForRDShaderSPIRV.fnGetStageBytecode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3816765404))
	}
	{
		methodName := StringNameFromStr("set_stage_compile_error")
		defer methodName.Destroy()
		ptrsForRDShaderSPIRV.fnSetStageCompileError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 620821314))
	}
	{
		methodName := StringNameFromStr("get_stage_compile_error")
		defer methodName.Destroy()
		ptrsForRDShaderSPIRV.fnGetStageCompileError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3354920045))
	}

}

type RDShaderSPIRV struct {
	Resource
}

func (me *RDShaderSPIRV) BaseClass() string {
	return "RDShaderSPIRV"
}

func NewRDShaderSPIRV() *RDShaderSPIRV {
	str := StringNameFromStr("RDShaderSPIRV") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDShaderSPIRV{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDShaderSPIRV) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDShaderSPIRV) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDShaderSPIRV) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDShaderSPIRV) SetStageBytecode(stage RenderingDeviceShaderStage, bytecode PackedByteArray) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage), bytecode.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderSPIRV.fnSetStageBytecode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDShaderSPIRV) GetStageBytecode(stage RenderingDeviceShaderStage) PackedByteArray {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()
	pinner.Pin(&stage)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderSPIRV.fnGetStageBytecode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RDShaderSPIRV) SetStageCompileError(stage RenderingDeviceShaderStage, compile_error String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage), compile_error.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderSPIRV.fnSetStageCompileError), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDShaderSPIRV) GetStageCompileError(stage RenderingDeviceShaderStage) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&stage)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderSPIRV.fnGetStageCompileError), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
