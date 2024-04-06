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

type ptrsForTileMapList struct {
	fnXUseTileDataRuntimeUpdate   gdc.MethodBindPtr
	fnXTileDataRuntimeUpdate      gdc.MethodBindPtr
	fnSetNavigationMap            gdc.MethodBindPtr
	fnGetNavigationMap            gdc.MethodBindPtr
	fnForceUpdate                 gdc.MethodBindPtr
	fnSetTileset                  gdc.MethodBindPtr
	fnGetTileset                  gdc.MethodBindPtr
	fnSetRenderingQuadrantSize    gdc.MethodBindPtr
	fnGetRenderingQuadrantSize    gdc.MethodBindPtr
	fnGetLayersCount              gdc.MethodBindPtr
	fnAddLayer                    gdc.MethodBindPtr
	fnMoveLayer                   gdc.MethodBindPtr
	fnRemoveLayer                 gdc.MethodBindPtr
	fnSetLayerName                gdc.MethodBindPtr
	fnGetLayerName                gdc.MethodBindPtr
	fnSetLayerEnabled             gdc.MethodBindPtr
	fnIsLayerEnabled              gdc.MethodBindPtr
	fnSetLayerModulate            gdc.MethodBindPtr
	fnGetLayerModulate            gdc.MethodBindPtr
	fnSetLayerYSortEnabled        gdc.MethodBindPtr
	fnIsLayerYSortEnabled         gdc.MethodBindPtr
	fnSetLayerYSortOrigin         gdc.MethodBindPtr
	fnGetLayerYSortOrigin         gdc.MethodBindPtr
	fnSetLayerZIndex              gdc.MethodBindPtr
	fnGetLayerZIndex              gdc.MethodBindPtr
	fnSetLayerNavigationEnabled   gdc.MethodBindPtr
	fnIsLayerNavigationEnabled    gdc.MethodBindPtr
	fnSetLayerNavigationMap       gdc.MethodBindPtr
	fnGetLayerNavigationMap       gdc.MethodBindPtr
	fnSetCollisionAnimatable      gdc.MethodBindPtr
	fnIsCollisionAnimatable       gdc.MethodBindPtr
	fnSetCollisionVisibilityMode  gdc.MethodBindPtr
	fnGetCollisionVisibilityMode  gdc.MethodBindPtr
	fnSetNavigationVisibilityMode gdc.MethodBindPtr
	fnGetNavigationVisibilityMode gdc.MethodBindPtr
	fnSetCell                     gdc.MethodBindPtr
	fnEraseCell                   gdc.MethodBindPtr
	fnGetCellSourceId             gdc.MethodBindPtr
	fnGetCellAtlasCoords          gdc.MethodBindPtr
	fnGetCellAlternativeTile      gdc.MethodBindPtr
	fnGetCellTileData             gdc.MethodBindPtr
	fnGetCoordsForBodyRid         gdc.MethodBindPtr
	fnGetLayerForBodyRid          gdc.MethodBindPtr
	fnGetPattern                  gdc.MethodBindPtr
	fnMapPattern                  gdc.MethodBindPtr
	fnSetPattern                  gdc.MethodBindPtr
	fnSetCellsTerrainConnect      gdc.MethodBindPtr
	fnSetCellsTerrainPath         gdc.MethodBindPtr
	fnFixInvalidTiles             gdc.MethodBindPtr
	fnClearLayer                  gdc.MethodBindPtr
	fnClear                       gdc.MethodBindPtr
	fnUpdateInternals             gdc.MethodBindPtr
	fnNotifyRuntimeTileDataUpdate gdc.MethodBindPtr
	fnGetSurroundingCells         gdc.MethodBindPtr
	fnGetUsedCells                gdc.MethodBindPtr
	fnGetUsedCellsById            gdc.MethodBindPtr
	fnGetUsedRect                 gdc.MethodBindPtr
	fnMapToLocal                  gdc.MethodBindPtr
	fnLocalToMap                  gdc.MethodBindPtr
	fnGetNeighborCell             gdc.MethodBindPtr
}

