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

type ptrsForNavigationMeshGeneratorList struct {
	fnBake                       gdc.MethodBindPtr
	fnClear                      gdc.MethodBindPtr
	fnParseSourceGeometryData    gdc.MethodBindPtr
	fnBakeFromSourceGeometryData gdc.MethodBindPtr
}

var ptrsForNavigationMeshGenerator ptrsForNavigationMeshGeneratorList

func initNavigationMeshGeneratorPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationMeshGenerator")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("bake")
		defer methodName.Destroy()
		ptrsForNavigationMeshGenerator.fnBake = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1401173477))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForNavigationMeshGenerator.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2923361153))
	}
	{
		methodName := StringNameFromStr("parse_source_geometry_data")
		defer methodName.Destroy()
		ptrsForNavigationMeshGenerator.fnParseSourceGeometryData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 685862123))
	}
	{
		methodName := StringNameFromStr("bake_from_source_geometry_data")
		defer methodName.Destroy()
		ptrsForNavigationMeshGenerator.fnBakeFromSourceGeometryData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2469318639))
	}

}

type NavigationMeshGenerator struct {
	Object
}

func (me *NavigationMeshGenerator) BaseClass() string {
	return "NavigationMeshGenerator"
}

func NewNavigationMeshGenerator() *NavigationMeshGenerator {
	str := StringNameFromStr("NavigationMeshGenerator") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationMeshGenerator{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *NavigationMeshGenerator) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationMeshGenerator) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationMeshGenerator) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationMeshGenerator) Bake(navigation_mesh NavigationMesh, root_node Node) {
	cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), root_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshGenerator.fnBake), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshGenerator) Clear(navigation_mesh NavigationMesh) {
	cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshGenerator.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshGenerator) ParseSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, root_node Node, callback Callable) {
	cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), source_geometry_data.AsCTypePtr(), root_node.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshGenerator.fnParseSourceGeometryData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshGenerator) BakeFromSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, callback Callable) {
	cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), source_geometry_data.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshGenerator.fnBakeFromSourceGeometryData), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
