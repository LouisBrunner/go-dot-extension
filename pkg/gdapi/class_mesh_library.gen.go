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

type ptrsForMeshLibraryList struct {
	fnCreateItem                     gdc.MethodBindPtr
	fnSetItemName                    gdc.MethodBindPtr
	fnSetItemMesh                    gdc.MethodBindPtr
	fnSetItemMeshTransform           gdc.MethodBindPtr
	fnSetItemNavigationMesh          gdc.MethodBindPtr
	fnSetItemNavigationMeshTransform gdc.MethodBindPtr
	fnSetItemNavigationLayers        gdc.MethodBindPtr
	fnSetItemShapes                  gdc.MethodBindPtr
	fnSetItemPreview                 gdc.MethodBindPtr
	fnGetItemName                    gdc.MethodBindPtr
	fnGetItemMesh                    gdc.MethodBindPtr
	fnGetItemMeshTransform           gdc.MethodBindPtr
	fnGetItemNavigationMesh          gdc.MethodBindPtr
	fnGetItemNavigationMeshTransform gdc.MethodBindPtr
	fnGetItemNavigationLayers        gdc.MethodBindPtr
	fnGetItemShapes                  gdc.MethodBindPtr
	fnGetItemPreview                 gdc.MethodBindPtr
	fnRemoveItem                     gdc.MethodBindPtr
	fnFindItemByName                 gdc.MethodBindPtr
	fnClear                          gdc.MethodBindPtr
	fnGetItemList                    gdc.MethodBindPtr
	fnGetLastUnusedItemId            gdc.MethodBindPtr
}

var ptrsForMeshLibrary ptrsForMeshLibraryList

func initMeshLibraryPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MeshLibrary")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("create_item")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnCreateItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_item_name")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnSetItemName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("set_item_mesh")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnSetItemMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969122797))
	}
	{
		methodName := StringNameFromStr("set_item_mesh_transform")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnSetItemMeshTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3616898986))
	}
	{
		methodName := StringNameFromStr("set_item_navigation_mesh")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnSetItemNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3483353960))
	}
	{
		methodName := StringNameFromStr("set_item_navigation_mesh_transform")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnSetItemNavigationMeshTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3616898986))
	}
	{
		methodName := StringNameFromStr("set_item_navigation_layers")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnSetItemNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_item_shapes")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnSetItemShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 537221740))
	}
	{
		methodName := StringNameFromStr("set_item_preview")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnSetItemPreview = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
	}
	{
		methodName := StringNameFromStr("get_item_name")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnGetItemName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_item_mesh")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnGetItemMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1576363275))
	}
	{
		methodName := StringNameFromStr("get_item_mesh_transform")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnGetItemMeshTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("get_item_navigation_mesh")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnGetItemNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2729647406))
	}
	{
		methodName := StringNameFromStr("get_item_navigation_mesh_transform")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnGetItemNavigationMeshTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("get_item_navigation_layers")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnGetItemNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_item_shapes")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnGetItemShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
	}
	{
		methodName := StringNameFromStr("get_item_preview")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnGetItemPreview = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
	}
	{
		methodName := StringNameFromStr("remove_item")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnRemoveItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("find_item_by_name")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnFindItemByName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_item_list")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnGetItemList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
	}
	{
		methodName := StringNameFromStr("get_last_unused_item_id")
		defer methodName.Destroy()
		ptrsForMeshLibrary.fnGetLastUnusedItemId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type MeshLibrary struct {
	Resource
}

func (me *MeshLibrary) BaseClass() string {
	return "MeshLibrary"
}

func NewMeshLibrary() *MeshLibrary {
	str := StringNameFromStr("MeshLibrary") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MeshLibrary{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MeshLibrary) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MeshLibrary) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MeshLibrary) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *MeshLibrary) CreateItem(id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnCreateItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) SetItemName(id int64, name String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnSetItemName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) SetItemMesh(id int64, mesh Mesh) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnSetItemMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) SetItemMeshTransform(id int64, mesh_transform Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), mesh_transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnSetItemMeshTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) SetItemNavigationMesh(id int64, navigation_mesh NavigationMesh) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), navigation_mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnSetItemNavigationMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) SetItemNavigationMeshTransform(id int64, navigation_mesh Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), navigation_mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnSetItemNavigationMeshTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) SetItemNavigationLayers(id int64, navigation_layers int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&navigation_layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnSetItemNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) SetItemShapes(id int64, shapes Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), shapes.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnSetItemShapes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) SetItemPreview(id int64, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnSetItemPreview), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) GetItemName(id int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnGetItemName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshLibrary) GetItemMesh(id int64) Mesh {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMesh()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnGetItemMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshLibrary) GetItemMeshTransform(id int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnGetItemMeshTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshLibrary) GetItemNavigationMesh(id int64) NavigationMesh {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNavigationMesh()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnGetItemNavigationMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshLibrary) GetItemNavigationMeshTransform(id int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnGetItemNavigationMeshTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshLibrary) GetItemNavigationLayers(id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnGetItemNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MeshLibrary) GetItemShapes(id int64) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnGetItemShapes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshLibrary) GetItemPreview(id int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnGetItemPreview), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshLibrary) RemoveItem(id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnRemoveItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) FindItemByName(name String) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnFindItemByName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MeshLibrary) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshLibrary) GetItemList() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnGetItemList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshLibrary) GetLastUnusedItemId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshLibrary.fnGetLastUnusedItemId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
