// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TileData struct {
  obj gdc.ObjectPtr
}

func (me *TileData) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TileData) BaseClass() string {
  return "TileData"
}



// Enums

func (me *TileData) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TileData) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TileData) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TileData) SetFlipH(flip_h bool, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_h), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetFlipH() bool {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flip_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetFlipV(flip_v bool, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_v), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetFlipV() bool {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flip_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetTranspose(transpose bool, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transpose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&transpose), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetTranspose() bool {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transpose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetMaterial() Material {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  var ret Material
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetTextureOrigin(texture_origin Vector2i, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture_origin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetTextureOrigin() Vector2i {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetModulate(modulate Color, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetModulate() Color {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetZIndex(z_index int, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_z_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&z_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetZIndex() int {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_z_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetYSortOrigin(y_sort_origin int, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_y_sort_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&y_sort_origin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetYSortOrigin() int {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_y_sort_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetOccluder(layer_id int, occluder_polygon OccluderPolygon2D, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_occluder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 914399637) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(occluder_polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetOccluder(layer_id int, ) OccluderPolygon2D {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_occluder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2458574231) // FIXME: should cache?
  var ret OccluderPolygon2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetConstantLinearVelocity(layer_id int, velocity Vector2, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetConstantLinearVelocity(layer_id int, ) Vector2 {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetConstantAngularVelocity(layer_id int, velocity float32, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&velocity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetConstantAngularVelocity(layer_id int, ) float32 {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetCollisionPolygonsCount(layer_id int, polygons_count int, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_polygons_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygons_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetCollisionPolygonsCount(layer_id int, ) int {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_polygons_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) AddCollisionPolygon(layer_id int, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_collision_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) RemoveCollisionPolygon(layer_id int, polygon_index int, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_collision_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) SetCollisionPolygonPoints(layer_id int, polygon_index int, polygon PackedVector2Array, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_polygon_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3230546541) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index), gdc.ConstTypePtr(polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetCollisionPolygonPoints(layer_id int, polygon_index int, ) PackedVector2Array {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_polygon_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 103942801) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetCollisionPolygonOneWay(layer_id int, polygon_index int, one_way bool, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_polygon_one_way")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1383440665) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index), gdc.ConstTypePtr(&one_way), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) IsCollisionPolygonOneWay(layer_id int, polygon_index int, ) bool {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collision_polygon_one_way")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetCollisionPolygonOneWayMargin(layer_id int, polygon_index int, one_way_margin float32, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_polygon_one_way_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index), gdc.ConstTypePtr(&one_way_margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetCollisionPolygonOneWayMargin(layer_id int, polygon_index int, ) float32 {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_polygon_one_way_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetTerrainSet(terrain_set int, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_terrain_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetTerrainSet() int {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_terrain_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetTerrain(terrain int, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_terrain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetTerrain() int {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_terrain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetTerrainPeeringBit(peering_bit TileSetCellNeighbor, terrain int, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_terrain_peering_bit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1084452308) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peering_bit), gdc.ConstTypePtr(&terrain), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetTerrainPeeringBit(peering_bit TileSetCellNeighbor, ) int {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_terrain_peering_bit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3831796792) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peering_bit), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetNavigationPolygon(layer_id int, navigation_polygon NavigationPolygon, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2224691167) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(navigation_polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetNavigationPolygon(layer_id int, ) NavigationPolygon {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3991786031) // FIXME: should cache?
  var ret NavigationPolygon
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetProbability(probability float32, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_probability")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&probability), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetProbability() float32 {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_probability")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetCustomData(layer_name String, value Variant, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 402577236) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(layer_name.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetCustomData(layer_name String, ) Variant {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1868160156) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(layer_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileData) SetCustomDataByLayerId(layer_id int, value Variant, )  {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_data_by_layer_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileData) GetCustomDataByLayerId(layer_id int, ) Variant {
  classNameV := StringNameFromStr("TileData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_data_by_layer_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *TileData) GetPropFlipH() bool {
  panic("TODO: implement")
}

func (me *TileData) SetPropFlipH(value bool) {
  panic("TODO: implement")
}

func (me *TileData) GetPropFlipV() bool {
  panic("TODO: implement")
}

func (me *TileData) SetPropFlipV(value bool) {
  panic("TODO: implement")
}

func (me *TileData) GetPropTranspose() bool {
  panic("TODO: implement")
}

func (me *TileData) SetPropTranspose(value bool) {
  panic("TODO: implement")
}

func (me *TileData) GetPropTextureOrigin() Vector2i {
  panic("TODO: implement")
}

func (me *TileData) SetPropTextureOrigin(value Vector2i) {
  panic("TODO: implement")
}

func (me *TileData) GetPropModulate() Color {
  panic("TODO: implement")
}

func (me *TileData) SetPropModulate(value Color) {
  panic("TODO: implement")
}

func (me *TileData) GetPropMaterial() any {
  panic("TODO: implement")
}

func (me *TileData) SetPropMaterial(value any) {
  panic("TODO: implement")
}

func (me *TileData) GetPropZIndex() int {
  panic("TODO: implement")
}

func (me *TileData) SetPropZIndex(value int) {
  panic("TODO: implement")
}

func (me *TileData) GetPropYSortOrigin() int {
  panic("TODO: implement")
}

func (me *TileData) SetPropYSortOrigin(value int) {
  panic("TODO: implement")
}

func (me *TileData) GetPropTerrainSet() int {
  panic("TODO: implement")
}

func (me *TileData) SetPropTerrainSet(value int) {
  panic("TODO: implement")
}

func (me *TileData) GetPropTerrain() int {
  panic("TODO: implement")
}

func (me *TileData) SetPropTerrain(value int) {
  panic("TODO: implement")
}

func (me *TileData) GetPropProbability() float32 {
  panic("TODO: implement")
}

func (me *TileData) SetPropProbability(value float32) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API