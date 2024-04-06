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

type ptrsForTileSetList struct {
	fnGetNextSourceId                 gdc.MethodBindPtr
	fnAddSource                       gdc.MethodBindPtr
	fnRemoveSource                    gdc.MethodBindPtr
	fnSetSourceId                     gdc.MethodBindPtr
	fnGetSourceCount                  gdc.MethodBindPtr
	fnGetSourceId                     gdc.MethodBindPtr
	fnHasSource                       gdc.MethodBindPtr
	fnGetSource                       gdc.MethodBindPtr
	fnSetTileShape                    gdc.MethodBindPtr
	fnGetTileShape                    gdc.MethodBindPtr
	fnSetTileLayout                   gdc.MethodBindPtr
	fnGetTileLayout                   gdc.MethodBindPtr
	fnSetTileOffsetAxis               gdc.MethodBindPtr
	fnGetTileOffsetAxis               gdc.MethodBindPtr
	fnSetTileSize                     gdc.MethodBindPtr
	fnGetTileSize                     gdc.MethodBindPtr
	fnSetUvClipping                   gdc.MethodBindPtr
	fnIsUvClipping                    gdc.MethodBindPtr
	fnGetOcclusionLayersCount         gdc.MethodBindPtr
	fnAddOcclusionLayer               gdc.MethodBindPtr
	fnMoveOcclusionLayer              gdc.MethodBindPtr
	fnRemoveOcclusionLayer            gdc.MethodBindPtr
	fnSetOcclusionLayerLightMask      gdc.MethodBindPtr
	fnGetOcclusionLayerLightMask      gdc.MethodBindPtr
	fnSetOcclusionLayerSdfCollision   gdc.MethodBindPtr
	fnGetOcclusionLayerSdfCollision   gdc.MethodBindPtr
	fnGetPhysicsLayersCount           gdc.MethodBindPtr
	fnAddPhysicsLayer                 gdc.MethodBindPtr
	fnMovePhysicsLayer                gdc.MethodBindPtr
	fnRemovePhysicsLayer              gdc.MethodBindPtr
	fnSetPhysicsLayerCollisionLayer   gdc.MethodBindPtr
	fnGetPhysicsLayerCollisionLayer   gdc.MethodBindPtr
	fnSetPhysicsLayerCollisionMask    gdc.MethodBindPtr
	fnGetPhysicsLayerCollisionMask    gdc.MethodBindPtr
	fnSetPhysicsLayerPhysicsMaterial  gdc.MethodBindPtr
	fnGetPhysicsLayerPhysicsMaterial  gdc.MethodBindPtr
	fnGetTerrainSetsCount             gdc.MethodBindPtr
	fnAddTerrainSet                   gdc.MethodBindPtr
	fnMoveTerrainSet                  gdc.MethodBindPtr
	fnRemoveTerrainSet                gdc.MethodBindPtr
	fnSetTerrainSetMode               gdc.MethodBindPtr
	fnGetTerrainSetMode               gdc.MethodBindPtr
	fnGetTerrainsCount                gdc.MethodBindPtr
	fnAddTerrain                      gdc.MethodBindPtr
	fnMoveTerrain                     gdc.MethodBindPtr
	fnRemoveTerrain                   gdc.MethodBindPtr
	fnSetTerrainName                  gdc.MethodBindPtr
	fnGetTerrainName                  gdc.MethodBindPtr
	fnSetTerrainColor                 gdc.MethodBindPtr
	fnGetTerrainColor                 gdc.MethodBindPtr
	fnGetNavigationLayersCount        gdc.MethodBindPtr
	fnAddNavigationLayer              gdc.MethodBindPtr
	fnMoveNavigationLayer             gdc.MethodBindPtr
	fnRemoveNavigationLayer           gdc.MethodBindPtr
	fnSetNavigationLayerLayers        gdc.MethodBindPtr
	fnGetNavigationLayerLayers        gdc.MethodBindPtr
	fnSetNavigationLayerLayerValue    gdc.MethodBindPtr
	fnGetNavigationLayerLayerValue    gdc.MethodBindPtr
	fnGetCustomDataLayersCount        gdc.MethodBindPtr
	fnAddCustomDataLayer              gdc.MethodBindPtr
	fnMoveCustomDataLayer             gdc.MethodBindPtr
	fnRemoveCustomDataLayer           gdc.MethodBindPtr
	fnGetCustomDataLayerByName        gdc.MethodBindPtr
	fnSetCustomDataLayerName          gdc.MethodBindPtr
	fnGetCustomDataLayerName          gdc.MethodBindPtr
	fnSetCustomDataLayerType          gdc.MethodBindPtr
	fnGetCustomDataLayerType          gdc.MethodBindPtr
	fnSetSourceLevelTileProxy         gdc.MethodBindPtr
	fnGetSourceLevelTileProxy         gdc.MethodBindPtr
	fnHasSourceLevelTileProxy         gdc.MethodBindPtr
	fnRemoveSourceLevelTileProxy      gdc.MethodBindPtr
	fnSetCoordsLevelTileProxy         gdc.MethodBindPtr
	fnGetCoordsLevelTileProxy         gdc.MethodBindPtr
	fnHasCoordsLevelTileProxy         gdc.MethodBindPtr
	fnRemoveCoordsLevelTileProxy      gdc.MethodBindPtr
	fnSetAlternativeLevelTileProxy    gdc.MethodBindPtr
	fnGetAlternativeLevelTileProxy    gdc.MethodBindPtr
	fnHasAlternativeLevelTileProxy    gdc.MethodBindPtr
	fnRemoveAlternativeLevelTileProxy gdc.MethodBindPtr
	fnMapTileProxy                    gdc.MethodBindPtr
	fnCleanupInvalidTileProxies       gdc.MethodBindPtr
	fnClearTileProxies                gdc.MethodBindPtr
	fnAddPattern                      gdc.MethodBindPtr
	fnGetPattern                      gdc.MethodBindPtr
	fnRemovePattern                   gdc.MethodBindPtr
	fnGetPatternsCount                gdc.MethodBindPtr
}

