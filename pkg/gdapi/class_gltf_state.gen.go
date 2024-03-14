// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFState struct {
  Resource
}

func (me *GLTFState) BaseClass() string {
  return "GLTFState"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(extension_name.AsCTypePtr()), gdc.ConstTypePtr(&required), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetJson() Dictionary {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_json")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2382534195) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetJson(json Dictionary, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_json")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(json.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetMajorVersion() int {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_major_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetMajorVersion(major_version int, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_major_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&major_version), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetMinorVersion() int {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minor_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetMinorVersion(minor_version int, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_minor_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&minor_version), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetCopyright() String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_copyright")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetCopyright(copyright String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_copyright")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(copyright.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetGlbData() PackedByteArray {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glb_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2115431945) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetGlbData(glb_data PackedByteArray, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glb_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2971499966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(glb_data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetUseNamedSkinBinds() bool {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_named_skin_binds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetUseNamedSkinBinds(use_named_skin_binds bool, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_named_skin_binds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_named_skin_binds), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetNodes() GLTFNode {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFNode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetNodes(nodes GLTFNode, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(nodes.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetBuffers() PackedByteArray {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetBuffers(buffers PackedByteArray, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffers.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetBufferViews() GLTFBufferView {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer_views")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFBufferView
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetBufferViews(buffer_views GLTFBufferView, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffer_views")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer_views.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetAccessors() GLTFAccessor {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_accessors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFAccessor
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetAccessors(accessors GLTFAccessor, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_accessors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(accessors.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetMeshes() GLTFMesh {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_meshes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFMesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetMeshes(meshes GLTFMesh, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_meshes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(meshes.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetAnimationPlayersCount(idx int, ) int {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_players_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) GetAnimationPlayer(idx int, ) AnimationPlayer {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_player")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 925043400) // FIXME: should cache?
  var ret AnimationPlayer
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) GetMaterials() Material {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_materials")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret Material
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetMaterials(materials Material, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_materials")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(materials.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetSceneName() String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetSceneName(scene_name String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scene_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scene_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetBasePath() String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_base_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetBasePath(base_path String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_base_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(base_path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetFilename() String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filename")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetFilename(filename String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filename")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(filename.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetRootNodes() PackedInt32Array {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetRootNodes(root_nodes PackedInt32Array, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(root_nodes.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetTextures() GLTFTexture {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_textures")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFTexture
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetTextures(textures GLTFTexture, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_textures")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(textures.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetTextureSamplers() GLTFTextureSampler {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_samplers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFTextureSampler
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetTextureSamplers(texture_samplers GLTFTextureSampler, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_samplers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture_samplers.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetImages() Texture2D {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_images")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetImages(images Texture2D, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_images")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(images.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetSkins() GLTFSkin {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skins")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFSkin
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetSkins(skins GLTFSkin, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skins")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(skins.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetCameras() GLTFCamera {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cameras")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFCamera
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetCameras(cameras GLTFCamera, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cameras")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(cameras.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetLights() GLTFLight {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFLight
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetLights(lights GLTFLight, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(lights.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetUniqueNames() String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unique_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetUniqueNames(unique_names String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unique_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(unique_names.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetUniqueAnimationNames() String {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unique_animation_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetUniqueAnimationNames(unique_animation_names String, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unique_animation_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(unique_animation_names.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetSkeletons() GLTFSkeleton {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeletons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFSkeleton
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetSkeletons(skeletons GLTFSkeleton, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skeletons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(skeletons.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetCreateAnimations() bool {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_create_animations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetCreateAnimations(create_animations bool, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_create_animations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&create_animations), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetAnimations() GLTFAnimation {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret GLTFAnimation
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetAnimations(animations GLTFAnimation, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(animations.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetSceneNode(idx int, ) Node {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4253421667) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) GetNodeIndex(scene_node Node, ) int {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1205807060) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scene_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) GetAdditionalData(extension_name StringName, ) Variant {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_additional_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2138907829) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(extension_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetAdditionalData(extension_name StringName, additional_data Variant, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_additional_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(extension_name.AsCTypePtr()), gdc.ConstTypePtr(additional_data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFState) GetHandleBinaryImage() int {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_handle_binary_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFState) SetHandleBinaryImage(method int, )  {
  classNameV := StringNameFromStr("GLTFState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_handle_binary_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&method), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