var ptrsForTileMap ptrsForTileMapList

func initTileMapPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TileMap")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_navigation_map")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4040184819))
	}
	{
		methodName := StringNameFromStr("get_navigation_map")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 495598643))
	}
	{
		methodName := StringNameFromStr("force_update")
		defer methodName.Destroy()
		ptrsForTileMap.fnForceUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("set_tileset")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetTileset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 774531446))
	}
	{
		methodName := StringNameFromStr("get_tileset")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetTileset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2678226422))
	}
	{
		methodName := StringNameFromStr("set_rendering_quadrant_size")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetRenderingQuadrantSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_rendering_quadrant_size")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetRenderingQuadrantSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_layers_count")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetLayersCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_layer")
		defer methodName.Destroy()
		ptrsForTileMap.fnAddLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("move_layer")
		defer methodName.Destroy()
		ptrsForTileMap.fnMoveLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("remove_layer")
		defer methodName.Destroy()
		ptrsForTileMap.fnRemoveLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_layer_name")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetLayerName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_layer_name")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetLayerName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_layer_enabled")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetLayerEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_layer_enabled")
		defer methodName.Destroy()
		ptrsForTileMap.fnIsLayerEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_layer_modulate")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetLayerModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_layer_modulate")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetLayerModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("set_layer_y_sort_enabled")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetLayerYSortEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_layer_y_sort_enabled")
		defer methodName.Destroy()
		ptrsForTileMap.fnIsLayerYSortEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_layer_y_sort_origin")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetLayerYSortOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_layer_y_sort_origin")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetLayerYSortOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_layer_z_index")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetLayerZIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_layer_z_index")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetLayerZIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_layer_navigation_enabled")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetLayerNavigationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_layer_navigation_enabled")
		defer methodName.Destroy()
		ptrsForTileMap.fnIsLayerNavigationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_layer_navigation_map")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetLayerNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4040184819))
	}
	{
		methodName := StringNameFromStr("get_layer_navigation_map")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetLayerNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 495598643))
	}
	{
		methodName := StringNameFromStr("set_collision_animatable")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetCollisionAnimatable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collision_animatable")
		defer methodName.Destroy()
		ptrsForTileMap.fnIsCollisionAnimatable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collision_visibility_mode")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetCollisionVisibilityMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3193440636))
	}
	{
		methodName := StringNameFromStr("get_collision_visibility_mode")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetCollisionVisibilityMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2026313073))
	}
	{
		methodName := StringNameFromStr("set_navigation_visibility_mode")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetNavigationVisibilityMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3193440636))
	}
	{
		methodName := StringNameFromStr("get_navigation_visibility_mode")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetNavigationVisibilityMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2026313073))
	}
	{
		methodName := StringNameFromStr("set_cell")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetCell = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 966713560))
	}
	{
		methodName := StringNameFromStr("erase_cell")
		defer methodName.Destroy()
		ptrsForTileMap.fnEraseCell = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2311374912))
	}
	{
		methodName := StringNameFromStr("get_cell_source_id")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetCellSourceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 551761942))
	}
	{
		methodName := StringNameFromStr("get_cell_atlas_coords")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetCellAtlasCoords = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1869815066))
	}
	{
		methodName := StringNameFromStr("get_cell_alternative_tile")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetCellAlternativeTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 551761942))
	}
	{
		methodName := StringNameFromStr("get_cell_tile_data")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetCellTileData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2849631287))
	}
	{
		methodName := StringNameFromStr("get_coords_for_body_rid")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetCoordsForBodyRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 291584212))
	}
	{
		methodName := StringNameFromStr("get_layer_for_body_rid")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetLayerForBodyRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3917799429))
	}
	{
		methodName := StringNameFromStr("get_pattern")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetPattern = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2833570986))
	}
	{
		methodName := StringNameFromStr("map_pattern")
		defer methodName.Destroy()
		ptrsForTileMap.fnMapPattern = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1864516957))
	}
	{
		methodName := StringNameFromStr("set_pattern")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetPattern = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1195853946))
	}
	{
		methodName := StringNameFromStr("set_cells_terrain_connect")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetCellsTerrainConnect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3578627656))
	}
	{
		methodName := StringNameFromStr("set_cells_terrain_path")
		defer methodName.Destroy()
		ptrsForTileMap.fnSetCellsTerrainPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3578627656))
	}
	{
		methodName := StringNameFromStr("fix_invalid_tiles")
		defer methodName.Destroy()
		ptrsForTileMap.fnFixInvalidTiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("clear_layer")
		defer methodName.Destroy()
		ptrsForTileMap.fnClearLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForTileMap.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("update_internals")
		defer methodName.Destroy()
		ptrsForTileMap.fnUpdateInternals = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("notify_runtime_tile_data_update")
		defer methodName.Destroy()
		ptrsForTileMap.fnNotifyRuntimeTileDataUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("get_surrounding_cells")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetSurroundingCells = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2673526557))
	}
	{
		methodName := StringNameFromStr("get_used_cells")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetUsedCells = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
	}
	{
		methodName := StringNameFromStr("get_used_cells_by_id")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetUsedCellsById = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2931012785))
	}
	{
		methodName := StringNameFromStr("get_used_rect")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetUsedRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 410525958))
	}
	{
		methodName := StringNameFromStr("map_to_local")
		defer methodName.Destroy()
		ptrsForTileMap.fnMapToLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 108438297))
	}
	{
		methodName := StringNameFromStr("local_to_map")
		defer methodName.Destroy()
		ptrsForTileMap.fnLocalToMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 837806996))
	}
	{
		methodName := StringNameFromStr("get_neighbor_cell")
		defer methodName.Destroy()
		ptrsForTileMap.fnGetNeighborCell = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 986575103))
	}
}

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
	TileMapVisibilityModeVisibilityModeDefault   TileMapVisibilityMode = 0
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

