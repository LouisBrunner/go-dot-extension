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

type ptrsForNavigationMeshSourceGeometryData2DList struct {
	fnClear                  gdc.MethodBindPtr
	fnHasData                gdc.MethodBindPtr
	fnSetTraversableOutlines gdc.MethodBindPtr
	fnGetTraversableOutlines gdc.MethodBindPtr
	fnSetObstructionOutlines gdc.MethodBindPtr
	fnGetObstructionOutlines gdc.MethodBindPtr
	fnAddTraversableOutline  gdc.MethodBindPtr
	fnAddObstructionOutline  gdc.MethodBindPtr
}

var ptrsForNavigationMeshSourceGeometryData2D ptrsForNavigationMeshSourceGeometryData2DList

func initNavigationMeshSourceGeometryData2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationMeshSourceGeometryData2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("has_data")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnHasData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_traversable_outlines")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnSetTraversableOutlines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_traversable_outlines")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnGetTraversableOutlines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_obstruction_outlines")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnSetObstructionOutlines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_obstruction_outlines")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnGetObstructionOutlines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("add_traversable_outline")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnAddTraversableOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("add_obstruction_outline")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnAddObstructionOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}

}

type NavigationMeshSourceGeometryData2D struct {
	Resource
}

func (me *NavigationMeshSourceGeometryData2D) BaseClass() string {
	return "NavigationMeshSourceGeometryData2D"
}

func NewNavigationMeshSourceGeometryData2D() *NavigationMeshSourceGeometryData2D {
	str := StringNameFromStr("NavigationMeshSourceGeometryData2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationMeshSourceGeometryData2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *NavigationMeshSourceGeometryData2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationMeshSourceGeometryData2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationMeshSourceGeometryData2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationMeshSourceGeometryData2D) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData2D) HasData() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnHasData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationMeshSourceGeometryData2D) SetTraversableOutlines(traversable_outlines []PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&traversable_outlines)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnSetTraversableOutlines), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData2D) GetTraversableOutlines() []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnGetTraversableOutlines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *NavigationMeshSourceGeometryData2D) SetObstructionOutlines(obstruction_outlines []PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&obstruction_outlines)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnSetObstructionOutlines), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData2D) GetObstructionOutlines() []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnGetObstructionOutlines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *NavigationMeshSourceGeometryData2D) AddTraversableOutline(shape_outline PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{shape_outline.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnAddTraversableOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData2D) AddObstructionOutline(shape_outline PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{shape_outline.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnAddObstructionOutline), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
