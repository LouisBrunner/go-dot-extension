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

type ptrsForMissingNodeList struct {
	fnSetOriginalClass       gdc.MethodBindPtr
	fnGetOriginalClass       gdc.MethodBindPtr
	fnSetOriginalScene       gdc.MethodBindPtr
	fnGetOriginalScene       gdc.MethodBindPtr
	fnSetRecordingProperties gdc.MethodBindPtr
	fnIsRecordingProperties  gdc.MethodBindPtr
}

var ptrsForMissingNode ptrsForMissingNodeList

func initMissingNodePtrs(iface gdc.Interface) {

	className := StringNameFromStr("MissingNode")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_original_class")
		defer methodName.Destroy()
		ptrsForMissingNode.fnSetOriginalClass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_original_class")
		defer methodName.Destroy()
		ptrsForMissingNode.fnGetOriginalClass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_original_scene")
		defer methodName.Destroy()
		ptrsForMissingNode.fnSetOriginalScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_original_scene")
		defer methodName.Destroy()
		ptrsForMissingNode.fnGetOriginalScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_recording_properties")
		defer methodName.Destroy()
		ptrsForMissingNode.fnSetRecordingProperties = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_recording_properties")
		defer methodName.Destroy()
		ptrsForMissingNode.fnIsRecordingProperties = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type MissingNode struct {
	Node
}

func (me *MissingNode) BaseClass() string {
	return "MissingNode"
}

func NewMissingNode() *MissingNode {
	str := StringNameFromStr("MissingNode") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MissingNode{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MissingNode) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MissingNode) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MissingNode) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *MissingNode) SetOriginalClass(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMissingNode.fnSetOriginalClass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MissingNode) GetOriginalClass() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMissingNode.fnGetOriginalClass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MissingNode) SetOriginalScene(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMissingNode.fnSetOriginalScene), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MissingNode) GetOriginalScene() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMissingNode.fnGetOriginalScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MissingNode) SetRecordingProperties(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMissingNode.fnSetRecordingProperties), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MissingNode) IsRecordingProperties() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMissingNode.fnIsRecordingProperties), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
