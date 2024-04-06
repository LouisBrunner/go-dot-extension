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

type ptrsForTileSetScenesCollectionSourceList struct {
	fnGetSceneTilesCount             gdc.MethodBindPtr
	fnGetSceneTileId                 gdc.MethodBindPtr
	fnHasSceneTileId                 gdc.MethodBindPtr
	fnCreateSceneTile                gdc.MethodBindPtr
	fnSetSceneTileId                 gdc.MethodBindPtr
	fnSetSceneTileScene              gdc.MethodBindPtr
	fnGetSceneTileScene              gdc.MethodBindPtr
	fnSetSceneTileDisplayPlaceholder gdc.MethodBindPtr
	fnGetSceneTileDisplayPlaceholder gdc.MethodBindPtr
	fnRemoveSceneTile                gdc.MethodBindPtr
	fnGetNextSceneTileId             gdc.MethodBindPtr
}

var ptrsForTileSetScenesCollectionSource ptrsForTileSetScenesCollectionSourceList

func initTileSetScenesCollectionSourcePtrs(iface gdc.Interface) {

	className := StringNameFromStr("TileSetScenesCollectionSource")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_scene_tiles_count")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnGetSceneTilesCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_scene_tile_id")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnGetSceneTileId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
	}
	{
		methodName := StringNameFromStr("has_scene_tile_id")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnHasSceneTileId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3067735520))
	}
	{
		methodName := StringNameFromStr("create_scene_tile")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnCreateSceneTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1117465415))
	}
	{
		methodName := StringNameFromStr("set_scene_tile_id")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnSetSceneTileId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_scene_tile_scene")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnSetSceneTileScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3435852839))
	}
	{
		methodName := StringNameFromStr("get_scene_tile_scene")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnGetSceneTileScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 511017218))
	}
	{
		methodName := StringNameFromStr("set_scene_tile_display_placeholder")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnSetSceneTileDisplayPlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_scene_tile_display_placeholder")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnGetSceneTileDisplayPlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("remove_scene_tile")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnRemoveSceneTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_next_scene_tile_id")
		defer methodName.Destroy()
		ptrsForTileSetScenesCollectionSource.fnGetNextSceneTileId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

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

func (me *TileSetScenesCollectionSource) GetSceneTilesCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnGetSceneTilesCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSetScenesCollectionSource) GetSceneTileId(index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnGetSceneTileId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSetScenesCollectionSource) HasSceneTileId(id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnHasSceneTileId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSetScenesCollectionSource) CreateSceneTile(packed_scene PackedScene, id_override int64) int64 {
	cargs := []gdc.ConstTypePtr{packed_scene.AsCTypePtr(), gdc.ConstTypePtr(&id_override)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&id_override)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnCreateSceneTile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSetScenesCollectionSource) SetSceneTileId(id int64, new_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&new_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnSetSceneTileId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSetScenesCollectionSource) SetSceneTileScene(id int64, packed_scene PackedScene) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), packed_scene.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnSetSceneTileScene), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSetScenesCollectionSource) GetSceneTileScene(id int64) PackedScene {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedScene()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnGetSceneTileScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSetScenesCollectionSource) SetSceneTileDisplayPlaceholder(id int64, display_placeholder bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&display_placeholder)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnSetSceneTileDisplayPlaceholder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSetScenesCollectionSource) GetSceneTileDisplayPlaceholder(id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnGetSceneTileDisplayPlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSetScenesCollectionSource) RemoveSceneTile(id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnRemoveSceneTile), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSetScenesCollectionSource) GetNextSceneTileId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetScenesCollectionSource.fnGetNextSceneTileId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
