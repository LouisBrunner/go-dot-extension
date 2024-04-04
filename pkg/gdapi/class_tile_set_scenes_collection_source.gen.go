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

type TileSetScenesCollectionSource struct {
  TileSetSource
}

func (me *TileSetScenesCollectionSource) BaseClass() string {
  return "TileSetScenesCollectionSource"
}

func NewTileSetScenesCollectionSource() *TileSetScenesCollectionSource {
  str := StringNameFromStr("TileSetScenesCollectionSource") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TileSetScenesCollectionSource{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *TileSetScenesCollectionSource) GetSceneTilesCount() int64 {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_tiles_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetScenesCollectionSource) GetSceneTileId(index int64, ) int64 {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetScenesCollectionSource) HasSceneTileId(id int64, ) bool {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_scene_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetScenesCollectionSource) CreateSceneTile(packed_scene PackedScene, id_override int64, ) int64 {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_scene_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1117465415) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{packed_scene.AsCTypePtr(), gdc.ConstTypePtr(&id_override) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&id_override)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetScenesCollectionSource) SetSceneTileId(id int64, new_id int64, )  {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scene_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&new_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetScenesCollectionSource) SetSceneTileScene(id int64, packed_scene PackedScene, )  {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scene_tile_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3435852839) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , packed_scene.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetScenesCollectionSource) GetSceneTileScene(id int64, ) PackedScene {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_tile_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 511017218) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedScene()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetScenesCollectionSource) SetSceneTileDisplayPlaceholder(id int64, display_placeholder bool, )  {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scene_tile_display_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&display_placeholder) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetScenesCollectionSource) GetSceneTileDisplayPlaceholder(id int64, ) bool {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_tile_display_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetScenesCollectionSource) RemoveSceneTile(id int64, )  {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_scene_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetScenesCollectionSource) GetNextSceneTileId() int64 {
  classNameV := StringNameFromStr("TileSetScenesCollectionSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_scene_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
