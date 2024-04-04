// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type TileMap struct {
  Node2D
}

func (me *TileMap) BaseClass() string {
  return "TileMap"
}

func NewTileMap() *TileMap {
  str := StringNameFromStr("TileMap") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TileMap{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type TileMapVisibilityMode int
const (
  TileMapVisibilityModeVisibilityModeDefault TileMapVisibilityMode = 0
  TileMapVisibilityModeVisibilityModeForceHide TileMapVisibilityMode = 2
  TileMapVisibilityModeVisibilityModeForceShow TileMapVisibilityMode = 1
)

func (me *TileMap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TileMap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TileMap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TileMap) SetNavigationMap(layer int64, map_ RID, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4040184819) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetNavigationMap(layer int64, ) RID {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 495598643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) ForceUpdate(layer int64, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) SetTileset(tileset TileSet, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tileset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 774531446) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{tileset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetTileset() TileSet {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tileset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678226422) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTileSet()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) SetRenderingQuadrantSize(size int64, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rendering_quadrant_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetRenderingQuadrantSize() int64 {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rendering_quadrant_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) GetLayersCount() int64 {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layers_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) AddLayer(to_position int64, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) MoveLayer(layer int64, to_position int64, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&to_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) RemoveLayer(layer int64, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) SetLayerName(layer int64, name String, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetLayerName(layer int64, ) String {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) SetLayerEnabled(layer int64, enabled bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) IsLayerEnabled(layer int64, ) bool {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_layer_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) SetLayerModulate(layer int64, modulate Color, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetLayerModulate(layer int64, ) Color {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) SetLayerYSortEnabled(layer int64, y_sort_enabled bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_y_sort_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&y_sort_enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) IsLayerYSortEnabled(layer int64, ) bool {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_layer_y_sort_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) SetLayerYSortOrigin(layer int64, y_sort_origin int64, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_y_sort_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&y_sort_origin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetLayerYSortOrigin(layer int64, ) int64 {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_y_sort_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) SetLayerZIndex(layer int64, z_index int64, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_z_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&z_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetLayerZIndex(layer int64, ) int64 {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_z_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) SetLayerNavigationEnabled(layer int64, enabled bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_navigation_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) IsLayerNavigationEnabled(layer int64, ) bool {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_layer_navigation_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) SetLayerNavigationMap(layer int64, map_ RID, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4040184819) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetLayerNavigationMap(layer int64, ) RID {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 495598643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) SetCollisionAnimatable(enabled bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_animatable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) IsCollisionAnimatable() bool {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collision_animatable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) SetCollisionVisibilityMode(collision_visibility_mode TileMapVisibilityMode, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_visibility_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3193440636) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_visibility_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetCollisionVisibilityMode() TileMapVisibilityMode {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_visibility_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2026313073) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TileMapVisibilityMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TileMap) SetNavigationVisibilityMode(navigation_visibility_mode TileMapVisibilityMode, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_visibility_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3193440636) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_visibility_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetNavigationVisibilityMode() TileMapVisibilityMode {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_visibility_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2026313073) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TileMapVisibilityMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TileMap) SetCell(layer int64, coords Vector2i, source_id int64, atlas_coords Vector2i, alternative_tile int64, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 966713560) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , coords.AsCTypePtr(), gdc.ConstTypePtr(&source_id) , atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) EraseCell(layer int64, coords Vector2i, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_cell")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311374912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetCellSourceId(layer int64, coords Vector2i, use_proxies bool, ) int64 {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_source_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 551761942) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , coords.AsCTypePtr(), gdc.ConstTypePtr(&use_proxies) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&layer)
  pinner.Pin(&use_proxies)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) GetCellAtlasCoords(layer int64, coords Vector2i, use_proxies bool, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_atlas_coords")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1869815066) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , coords.AsCTypePtr(), gdc.ConstTypePtr(&use_proxies) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()
  pinner.Pin(&layer)
  pinner.Pin(&use_proxies)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) GetCellAlternativeTile(layer int64, coords Vector2i, use_proxies bool, ) int64 {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_alternative_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 551761942) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , coords.AsCTypePtr(), gdc.ConstTypePtr(&use_proxies) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&layer)
  pinner.Pin(&use_proxies)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) GetCellTileData(layer int64, coords Vector2i, use_proxies bool, ) TileData {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_tile_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2849631287) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , coords.AsCTypePtr(), gdc.ConstTypePtr(&use_proxies) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTileData()
  pinner.Pin(&layer)
  pinner.Pin(&use_proxies)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) GetCoordsForBodyRid(body RID, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_coords_for_body_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 291584212) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) GetLayerForBodyRid(body RID, ) int64 {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_for_body_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3917799429) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMap) GetPattern(layer int64, coords_array []Vector2i, ) TileMapPattern {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pattern")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2833570986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&coords_array) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTileMapPattern()
  pinner.Pin(&layer)
  pinner.Pin(&coords_array)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) MapPattern(position_in_tilemap Vector2i, coords_in_pattern Vector2i, pattern TileMapPattern, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_pattern")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1864516957) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position_in_tilemap.AsCTypePtr(), coords_in_pattern.AsCTypePtr(), pattern.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) SetPattern(layer int64, position Vector2i, pattern TileMapPattern, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pattern")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1195853946) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , position.AsCTypePtr(), pattern.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) SetCellsTerrainConnect(layer int64, cells []Vector2i, terrain_set int64, terrain int64, ignore_empty_terrains bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cells_terrain_connect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3578627656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&cells) , gdc.ConstTypePtr(&terrain_set) , gdc.ConstTypePtr(&terrain) , gdc.ConstTypePtr(&ignore_empty_terrains) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) SetCellsTerrainPath(layer int64, path []Vector2i, terrain_set int64, terrain int64, ignore_empty_terrains bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cells_terrain_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3578627656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&path) , gdc.ConstTypePtr(&terrain_set) , gdc.ConstTypePtr(&terrain) , gdc.ConstTypePtr(&ignore_empty_terrains) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) FixInvalidTiles()  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fix_invalid_tiles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) ClearLayer(layer int64, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) Clear()  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) UpdateInternals()  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_internals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) NotifyRuntimeTileDataUpdate(layer int64, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("notify_runtime_tile_data_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMap) GetSurroundingCells(coords Vector2i, ) []Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surrounding_cells")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2673526557) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Vector2i](ret)
}

func  (me *TileMap) GetUsedCells(layer int64, ) []Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_used_cells")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Vector2i](ret)
}

func  (me *TileMap) GetUsedCellsById(layer int64, source_id int64, atlas_coords Vector2i, alternative_tile int64, ) []Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_used_cells_by_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2931012785) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&source_id) , atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&layer)
  pinner.Pin(&source_id)
  pinner.Pin(&alternative_tile)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Vector2i](ret)
}

func  (me *TileMap) GetUsedRect() Rect2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_used_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 410525958) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) MapToLocal(map_position Vector2i, ) Vector2 {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_to_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 108438297) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) LocalToMap(local_position Vector2, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("local_to_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 837806996) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{local_position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMap) GetNeighborCell(coords Vector2i, neighbor TileSetCellNeighbor, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_neighbor_cell")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 986575103) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), gdc.ConstTypePtr(&neighbor) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()
  pinner.Pin(&neighbor)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TileMapChangedSignalFn func()

func (me *TileMap) ConnectChanged(subs SignalSubscribers, fn TileMapChangedSignalFn) {
  sig := StringNameFromStr("changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *TileMap) DisconnectChanged(subs SignalSubscribers, fn TileMapChangedSignalFn) {
  sig := StringNameFromStr("changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
