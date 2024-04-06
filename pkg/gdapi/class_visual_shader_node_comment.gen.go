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

type ptrsForVisualShaderNodeCommentList struct {
	fnSetTitle       gdc.MethodBindPtr
	fnGetTitle       gdc.MethodBindPtr
	fnSetDescription gdc.MethodBindPtr
	fnGetDescription gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeComment ptrsForVisualShaderNodeCommentList

func initVisualShaderNodeCommentPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeComment")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_title")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeComment.fnSetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_title")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeComment.fnGetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_description")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeComment.fnSetDescription = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_description")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeComment.fnGetDescription = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}

}

type VisualShaderNodeComment struct {
	VisualShaderNodeResizableBase
}

func (me *VisualShaderNodeComment) BaseClass() string {
	return "VisualShaderNodeComment"
}

func NewVisualShaderNodeComment() *VisualShaderNodeComment {
	str := StringNameFromStr("VisualShaderNodeComment") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeComment{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeComment) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeComment) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeComment) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeComment) SetTitle(title String) {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeComment.fnSetTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeComment) GetTitle() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeComment.fnGetTitle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShaderNodeComment) SetDescription(description String) {
	cargs := []gdc.ConstTypePtr{description.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeComment.fnSetDescription), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeComment) GetDescription() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeComment.fnGetDescription), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
