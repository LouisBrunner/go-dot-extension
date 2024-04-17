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

type ptrsForVisualShaderNodeFrameList struct {
	fnSetTitle             gdc.MethodBindPtr
	fnGetTitle             gdc.MethodBindPtr
	fnSetTintColorEnabled  gdc.MethodBindPtr
	fnIsTintColorEnabled   gdc.MethodBindPtr
	fnSetTintColor         gdc.MethodBindPtr
	fnGetTintColor         gdc.MethodBindPtr
	fnSetAutoshrinkEnabled gdc.MethodBindPtr
	fnIsAutoshrinkEnabled  gdc.MethodBindPtr
	fnAddAttachedNode      gdc.MethodBindPtr
	fnRemoveAttachedNode   gdc.MethodBindPtr
	fnSetAttachedNodes     gdc.MethodBindPtr
	fnGetAttachedNodes     gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeFrame ptrsForVisualShaderNodeFrameList

func initVisualShaderNodeFramePtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeFrame")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_title")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnSetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_title")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnGetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_tint_color_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnSetTintColorEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_tint_color_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnIsTintColorEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tint_color")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnSetTintColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_tint_color")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnGetTintColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_autoshrink_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnSetAutoshrinkEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_autoshrink_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnIsAutoshrinkEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("add_attached_node")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnAddAttachedNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("remove_attached_node")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnRemoveAttachedNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_attached_nodes")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnSetAttachedNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
	}
	{
		methodName := StringNameFromStr("get_attached_nodes")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeFrame.fnGetAttachedNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
	}

}

type VisualShaderNodeFrame struct {
	VisualShaderNodeResizableBase
}

func (me *VisualShaderNodeFrame) BaseClass() string {
	return "VisualShaderNodeFrame"
}

func NewVisualShaderNodeFrame() *VisualShaderNodeFrame {
	str := StringNameFromStr("VisualShaderNodeFrame") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeFrame{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeFrame) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFrame) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFrame) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeFrame) SetTitle(title String) {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnSetTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeFrame) GetTitle() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnGetTitle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShaderNodeFrame) SetTintColorEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnSetTintColorEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeFrame) IsTintColorEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnIsTintColorEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeFrame) SetTintColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnSetTintColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeFrame) GetTintColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnGetTintColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShaderNodeFrame) SetAutoshrinkEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnSetAutoshrinkEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeFrame) IsAutoshrinkEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnIsAutoshrinkEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeFrame) AddAttachedNode(node int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&node)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnAddAttachedNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeFrame) RemoveAttachedNode(node int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&node)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnRemoveAttachedNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeFrame) SetAttachedNodes(attached_nodes PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{attached_nodes.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnSetAttachedNodes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeFrame) GetAttachedNodes() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFrame.fnGetAttachedNodes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
