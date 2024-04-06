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

type ptrsForEditorNode3DGizmoPluginList struct {
	fnXHasGizmo                  gdc.MethodBindPtr
	fnXCreateGizmo               gdc.MethodBindPtr
	fnXGetGizmoName              gdc.MethodBindPtr
	fnXGetPriority               gdc.MethodBindPtr
	fnXCanBeHidden               gdc.MethodBindPtr
	fnXIsSelectableWhenHidden    gdc.MethodBindPtr
	fnXRedraw                    gdc.MethodBindPtr
	fnXGetHandleName             gdc.MethodBindPtr
	fnXIsHandleHighlighted       gdc.MethodBindPtr
	fnXGetHandleValue            gdc.MethodBindPtr
	fnXSetHandle                 gdc.MethodBindPtr
	fnXCommitHandle              gdc.MethodBindPtr
	fnXSubgizmosIntersectRay     gdc.MethodBindPtr
	fnXSubgizmosIntersectFrustum gdc.MethodBindPtr
	fnXGetSubgizmoTransform      gdc.MethodBindPtr
	fnXSetSubgizmoTransform      gdc.MethodBindPtr
	fnXCommitSubgizmos           gdc.MethodBindPtr
	fnCreateMaterial             gdc.MethodBindPtr
	fnCreateIconMaterial         gdc.MethodBindPtr
	fnCreateHandleMaterial       gdc.MethodBindPtr
	fnAddMaterial                gdc.MethodBindPtr
	fnGetMaterial                gdc.MethodBindPtr
}

var ptrsForEditorNode3DGizmoPlugin ptrsForEditorNode3DGizmoPluginList

func initEditorNode3DGizmoPluginPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorNode3DGizmoPlugin")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("create_material")
		defer methodName.Destroy()
		ptrsForEditorNode3DGizmoPlugin.fnCreateMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3486012546))
	}
	{
		methodName := StringNameFromStr("create_icon_material")
		defer methodName.Destroy()
		ptrsForEditorNode3DGizmoPlugin.fnCreateIconMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3804976916))
	}
	{
		methodName := StringNameFromStr("create_handle_material")
		defer methodName.Destroy()
		ptrsForEditorNode3DGizmoPlugin.fnCreateHandleMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2486475223))
	}
	{
		methodName := StringNameFromStr("add_material")
		defer methodName.Destroy()
		ptrsForEditorNode3DGizmoPlugin.fnAddMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1374068695))
	}
	{
		methodName := StringNameFromStr("get_material")
		defer methodName.Destroy()
		ptrsForEditorNode3DGizmoPlugin.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 974464017))
	}

}

type EditorNode3DGizmoPlugin struct {
	Resource
}

func (me *EditorNode3DGizmoPlugin) BaseClass() string {
	return "EditorNode3DGizmoPlugin"
}

func NewEditorNode3DGizmoPlugin() *EditorNode3DGizmoPlugin {
	str := StringNameFromStr("EditorNode3DGizmoPlugin") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorNode3DGizmoPlugin{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorNode3DGizmoPlugin) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorNode3DGizmoPlugin) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorNode3DGizmoPlugin) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorNode3DGizmoPlugin) CreateMaterial(name String, color Color, billboard bool, on_top bool, use_vertex_color bool) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&billboard), gdc.ConstTypePtr(&on_top), gdc.ConstTypePtr(&use_vertex_color)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmoPlugin.fnCreateMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorNode3DGizmoPlugin) CreateIconMaterial(name String, texture Texture2D, on_top bool, color Color) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), texture.AsCTypePtr(), gdc.ConstTypePtr(&on_top), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmoPlugin.fnCreateIconMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorNode3DGizmoPlugin) CreateHandleMaterial(name String, billboard bool, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&billboard), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmoPlugin.fnCreateHandleMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorNode3DGizmoPlugin) AddMaterial(name String, material StandardMaterial3D) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmoPlugin.fnAddMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorNode3DGizmoPlugin) GetMaterial(name String, gizmo EditorNode3DGizmo) StandardMaterial3D {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gizmo.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStandardMaterial3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmoPlugin.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
