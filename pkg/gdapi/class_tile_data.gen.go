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

type ptrsForTileDataList struct {
	fnSetFlipH                        gdc.MethodBindPtr
	fnGetFlipH                        gdc.MethodBindPtr
	fnSetFlipV                        gdc.MethodBindPtr
	fnGetFlipV                        gdc.MethodBindPtr
	fnSetTranspose                    gdc.MethodBindPtr
	fnGetTranspose                    gdc.MethodBindPtr
	fnSetMaterial                     gdc.MethodBindPtr
	fnGetMaterial                     gdc.MethodBindPtr
	fnSetTextureOrigin                gdc.MethodBindPtr
	fnGetTextureOrigin                gdc.MethodBindPtr
	fnSetModulate                     gdc.MethodBindPtr
	fnGetModulate                     gdc.MethodBindPtr
	fnSetZIndex                       gdc.MethodBindPtr
	fnGetZIndex                       gdc.MethodBindPtr
	fnSetYSortOrigin                  gdc.MethodBindPtr
	fnGetYSortOrigin                  gdc.MethodBindPtr
	fnSetOccluder                     gdc.MethodBindPtr
	fnGetOccluder                     gdc.MethodBindPtr
	fnSetConstantLinearVelocity       gdc.MethodBindPtr
	fnGetConstantLinearVelocity       gdc.MethodBindPtr
	fnSetConstantAngularVelocity      gdc.MethodBindPtr
	fnGetConstantAngularVelocity      gdc.MethodBindPtr
	fnSetCollisionPolygonsCount       gdc.MethodBindPtr
	fnGetCollisionPolygonsCount       gdc.MethodBindPtr
	fnAddCollisionPolygon             gdc.MethodBindPtr
	fnRemoveCollisionPolygon          gdc.MethodBindPtr
	fnSetCollisionPolygonPoints       gdc.MethodBindPtr
	fnGetCollisionPolygonPoints       gdc.MethodBindPtr
	fnSetCollisionPolygonOneWay       gdc.MethodBindPtr
	fnIsCollisionPolygonOneWay        gdc.MethodBindPtr
	fnSetCollisionPolygonOneWayMargin gdc.MethodBindPtr
	fnGetCollisionPolygonOneWayMargin gdc.MethodBindPtr
	fnSetTerrainSet                   gdc.MethodBindPtr
	fnGetTerrainSet                   gdc.MethodBindPtr
	fnSetTerrain                      gdc.MethodBindPtr
	fnGetTerrain                      gdc.MethodBindPtr
	fnSetTerrainPeeringBit            gdc.MethodBindPtr
	fnGetTerrainPeeringBit            gdc.MethodBindPtr
	fnSetNavigationPolygon            gdc.MethodBindPtr
	fnGetNavigationPolygon            gdc.MethodBindPtr
	fnSetProbability                  gdc.MethodBindPtr
	fnGetProbability                  gdc.MethodBindPtr
	fnSetCustomData                   gdc.MethodBindPtr
	fnGetCustomData                   gdc.MethodBindPtr
	fnSetCustomDataByLayerId          gdc.MethodBindPtr
	fnGetCustomDataByLayerId          gdc.MethodBindPtr
}

var ptrsForTileData ptrsForTileDataList

func initTileDataPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TileData")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_flip_h")
		defer methodName.Destroy()
		ptrsForTileData.fnSetFlipH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_flip_h")
		defer methodName.Destroy()
		ptrsForTileData.fnGetFlipH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_flip_v")
		defer methodName.Destroy()
		ptrsForTileData.fnSetFlipV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_flip_v")
		defer methodName.Destroy()
		ptrsForTileData.fnGetFlipV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_transpose")
		defer methodName.Destroy()
		ptrsForTileData.fnSetTranspose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_transpose")
		defer methodName.Destroy()
		ptrsForTileData.fnGetTranspose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_material")
		defer methodName.Destroy()
		ptrsForTileData.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material")
		defer methodName.Destroy()
		ptrsForTileData.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
	{
		methodName := StringNameFromStr("set_texture_origin")
		defer methodName.Destroy()
		ptrsForTileData.fnSetTextureOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_texture_origin")
		defer methodName.Destroy()
		ptrsForTileData.fnGetTextureOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("set_modulate")
		defer methodName.Destroy()
		ptrsForTileData.fnSetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_modulate")
		defer methodName.Destroy()
		ptrsForTileData.fnGetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_z_index")
		defer methodName.Destroy()
		ptrsForTileData.fnSetZIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_z_index")
		defer methodName.Destroy()
		ptrsForTileData.fnGetZIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_y_sort_origin")
		defer methodName.Destroy()
		ptrsForTileData.fnSetYSortOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_y_sort_origin")
		defer methodName.Destroy()
		ptrsForTileData.fnGetYSortOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_occluder")
		defer methodName.Destroy()
		ptrsForTileData.fnSetOccluder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 914399637))
	}
	{
		methodName := StringNameFromStr("get_occluder")
		defer methodName.Destroy()
		ptrsForTileData.fnGetOccluder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2458574231))
	}
	{
		methodName := StringNameFromStr("set_constant_linear_velocity")
		defer methodName.Destroy()
		ptrsForTileData.fnSetConstantLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
	}
	{
		methodName := StringNameFromStr("get_constant_linear_velocity")
		defer methodName.Destroy()
		ptrsForTileData.fnGetConstantLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("set_constant_angular_velocity")
		defer methodName.Destroy()
		ptrsForTileData.fnSetConstantAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_constant_angular_velocity")
		defer methodName.Destroy()
		ptrsForTileData.fnGetConstantAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_collision_polygons_count")
		defer methodName.Destroy()
		ptrsForTileData.fnSetCollisionPolygonsCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_collision_polygons_count")
		defer methodName.Destroy()
		ptrsForTileData.fnGetCollisionPolygonsCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("add_collision_polygon")
		defer methodName.Destroy()
		ptrsForTileData.fnAddCollisionPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("remove_collision_polygon")
		defer methodName.Destroy()
		ptrsForTileData.fnRemoveCollisionPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_collision_polygon_points")
		defer methodName.Destroy()
		ptrsForTileData.fnSetCollisionPolygonPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3230546541))
	}
	{
		methodName := StringNameFromStr("get_collision_polygon_points")
		defer methodName.Destroy()
		ptrsForTileData.fnGetCollisionPolygonPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 103942801))
	}
	{
		methodName := StringNameFromStr("set_collision_polygon_one_way")
		defer methodName.Destroy()
		ptrsForTileData.fnSetCollisionPolygonOneWay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1383440665))
	}
	{
		methodName := StringNameFromStr("is_collision_polygon_one_way")
		defer methodName.Destroy()
		ptrsForTileData.fnIsCollisionPolygonOneWay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
	}
	{
		methodName := StringNameFromStr("set_collision_polygon_one_way_margin")
		defer methodName.Destroy()
		ptrsForTileData.fnSetCollisionPolygonOneWayMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
	}
	{
		methodName := StringNameFromStr("get_collision_polygon_one_way_margin")
		defer methodName.Destroy()
		ptrsForTileData.fnGetCollisionPolygonOneWayMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
	}
	{
		methodName := StringNameFromStr("set_terrain_set")
		defer methodName.Destroy()
		ptrsForTileData.fnSetTerrainSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_terrain_set")
		defer methodName.Destroy()
		ptrsForTileData.fnGetTerrainSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_terrain")
		defer methodName.Destroy()
		ptrsForTileData.fnSetTerrain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_terrain")
		defer methodName.Destroy()
		ptrsForTileData.fnGetTerrain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_terrain_peering_bit")
		defer methodName.Destroy()
		ptrsForTileData.fnSetTerrainPeeringBit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1084452308))
	}
	{
		methodName := StringNameFromStr("get_terrain_peering_bit")
		defer methodName.Destroy()
		ptrsForTileData.fnGetTerrainPeeringBit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3831796792))
	}
	{
		methodName := StringNameFromStr("set_navigation_polygon")
		defer methodName.Destroy()
		ptrsForTileData.fnSetNavigationPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2224691167))
	}
	{
		methodName := StringNameFromStr("get_navigation_polygon")
		defer methodName.Destroy()
		ptrsForTileData.fnGetNavigationPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3991786031))
	}
	{
		methodName := StringNameFromStr("set_probability")
		defer methodName.Destroy()
		ptrsForTileData.fnSetProbability = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_probability")
		defer methodName.Destroy()
		ptrsForTileData.fnGetProbability = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_custom_data")
		defer methodName.Destroy()
		ptrsForTileData.fnSetCustomData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 402577236))
	}
	{
		methodName := StringNameFromStr("get_custom_data")
		defer methodName.Destroy()
		ptrsForTileData.fnGetCustomData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1868160156))
	}
	{
		methodName := StringNameFromStr("set_custom_data_by_layer_id")
		defer methodName.Destroy()
		ptrsForTileData.fnSetCustomDataByLayerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2152698145))
	}
	{
		methodName := StringNameFromStr("get_custom_data_by_layer_id")
		defer methodName.Destroy()
		ptrsForTileData.fnGetCustomDataByLayerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
	}

}

type TileData struct {
	Object
}

func (me *TileData) BaseClass() string {
	return "TileData"
}

