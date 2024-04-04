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

type GLTFState struct {
  Resource
}

func (me *GLTFState) BaseClass() string {
  return "GLTFState"
}

func NewGLTFState() *GLTFState {
  str := StringNameFromStr("GLTFState") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFState{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  GLTFStateHandleBinaryDiscardTextures = "0" // TODO: construct correctly
  GLTFStateHandleBinaryExtractTextures = "1" // TODO: construct correctly
  GLTFStateHandleBinaryEmbedAsBasisu = "2" // TODO: construct correctly
  GLTFStateHandleBinaryEmbedAsUncompressed = "3" // TODO: construct correctly
)

// Enums

func (me *GLTFState) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFState) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFState) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GLTFState) AddUsedExtension(extension_name String, required bool, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_used_extension")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678287736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{extension_name.AsCTypePtr(), gdc.ConstTypePtr(&required) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetJson() Dictionary {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_json")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2382534195) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetJson(json Dictionary, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_json")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{json.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetMajorVersion() int64 {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_major_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) SetMajorVersion(major_version int64, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_major_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&major_version) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetMinorVersion() int64 {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minor_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) SetMinorVersion(minor_version int64, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_minor_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&minor_version) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetCopyright() String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_copyright")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetCopyright(copyright String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_copyright")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{copyright.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetGlbData() PackedByteArray {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glb_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2115431945) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetGlbData(glb_data PackedByteArray, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glb_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2971499966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{glb_data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetUseNamedSkinBinds() bool {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_named_skin_binds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) SetUseNamedSkinBinds(use_named_skin_binds bool, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_named_skin_binds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_named_skin_binds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetNodes() []GLTFNode {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFNode](ret)
}

func  (me *GLTFState) SetNodes(nodes []GLTFNode, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&nodes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetBuffers() []PackedByteArray {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PackedByteArray](ret)
}

func  (me *GLTFState) SetBuffers(buffers []PackedByteArray, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetBufferViews() []GLTFBufferView {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer_views")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFBufferView](ret)
}

func  (me *GLTFState) SetBufferViews(buffer_views []GLTFBufferView, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffer_views")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_views) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetAccessors() []GLTFAccessor {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_accessors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFAccessor](ret)
}

func  (me *GLTFState) SetAccessors(accessors []GLTFAccessor, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_accessors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&accessors) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetMeshes() []GLTFMesh {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_meshes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFMesh](ret)
}

func  (me *GLTFState) SetMeshes(meshes []GLTFMesh, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_meshes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&meshes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetAnimationPlayersCount(idx int64, ) int64 {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_players_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) GetAnimationPlayer(idx int64, ) AnimationPlayer {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_player")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 925043400) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationPlayer()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) GetMaterials() []Material {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_materials")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Material](ret)
}

func  (me *GLTFState) SetMaterials(materials []Material, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_materials")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&materials) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetSceneName() String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetSceneName(scene_name String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scene_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scene_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetBasePath() String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_base_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetBasePath(base_path String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_base_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{base_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetFilename() String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filename")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetFilename(filename String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filename")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{filename.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetRootNodes() PackedInt32Array {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetRootNodes(root_nodes PackedInt32Array, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{root_nodes.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetTextures() []GLTFTexture {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_textures")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFTexture](ret)
}

func  (me *GLTFState) SetTextures(textures []GLTFTexture, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_textures")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&textures) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetTextureSamplers() []GLTFTextureSampler {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_samplers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFTextureSampler](ret)
}

func  (me *GLTFState) SetTextureSamplers(texture_samplers []GLTFTextureSampler, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_samplers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_samplers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetImages() []Texture2D {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_images")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Texture2D](ret)
}

func  (me *GLTFState) SetImages(images []Texture2D, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_images")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&images) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetSkins() []GLTFSkin {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skins")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFSkin](ret)
}

func  (me *GLTFState) SetSkins(skins []GLTFSkin, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skins")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skins) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetCameras() []GLTFCamera {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cameras")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFCamera](ret)
}

func  (me *GLTFState) SetCameras(cameras []GLTFCamera, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cameras")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cameras) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetLights() []GLTFLight {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFLight](ret)
}

func  (me *GLTFState) SetLights(lights []GLTFLight, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lights) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetUniqueNames() []String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unique_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[String](ret)
}

func  (me *GLTFState) SetUniqueNames(unique_names []String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unique_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unique_names) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetUniqueAnimationNames() []String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unique_animation_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[String](ret)
}

func  (me *GLTFState) SetUniqueAnimationNames(unique_animation_names []String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unique_animation_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unique_animation_names) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetSkeletons() []GLTFSkeleton {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeletons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFSkeleton](ret)
}

func  (me *GLTFState) SetSkeletons(skeletons []GLTFSkeleton, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skeletons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skeletons) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetCreateAnimations() bool {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_create_animations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) SetCreateAnimations(create_animations bool, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_create_animations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&create_animations) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetAnimations() []GLTFAnimation {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[GLTFAnimation](ret)
}

func  (me *GLTFState) SetAnimations(animations []GLTFAnimation, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&animations) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetSceneNode(idx int64, ) Node {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4253421667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) GetNodeIndex(scene_node Node, ) int64 {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1205807060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scene_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) GetAdditionalData(extension_name StringName, ) Variant {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_additional_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2138907829) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{extension_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetAdditionalData(extension_name StringName, additional_data Variant, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_additional_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{extension_name.AsCTypePtr(), additional_data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetHandleBinaryImage() int64 {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_handle_binary_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) SetHandleBinaryImage(method int64, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_handle_binary_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&method) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
