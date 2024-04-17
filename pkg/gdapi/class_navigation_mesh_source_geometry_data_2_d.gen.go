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
	fnClear                      gdc.MethodBindPtr
	fnHasData                    gdc.MethodBindPtr
	fnSetTraversableOutlines     gdc.MethodBindPtr
	fnGetTraversableOutlines     gdc.MethodBindPtr
	fnSetObstructionOutlines     gdc.MethodBindPtr
	fnGetObstructionOutlines     gdc.MethodBindPtr
	fnAddTraversableOutline      gdc.MethodBindPtr
	fnAddObstructionOutline      gdc.MethodBindPtr
	fnMerge                      gdc.MethodBindPtr
	fnAddProjectedObstruction    gdc.MethodBindPtr
	fnClearProjectedObstructions gdc.MethodBindPtr
	fnSetProjectedObstructions   gdc.MethodBindPtr
	fnGetProjectedObstructions   gdc.MethodBindPtr
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
	{
		methodName := StringNameFromStr("merge")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnMerge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 742424872))
	}
	{
		methodName := StringNameFromStr("add_projected_obstruction")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnAddProjectedObstruction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3882407395))
	}
	{
		methodName := StringNameFromStr("clear_projected_obstructions")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnClearProjectedObstructions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_projected_obstructions")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnSetProjectedObstructions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_projected_obstructions")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData2D.fnGetProjectedObstructions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
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

func (me *NavigationMeshSourceGeometryData2D) Merge(other_geometry NavigationMeshSourceGeometryData2D) {
	cargs := []gdc.ConstTypePtr{other_geometry.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnMerge), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData2D) AddProjectedObstruction(vertices PackedVector2Array, carve bool) {
	cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), gdc.ConstTypePtr(&carve)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnAddProjectedObstruction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData2D) ClearProjectedObstructions() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnClearProjectedObstructions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData2D) SetProjectedObstructions(projected_obstructions Array) {
	cargs := []gdc.ConstTypePtr{projected_obstructions.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnSetProjectedObstructions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData2D) GetProjectedObstructions() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData2D.fnGetProjectedObstructions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
