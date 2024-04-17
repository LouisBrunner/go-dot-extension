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

type ptrsForTileMapLayerList struct {
	fnXUseTileDataRuntimeUpdate   gdc.MethodBindPtr
	fnXTileDataRuntimeUpdate      gdc.MethodBindPtr
	fnSetCell                     gdc.MethodBindPtr
	fnEraseCell                   gdc.MethodBindPtr
	fnFixInvalidTiles             gdc.MethodBindPtr
	fnClear                       gdc.MethodBindPtr
	fnGetCellSourceId             gdc.MethodBindPtr
	fnGetCellAtlasCoords          gdc.MethodBindPtr
	fnGetCellAlternativeTile      gdc.MethodBindPtr
	fnGetCellTileData             gdc.MethodBindPtr
	fnGetUsedCells                gdc.MethodBindPtr
	fnGetUsedCellsById            gdc.MethodBindPtr
	fnGetUsedRect                 gdc.MethodBindPtr
	fnGetPattern                  gdc.MethodBindPtr
	fnSetPattern                  gdc.MethodBindPtr
	fnSetCellsTerrainConnect      gdc.MethodBindPtr
	fnSetCellsTerrainPath         gdc.MethodBindPtr
	fnHasBodyRid                  gdc.MethodBindPtr
	fnGetCoordsForBodyRid         gdc.MethodBindPtr
	fnUpdateInternals             gdc.MethodBindPtr
	fnNotifyRuntimeTileDataUpdate gdc.MethodBindPtr
	fnMapPattern                  gdc.MethodBindPtr
	fnGetSurroundingCells         gdc.MethodBindPtr
	fnGetNeighborCell             gdc.MethodBindPtr
	fnMapToLocal                  gdc.MethodBindPtr
	fnLocalToMap                  gdc.MethodBindPtr
	fnSetTileMapDataFromArray     gdc.MethodBindPtr
	fnGetTileMapDataAsArray       gdc.MethodBindPtr
	fnSetEnabled                  gdc.MethodBindPtr
	fnIsEnabled                   gdc.MethodBindPtr
	fnSetTileSet                  gdc.MethodBindPtr
	fnGetTileSet                  gdc.MethodBindPtr
	fnSetYSortOrigin              gdc.MethodBindPtr
	fnGetYSortOrigin              gdc.MethodBindPtr
	fnSetRenderingQuadrantSize    gdc.MethodBindPtr
	fnGetRenderingQuadrantSize    gdc.MethodBindPtr
	fnSetCollisionEnabled         gdc.MethodBindPtr
	fnIsCollisionEnabled          gdc.MethodBindPtr
	fnSetUseKinematicBodies       gdc.MethodBindPtr
	fnIsUsingKinematicBodies      gdc.MethodBindPtr
	fnSetCollisionVisibilityMode  gdc.MethodBindPtr
	fnGetCollisionVisibilityMode  gdc.MethodBindPtr
	fnSetNavigationEnabled        gdc.MethodBindPtr
	fnIsNavigationEnabled         gdc.MethodBindPtr
	fnSetNavigationMap            gdc.MethodBindPtr
	fnGetNavigationMap            gdc.MethodBindPtr
	fnSetNavigationVisibilityMode gdc.MethodBindPtr
	fnGetNavigationVisibilityMode gdc.MethodBindPtr
}

var ptrsForTileMapLayer ptrsForTileMapLayerList

func initTileMapLayerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TileMapLayer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_cell")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetCell = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2428518503))
	}
	{
		methodName := StringNameFromStr("erase_cell")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnEraseCell = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("fix_invalid_tiles")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnFixInvalidTiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_cell_source_id")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetCellSourceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2485466453))
	}
	{
		methodName := StringNameFromStr("get_cell_atlas_coords")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetCellAtlasCoords = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3050897911))
	}
	{
		methodName := StringNameFromStr("get_cell_alternative_tile")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetCellAlternativeTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2485466453))
	}
	{
		methodName := StringNameFromStr("get_cell_tile_data")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetCellTileData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 205084707))
	}
	{
		methodName := StringNameFromStr("get_used_cells")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetUsedCells = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_used_cells_by_id")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetUsedCellsById = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4175304538))
	}
	{
		methodName := StringNameFromStr("get_used_rect")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetUsedRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 410525958))
	}
	{
		methodName := StringNameFromStr("get_pattern")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetPattern = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3820813253))
	}
	{
		methodName := StringNameFromStr("set_pattern")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetPattern = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1491151770))
	}
	{
		methodName := StringNameFromStr("set_cells_terrain_connect")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetCellsTerrainConnect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 748968311))
	}
	{
		methodName := StringNameFromStr("set_cells_terrain_path")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetCellsTerrainPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 748968311))
	}
	{
		methodName := StringNameFromStr("has_body_rid")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnHasBodyRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("get_coords_for_body_rid")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetCoordsForBodyRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 733700038))
	}
	{
		methodName := StringNameFromStr("update_internals")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnUpdateInternals = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("notify_runtime_tile_data_update")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnNotifyRuntimeTileDataUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2275361663))
	}
	{
		methodName := StringNameFromStr("map_pattern")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnMapPattern = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1864516957))
	}
	{
		methodName := StringNameFromStr("get_surrounding_cells")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetSurroundingCells = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2673526557))
	}
	{
		methodName := StringNameFromStr("get_neighbor_cell")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetNeighborCell = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 986575103))
	}
	{
		methodName := StringNameFromStr("map_to_local")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnMapToLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 108438297))
	}
	{
		methodName := StringNameFromStr("local_to_map")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnLocalToMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 837806996))
	}
	{
		methodName := StringNameFromStr("set_tile_map_data_from_array")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetTileMapDataFromArray = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2971499966))
	}
	{
		methodName := StringNameFromStr("get_tile_map_data_as_array")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetTileMapDataAsArray = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2362200018))
	}
	{
		methodName := StringNameFromStr("set_enabled")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_enabled")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tile_set")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetTileSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 774531446))
	}
	{
		methodName := StringNameFromStr("get_tile_set")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetTileSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2678226422))
	}
	{
		methodName := StringNameFromStr("set_y_sort_origin")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetYSortOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_y_sort_origin")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetYSortOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_rendering_quadrant_size")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetRenderingQuadrantSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_rendering_quadrant_size")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetRenderingQuadrantSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_enabled")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collision_enabled")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnIsCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_kinematic_bodies")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetUseKinematicBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_kinematic_bodies")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnIsUsingKinematicBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collision_visibility_mode")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetCollisionVisibilityMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3508099847))
	}
	{
		methodName := StringNameFromStr("get_collision_visibility_mode")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetCollisionVisibilityMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 338220793))
	}
	{
		methodName := StringNameFromStr("set_navigation_enabled")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetNavigationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_navigation_enabled")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnIsNavigationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_navigation_map")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_navigation_map")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_navigation_visibility_mode")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnSetNavigationVisibilityMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3508099847))
	}
	{
		methodName := StringNameFromStr("get_navigation_visibility_mode")
		defer methodName.Destroy()
		ptrsForTileMapLayer.fnGetNavigationVisibilityMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 338220793))
	}

}

type TileMapLayer struct {
	Node2D
}

func (me *TileMapLayer) BaseClass() string {
	return "TileMapLayer"
}

