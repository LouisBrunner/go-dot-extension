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

type ptrsForVisualShaderNodeInputList struct {
	fnSetInputName     gdc.MethodBindPtr
	fnGetInputName     gdc.MethodBindPtr
	fnGetInputRealName gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeInput ptrsForVisualShaderNodeInputList

func initVisualShaderNodeInputPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeInput")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_input_name")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeInput.fnSetInputName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_input_name")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeInput.fnGetInputName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_input_real_name")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeInput.fnGetInputRealName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}

}

type VisualShaderNodeInput struct {
	VisualShaderNode
}

func (me *VisualShaderNodeInput) BaseClass() string {
	return "VisualShaderNodeInput"
}

func NewVisualShaderNodeInput() *VisualShaderNodeInput {
	str := StringNameFromStr("VisualShaderNodeInput") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeInput{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeInput) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeInput) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeInput) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeInput) SetInputName(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeInput.fnSetInputName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeInput) GetInputName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeInput.fnGetInputName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShaderNodeInput) GetInputRealName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeInput.fnGetInputRealName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type VisualShaderNodeInputInputTypeChangedSignalFn func()

func (me *VisualShaderNodeInput) ConnectInputTypeChanged(subs SignalSubscribers, fn VisualShaderNodeInputInputTypeChangedSignalFn) {
	sig := StringNameFromStr("input_type_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *VisualShaderNodeInput) DisconnectInputTypeChanged(subs SignalSubscribers, fn VisualShaderNodeInputInputTypeChangedSignalFn) {
	sig := StringNameFromStr("input_type_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
