// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *GridMap) GetCollisionLayer()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetCollisionMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetCollisionMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCollisionMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *GridMap) SetCollisionLayerValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCollisionLayerValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *GridMap) SetCollisionPriority(priority float32, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCollisionPriority()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetPhysicsMaterial(material PhysicsMaterial, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetPhysicsMaterial()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetBakeNavigation(bake_navigation bool, )  {
  panic("TODO: implement")
}

func  (me *GridMap) IsBakingNavigation()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetNavigationMap(navigation_map RID, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetNavigationMap()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetMeshLibrary(mesh_library MeshLibrary, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetMeshLibrary()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetCellSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCellSize()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetCellScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCellScale()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetOctantSize(size int, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetOctantSize()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetCellItem(position Vector3i, item int, orientation int, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCellItem(position Vector3i, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCellItemOrientation(position Vector3i, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCellItemBasis(position Vector3i, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetBasisWithOrthogonalIndex(index int, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetOrthogonalIndexFromBasis(basis Basis, )  {
  panic("TODO: implement")
}

func  (me *GridMap) LocalToMap(local_position Vector3, )  {
  panic("TODO: implement")
}

func  (me *GridMap) MapToLocal(map_position Vector3i, )  {
  panic("TODO: implement")
}

func  (me *GridMap) ResourceChanged(resource Resource, )  {
  panic("TODO: implement")
}

func  (me *GridMap) SetCenterX(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCenterX()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetCenterY(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCenterY()  {
  panic("TODO: implement")
}

func  (me *GridMap) SetCenterZ(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetCenterZ()  {
  panic("TODO: implement")
}

func  (me *GridMap) Clear()  {
  panic("TODO: implement")
}

func  (me *GridMap) GetUsedCells()  {
  panic("TODO: implement")
}

func  (me *GridMap) GetUsedCellsByItem(item int, )  {
  panic("TODO: implement")
}

func  (me *GridMap) GetMeshes()  {
  panic("TODO: implement")
}

func  (me *GridMap) GetBakeMeshes()  {
  panic("TODO: implement")
}

func  (me *GridMap) GetBakeMeshInstance(idx int, )  {
  panic("TODO: implement")
}

func  (me *GridMap) ClearBakedMeshes()  {
  panic("TODO: implement")
}

func  (me *GridMap) MakeBakedMeshes(gen_lightmap_uv bool, lightmap_uv_texel_size float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
