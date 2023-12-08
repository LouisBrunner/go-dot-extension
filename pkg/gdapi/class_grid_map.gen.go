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

// TODO: needed?
// const (
// // )

var (
  GridMapInvalidCellItem = "-1" // TODO: construct correctly
)

func  (me *GridMap) SetCollisionLayer(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCollisionLayer() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetCollisionMask(mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCollisionMask() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetCollisionMaskValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCollisionMaskValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetCollisionLayerValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCollisionLayerValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetCollisionPriority(priority float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCollisionPriority() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetPhysicsMaterial(material PhysicsMaterial, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetPhysicsMaterial() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetBakeNavigation(bake_navigation bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) IsBakingNavigation() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetNavigationMap(navigation_map RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetNavigationMap() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetMeshLibrary(mesh_library MeshLibrary, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetMeshLibrary() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetCellSize(size Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCellSize() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetCellScale(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCellScale() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetOctantSize(size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetOctantSize() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetCellItem(position Vector3i, item int, orientation int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCellItem(position Vector3i, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCellItemOrientation(position Vector3i, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCellItemBasis(position Vector3i, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetBasisWithOrthogonalIndex(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetOrthogonalIndexFromBasis(basis Basis, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) LocalToMap(local_position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) MapToLocal(map_position Vector3i, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) ResourceChanged(resource Resource, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetCenterX(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCenterX() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetCenterY(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCenterY() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) SetCenterZ(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetCenterZ() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetUsedCells() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetUsedCellsByItem(item int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetMeshes() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetBakeMeshes() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) GetBakeMeshInstance(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) ClearBakedMeshes() { // TODO: return value
  // TODO: implement
}

func  (me *GridMap) MakeBakedMeshes(gen_lightmap_uv bool, lightmap_uv_texel_size float32, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
