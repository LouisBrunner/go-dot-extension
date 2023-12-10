// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TileSet struct {
  obj gdc.ObjectPtr
}

func (me *TileSet) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TileSet) BaseClass() string {
  return "TileSet"
}



// Enums

type TileSetTileShape int
const (
  TileSetTileShapeTileShapeSquare TileSetTileShape = 0
  TileSetTileShapeTileShapeIsometric TileSetTileShape = 1
  TileSetTileShapeTileShapeHalfOffsetSquare TileSetTileShape = 2
  TileSetTileShapeTileShapeHexagon TileSetTileShape = 3
)

type TileSetTileLayout int
const (
  TileSetTileLayoutTileLayoutStacked TileSetTileLayout = 0
  TileSetTileLayoutTileLayoutStackedOffset TileSetTileLayout = 1
  TileSetTileLayoutTileLayoutStairsRight TileSetTileLayout = 2
  TileSetTileLayoutTileLayoutStairsDown TileSetTileLayout = 3
  TileSetTileLayoutTileLayoutDiamondRight TileSetTileLayout = 4
  TileSetTileLayoutTileLayoutDiamondDown TileSetTileLayout = 5
)

type TileSetTileOffsetAxis int
const (
  TileSetTileOffsetAxisTileOffsetAxisHorizontal TileSetTileOffsetAxis = 0
  TileSetTileOffsetAxisTileOffsetAxisVertical TileSetTileOffsetAxis = 1
)

type TileSetCellNeighbor int
const (
  TileSetCellNeighborCellNeighborRightSide TileSetCellNeighbor = 0
  TileSetCellNeighborCellNeighborRightCorner TileSetCellNeighbor = 1
  TileSetCellNeighborCellNeighborBottomRightSide TileSetCellNeighbor = 2
  TileSetCellNeighborCellNeighborBottomRightCorner TileSetCellNeighbor = 3
  TileSetCellNeighborCellNeighborBottomSide TileSetCellNeighbor = 4
  TileSetCellNeighborCellNeighborBottomCorner TileSetCellNeighbor = 5
  TileSetCellNeighborCellNeighborBottomLeftSide TileSetCellNeighbor = 6
  TileSetCellNeighborCellNeighborBottomLeftCorner TileSetCellNeighbor = 7
  TileSetCellNeighborCellNeighborLeftSide TileSetCellNeighbor = 8
  TileSetCellNeighborCellNeighborLeftCorner TileSetCellNeighbor = 9
  TileSetCellNeighborCellNeighborTopLeftSide TileSetCellNeighbor = 10
  TileSetCellNeighborCellNeighborTopLeftCorner TileSetCellNeighbor = 11
  TileSetCellNeighborCellNeighborTopSide TileSetCellNeighbor = 12
  TileSetCellNeighborCellNeighborTopCorner TileSetCellNeighbor = 13
  TileSetCellNeighborCellNeighborTopRightSide TileSetCellNeighbor = 14
  TileSetCellNeighborCellNeighborTopRightCorner TileSetCellNeighbor = 15
)

type TileSetTerrainMode int
const (
  TileSetTerrainModeTerrainModeMatchCornersAndSides TileSetTerrainMode = 0
  TileSetTerrainModeTerrainModeMatchCorners TileSetTerrainMode = 1
  TileSetTerrainModeTerrainModeMatchSides TileSetTerrainMode = 2
)

