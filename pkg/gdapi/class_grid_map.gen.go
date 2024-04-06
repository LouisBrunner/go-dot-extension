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

type ptrsForGridMapList struct {
	fnSetCollisionLayer           gdc.MethodBindPtr
	fnGetCollisionLayer           gdc.MethodBindPtr
	fnSetCollisionMask            gdc.MethodBindPtr
	fnGetCollisionMask            gdc.MethodBindPtr
	fnSetCollisionMaskValue       gdc.MethodBindPtr
	fnGetCollisionMaskValue       gdc.MethodBindPtr
	fnSetCollisionLayerValue      gdc.MethodBindPtr
	fnGetCollisionLayerValue      gdc.MethodBindPtr
	fnSetCollisionPriority        gdc.MethodBindPtr
	fnGetCollisionPriority        gdc.MethodBindPtr
	fnSetPhysicsMaterial          gdc.MethodBindPtr
	fnGetPhysicsMaterial          gdc.MethodBindPtr
	fnSetBakeNavigation           gdc.MethodBindPtr
	fnIsBakingNavigation          gdc.MethodBindPtr
	fnSetNavigationMap            gdc.MethodBindPtr
	fnGetNavigationMap            gdc.MethodBindPtr
	fnSetMeshLibrary              gdc.MethodBindPtr
	fnGetMeshLibrary              gdc.MethodBindPtr
	fnSetCellSize                 gdc.MethodBindPtr
	fnGetCellSize                 gdc.MethodBindPtr
	fnSetCellScale                gdc.MethodBindPtr
	fnGetCellScale                gdc.MethodBindPtr
	fnSetOctantSize               gdc.MethodBindPtr
	fnGetOctantSize               gdc.MethodBindPtr
	fnSetCellItem                 gdc.MethodBindPtr
	fnGetCellItem                 gdc.MethodBindPtr
	fnGetCellItemOrientation      gdc.MethodBindPtr
	fnGetCellItemBasis            gdc.MethodBindPtr
	fnGetBasisWithOrthogonalIndex gdc.MethodBindPtr
	fnGetOrthogonalIndexFromBasis gdc.MethodBindPtr
	fnLocalToMap                  gdc.MethodBindPtr
	fnMapToLocal                  gdc.MethodBindPtr
	fnResourceChanged             gdc.MethodBindPtr
	fnSetCenterX                  gdc.MethodBindPtr
	fnGetCenterX                  gdc.MethodBindPtr
	fnSetCenterY                  gdc.MethodBindPtr
	fnGetCenterY                  gdc.MethodBindPtr
	fnSetCenterZ                  gdc.MethodBindPtr
	fnGetCenterZ                  gdc.MethodBindPtr
	fnClear                       gdc.MethodBindPtr
	fnGetUsedCells                gdc.MethodBindPtr
	fnGetUsedCellsByItem          gdc.MethodBindPtr
	fnGetMeshes                   gdc.MethodBindPtr
	fnGetBakeMeshes               gdc.MethodBindPtr
	fnGetBakeMeshInstance         gdc.MethodBindPtr
	fnClearBakedMeshes            gdc.MethodBindPtr
	fnMakeBakedMeshes             gdc.MethodBindPtr
}

var ptrsForGridMap ptrsForGridMapList

func initGridMapPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GridMap")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_collision_layer")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_layer")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_mask_value")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_collision_mask_value")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_collision_layer_value")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCollisionLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_collision_layer_value")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCollisionLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_collision_priority")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_collision_priority")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_physics_material")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetPhysicsMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1784508650))
	}
	{
		methodName := StringNameFromStr("get_physics_material")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetPhysicsMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2521850424))
	}
	{
		methodName := StringNameFromStr("set_bake_navigation")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetBakeNavigation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_baking_navigation")
		defer methodName.Destroy()
		ptrsForGridMap.fnIsBakingNavigation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_navigation_map")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_navigation_map")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_mesh_library")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetMeshLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1488083439))
	}
	{
		methodName := StringNameFromStr("get_mesh_library")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetMeshLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3350993772))
	}
	{
		methodName := StringNameFromStr("set_cell_size")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_cell_size")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_cell_scale")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCellScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_cell_scale")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCellScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_octant_size")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetOctantSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_octant_size")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetOctantSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_cell_item")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCellItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3449088946))
	}
	{
		methodName := StringNameFromStr("get_cell_item")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCellItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3724960147))
	}
	{
		methodName := StringNameFromStr("get_cell_item_orientation")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCellItemOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3724960147))
	}
	{
		methodName := StringNameFromStr("get_cell_item_basis")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCellItemBasis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3493604918))
	}
	{
		methodName := StringNameFromStr("get_basis_with_orthogonal_index")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetBasisWithOrthogonalIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2816196998))
	}
	{
		methodName := StringNameFromStr("get_orthogonal_index_from_basis")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetOrthogonalIndexFromBasis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4210359952))
	}
	{
		methodName := StringNameFromStr("local_to_map")
		defer methodName.Destroy()
		ptrsForGridMap.fnLocalToMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1257687843))
	}
	{
		methodName := StringNameFromStr("map_to_local")
		defer methodName.Destroy()
		ptrsForGridMap.fnMapToLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1088329196))
	}
	{
		methodName := StringNameFromStr("resource_changed")
		defer methodName.Destroy()
		ptrsForGridMap.fnResourceChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 968641751))
	}
	{
		methodName := StringNameFromStr("set_center_x")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCenterX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_center_x")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCenterX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_center_y")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCenterY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_center_y")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCenterY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_center_z")
		defer methodName.Destroy()
		ptrsForGridMap.fnSetCenterZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_center_z")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetCenterZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForGridMap.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_used_cells")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetUsedCells = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_used_cells_by_item")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetUsedCellsByItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
	}
	{
		methodName := StringNameFromStr("get_meshes")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetMeshes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_bake_meshes")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetBakeMeshes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("get_bake_mesh_instance")
		defer methodName.Destroy()
		ptrsForGridMap.fnGetBakeMeshInstance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 937000113))
	}
	{
		methodName := StringNameFromStr("clear_baked_meshes")
		defer methodName.Destroy()
		ptrsForGridMap.fnClearBakedMeshes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("make_baked_meshes")
		defer methodName.Destroy()
		ptrsForGridMap.fnMakeBakedMeshes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3609286057))
	}

}

