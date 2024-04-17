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

type ptrsForNavigationMeshSourceGeometryData3DList struct {
	fnSetVertices                gdc.MethodBindPtr
	fnGetVertices                gdc.MethodBindPtr
	fnSetIndices                 gdc.MethodBindPtr
	fnGetIndices                 gdc.MethodBindPtr
	fnClear                      gdc.MethodBindPtr
	fnHasData                    gdc.MethodBindPtr
	fnAddMesh                    gdc.MethodBindPtr
	fnAddMeshArray               gdc.MethodBindPtr
	fnAddFaces                   gdc.MethodBindPtr
	fnMerge                      gdc.MethodBindPtr
	fnAddProjectedObstruction    gdc.MethodBindPtr
	fnClearProjectedObstructions gdc.MethodBindPtr
	fnSetProjectedObstructions   gdc.MethodBindPtr
	fnGetProjectedObstructions   gdc.MethodBindPtr
}

var ptrsForNavigationMeshSourceGeometryData3D ptrsForNavigationMeshSourceGeometryData3DList

func initNavigationMeshSourceGeometryData3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationMeshSourceGeometryData3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_vertices")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnSetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2899603908))
	}
	{
		methodName := StringNameFromStr("get_vertices")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnGetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675695659))
	}
	{
		methodName := StringNameFromStr("set_indices")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnSetIndices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
	}
	{
		methodName := StringNameFromStr("get_indices")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnGetIndices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("has_data")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnHasData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("add_mesh")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnAddMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 975462459))
	}
	{
		methodName := StringNameFromStr("add_mesh_array")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnAddMeshArray = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4235710913))
	}
	{
		methodName := StringNameFromStr("add_faces")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnAddFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1440358797))
	}
	{
		methodName := StringNameFromStr("merge")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnMerge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 655828145))
	}
	{
		methodName := StringNameFromStr("add_projected_obstruction")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnAddProjectedObstruction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3351846707))
	}
	{
		methodName := StringNameFromStr("clear_projected_obstructions")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnClearProjectedObstructions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_projected_obstructions")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnSetProjectedObstructions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_projected_obstructions")
		defer methodName.Destroy()
		ptrsForNavigationMeshSourceGeometryData3D.fnGetProjectedObstructions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}

}

type NavigationMeshSourceGeometryData3D struct {
	Resource
}

func (me *NavigationMeshSourceGeometryData3D) BaseClass() string {
	return "NavigationMeshSourceGeometryData3D"
}

func NewNavigationMeshSourceGeometryData3D() *NavigationMeshSourceGeometryData3D {
	str := StringNameFromStr("NavigationMeshSourceGeometryData3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationMeshSourceGeometryData3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *NavigationMeshSourceGeometryData3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationMeshSourceGeometryData3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationMeshSourceGeometryData3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationMeshSourceGeometryData3D) SetVertices(vertices PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnSetVertices), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData3D) GetVertices() PackedFloat32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedFloat32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnGetVertices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationMeshSourceGeometryData3D) SetIndices(indices PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{indices.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnSetIndices), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData3D) GetIndices() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnGetIndices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationMeshSourceGeometryData3D) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData3D) HasData() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnHasData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationMeshSourceGeometryData3D) AddMesh(mesh Mesh, xform Transform3D) {
	cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnAddMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData3D) AddMeshArray(mesh_array Array, xform Transform3D) {
	cargs := []gdc.ConstTypePtr{mesh_array.AsCTypePtr(), xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnAddMeshArray), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData3D) AddFaces(faces PackedVector3Array, xform Transform3D) {
	cargs := []gdc.ConstTypePtr{faces.AsCTypePtr(), xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnAddFaces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData3D) Merge(other_geometry NavigationMeshSourceGeometryData3D) {
	cargs := []gdc.ConstTypePtr{other_geometry.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnMerge), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData3D) AddProjectedObstruction(vertices PackedVector3Array, elevation float64, height float64, carve bool) {
	cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), gdc.ConstTypePtr(&elevation), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&carve)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnAddProjectedObstruction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData3D) ClearProjectedObstructions() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnClearProjectedObstructions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData3D) SetProjectedObstructions(projected_obstructions Array) {
	cargs := []gdc.ConstTypePtr{projected_obstructions.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnSetProjectedObstructions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationMeshSourceGeometryData3D) GetProjectedObstructions() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMeshSourceGeometryData3D.fnGetProjectedObstructions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
