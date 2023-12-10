// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Polygon2D struct {
  obj gdc.ObjectPtr
}

func (me *Polygon2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Polygon2D) BaseClass() string {
  return "Polygon2D"
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

func  (me *Polygon2D) SetPolygon(polygon PackedVector2Array, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetPolygon() PackedVector2Array {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetUv(uv PackedVector2Array, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(uv.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetUv() PackedVector2Array {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetColor(color Color, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetColor() Color {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetPolygons(polygons Array, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(polygons.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetPolygons() Array {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetVertexColors(vertex_colors PackedColorArray, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3546319833) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(vertex_colors.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetVertexColors() PackedColorArray {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1392750486) // FIXME: should cache?
  var ret PackedColorArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetTexture() Texture2D {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetTextureOffset(texture_offset Vector2, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture_offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetTextureOffset() Vector2 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetTextureRotation(texture_rotation float32, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_rotation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetTextureRotation() float32 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetTextureScale(texture_scale Vector2, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture_scale.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetTextureScale() Vector2 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetInvertEnabled(invert bool, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_invert_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetInvertEnabled() bool {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_invert_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetAntialiased(antialiased bool, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_antialiased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&antialiased), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetAntialiased() bool {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_antialiased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetInvertBorder(invert_border float32, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_invert_border")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert_border), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetInvertBorder() float32 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_invert_border")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetOffset() Vector2 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) AddBone(path NodePath, weights PackedFloat32Array, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 703042815) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(weights.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetBoneCount() int {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) GetBonePath(index int, ) NodePath {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) GetBoneWeights(index int, ) PackedFloat32Array {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_weights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1542882410) // FIXME: should cache?
  var ret PackedFloat32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) EraseBone(index int, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) ClearBones()  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) SetBonePath(index int, path NodePath, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761262315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) SetBoneWeights(index int, weights PackedFloat32Array, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_weights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1345852415) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(weights.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) SetSkeleton(skeleton NodePath, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(skeleton.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetSkeleton() NodePath {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Polygon2D) SetInternalVertexCount(internal_vertex_count int, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_internal_vertex_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&internal_vertex_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Polygon2D) GetInternalVertexCount() int {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_internal_vertex_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Polygon2D) GetPropColor() Color {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropColor(value Color) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropOffset() Vector2 {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropAntialiased() bool {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropAntialiased(value bool) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropTexture() Texture2D {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropTexture(value Texture2D) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropTextureOffset() Vector2 {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropTextureOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropTextureScale() Vector2 {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropTextureScale(value Vector2) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropTextureRotation() float32 {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropTextureRotation(value float32) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropSkeleton() NodePath {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropSkeleton(value NodePath) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropInvertEnabled() bool {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropInvertEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropInvertBorder() float32 {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropInvertBorder(value float32) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropPolygon() PackedVector2Array {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropPolygon(value PackedVector2Array) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropUv() PackedVector2Array {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropUv(value PackedVector2Array) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropVertexColors() PackedColorArray {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropVertexColors(value PackedColorArray) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropPolygons() Array {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropPolygons(value Array) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropBones() Array {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropBones(value Array) {
  panic("TODO: implement")
}

func (me *Polygon2D) GetPropInternalVertexCount() int {
  panic("TODO: implement")
}

func (me *Polygon2D) SetPropInternalVertexCount(value int) {
  panic("TODO: implement")
}