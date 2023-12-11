// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GridMap struct {
  obj gdc.ObjectPtr
}

func (me *GridMap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GridMap) BaseClass() string {
  return "GridMap"
}



// Constants

var (
  GridMapInvalidCellItem = "-1" // TODO: construct correctly
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

func  (me *GridMap) SetCollisionLayer(layer int, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCollisionLayer() int {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetCollisionMask(mask int, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCollisionMask() int {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetCollisionMaskValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCollisionMaskValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetCollisionLayerValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCollisionLayerValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetCollisionPriority(priority float32, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCollisionPriority() float32 {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetPhysicsMaterial(material PhysicsMaterial, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1784508650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetPhysicsMaterial() PhysicsMaterial {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2521850424) // FIXME: should cache?
  var ret PhysicsMaterial
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetBakeNavigation(bake_navigation bool, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_navigation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bake_navigation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) IsBakingNavigation() bool {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_baking_navigation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetNavigationMap(navigation_map RID, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(navigation_map.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetNavigationMap() RID {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetMeshLibrary(mesh_library MeshLibrary, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1488083439) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh_library.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetMeshLibrary() MeshLibrary {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3350993772) // FIXME: should cache?
  var ret MeshLibrary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetCellSize(size Vector3, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCellSize() Vector3 {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetCellScale(scale float32, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCellScale() float32 {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetOctantSize(size int, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_octant_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetOctantSize() int {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_octant_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetCellItem(position Vector3i, item int, orientation int, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4177201334) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&item), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCellItem(position Vector3i, ) int {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3724960147) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) GetCellItemOrientation(position Vector3i, ) int {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_item_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3724960147) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) GetCellItemBasis(position Vector3i, ) Basis {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_item_basis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3493604918) // FIXME: should cache?
  var ret Basis
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) GetBasisWithOrthogonalIndex(index int, ) Basis {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_basis_with_orthogonal_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2816196998) // FIXME: should cache?
  var ret Basis
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) GetOrthogonalIndexFromBasis(basis Basis, ) int {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_orthogonal_index_from_basis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4210359952) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(basis.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) LocalToMap(local_position Vector3, ) Vector3i {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("local_to_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1257687843) // FIXME: should cache?
  var ret Vector3i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) MapToLocal(map_position Vector3i, ) Vector3 {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_to_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1088329196) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) ResourceChanged(resource Resource, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resource_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 968641751) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(resource.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) SetCenterX(enable bool, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_center_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCenterX() bool {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetCenterY(enable bool, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_center_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCenterY() bool {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) SetCenterZ(enable bool, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_center_z")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetCenterZ() bool {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_z")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) Clear()  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) GetUsedCells() Vector3i {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_used_cells")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Vector3i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) GetUsedCellsByItem(item int, ) Vector3i {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_used_cells_by_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  var ret Vector3i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&item), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) GetMeshes() Array {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_meshes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) GetBakeMeshes() Array {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_meshes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) GetBakeMeshInstance(idx int, ) RID {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_mesh_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 937000113) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GridMap) ClearBakedMeshes()  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_baked_meshes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GridMap) MakeBakedMeshes(gen_lightmap_uv bool, lightmap_uv_texel_size float32, )  {
  classNameV := StringNameFromStr("GridMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_baked_meshes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3609286057) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gen_lightmap_uv), gdc.ConstTypePtr(&lightmap_uv_texel_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GridMapCellSizeChangedSignalFn func(cell_size Vector3, )

func (me *GridMap) ConnectCellSizeChanged(subs SignalSubscribers, fn GridMapCellSizeChangedSignalFn) {
  sig := StringNameFromStr("cell_size_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GridMap) DisconnectCellSizeChanged(subs SignalSubscribers, fn GridMapCellSizeChangedSignalFn) {
  sig := StringNameFromStr("cell_size_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