func (me *TileSet) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TileSet) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TileSet) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TileSet) GetNextSourceId() int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_source_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) AddSource(source TileSetSource, atlas_source_id_override int, ) int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 276991387) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(source.AsCTypePtr()), gdc.ConstTypePtr(&atlas_source_id_override), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) RemoveSource(source_id int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) SetSourceId(source_id int, new_source_id int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_id), gdc.ConstTypePtr(&new_source_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetSourceCount() int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) GetSourceId(index int, ) int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) HasSource(source_id int, ) bool {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) GetSource(source_id int, ) TileSetSource {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1763540252) // FIXME: should cache?
  var ret TileSetSource
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetTileShape(shape TileSetTileShape, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2131427112) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetTileShape() TileSetTileShape {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 716918169) // FIXME: should cache?
  var ret TileSetTileShape
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetTileLayout(layout TileSetTileLayout, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_layout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1071216679) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layout), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetTileLayout() TileSetTileLayout {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_layout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 194628839) // FIXME: should cache?
  var ret TileSetTileLayout
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetTileOffsetAxis(alignment TileSetTileOffsetAxis, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_offset_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3300198521) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetTileOffsetAxis() TileSetTileOffsetAxis {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_offset_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 762494114) // FIXME: should cache?
  var ret TileSetTileOffsetAxis
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetTileSize(size Vector2i, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetTileSize() Vector2i {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetUvClipping(uv_clipping bool, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv_clipping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&uv_clipping), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) IsUvClipping() bool {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_uv_clipping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) GetOcclusionLayersCount() int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_occlusion_layers_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) AddOcclusionLayer(to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_occlusion_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) MoveOcclusionLayer(layer_index int, to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_occlusion_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) RemoveOcclusionLayer(layer_index int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_occlusion_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) SetOcclusionLayerLightMask(layer_index int, light_mask int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_occlusion_layer_light_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&light_mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetOcclusionLayerLightMask(layer_index int, ) int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_occlusion_layer_light_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetOcclusionLayerSdfCollision(layer_index int, sdf_collision bool, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_occlusion_layer_sdf_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&sdf_collision), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetOcclusionLayerSdfCollision(layer_index int, ) bool {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_occlusion_layer_sdf_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) GetPhysicsLayersCount() int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_layers_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) AddPhysicsLayer(to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_physics_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) MovePhysicsLayer(layer_index int, to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_physics_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) RemovePhysicsLayer(layer_index int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_physics_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) SetPhysicsLayerCollisionLayer(layer_index int, layer int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_layer_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetPhysicsLayerCollisionLayer(layer_index int, ) int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_layer_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetPhysicsLayerCollisionMask(layer_index int, mask int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_layer_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetPhysicsLayerCollisionMask(layer_index int, ) int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_layer_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetPhysicsLayerPhysicsMaterial(layer_index int, physics_material PhysicsMaterial, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_layer_physics_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1018687357) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(physics_material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetPhysicsLayerPhysicsMaterial(layer_index int, ) PhysicsMaterial {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_layer_physics_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 788318639) // FIXME: should cache?
  var ret PhysicsMaterial
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) GetTerrainSetsCount() int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_terrain_sets_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) AddTerrainSet(to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_terrain_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) MoveTerrainSet(terrain_set int, to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_terrain_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) RemoveTerrainSet(terrain_set int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_terrain_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) SetTerrainSetMode(terrain_set int, mode TileSetTerrainMode, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_terrain_set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3943003916) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetTerrainSetMode(terrain_set int, ) TileSetTerrainMode {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_terrain_set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2084469411) // FIXME: should cache?
  var ret TileSetTerrainMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) GetTerrainsCount(terrain_set int, ) int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_terrains_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) AddTerrain(terrain_set int, to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_terrain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3023605688) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) MoveTerrain(terrain_set int, terrain_index int, to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_terrain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1649997291) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index), gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) RemoveTerrain(terrain_set int, terrain_index int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_terrain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) SetTerrainName(terrain_set int, terrain_index int, name String, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_terrain_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2285447957) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetTerrainName(terrain_set int, terrain_index int, ) String {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_terrain_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1391810591) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetTerrainColor(terrain_set int, terrain_index int, color Color, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_terrain_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3733378741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetTerrainColor(terrain_set int, terrain_index int, ) Color {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_terrain_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2165839948) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) GetNavigationLayersCount() int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layers_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) AddNavigationLayer(to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_navigation_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) MoveNavigationLayer(layer_index int, to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_navigation_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) RemoveNavigationLayer(layer_index int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_navigation_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) SetNavigationLayerLayers(layer_index int, layers int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layer_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetNavigationLayerLayers(layer_index int, ) int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layer_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetNavigationLayerLayerValue(layer_index int, layer_number int, value bool, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layer_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1383440665) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetNavigationLayerLayerValue(layer_index int, layer_number int, ) bool {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layer_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) GetCustomDataLayersCount() int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_data_layers_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) AddCustomDataLayer(to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_custom_data_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) MoveCustomDataLayer(layer_index int, to_position int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_custom_data_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) RemoveCustomDataLayer(layer_index int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_custom_data_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetCustomDataLayerByName(layer_name String, ) int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_data_layer_by_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(layer_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetCustomDataLayerName(layer_index int, layer_name String, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_data_layer_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(layer_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetCustomDataLayerName(layer_index int, ) String {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_data_layer_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetCustomDataLayerType(layer_index int, layer_type VariantType, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_data_layer_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3492912874) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&layer_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetCustomDataLayerType(layer_index int, ) VariantType {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_data_layer_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2990820875) // FIXME: should cache?
  var ret VariantType
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) SetSourceLevelTileProxy(source_from int, source_to int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), gdc.ConstTypePtr(&source_to), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetSourceLevelTileProxy(source_from int, ) int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) HasSourceLevelTileProxy(source_from int, ) bool {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_source_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) RemoveSourceLevelTileProxy(source_from int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_source_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) SetCoordsLevelTileProxy(p_source_from int, coords_from Vector2i, source_to int, coords_to Vector2i, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_coords_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1769939278) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_source_from), gdc.ConstTypePtr(coords_from.AsCTypePtr()), gdc.ConstTypePtr(&source_to), gdc.ConstTypePtr(coords_to.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetCoordsLevelTileProxy(source_from int, coords_from Vector2i, ) Array {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_coords_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2856536371) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), gdc.ConstTypePtr(coords_from.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) HasCoordsLevelTileProxy(source_from int, coords_from Vector2i, ) bool {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_coords_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3957903770) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), gdc.ConstTypePtr(coords_from.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) RemoveCoordsLevelTileProxy(source_from int, coords_from Vector2i, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_coords_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311374912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), gdc.ConstTypePtr(coords_from.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) SetAlternativeLevelTileProxy(source_from int, coords_from Vector2i, alternative_from int, source_to int, coords_to Vector2i, alternative_to int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alternative_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3862385460) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), gdc.ConstTypePtr(coords_from.AsCTypePtr()), gdc.ConstTypePtr(&alternative_from), gdc.ConstTypePtr(&source_to), gdc.ConstTypePtr(coords_to.AsCTypePtr()), gdc.ConstTypePtr(&alternative_to), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetAlternativeLevelTileProxy(source_from int, coords_from Vector2i, alternative_from int, ) Array {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alternative_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2303761075) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), gdc.ConstTypePtr(coords_from.AsCTypePtr()), gdc.ConstTypePtr(&alternative_from), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) HasAlternativeLevelTileProxy(source_from int, coords_from Vector2i, alternative_from int, ) bool {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_alternative_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 180086755) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), gdc.ConstTypePtr(coords_from.AsCTypePtr()), gdc.ConstTypePtr(&alternative_from), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) RemoveAlternativeLevelTileProxy(source_from int, coords_from Vector2i, alternative_from int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_alternative_level_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2328951467) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), gdc.ConstTypePtr(coords_from.AsCTypePtr()), gdc.ConstTypePtr(&alternative_from), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) MapTileProxy(source_from int, coords_from Vector2i, alternative_from int, ) Array {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_tile_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4267935328) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), gdc.ConstTypePtr(coords_from.AsCTypePtr()), gdc.ConstTypePtr(&alternative_from), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) CleanupInvalidTileProxies()  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cleanup_invalid_tile_proxies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) ClearTileProxies()  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_tile_proxies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) AddPattern(pattern TileMapPattern, index int, ) int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_pattern")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3009264082) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pattern.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) GetPattern(index int, ) TileMapPattern {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pattern")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4207737510) // FIXME: should cache?
  var ret TileMapPattern
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSet) RemovePattern(index int, )  {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_pattern")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSet) GetPatternsCount() int {
  classNameV := StringNameFromStr("TileSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_patterns_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *TileSet) GetPropTileShape() int {
  panic("TODO: implement")
}

func (me *TileSet) SetPropTileShape(value int) {
  panic("TODO: implement")
}

func (me *TileSet) GetPropTileLayout() int {
  panic("TODO: implement")
}

func (me *TileSet) SetPropTileLayout(value int) {
  panic("TODO: implement")
}

func (me *TileSet) GetPropTileOffsetAxis() int {
  panic("TODO: implement")
}

func (me *TileSet) SetPropTileOffsetAxis(value int) {
  panic("TODO: implement")
}

func (me *TileSet) GetPropTileSize() Vector2i {
  panic("TODO: implement")
}

func (me *TileSet) SetPropTileSize(value Vector2i) {
  panic("TODO: implement")
}

func (me *TileSet) GetPropUvClipping() bool {
  panic("TODO: implement")
}

func (me *TileSet) SetPropUvClipping(value bool) {
  panic("TODO: implement")
}