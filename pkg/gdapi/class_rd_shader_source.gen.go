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

type ptrsForRDShaderSourceList struct {
	fnSetStageSource gdc.MethodBindPtr
	fnGetStageSource gdc.MethodBindPtr
	fnSetLanguage    gdc.MethodBindPtr
	fnGetLanguage    gdc.MethodBindPtr
}

var ptrsForRDShaderSource ptrsForRDShaderSourceList

func initRDShaderSourcePtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDShaderSource")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_stage_source")
		defer methodName.Destroy()
		ptrsForRDShaderSource.fnSetStageSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 620821314))
	}
	{
		methodName := StringNameFromStr("get_stage_source")
		defer methodName.Destroy()
		ptrsForRDShaderSource.fnGetStageSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3354920045))
	}
	{
		methodName := StringNameFromStr("set_language")
		defer methodName.Destroy()
		ptrsForRDShaderSource.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3422186742))
	}
	{
		methodName := StringNameFromStr("get_language")
		defer methodName.Destroy()
		ptrsForRDShaderSource.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1063538261))
	}

}

type RDShaderSource struct {
	RefCounted
}

func (me *RDShaderSource) BaseClass() string {
	return "RDShaderSource"
}

func NewRDShaderSource() *RDShaderSource {
	str := StringNameFromStr("RDShaderSource") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDShaderSource{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDShaderSource) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDShaderSource) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDShaderSource) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDShaderSource) SetStageSource(stage RenderingDeviceShaderStage, source String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage), source.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderSource.fnSetStageSource), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDShaderSource) GetStageSource(stage RenderingDeviceShaderStage) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&stage)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderSource.fnGetStageSource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RDShaderSource) SetLanguage(language RenderingDeviceShaderLanguage) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&language)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderSource.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDShaderSource) GetLanguage() RenderingDeviceShaderLanguage {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceShaderLanguage

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderSource.fnGetLanguage), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
