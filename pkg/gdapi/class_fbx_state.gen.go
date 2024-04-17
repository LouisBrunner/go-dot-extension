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

type ptrsForFBXStateList struct {
	fnGetAllowGeometryHelperNodes gdc.MethodBindPtr
	fnSetAllowGeometryHelperNodes gdc.MethodBindPtr
}

var ptrsForFBXState ptrsForFBXStateList

func initFBXStatePtrs(iface gdc.Interface) {

	className := StringNameFromStr("FBXState")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_allow_geometry_helper_nodes")
		defer methodName.Destroy()
		ptrsForFBXState.fnGetAllowGeometryHelperNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_allow_geometry_helper_nodes")
		defer methodName.Destroy()
		ptrsForFBXState.fnSetAllowGeometryHelperNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}

}

type FBXState struct {
	GLTFState
}

func (me *FBXState) BaseClass() string {
	return "FBXState"
}

func NewFBXState() *FBXState {
	str := StringNameFromStr("FBXState") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &FBXState{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *FBXState) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *FBXState) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *FBXState) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *FBXState) GetAllowGeometryHelperNodes() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFBXState.fnGetAllowGeometryHelperNodes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FBXState) SetAllowGeometryHelperNodes(allow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFBXState.fnSetAllowGeometryHelperNodes), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
