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

type ptrsForRDPipelineColorBlendStateList struct {
	fnSetEnableLogicOp gdc.MethodBindPtr
	fnGetEnableLogicOp gdc.MethodBindPtr
	fnSetLogicOp       gdc.MethodBindPtr
	fnGetLogicOp       gdc.MethodBindPtr
	fnSetBlendConstant gdc.MethodBindPtr
	fnGetBlendConstant gdc.MethodBindPtr
	fnSetAttachments   gdc.MethodBindPtr
	fnGetAttachments   gdc.MethodBindPtr
}

var ptrsForRDPipelineColorBlendState ptrsForRDPipelineColorBlendStateList

func initRDPipelineColorBlendStatePtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDPipelineColorBlendState")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_enable_logic_op")
		defer methodName.Destroy()
		ptrsForRDPipelineColorBlendState.fnSetEnableLogicOp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_logic_op")
		defer methodName.Destroy()
		ptrsForRDPipelineColorBlendState.fnGetEnableLogicOp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_logic_op")
		defer methodName.Destroy()
		ptrsForRDPipelineColorBlendState.fnSetLogicOp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3610841058))
	}
	{
		methodName := StringNameFromStr("get_logic_op")
		defer methodName.Destroy()
		ptrsForRDPipelineColorBlendState.fnGetLogicOp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 988254690))
	}
	{
		methodName := StringNameFromStr("set_blend_constant")
		defer methodName.Destroy()
		ptrsForRDPipelineColorBlendState.fnSetBlendConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_blend_constant")
		defer methodName.Destroy()
		ptrsForRDPipelineColorBlendState.fnGetBlendConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_attachments")
		defer methodName.Destroy()
		ptrsForRDPipelineColorBlendState.fnSetAttachments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_attachments")
		defer methodName.Destroy()
		ptrsForRDPipelineColorBlendState.fnGetAttachments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
}

type RDPipelineColorBlendState struct {
	RefCounted
}

func (me *RDPipelineColorBlendState) BaseClass() string {
	return "RDPipelineColorBlendState"
}

func NewRDPipelineColorBlendState() *RDPipelineColorBlendState {
	str := StringNameFromStr("RDPipelineColorBlendState") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDPipelineColorBlendState{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDPipelineColorBlendState) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDPipelineColorBlendState) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDPipelineColorBlendState) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDPipelineColorBlendState) SetEnableLogicOp(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendState.fnSetEnableLogicOp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineColorBlendState) GetEnableLogicOp() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendState.fnGetEnableLogicOp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineColorBlendState) SetLogicOp(p_member RenderingDeviceLogicOperation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendState.fnSetLogicOp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineColorBlendState) GetLogicOp() RenderingDeviceLogicOperation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceLogicOperation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendState.fnGetLogicOp), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineColorBlendState) SetBlendConstant(p_member Color) {
	cargs := []gdc.ConstTypePtr{p_member.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendState.fnSetBlendConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineColorBlendState) GetBlendConstant() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendState.fnGetBlendConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RDPipelineColorBlendState) SetAttachments(attachments []RDPipelineColorBlendStateAttachment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&attachments)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendState.fnSetAttachments), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineColorBlendState) GetAttachments() []RDPipelineColorBlendStateAttachment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendState.fnGetAttachments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RDPipelineColorBlendStateAttachment](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