func NewTileMapLayer() *TileMapLayer {
	str := StringNameFromStr("TileMapLayer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TileMapLayer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TileMapLayerDebugVisibilityMode int

const (
	TileMapLayerDebugVisibilityModeDebugVisibilityModeDefault   TileMapLayerDebugVisibilityMode = 0
	TileMapLayerDebugVisibilityModeDebugVisibilityModeForceHide TileMapLayerDebugVisibilityMode = 2
	TileMapLayerDebugVisibilityModeDebugVisibilityModeForceShow TileMapLayerDebugVisibilityMode = 1
)

func (me *TileMapLayer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TileMapLayer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TileMapLayer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TileMapLayer) SetCell(coords Vector2i, source_id int64, atlas_coords Vector2i, alternative_tile int64) {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), gdc.ConstTypePtr(&source_id), atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetCell), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) EraseCell(coords Vector2i) {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnEraseCell), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) FixInvalidTiles() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnFixInvalidTiles), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) GetCellSourceId(coords Vector2i) int64 {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetCellSourceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMapLayer) GetCellAtlasCoords(coords Vector2i) Vector2i {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetCellAtlasCoords), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) GetCellAlternativeTile(coords Vector2i) int64 {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetCellAlternativeTile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMapLayer) GetCellTileData(coords Vector2i) TileData {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTileData()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetCellTileData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) GetUsedCells() []Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetUsedCells), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TileMapLayer) GetUsedCellsById(source_id int64, atlas_coords Vector2i, alternative_tile int64) []Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_id), atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&source_id)
	pinner.Pin(&alternative_tile)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetUsedCellsById), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TileMapLayer) GetUsedRect() Rect2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetUsedRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) GetPattern(coords_array []Vector2i) TileMapPattern {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&coords_array)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTileMapPattern()
	pinner.Pin(&coords_array)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetPattern), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) SetPattern(position Vector2i, pattern TileMapPattern) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), pattern.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetPattern), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) SetCellsTerrainConnect(cells []Vector2i, terrain_set int64, terrain int64, ignore_empty_terrains bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cells), gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain), gdc.ConstTypePtr(&ignore_empty_terrains)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetCellsTerrainConnect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) SetCellsTerrainPath(path []Vector2i, terrain_set int64, terrain int64, ignore_empty_terrains bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path), gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain), gdc.ConstTypePtr(&ignore_empty_terrains)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetCellsTerrainPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) HasBodyRid(body RID) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnHasBodyRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMapLayer) GetCoordsForBodyRid(body RID) Vector2i {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetCoordsForBodyRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) UpdateInternals() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnUpdateInternals), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) NotifyRuntimeTileDataUpdate() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnNotifyRuntimeTileDataUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) MapPattern(position_in_tilemap Vector2i, coords_in_pattern Vector2i, pattern TileMapPattern) Vector2i {
	cargs := []gdc.ConstTypePtr{position_in_tilemap.AsCTypePtr(), coords_in_pattern.AsCTypePtr(), pattern.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnMapPattern), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) GetSurroundingCells(coords Vector2i) []Vector2i {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetSurroundingCells), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TileMapLayer) GetNeighborCell(coords Vector2i, neighbor TileSetCellNeighbor) Vector2i {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), gdc.ConstTypePtr(&neighbor)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&neighbor)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetNeighborCell), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) MapToLocal(map_position Vector2i) Vector2 {
	cargs := []gdc.ConstTypePtr{map_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnMapToLocal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) LocalToMap(local_position Vector2) Vector2i {
	cargs := []gdc.ConstTypePtr{local_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnLocalToMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) SetTileMapDataFromArray(tile_map_layer_data PackedByteArray) {
	cargs := []gdc.ConstTypePtr{tile_map_layer_data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetTileMapDataFromArray), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) GetTileMapDataAsArray() PackedByteArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetTileMapDataAsArray), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) SetEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) IsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMapLayer) SetTileSet(tile_set TileSet) {
	cargs := []gdc.ConstTypePtr{tile_set.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetTileSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) GetTileSet() TileSet {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTileSet()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetTileSet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) SetYSortOrigin(y_sort_origin int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&y_sort_origin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetYSortOrigin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) GetYSortOrigin() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetYSortOrigin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMapLayer) SetRenderingQuadrantSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetRenderingQuadrantSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) GetRenderingQuadrantSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetRenderingQuadrantSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMapLayer) SetCollisionEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetCollisionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) IsCollisionEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnIsCollisionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMapLayer) SetUseKinematicBodies(use_kinematic_bodies bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_kinematic_bodies)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetUseKinematicBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) IsUsingKinematicBodies() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnIsUsingKinematicBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMapLayer) SetCollisionVisibilityMode(visibility_mode TileMapLayerDebugVisibilityMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visibility_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetCollisionVisibilityMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) GetCollisionVisibilityMode() TileMapLayerDebugVisibilityMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TileMapLayerDebugVisibilityMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetCollisionVisibilityMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TileMapLayer) SetNavigationEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetNavigationEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) IsNavigationEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnIsNavigationEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileMapLayer) SetNavigationMap(map_ RID) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetNavigationMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) GetNavigationMap() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileMapLayer) SetNavigationVisibilityMode(show_navigation TileMapLayerDebugVisibilityMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show_navigation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnSetNavigationVisibilityMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileMapLayer) GetNavigationVisibilityMode() TileMapLayerDebugVisibilityMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TileMapLayerDebugVisibilityMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapLayer.fnGetNavigationVisibilityMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TileMapLayerChangedSignalFn func()

func (me *TileMapLayer) ConnectChanged(subs SignalSubscribers, fn TileMapLayerChangedSignalFn) {
	sig := StringNameFromStr("changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TileMapLayer) DisconnectChanged(subs SignalSubscribers, fn TileMapLayerChangedSignalFn) {
	sig := StringNameFromStr("changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
