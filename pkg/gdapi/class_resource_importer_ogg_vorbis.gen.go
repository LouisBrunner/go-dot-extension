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

type ptrsForResourceImporterOggVorbisList struct {
	fnLoadFromBuffer gdc.MethodBindPtr
	fnLoadFromFile   gdc.MethodBindPtr
}

var ptrsForResourceImporterOggVorbis ptrsForResourceImporterOggVorbisList

func initResourceImporterOggVorbisPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ResourceImporterOggVorbis")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("load_from_buffer")
		defer methodName.Destroy()
		ptrsForResourceImporterOggVorbis.fnLoadFromBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 354904730))
	}
	{
		methodName := StringNameFromStr("load_from_file")
		defer methodName.Destroy()
		ptrsForResourceImporterOggVorbis.fnLoadFromFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 797568536))
	}
}

type ResourceImporterOggVorbis struct {
	ResourceImporter
}

func (me *ResourceImporterOggVorbis) BaseClass() string {
	return "ResourceImporterOggVorbis"
}

func NewResourceImporterOggVorbis() *ResourceImporterOggVorbis {
	str := StringNameFromStr("ResourceImporterOggVorbis") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ResourceImporterOggVorbis{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ResourceImporterOggVorbis) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ResourceImporterOggVorbis) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterOggVorbis) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func ResourceImporterOggVorbisLoadFromBuffer(buffer PackedByteArray) AudioStreamOggVorbis {
	cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAudioStreamOggVorbis()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceImporterOggVorbis.fnLoadFromBuffer), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func ResourceImporterOggVorbisLoadFromFile(path String) AudioStreamOggVorbis {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAudioStreamOggVorbis()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceImporterOggVorbis.fnLoadFromFile), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
