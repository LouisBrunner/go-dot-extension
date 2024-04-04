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

func  (me *Polygon2D) SetPolygon(polygon PackedVector2Array, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetPolygon() PackedVector2Array {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) SetUv(uv PackedVector2Array, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{uv.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetUv() PackedVector2Array {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) SetColor(color Color, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetColor() Color {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) SetPolygons(polygons Array, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygons.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetPolygons() Array {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) SetVertexColors(vertex_colors PackedColorArray, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3546319833) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{vertex_colors.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetVertexColors() PackedColorArray {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1392750486) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedColorArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetTexture() Texture2D {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) SetTextureOffset(texture_offset Vector2, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture_offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetTextureOffset() Vector2 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) SetTextureRotation(texture_rotation float64, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_rotation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetTextureRotation() float64 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Polygon2D) SetTextureScale(texture_scale Vector2, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture_scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetTextureScale() Vector2 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) SetInvertEnabled(invert bool, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_invert_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetInvertEnabled() bool {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_invert_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Polygon2D) SetAntialiased(antialiased bool, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_antialiased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&antialiased) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetAntialiased() bool {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_antialiased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Polygon2D) SetInvertBorder(invert_border float64, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_invert_border")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert_border) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetInvertBorder() float64 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_invert_border")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Polygon2D) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetOffset() Vector2 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) AddBone(path NodePath, weights PackedFloat32Array, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 703042815) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), weights.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetBoneCount() int64 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Polygon2D) GetBonePath(index int64, ) NodePath {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) GetBoneWeights(index int64, ) PackedFloat32Array {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_weights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1542882410) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) EraseBone(index int64, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) ClearBones()  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) SetBonePath(index int64, path NodePath, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761262315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) SetBoneWeights(index int64, weights PackedFloat32Array, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_weights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1345852415) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , weights.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) SetSkeleton(skeleton NodePath, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{skeleton.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetSkeleton() NodePath {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Polygon2D) SetInternalVertexCount(internal_vertex_count int64, )  {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_internal_vertex_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&internal_vertex_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Polygon2D) GetInternalVertexCount() int64 {
  classNameV := StringNameFromStr("Polygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_internal_vertex_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
