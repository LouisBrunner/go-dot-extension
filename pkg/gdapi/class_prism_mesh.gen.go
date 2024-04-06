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

type ptrsForPrismMeshList struct {
	fnSetLeftToRight     gdc.MethodBindPtr
	fnGetLeftToRight     gdc.MethodBindPtr
	fnSetSize            gdc.MethodBindPtr
	fnGetSize            gdc.MethodBindPtr
	fnSetSubdivideWidth  gdc.MethodBindPtr
	fnGetSubdivideWidth  gdc.MethodBindPtr
	fnSetSubdivideHeight gdc.MethodBindPtr
	fnGetSubdivideHeight gdc.MethodBindPtr
	fnSetSubdivideDepth  gdc.MethodBindPtr
	fnGetSubdivideDepth  gdc.MethodBindPtr
}

var ptrsForPrismMesh ptrsForPrismMeshList

func initPrismMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PrismMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_left_to_right")
		defer methodName.Destroy()
		ptrsForPrismMesh.fnSetLeftToRight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_left_to_right")
		defer methodName.Destroy()
		ptrsForPrismMesh.fnGetLeftToRight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForPrismMesh.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForPrismMesh.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_subdivide_width")
		defer methodName.Destroy()
		ptrsForPrismMesh.fnSetSubdivideWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_subdivide_width")
		defer methodName.Destroy()
		ptrsForPrismMesh.fnGetSubdivideWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_subdivide_height")
		defer methodName.Destroy()
		ptrsForPrismMesh.fnSetSubdivideHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_subdivide_height")
		defer methodName.Destroy()
		ptrsForPrismMesh.fnGetSubdivideHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_subdivide_depth")
		defer methodName.Destroy()
		ptrsForPrismMesh.fnSetSubdivideDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_subdivide_depth")
		defer methodName.Destroy()
		ptrsForPrismMesh.fnGetSubdivideDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type PrismMesh struct {
	PrimitiveMesh
}

func (me *PrismMesh) BaseClass() string {
	return "PrismMesh"
}

func NewPrismMesh() *PrismMesh {
	str := StringNameFromStr("PrismMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PrismMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PrismMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PrismMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PrismMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PrismMesh) SetLeftToRight(left_to_right float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&left_to_right)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrismMesh.fnSetLeftToRight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PrismMesh) GetLeftToRight() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrismMesh.fnGetLeftToRight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PrismMesh) SetSize(size Vector3) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrismMesh.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PrismMesh) GetSize() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrismMesh.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PrismMesh) SetSubdivideWidth(segments int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&segments)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrismMesh.fnSetSubdivideWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PrismMesh) GetSubdivideWidth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrismMesh.fnGetSubdivideWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PrismMesh) SetSubdivideHeight(segments int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&segments)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrismMesh.fnSetSubdivideHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PrismMesh) GetSubdivideHeight() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrismMesh.fnGetSubdivideHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PrismMesh) SetSubdivideDepth(segments int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&segments)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrismMesh.fnSetSubdivideDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PrismMesh) GetSubdivideDepth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrismMesh.fnGetSubdivideDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
