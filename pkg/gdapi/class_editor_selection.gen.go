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

type ptrsForEditorSelectionList struct {
	fnClear                         gdc.MethodBindPtr
	fnAddNode                       gdc.MethodBindPtr
	fnRemoveNode                    gdc.MethodBindPtr
	fnGetSelectedNodes              gdc.MethodBindPtr
	fnGetTransformableSelectedNodes gdc.MethodBindPtr
}

var ptrsForEditorSelection ptrsForEditorSelectionList

func initEditorSelectionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorSelection")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForEditorSelection.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("add_node")
		defer methodName.Destroy()
		ptrsForEditorSelection.fnAddNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("remove_node")
		defer methodName.Destroy()
		ptrsForEditorSelection.fnRemoveNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("get_selected_nodes")
		defer methodName.Destroy()
		ptrsForEditorSelection.fnGetSelectedNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("get_transformable_selected_nodes")
		defer methodName.Destroy()
		ptrsForEditorSelection.fnGetTransformableSelectedNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
}

type EditorSelection struct {
	Object
}

func (me *EditorSelection) BaseClass() string {
	return "EditorSelection"
}

func NewEditorSelection() *EditorSelection {
	str := StringNameFromStr("EditorSelection") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorSelection{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorSelection) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorSelection) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorSelection) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorSelection) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSelection.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSelection) AddNode(node Node) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSelection.fnAddNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSelection) RemoveNode(node Node) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSelection.fnRemoveNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorSelection) GetSelectedNodes() []Node {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSelection.fnGetSelectedNodes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Node](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *EditorSelection) GetTransformableSelectedNodes() []Node {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSelection.fnGetTransformableSelectedNodes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Node](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Signals

type EditorSelectionSelectionChangedSignalFn func()

func (me *EditorSelection) ConnectSelectionChanged(subs SignalSubscribers, fn EditorSelectionSelectionChangedSignalFn) {
	sig := StringNameFromStr("selection_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorSelection) DisconnectSelectionChanged(subs SignalSubscribers, fn EditorSelectionSelectionChangedSignalFn) {
	sig := StringNameFromStr("selection_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
