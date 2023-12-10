// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TileSetScenesCollectionSource struct {
  obj gdc.ObjectPtr
}

func (me *TileSetScenesCollectionSource) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TileSetScenesCollectionSource) BaseClass() string {
  return "TileSetScenesCollectionSource"
}



// Enums

func (me *TileSetScenesCollectionSource) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TileSetScenesCollectionSource) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TileSetScenesCollectionSource) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TileSetScenesCollectionSource) GetSceneTilesCount() int {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_tiles_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetScenesCollectionSource) GetSceneTileId(index int, ) int {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetScenesCollectionSource) HasSceneTileId(id int, ) bool {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_scene_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetScenesCollectionSource) CreateSceneTile(packed_scene PackedScene, id_override int, ) int {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_scene_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2633389122) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(packed_scene.AsCTypePtr()), gdc.ConstTypePtr(&id_override), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetScenesCollectionSource) SetSceneTileId(id int, new_id int, )  {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scene_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&new_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetScenesCollectionSource) SetSceneTileScene(id int, packed_scene PackedScene, )  {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scene_tile_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3435852839) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(packed_scene.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetScenesCollectionSource) GetSceneTileScene(id int, ) PackedScene {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_tile_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 511017218) // FIXME: should cache?
  var ret PackedScene
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetScenesCollectionSource) SetSceneTileDisplayPlaceholder(id int, display_placeholder bool, )  {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scene_tile_display_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&display_placeholder), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetScenesCollectionSource) GetSceneTileDisplayPlaceholder(id int, ) bool {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_tile_display_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetScenesCollectionSource) RemoveSceneTile(id int, )  {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_scene_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetScenesCollectionSource) GetNextSceneTileId() int {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_scene_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties