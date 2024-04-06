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

type ptrsForVisualShaderNodeSample3DList struct {
	fnSetSource gdc.MethodBindPtr
	fnGetSource gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeSample3D ptrsForVisualShaderNodeSample3DList

func initVisualShaderNodeSample3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeSample3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_source")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeSample3D.fnSetSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3315130991))
	}
	{
		methodName := StringNameFromStr("get_source")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeSample3D.fnGetSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1079494121))
	}
}

type VisualShaderNodeSample3D struct {
	VisualShaderNode
}

func (me *VisualShaderNodeSample3D) BaseClass() string {
	return "VisualShaderNodeSample3D"
}

func NewVisualShaderNodeSample3D() *VisualShaderNodeSample3D {
	str := StringNameFromStr("VisualShaderNodeSample3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeSample3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeSample3DSource int

const (
	VisualShaderNodeSample3DSourceSourceTexture VisualShaderNodeSample3DSource = 0
	VisualShaderNodeSample3DSourceSourcePort    VisualShaderNodeSample3DSource = 1
	VisualShaderNodeSample3DSourceSourceMax     VisualShaderNodeSample3DSource = 2
)

func (me *VisualShaderNodeSample3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSample3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSample3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeSample3D) SetSource(value VisualShaderNodeSample3DSource) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeSample3D.fnSetSource), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeSample3D) GetSource() VisualShaderNodeSample3DSource {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeSample3DSource

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeSample3D.fnGetSource), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
