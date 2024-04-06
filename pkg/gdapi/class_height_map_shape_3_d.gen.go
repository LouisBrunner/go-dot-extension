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

type ptrsForHeightMapShape3DList struct {
	fnSetMapWidth gdc.MethodBindPtr
	fnGetMapWidth gdc.MethodBindPtr
	fnSetMapDepth gdc.MethodBindPtr
	fnGetMapDepth gdc.MethodBindPtr
	fnSetMapData  gdc.MethodBindPtr
	fnGetMapData  gdc.MethodBindPtr
}

var ptrsForHeightMapShape3D ptrsForHeightMapShape3DList

func initHeightMapShape3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("HeightMapShape3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_map_width")
		defer methodName.Destroy()
		ptrsForHeightMapShape3D.fnSetMapWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_map_width")
		defer methodName.Destroy()
		ptrsForHeightMapShape3D.fnGetMapWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_map_depth")
		defer methodName.Destroy()
		ptrsForHeightMapShape3D.fnSetMapDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_map_depth")
		defer methodName.Destroy()
		ptrsForHeightMapShape3D.fnGetMapDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_map_data")
		defer methodName.Destroy()
		ptrsForHeightMapShape3D.fnSetMapData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2899603908))
	}
	{
		methodName := StringNameFromStr("get_map_data")
		defer methodName.Destroy()
		ptrsForHeightMapShape3D.fnGetMapData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675695659))
	}
}

type HeightMapShape3D struct {
	Shape3D
}

func (me *HeightMapShape3D) BaseClass() string {
	return "HeightMapShape3D"
}

func NewHeightMapShape3D() *HeightMapShape3D {
	str := StringNameFromStr("HeightMapShape3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &HeightMapShape3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *HeightMapShape3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *HeightMapShape3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *HeightMapShape3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *HeightMapShape3D) SetMapWidth(width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHeightMapShape3D.fnSetMapWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *HeightMapShape3D) GetMapWidth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHeightMapShape3D.fnGetMapWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *HeightMapShape3D) SetMapDepth(height int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHeightMapShape3D.fnSetMapDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *HeightMapShape3D) GetMapDepth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHeightMapShape3D.fnGetMapDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *HeightMapShape3D) SetMapData(data PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHeightMapShape3D.fnSetMapData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *HeightMapShape3D) GetMapData() PackedFloat32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedFloat32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHeightMapShape3D.fnGetMapData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