func NewTileData() *TileData {
	str := StringNameFromStr("TileData") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TileData{}
	obj.SetBaseObject(objPtr)
	return obj
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

func (me *TileData) SetFlipH(flip_h bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_h)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetFlipH), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetFlipH() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetFlipH), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetFlipV(flip_v bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_v)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetFlipV), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetFlipV() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetFlipV), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetTranspose(transpose bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&transpose)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetTranspose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetTranspose() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetTranspose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetMaterial(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetMaterial() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileData) SetTextureOrigin(texture_origin Vector2i) {
	cargs := []gdc.ConstTypePtr{texture_origin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetTextureOrigin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetTextureOrigin() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetTextureOrigin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileData) SetModulate(modulate Color) {
	cargs := []gdc.ConstTypePtr{modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetModulate() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileData) SetZIndex(z_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&z_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetZIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetZIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetZIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetYSortOrigin(y_sort_origin int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&y_sort_origin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetYSortOrigin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetYSortOrigin() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetYSortOrigin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetOccluder(layer_id int64, occluder_polygon OccluderPolygon2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), occluder_polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetOccluder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetOccluder(layer_id int64) OccluderPolygon2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewOccluderPolygon2D()
	pinner.Pin(&layer_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetOccluder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileData) SetConstantLinearVelocity(layer_id int64, velocity Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetConstantLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetConstantLinearVelocity(layer_id int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&layer_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetConstantLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileData) SetConstantAngularVelocity(layer_id int64, velocity float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&velocity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetConstantAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetConstantAngularVelocity(layer_id int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&layer_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetConstantAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetCollisionPolygonsCount(layer_id int64, polygons_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygons_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetCollisionPolygonsCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetCollisionPolygonsCount(layer_id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&layer_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetCollisionPolygonsCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) AddCollisionPolygon(layer_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnAddCollisionPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) RemoveCollisionPolygon(layer_id int64, polygon_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnRemoveCollisionPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) SetCollisionPolygonPoints(layer_id int64, polygon_index int64, polygon PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index), polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetCollisionPolygonPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetCollisionPolygonPoints(layer_id int64, polygon_index int64) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()
	pinner.Pin(&layer_id)
	pinner.Pin(&polygon_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetCollisionPolygonPoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileData) SetCollisionPolygonOneWay(layer_id int64, polygon_index int64, one_way bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index), gdc.ConstTypePtr(&one_way)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetCollisionPolygonOneWay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) IsCollisionPolygonOneWay(layer_id int64, polygon_index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_id)
	pinner.Pin(&polygon_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnIsCollisionPolygonOneWay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetCollisionPolygonOneWayMargin(layer_id int64, polygon_index int64, one_way_margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index), gdc.ConstTypePtr(&one_way_margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetCollisionPolygonOneWayMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetCollisionPolygonOneWayMargin(layer_id int64, polygon_index int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), gdc.ConstTypePtr(&polygon_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&layer_id)
	pinner.Pin(&polygon_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetCollisionPolygonOneWayMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetTerrainSet(terrain_set int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain_set)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetTerrainSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetTerrainSet() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetTerrainSet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetTerrain(terrain int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&terrain)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetTerrain), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetTerrain() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetTerrain), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetTerrainPeeringBit(peering_bit TileSetCellNeighbor, terrain int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peering_bit), gdc.ConstTypePtr(&terrain)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetTerrainPeeringBit), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetTerrainPeeringBit(peering_bit TileSetCellNeighbor) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peering_bit)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&peering_bit)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetTerrainPeeringBit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetNavigationPolygon(layer_id int64, navigation_polygon NavigationPolygon) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), navigation_polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetNavigationPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetNavigationPolygon(layer_id int64) NavigationPolygon {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNavigationPolygon()
	pinner.Pin(&layer_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetNavigationPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileData) SetProbability(probability float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&probability)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetProbability), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetProbability() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetProbability), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TileData) SetCustomData(layer_name String, value Variant) {
	cargs := []gdc.ConstTypePtr{layer_name.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetCustomData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetCustomData(layer_name String) Variant {
	cargs := []gdc.ConstTypePtr{layer_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetCustomData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TileData) SetCustomDataByLayerId(layer_id int64, value Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnSetCustomDataByLayerId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TileData) GetCustomDataByLayerId(layer_id int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&layer_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileData.fnGetCustomDataByLayerId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TileDataChangedSignalFn func()

func (me *TileData) ConnectChanged(subs SignalSubscribers, fn TileDataChangedSignalFn) {
	sig := StringNameFromStr("changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TileData) DisconnectChanged(subs SignalSubscribers, fn TileDataChangedSignalFn) {
	sig := StringNameFromStr("changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