type GridMap struct {
	Node3D
}

func (me *GridMap) BaseClass() string {
	return "GridMap"
}

func NewGridMap() *GridMap {
	str := StringNameFromStr("GridMap") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GridMap{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	GridMapInvalidCellItem = -1
)

// Enums

func (me *GridMap) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GridMap) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GridMap) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GridMap) SetCollisionLayer(layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCollisionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCollisionLayer() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCollisionLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) SetCollisionMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) SetCollisionMaskValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCollisionMaskValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) SetCollisionLayerValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCollisionLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCollisionLayerValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCollisionLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) SetCollisionPriority(priority float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCollisionPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCollisionPriority() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCollisionPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) SetPhysicsMaterial(material PhysicsMaterial) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetPhysicsMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetPhysicsMaterial() PhysicsMaterial {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetPhysicsMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) SetBakeNavigation(bake_navigation bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bake_navigation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetBakeNavigation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) IsBakingNavigation() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnIsBakingNavigation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) SetNavigationMap(navigation_map RID) {
	cargs := []gdc.ConstTypePtr{navigation_map.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetNavigationMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetNavigationMap() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) SetMeshLibrary(mesh_library MeshLibrary) {
	cargs := []gdc.ConstTypePtr{mesh_library.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetMeshLibrary), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetMeshLibrary() MeshLibrary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMeshLibrary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetMeshLibrary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) SetCellSize(size Vector3) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCellSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCellSize() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCellSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) SetCellScale(scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCellScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCellScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCellScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) SetOctantSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetOctantSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetOctantSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetOctantSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) SetCellItem(position Vector3i, item int64, orientation int64) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&item), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCellItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCellItem(position Vector3i) int64 {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCellItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) GetCellItemOrientation(position Vector3i) int64 {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCellItemOrientation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) GetCellItemBasis(position Vector3i) Basis {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBasis()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCellItemBasis), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) GetBasisWithOrthogonalIndex(index int64) Basis {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBasis()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetBasisWithOrthogonalIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) GetOrthogonalIndexFromBasis(basis Basis) int64 {
	cargs := []gdc.ConstTypePtr{basis.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetOrthogonalIndexFromBasis), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) LocalToMap(local_position Vector3) Vector3i {
	cargs := []gdc.ConstTypePtr{local_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnLocalToMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) MapToLocal(map_position Vector3i) Vector3 {
	cargs := []gdc.ConstTypePtr{map_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnMapToLocal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) ResourceChanged(resource Resource) {
	cargs := []gdc.ConstTypePtr{resource.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnResourceChanged), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) SetCenterX(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCenterX), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCenterX() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCenterX), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) SetCenterY(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCenterY), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCenterY() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCenterY), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) SetCenterZ(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnSetCenterZ), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetCenterZ() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetCenterZ), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GridMap) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) GetUsedCells() []Vector3i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetUsedCells), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector3i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *GridMap) GetUsedCellsByItem(item int64) []Vector3i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&item)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&item)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetUsedCellsByItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector3i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *GridMap) GetMeshes() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetMeshes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) GetBakeMeshes() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetBakeMeshes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) GetBakeMeshInstance(idx int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnGetBakeMeshInstance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GridMap) ClearBakedMeshes() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnClearBakedMeshes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridMap) MakeBakedMeshes(gen_lightmap_uv bool, lightmap_uv_texel_size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gen_lightmap_uv), gdc.ConstTypePtr(&lightmap_uv_texel_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridMap.fnMakeBakedMeshes), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GridMapCellSizeChangedSignalFn func(cell_size Vector3)

func (me *GridMap) ConnectCellSizeChanged(subs SignalSubscribers, fn GridMapCellSizeChangedSignalFn) {
	sig := StringNameFromStr("cell_size_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GridMap) DisconnectCellSizeChanged(subs SignalSubscribers, fn GridMapCellSizeChangedSignalFn) {
	sig := StringNameFromStr("cell_size_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GridMapChangedSignalFn func()

func (me *GridMap) ConnectChanged(subs SignalSubscribers, fn GridMapChangedSignalFn) {
	sig := StringNameFromStr("changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GridMap) DisconnectChanged(subs SignalSubscribers, fn GridMapChangedSignalFn) {
	sig := StringNameFromStr("changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