var ptrsForTileSet ptrsForTileSetList

func initTileSetPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TileSet")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_next_source_id")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetNextSourceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_source")
		defer methodName.Destroy()
		ptrsForTileSet.fnAddSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1059186179))
	}
	{
		methodName := StringNameFromStr("remove_source")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemoveSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_source_id")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetSourceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_source_count")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetSourceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_source_id")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetSourceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("has_source")
		defer methodName.Destroy()
		ptrsForTileSet.fnHasSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_source")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1763540252))
	}
	{
		methodName := StringNameFromStr("set_tile_shape")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetTileShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2131427112))
	}
	{
		methodName := StringNameFromStr("get_tile_shape")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetTileShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 716918169))
	}
	{
		methodName := StringNameFromStr("set_tile_layout")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetTileLayout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1071216679))
	}
	{
		methodName := StringNameFromStr("get_tile_layout")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetTileLayout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 194628839))
	}
	{
		methodName := StringNameFromStr("set_tile_offset_axis")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetTileOffsetAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3300198521))
	}
	{
		methodName := StringNameFromStr("get_tile_offset_axis")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetTileOffsetAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 762494114))
	}
	{
		methodName := StringNameFromStr("set_tile_size")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetTileSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_tile_size")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetTileSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("set_uv_clipping")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetUvClipping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_uv_clipping")
		defer methodName.Destroy()
		ptrsForTileSet.fnIsUvClipping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_occlusion_layers_count")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetOcclusionLayersCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_occlusion_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnAddOcclusionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("move_occlusion_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnMoveOcclusionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("remove_occlusion_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemoveOcclusionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_occlusion_layer_light_mask")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetOcclusionLayerLightMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_occlusion_layer_light_mask")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetOcclusionLayerLightMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_occlusion_layer_sdf_collision")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetOcclusionLayerSdfCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_occlusion_layer_sdf_collision")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetOcclusionLayerSdfCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_physics_layers_count")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetPhysicsLayersCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_physics_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnAddPhysicsLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("move_physics_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnMovePhysicsLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("remove_physics_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemovePhysicsLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_physics_layer_collision_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetPhysicsLayerCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_physics_layer_collision_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetPhysicsLayerCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_physics_layer_collision_mask")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetPhysicsLayerCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_physics_layer_collision_mask")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetPhysicsLayerCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_physics_layer_physics_material")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetPhysicsLayerPhysicsMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1018687357))
	}
	{
		methodName := StringNameFromStr("get_physics_layer_physics_material")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetPhysicsLayerPhysicsMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 788318639))
	}
	{
		methodName := StringNameFromStr("get_terrain_sets_count")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetTerrainSetsCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_terrain_set")
		defer methodName.Destroy()
		ptrsForTileSet.fnAddTerrainSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("move_terrain_set")
		defer methodName.Destroy()
		ptrsForTileSet.fnMoveTerrainSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("remove_terrain_set")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemoveTerrainSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_terrain_set_mode")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetTerrainSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3943003916))
	}
	{
		methodName := StringNameFromStr("get_terrain_set_mode")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetTerrainSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2084469411))
	}
	{
		methodName := StringNameFromStr("get_terrains_count")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetTerrainsCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("add_terrain")
		defer methodName.Destroy()
		ptrsForTileSet.fnAddTerrain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1230568737))
	}
	{
		methodName := StringNameFromStr("move_terrain")
		defer methodName.Destroy()
		ptrsForTileSet.fnMoveTerrain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1649997291))
	}
	{
		methodName := StringNameFromStr("remove_terrain")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemoveTerrain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_terrain_name")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetTerrainName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2285447957))
	}
	{
		methodName := StringNameFromStr("get_terrain_name")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetTerrainName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1391810591))
	}
	{
		methodName := StringNameFromStr("set_terrain_color")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetTerrainColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3733378741))
	}
	{
		methodName := StringNameFromStr("get_terrain_color")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetTerrainColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2165839948))
	}
	{
		methodName := StringNameFromStr("get_navigation_layers_count")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetNavigationLayersCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_navigation_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnAddNavigationLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("move_navigation_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnMoveNavigationLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("remove_navigation_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemoveNavigationLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_navigation_layer_layers")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetNavigationLayerLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_navigation_layer_layers")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetNavigationLayerLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_navigation_layer_layer_value")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetNavigationLayerLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1383440665))
	}
	{
		methodName := StringNameFromStr("get_navigation_layer_layer_value")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetNavigationLayerLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
	}
	{
		methodName := StringNameFromStr("get_custom_data_layers_count")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetCustomDataLayersCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_custom_data_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnAddCustomDataLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("move_custom_data_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnMoveCustomDataLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("remove_custom_data_layer")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemoveCustomDataLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_custom_data_layer_by_name")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetCustomDataLayerByName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}
	{
		methodName := StringNameFromStr("set_custom_data_layer_name")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetCustomDataLayerName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_custom_data_layer_name")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetCustomDataLayerName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_custom_data_layer_type")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetCustomDataLayerType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3492912874))
	}
	{
		methodName := StringNameFromStr("get_custom_data_layer_type")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetCustomDataLayerType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2990820875))
	}
	{
		methodName := StringNameFromStr("set_source_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetSourceLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_source_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetSourceLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
	}
	{
		methodName := StringNameFromStr("has_source_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnHasSourceLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3067735520))
	}
	{
		methodName := StringNameFromStr("remove_source_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemoveSourceLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_coords_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetCoordsLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1769939278))
	}
	{
		methodName := StringNameFromStr("get_coords_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetCoordsLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2856536371))
	}
	{
		methodName := StringNameFromStr("has_coords_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnHasCoordsLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3957903770))
	}
	{
		methodName := StringNameFromStr("remove_coords_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemoveCoordsLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2311374912))
	}
	{
		methodName := StringNameFromStr("set_alternative_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnSetAlternativeLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3862385460))
	}
	{
		methodName := StringNameFromStr("get_alternative_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetAlternativeLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2303761075))
	}
	{
		methodName := StringNameFromStr("has_alternative_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnHasAlternativeLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 180086755))
	}
	{
		methodName := StringNameFromStr("remove_alternative_level_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemoveAlternativeLevelTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2328951467))
	}
	{
		methodName := StringNameFromStr("map_tile_proxy")
		defer methodName.Destroy()
		ptrsForTileSet.fnMapTileProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4267935328))
	}
	{
		methodName := StringNameFromStr("cleanup_invalid_tile_proxies")
		defer methodName.Destroy()
		ptrsForTileSet.fnCleanupInvalidTileProxies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("clear_tile_proxies")
		defer methodName.Destroy()
		ptrsForTileSet.fnClearTileProxies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("add_pattern")
		defer methodName.Destroy()
		ptrsForTileSet.fnAddPattern = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 763712015))
	}
	{
		methodName := StringNameFromStr("get_pattern")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetPattern = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4207737510))
	}
	{
		methodName := StringNameFromStr("remove_pattern")
		defer methodName.Destroy()
		ptrsForTileSet.fnRemovePattern = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_patterns_count")
		defer methodName.Destroy()
		ptrsForTileSet.fnGetPatternsCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}

}

type TileSet struct {
	Resource
}

func (me *TileSet) BaseClass() string {
	return "TileSet"
}

func NewTileSet() *TileSet {
	str := StringNameFromStr("TileSet") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TileSet{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TileSetTileShape int

const (
	TileSetTileShapeTileShapeSquare           TileSetTileShape = 0
	TileSetTileShapeTileShapeIsometric        TileSetTileShape = 1
	TileSetTileShapeTileShapeHalfOffsetSquare TileSetTileShape = 2
	TileSetTileShapeTileShapeHexagon          TileSetTileShape = 3
)

type TileSetTileLayout int

const (
	TileSetTileLayoutTileLayoutStacked       TileSetTileLayout = 0
	TileSetTileLayoutTileLayoutStackedOffset TileSetTileLayout = 1
	TileSetTileLayoutTileLayoutStairsRight   TileSetTileLayout = 2
	TileSetTileLayoutTileLayoutStairsDown    TileSetTileLayout = 3
	TileSetTileLayoutTileLayoutDiamondRight  TileSetTileLayout = 4
	TileSetTileLayoutTileLayoutDiamondDown   TileSetTileLayout = 5
)

type TileSetTileOffsetAxis int

const (
	TileSetTileOffsetAxisTileOffsetAxisHorizontal TileSetTileOffsetAxis = 0
	TileSetTileOffsetAxisTileOffsetAxisVertical   TileSetTileOffsetAxis = 1
)

type TileSetCellNeighbor int

const (
	TileSetCellNeighborCellNeighborRightSide         TileSetCellNeighbor = 0
	TileSetCellNeighborCellNeighborRightCorner       TileSetCellNeighbor = 1
	TileSetCellNeighborCellNeighborBottomRightSide   TileSetCellNeighbor = 2
	TileSetCellNeighborCellNeighborBottomRightCorner TileSetCellNeighbor = 3
	TileSetCellNeighborCellNeighborBottomSide        TileSetCellNeighbor = 4
	TileSetCellNeighborCellNeighborBottomCorner      TileSetCellNeighbor = 5
	TileSetCellNeighborCellNeighborBottomLeftSide    TileSetCellNeighbor = 6
	TileSetCellNeighborCellNeighborBottomLeftCorner  TileSetCellNeighbor = 7
	TileSetCellNeighborCellNeighborLeftSide          TileSetCellNeighbor = 8
	TileSetCellNeighborCellNeighborLeftCorner        TileSetCellNeighbor = 9
	TileSetCellNeighborCellNeighborTopLeftSide       TileSetCellNeighbor = 10
	TileSetCellNeighborCellNeighborTopLeftCorner     TileSetCellNeighbor = 11
	TileSetCellNeighborCellNeighborTopSide           TileSetCellNeighbor = 12
	TileSetCellNeighborCellNeighborTopCorner         TileSetCellNeighbor = 13
	TileSetCellNeighborCellNeighborTopRightSide      TileSetCellNeighbor = 14
	TileSetCellNeighborCellNeighborTopRightCorner    TileSetCellNeighbor = 15
)

type TileSetTerrainMode int

const (
	TileSetTerrainModeTerrainModeMatchCornersAndSides TileSetTerrainMode = 0
	TileSetTerrainModeTerrainModeMatchCorners         TileSetTerrainMode = 1
	TileSetTerrainModeTerrainModeMatchSides           TileSetTerrainMode = 2
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

func (me *TileSet) GetNextSourceId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetNextSourceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) AddSource(source TileSetSource, atlas_source_id_override int64) int64 {
	cargs := []gdc.ConstTypePtr{source.AsCTypePtr(), gdc.ConstTypePtr(&atlas_source_id_override)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&atlas_source_id_override)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnAddSource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) RemoveSource(source_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemoveSource), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) SetSourceId(source_id int64, new_source_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_id), gdc.ConstTypePtr(&new_source_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetSourceId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetSourceCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetSourceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) GetSourceId(index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetSourceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) HasSource(source_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&source_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnHasSource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) GetSource(source_id int64) TileSetSource {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTileSetSource()
	pinner.Pin(&source_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetSource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSet) SetTileShape(shape TileSetTileShape) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetTileShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetTileShape() TileSetTileShape {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TileSetTileShape

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetTileShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TileSet) SetTileLayout(layout TileSetTileLayout) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layout)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetTileLayout), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetTileLayout() TileSetTileLayout {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TileSetTileLayout

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetTileLayout), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TileSet) SetTileOffsetAxis(alignment TileSetTileOffsetAxis) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetTileOffsetAxis), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetTileOffsetAxis() TileSetTileOffsetAxis {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TileSetTileOffsetAxis

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetTileOffsetAxis), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TileSet) SetTileSize(size Vector2i) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetTileSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetTileSize() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetTileSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSet) SetUvClipping(uv_clipping bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&uv_clipping)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetUvClipping), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) IsUvClipping() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnIsUvClipping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) GetOcclusionLayersCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetOcclusionLayersCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) AddOcclusionLayer(to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnAddOcclusionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) MoveOcclusionLayer(layer_index int64, to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnMoveOcclusionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) RemoveOcclusionLayer(layer_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemoveOcclusionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) SetOcclusionLayerLightMask(layer_index int64, light_mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&light_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetOcclusionLayerLightMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetOcclusionLayerLightMask(layer_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&layer_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetOcclusionLayerLightMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) SetOcclusionLayerSdfCollision(layer_index int64, sdf_collision bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&sdf_collision)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetOcclusionLayerSdfCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetOcclusionLayerSdfCollision(layer_index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetOcclusionLayerSdfCollision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) GetPhysicsLayersCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetPhysicsLayersCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) AddPhysicsLayer(to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnAddPhysicsLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) MovePhysicsLayer(layer_index int64, to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnMovePhysicsLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) RemovePhysicsLayer(layer_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemovePhysicsLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) SetPhysicsLayerCollisionLayer(layer_index int64, layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetPhysicsLayerCollisionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetPhysicsLayerCollisionLayer(layer_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&layer_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetPhysicsLayerCollisionLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) SetPhysicsLayerCollisionMask(layer_index int64, mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetPhysicsLayerCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetPhysicsLayerCollisionMask(layer_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&layer_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetPhysicsLayerCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) SetPhysicsLayerPhysicsMaterial(layer_index int64, physics_material PhysicsMaterial) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), physics_material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetPhysicsLayerPhysicsMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetPhysicsLayerPhysicsMaterial(layer_index int64) PhysicsMaterial {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsMaterial()
	pinner.Pin(&layer_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetPhysicsLayerPhysicsMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSet) GetTerrainSetsCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetTerrainSetsCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) AddTerrainSet(to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnAddTerrainSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) MoveTerrainSet(terrain_set int64, to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnMoveTerrainSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) RemoveTerrainSet(terrain_set int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemoveTerrainSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) SetTerrainSetMode(terrain_set int64, mode TileSetTerrainMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetTerrainSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetTerrainSetMode(terrain_set int64) TileSetTerrainMode {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TileSetTerrainMode
	pinner.Pin(&terrain_set)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetTerrainSetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TileSet) GetTerrainsCount(terrain_set int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&terrain_set)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetTerrainsCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) AddTerrain(terrain_set int64, to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnAddTerrain), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) MoveTerrain(terrain_set int64, terrain_index int64, to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index), gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnMoveTerrain), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) RemoveTerrain(terrain_set int64, terrain_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemoveTerrain), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) SetTerrainName(terrain_set int64, terrain_index int64, name String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetTerrainName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetTerrainName(terrain_set int64, terrain_index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&terrain_set)
	pinner.Pin(&terrain_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetTerrainName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSet) SetTerrainColor(terrain_set int64, terrain_index int64, color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetTerrainColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetTerrainColor(terrain_set int64, terrain_index int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&terrain_set)
	pinner.Pin(&terrain_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetTerrainColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSet) GetNavigationLayersCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetNavigationLayersCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) AddNavigationLayer(to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnAddNavigationLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) MoveNavigationLayer(layer_index int64, to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnMoveNavigationLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) RemoveNavigationLayer(layer_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemoveNavigationLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) SetNavigationLayerLayers(layer_index int64, layers int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetNavigationLayerLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetNavigationLayerLayers(layer_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&layer_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetNavigationLayerLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) SetNavigationLayerLayerValue(layer_index int64, layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetNavigationLayerLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetNavigationLayerLayerValue(layer_index int64, layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_index)
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetNavigationLayerLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) GetCustomDataLayersCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetCustomDataLayersCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) AddCustomDataLayer(to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnAddCustomDataLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) MoveCustomDataLayer(layer_index int64, to_position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&to_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnMoveCustomDataLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) RemoveCustomDataLayer(layer_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemoveCustomDataLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetCustomDataLayerByName(layer_name String) int64 {
	cargs := []gdc.ConstTypePtr{layer_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetCustomDataLayerByName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) SetCustomDataLayerName(layer_index int64, layer_name String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), layer_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetCustomDataLayerName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetCustomDataLayerName(layer_index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&layer_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetCustomDataLayerName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSet) SetCustomDataLayerType(layer_index int64, layer_type VariantType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index), gdc.ConstTypePtr(&layer_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetCustomDataLayerType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetCustomDataLayerType(layer_index int64) VariantType {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VariantType
	pinner.Pin(&layer_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetCustomDataLayerType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TileSet) SetSourceLevelTileProxy(source_from int64, source_to int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), gdc.ConstTypePtr(&source_to)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetSourceLevelTileProxy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetSourceLevelTileProxy(source_from int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&source_from)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetSourceLevelTileProxy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) HasSourceLevelTileProxy(source_from int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&source_from)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnHasSourceLevelTileProxy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) RemoveSourceLevelTileProxy(source_from int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemoveSourceLevelTileProxy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) SetCoordsLevelTileProxy(p_source_from int64, coords_from Vector2i, source_to int64, coords_to Vector2i) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_source_from), coords_from.AsCTypePtr(), gdc.ConstTypePtr(&source_to), coords_to.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetCoordsLevelTileProxy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetCoordsLevelTileProxy(source_from int64, coords_from Vector2i) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), coords_from.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&source_from)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetCoordsLevelTileProxy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSet) HasCoordsLevelTileProxy(source_from int64, coords_from Vector2i) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), coords_from.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&source_from)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnHasCoordsLevelTileProxy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) RemoveCoordsLevelTileProxy(source_from int64, coords_from Vector2i) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), coords_from.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemoveCoordsLevelTileProxy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) SetAlternativeLevelTileProxy(source_from int64, coords_from Vector2i, alternative_from int64, source_to int64, coords_to Vector2i, alternative_to int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), coords_from.AsCTypePtr(), gdc.ConstTypePtr(&alternative_from), gdc.ConstTypePtr(&source_to), coords_to.AsCTypePtr(), gdc.ConstTypePtr(&alternative_to)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnSetAlternativeLevelTileProxy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetAlternativeLevelTileProxy(source_from int64, coords_from Vector2i, alternative_from int64) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), coords_from.AsCTypePtr(), gdc.ConstTypePtr(&alternative_from)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&source_from)
	pinner.Pin(&alternative_from)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetAlternativeLevelTileProxy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSet) HasAlternativeLevelTileProxy(source_from int64, coords_from Vector2i, alternative_from int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), coords_from.AsCTypePtr(), gdc.ConstTypePtr(&alternative_from)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&source_from)
	pinner.Pin(&alternative_from)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnHasAlternativeLevelTileProxy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) RemoveAlternativeLevelTileProxy(source_from int64, coords_from Vector2i, alternative_from int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), coords_from.AsCTypePtr(), gdc.ConstTypePtr(&alternative_from)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemoveAlternativeLevelTileProxy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) MapTileProxy(source_from int64, coords_from Vector2i, alternative_from int64) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source_from), coords_from.AsCTypePtr(), gdc.ConstTypePtr(&alternative_from)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&source_from)
	pinner.Pin(&alternative_from)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnMapTileProxy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSet) CleanupInvalidTileProxies() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnCleanupInvalidTileProxies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) ClearTileProxies() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnClearTileProxies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) AddPattern(pattern TileMapPattern, index int64) int64 {
	cargs := []gdc.ConstTypePtr{pattern.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnAddPattern), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileSet) GetPattern(index int64) TileMapPattern {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTileMapPattern()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetPattern), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileSet) RemovePattern(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnRemovePattern), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileSet) GetPatternsCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSet.fnGetPatternsCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
