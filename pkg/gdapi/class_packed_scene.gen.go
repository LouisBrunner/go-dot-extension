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

type ptrsForPackedSceneList struct {
	fnPack           gdc.MethodBindPtr
	fnInstantiate    gdc.MethodBindPtr
	fnCanInstantiate gdc.MethodBindPtr
	fnGetState       gdc.MethodBindPtr
}

var ptrsForPackedScene ptrsForPackedSceneList

func initPackedScenePtrs(iface gdc.Interface) {

	className := StringNameFromStr("PackedScene")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("pack")
		defer methodName.Destroy()
		ptrsForPackedScene.fnPack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2584678054))
	}
	{
		methodName := StringNameFromStr("instantiate")
		defer methodName.Destroy()
		ptrsForPackedScene.fnInstantiate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2628778455))
	}
	{
		methodName := StringNameFromStr("can_instantiate")
		defer methodName.Destroy()
		ptrsForPackedScene.fnCanInstantiate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_state")
		defer methodName.Destroy()
		ptrsForPackedScene.fnGetState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3479783971))
	}

}

type PackedScene struct {
	Resource
}

func (me *PackedScene) BaseClass() string {
	return "PackedScene"
}

func NewPackedScene() *PackedScene {
	str := StringNameFromStr("PackedScene") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PackedScene{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type PackedSceneGenEditState int

const (
	PackedSceneGenEditStateGenEditStateDisabled      PackedSceneGenEditState = 0
	PackedSceneGenEditStateGenEditStateInstance      PackedSceneGenEditState = 1
	PackedSceneGenEditStateGenEditStateMain          PackedSceneGenEditState = 2
	PackedSceneGenEditStateGenEditStateMainInherited PackedSceneGenEditState = 3
)

func (me *PackedScene) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PackedScene) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PackedScene) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PackedScene) Pack(path Node) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPackedScene.fnPack), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PackedScene) Instantiate(edit_state PackedSceneGenEditState) Node {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edit_state)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()
	pinner.Pin(&edit_state)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPackedScene.fnInstantiate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PackedScene) CanInstantiate() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPackedScene.fnCanInstantiate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedScene) GetState() SceneState {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSceneState()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPackedScene.fnGetState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
