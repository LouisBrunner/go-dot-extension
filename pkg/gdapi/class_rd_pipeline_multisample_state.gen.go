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

type ptrsForRDPipelineMultisampleStateList struct {
	fnSetSampleCount           gdc.MethodBindPtr
	fnGetSampleCount           gdc.MethodBindPtr
	fnSetEnableSampleShading   gdc.MethodBindPtr
	fnGetEnableSampleShading   gdc.MethodBindPtr
	fnSetMinSampleShading      gdc.MethodBindPtr
	fnGetMinSampleShading      gdc.MethodBindPtr
	fnSetEnableAlphaToCoverage gdc.MethodBindPtr
	fnGetEnableAlphaToCoverage gdc.MethodBindPtr
	fnSetEnableAlphaToOne      gdc.MethodBindPtr
	fnGetEnableAlphaToOne      gdc.MethodBindPtr
	fnSetSampleMasks           gdc.MethodBindPtr
	fnGetSampleMasks           gdc.MethodBindPtr
}

var ptrsForRDPipelineMultisampleState ptrsForRDPipelineMultisampleStateList

func initRDPipelineMultisampleStatePtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDPipelineMultisampleState")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_sample_count")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnSetSampleCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3774171498))
	}
	{
		methodName := StringNameFromStr("get_sample_count")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnGetSampleCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 407791724))
	}
	{
		methodName := StringNameFromStr("set_enable_sample_shading")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnSetEnableSampleShading = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_sample_shading")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnGetEnableSampleShading = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_min_sample_shading")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnSetMinSampleShading = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_min_sample_shading")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnGetMinSampleShading = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_enable_alpha_to_coverage")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnSetEnableAlphaToCoverage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_alpha_to_coverage")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnGetEnableAlphaToCoverage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_enable_alpha_to_one")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnSetEnableAlphaToOne = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_alpha_to_one")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnGetEnableAlphaToOne = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_sample_masks")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnSetSampleMasks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_sample_masks")
		defer methodName.Destroy()
		ptrsForRDPipelineMultisampleState.fnGetSampleMasks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
}

type RDPipelineMultisampleState struct {
	RefCounted
}

func (me *RDPipelineMultisampleState) BaseClass() string {
	return "RDPipelineMultisampleState"
}

func NewRDPipelineMultisampleState() *RDPipelineMultisampleState {
	str := StringNameFromStr("RDPipelineMultisampleState") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDPipelineMultisampleState{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDPipelineMultisampleState) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDPipelineMultisampleState) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDPipelineMultisampleState) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDPipelineMultisampleState) SetSampleCount(p_member RenderingDeviceTextureSamples) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnSetSampleCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineMultisampleState) GetSampleCount() RenderingDeviceTextureSamples {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceTextureSamples

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnGetSampleCount), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineMultisampleState) SetEnableSampleShading(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnSetEnableSampleShading), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineMultisampleState) GetEnableSampleShading() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnGetEnableSampleShading), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineMultisampleState) SetMinSampleShading(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnSetMinSampleShading), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineMultisampleState) GetMinSampleShading() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnGetMinSampleShading), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineMultisampleState) SetEnableAlphaToCoverage(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnSetEnableAlphaToCoverage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineMultisampleState) GetEnableAlphaToCoverage() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnGetEnableAlphaToCoverage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineMultisampleState) SetEnableAlphaToOne(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnSetEnableAlphaToOne), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineMultisampleState) GetEnableAlphaToOne() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnGetEnableAlphaToOne), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineMultisampleState) SetSampleMasks(masks []int) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&masks)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnSetSampleMasks), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineMultisampleState) GetSampleMasks() []int {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineMultisampleState.fnGetSampleMasks), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[int](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
