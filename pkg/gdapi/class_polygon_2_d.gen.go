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

type ptrsForPolygon2DList struct {
	fnSetPolygon             gdc.MethodBindPtr
	fnGetPolygon             gdc.MethodBindPtr
	fnSetUv                  gdc.MethodBindPtr
	fnGetUv                  gdc.MethodBindPtr
	fnSetColor               gdc.MethodBindPtr
	fnGetColor               gdc.MethodBindPtr
	fnSetPolygons            gdc.MethodBindPtr
	fnGetPolygons            gdc.MethodBindPtr
	fnSetVertexColors        gdc.MethodBindPtr
	fnGetVertexColors        gdc.MethodBindPtr
	fnSetTexture             gdc.MethodBindPtr
	fnGetTexture             gdc.MethodBindPtr
	fnSetTextureOffset       gdc.MethodBindPtr
	fnGetTextureOffset       gdc.MethodBindPtr
	fnSetTextureRotation     gdc.MethodBindPtr
	fnGetTextureRotation     gdc.MethodBindPtr
	fnSetTextureScale        gdc.MethodBindPtr
	fnGetTextureScale        gdc.MethodBindPtr
	fnSetInvertEnabled       gdc.MethodBindPtr
	fnGetInvertEnabled       gdc.MethodBindPtr
	fnSetAntialiased         gdc.MethodBindPtr
	fnGetAntialiased         gdc.MethodBindPtr
	fnSetInvertBorder        gdc.MethodBindPtr
	fnGetInvertBorder        gdc.MethodBindPtr
	fnSetOffset              gdc.MethodBindPtr
	fnGetOffset              gdc.MethodBindPtr
	fnAddBone                gdc.MethodBindPtr
	fnGetBoneCount           gdc.MethodBindPtr
	fnGetBonePath            gdc.MethodBindPtr
	fnGetBoneWeights         gdc.MethodBindPtr
	fnEraseBone              gdc.MethodBindPtr
	fnClearBones             gdc.MethodBindPtr
	fnSetBonePath            gdc.MethodBindPtr
	fnSetBoneWeights         gdc.MethodBindPtr
	fnSetSkeleton            gdc.MethodBindPtr
	fnGetSkeleton            gdc.MethodBindPtr
	fnSetInternalVertexCount gdc.MethodBindPtr
	fnGetInternalVertexCount gdc.MethodBindPtr
}

var ptrsForPolygon2D ptrsForPolygon2DList

func initPolygon2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Polygon2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_polygon")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("get_polygon")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
	{
		methodName := StringNameFromStr("set_uv")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetUv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("get_uv")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetUv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
	{
		methodName := StringNameFromStr("set_color")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_color")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_polygons")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetPolygons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_polygons")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetPolygons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_vertex_colors")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetVertexColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3546319833))
	}
	{
		methodName := StringNameFromStr("get_vertex_colors")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetVertexColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1392750486))
	}
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_texture_offset")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetTextureOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_texture_offset")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetTextureOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_texture_rotation")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetTextureRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_texture_rotation")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetTextureRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_texture_scale")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetTextureScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_texture_scale")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetTextureScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_invert_enabled")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetInvertEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_invert_enabled")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetInvertEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_antialiased")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetAntialiased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_antialiased")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetAntialiased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_invert_border")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetInvertBorder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_invert_border")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetInvertBorder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_offset")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_offset")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("add_bone")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnAddBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 703042815))
	}
	{
		methodName := StringNameFromStr("get_bone_count")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetBoneCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_bone_path")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetBonePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408788394))
	}
	{
		methodName := StringNameFromStr("get_bone_weights")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetBoneWeights = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1542882410))
	}
	{
		methodName := StringNameFromStr("erase_bone")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnEraseBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("clear_bones")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnClearBones = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_bone_path")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetBonePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761262315))
	}
	{
		methodName := StringNameFromStr("set_bone_weights")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetBoneWeights = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1345852415))
	}
	{
		methodName := StringNameFromStr("set_skeleton")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_skeleton")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_internal_vertex_count")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnSetInternalVertexCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_internal_vertex_count")
		defer methodName.Destroy()
		ptrsForPolygon2D.fnGetInternalVertexCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type Polygon2D struct {
	Node2D
}

func (me *Polygon2D) BaseClass() string {
	return "Polygon2D"
}

func NewPolygon2D() *Polygon2D {
	str := StringNameFromStr("Polygon2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Polygon2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Polygon2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Polygon2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Polygon2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Polygon2D) SetPolygon(polygon PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetPolygon() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) SetUv(uv PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{uv.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetUv), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetUv() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetUv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) SetColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) SetPolygons(polygons Array) {
	cargs := []gdc.ConstTypePtr{polygons.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetPolygons), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetPolygons() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetPolygons), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) SetVertexColors(vertex_colors PackedColorArray) {
	cargs := []gdc.ConstTypePtr{vertex_colors.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetVertexColors), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetVertexColors() PackedColorArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedColorArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetVertexColors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) SetTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) SetTextureOffset(texture_offset Vector2) {
	cargs := []gdc.ConstTypePtr{texture_offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetTextureOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetTextureOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetTextureOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) SetTextureRotation(texture_rotation float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_rotation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetTextureRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetTextureRotation() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetTextureRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Polygon2D) SetTextureScale(texture_scale Vector2) {
	cargs := []gdc.ConstTypePtr{texture_scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetTextureScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetTextureScale() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetTextureScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) SetInvertEnabled(invert bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetInvertEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetInvertEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetInvertEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Polygon2D) SetAntialiased(antialiased bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&antialiased)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetAntialiased), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetAntialiased() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetAntialiased), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Polygon2D) SetInvertBorder(invert_border float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert_border)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetInvertBorder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetInvertBorder() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetInvertBorder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Polygon2D) SetOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) AddBone(path NodePath, weights PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), weights.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnAddBone), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetBoneCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetBoneCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Polygon2D) GetBonePath(index int64) NodePath {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetBonePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) GetBoneWeights(index int64) PackedFloat32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedFloat32Array()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetBoneWeights), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) EraseBone(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnEraseBone), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) ClearBones() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnClearBones), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) SetBonePath(index int64, path NodePath) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetBonePath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) SetBoneWeights(index int64, weights PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), weights.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetBoneWeights), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) SetSkeleton(skeleton NodePath) {
	cargs := []gdc.ConstTypePtr{skeleton.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetSkeleton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetSkeleton() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Polygon2D) SetInternalVertexCount(internal_vertex_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&internal_vertex_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnSetInternalVertexCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Polygon2D) GetInternalVertexCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygon2D.fnGetInternalVertexCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