func (me *TileMap) SetNavigationMap(layer int64, map_ RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetNavigationMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetNavigationMap(layer int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) ForceUpdate(layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnForceUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) SetTileset(tileset TileSet) {
	cargs := []gdc.ConstTypePtr{tileset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetTileset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetTileset() TileSet {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTileSet()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetTileset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) SetRenderingQuadrantSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetRenderingQuadrantSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetRenderingQuadrantSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetRenderingQuadrantSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) GetLayersCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetLayersCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) AddLayer(to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnAddLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) MoveLayer(layer int64, to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnMoveLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) RemoveLayer(layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnRemoveLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) SetLayerName(layer int64, name String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetLayerName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetLayerName(layer int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetLayerName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) SetLayerEnabled(layer int64, enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetLayerEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) IsLayerEnabled(layer int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnIsLayerEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) SetLayerModulate(layer int64, modulate Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetLayerModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetLayerModulate(layer int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetLayerModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) SetLayerYSortEnabled(layer int64, y_sort_enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&y_sort_enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetLayerYSortEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) IsLayerYSortEnabled(layer int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnIsLayerYSortEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) SetLayerYSortOrigin(layer int64, y_sort_origin int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&y_sort_origin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetLayerYSortOrigin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetLayerYSortOrigin(layer int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetLayerYSortOrigin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) SetLayerZIndex(layer int64, z_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&z_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetLayerZIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetLayerZIndex(layer int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetLayerZIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) SetLayerNavigationEnabled(layer int64, enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetLayerNavigationEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) IsLayerNavigationEnabled(layer int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnIsLayerNavigationEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) SetLayerNavigationMap(layer int64, map_ RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetLayerNavigationMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetLayerNavigationMap(layer int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetLayerNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) SetCollisionAnimatable(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetCollisionAnimatable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) IsCollisionAnimatable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnIsCollisionAnimatable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) SetCollisionVisibilityMode(collision_visibility_mode TileMapVisibilityMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_visibility_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetCollisionVisibilityMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetCollisionVisibilityMode() TileMapVisibilityMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TileMapVisibilityMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetCollisionVisibilityMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TileMap) SetNavigationVisibilityMode(navigation_visibility_mode TileMapVisibilityMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_visibility_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetNavigationVisibilityMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetNavigationVisibilityMode() TileMapVisibilityMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TileMapVisibilityMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetNavigationVisibilityMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TileMap) SetCell(layer int64, coords Vector2i, source_id int64, atlas_coords Vector2i, alternative_tile int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), coords.AsCTypePtr(), gdc.ConstTypePtr(&source_id), atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetCell), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) EraseCell(layer int64, coords Vector2i) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), coords.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnEraseCell), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetCellSourceId(layer int64, coords Vector2i, use_proxies bool) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), coords.AsCTypePtr(), gdc.ConstTypePtr(&use_proxies)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&layer)
	pinner.Pin(&use_proxies)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetCellSourceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) GetCellAtlasCoords(layer int64, coords Vector2i, use_proxies bool) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), coords.AsCTypePtr(), gdc.ConstTypePtr(&use_proxies)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&layer)
	pinner.Pin(&use_proxies)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetCellAtlasCoords), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) GetCellAlternativeTile(layer int64, coords Vector2i, use_proxies bool) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), coords.AsCTypePtr(), gdc.ConstTypePtr(&use_proxies)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&layer)
	pinner.Pin(&use_proxies)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetCellAlternativeTile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) GetCellTileData(layer int64, coords Vector2i, use_proxies bool) TileData {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), coords.AsCTypePtr(), gdc.ConstTypePtr(&use_proxies)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTileData()
	pinner.Pin(&layer)
	pinner.Pin(&use_proxies)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetCellTileData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) GetCoordsForBodyRid(body RID) Vector2i {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetCoordsForBodyRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) GetLayerForBodyRid(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetLayerForBodyRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMap) GetPattern(layer int64, coords_array []Vector2i) TileMapPattern {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&coords_array)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTileMapPattern()
	pinner.Pin(&layer)
	pinner.Pin(&coords_array)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetPattern), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) MapPattern(position_in_tilemap Vector2i, coords_in_pattern Vector2i, pattern TileMapPattern) Vector2i {
	cargs := []gdc.ConstTypePtr{position_in_tilemap.AsCTypePtr(), coords_in_pattern.AsCTypePtr(), pattern.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnMapPattern), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) SetPattern(layer int64, position Vector2i, pattern TileMapPattern) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), position.AsCTypePtr(), pattern.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetPattern), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) SetCellsTerrainConnect(layer int64, cells []Vector2i, terrain_set int64, terrain int64, ignore_empty_terrains bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&cells), gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain), gdc.ConstTypePtr(&ignore_empty_terrains)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetCellsTerrainConnect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) SetCellsTerrainPath(layer int64, path []Vector2i, terrain_set int64, terrain int64, ignore_empty_terrains bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&path), gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain), gdc.ConstTypePtr(&ignore_empty_terrains)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnSetCellsTerrainPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) FixInvalidTiles() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnFixInvalidTiles), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) ClearLayer(layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnClearLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) UpdateInternals() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnUpdateInternals), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) NotifyRuntimeTileDataUpdate(layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnNotifyRuntimeTileDataUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMap) GetSurroundingCells(coords Vector2i) []Vector2i {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetSurroundingCells), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TileMap) GetUsedCells(layer int64) []Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetUsedCells), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TileMap) GetUsedCellsById(layer int64, source_id int64, atlas_coords Vector2i, alternative_tile int64) []Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&source_id), atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&layer)
	pinner.Pin(&source_id)
	pinner.Pin(&alternative_tile)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetUsedCellsById), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TileMap) GetUsedRect() Rect2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetUsedRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) MapToLocal(map_position Vector2i) Vector2 {
	cargs := []gdc.ConstTypePtr{map_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnMapToLocal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) LocalToMap(local_position Vector2) Vector2i {
	cargs := []gdc.ConstTypePtr{local_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnLocalToMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMap) GetNeighborCell(coords Vector2i, neighbor TileSetCellNeighbor) Vector2i {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), gdc.ConstTypePtr(&neighbor)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&neighbor)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMap.fnGetNeighborCell), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
