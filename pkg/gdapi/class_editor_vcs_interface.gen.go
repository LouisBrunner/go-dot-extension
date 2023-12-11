// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorVCSInterface struct {
  obj gdc.ObjectPtr
}

func (me *EditorVCSInterface) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorVCSInterface) BaseClass() string {
  return "EditorVCSInterface"
}



// Enums

type EditorVCSInterfaceChangeType int
const (
  EditorVCSInterfaceChangeTypeChangeTypeNew EditorVCSInterfaceChangeType = 0
  EditorVCSInterfaceChangeTypeChangeTypeModified EditorVCSInterfaceChangeType = 1
  EditorVCSInterfaceChangeTypeChangeTypeRenamed EditorVCSInterfaceChangeType = 2
  EditorVCSInterfaceChangeTypeChangeTypeDeleted EditorVCSInterfaceChangeType = 3
  EditorVCSInterfaceChangeTypeChangeTypeTypechange EditorVCSInterfaceChangeType = 4
  EditorVCSInterfaceChangeTypeChangeTypeUnmerged EditorVCSInterfaceChangeType = 5
)

type EditorVCSInterfaceTreeArea int
const (
  EditorVCSInterfaceTreeAreaTreeAreaCommit EditorVCSInterfaceTreeArea = 0
  EditorVCSInterfaceTreeAreaTreeAreaStaged EditorVCSInterfaceTreeArea = 1
  EditorVCSInterfaceTreeAreaTreeAreaUnstaged EditorVCSInterfaceTreeArea = 2
)

func (me *EditorVCSInterface) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorVCSInterface) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorVCSInterface) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorVCSInterface) CreateDiffLine(new_line_no int, old_line_no int, content String, status String, ) Dictionary {
  classNameV := StringNameFromStr("EditorVCSInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_diff_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2901184053) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&new_line_no), gdc.ConstTypePtr(&old_line_no), gdc.ConstTypePtr(content.AsCTypePtr()), gdc.ConstTypePtr(status.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorVCSInterface) CreateDiffHunk(old_start int, new_start int, old_lines int, new_lines int, ) Dictionary {
  classNameV := StringNameFromStr("EditorVCSInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_diff_hunk")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3784842090) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&old_start), gdc.ConstTypePtr(&new_start), gdc.ConstTypePtr(&old_lines), gdc.ConstTypePtr(&new_lines), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorVCSInterface) CreateDiffFile(new_file String, old_file String, ) Dictionary {
  classNameV := StringNameFromStr("EditorVCSInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_diff_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2723227684) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(new_file.AsCTypePtr()), gdc.ConstTypePtr(old_file.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorVCSInterface) CreateCommit(msg String, author String, id String, unix_timestamp int, offset_minutes int, ) Dictionary {
  classNameV := StringNameFromStr("EditorVCSInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_commit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1075983584) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(msg.AsCTypePtr()), gdc.ConstTypePtr(author.AsCTypePtr()), gdc.ConstTypePtr(id.AsCTypePtr()), gdc.ConstTypePtr(&unix_timestamp), gdc.ConstTypePtr(&offset_minutes), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorVCSInterface) CreateStatusFile(file_path String, change_type EditorVCSInterfaceChangeType, area EditorVCSInterfaceTreeArea, ) Dictionary {
  classNameV := StringNameFromStr("EditorVCSInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_status_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1083471673) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file_path.AsCTypePtr()), gdc.ConstTypePtr(&change_type), gdc.ConstTypePtr(&area), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorVCSInterface) AddDiffHunksIntoDiffFile(diff_file Dictionary, diff_hunks Dictionary, ) Dictionary {
  classNameV := StringNameFromStr("EditorVCSInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_diff_hunks_into_diff_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015243225) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(diff_file.AsCTypePtr()), gdc.ConstTypePtr(diff_hunks.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorVCSInterface) AddLineDiffsIntoDiffHunk(diff_hunk Dictionary, line_diffs Dictionary, ) Dictionary {
  classNameV := StringNameFromStr("EditorVCSInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_line_diffs_into_diff_hunk")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015243225) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(diff_hunk.AsCTypePtr()), gdc.ConstTypePtr(line_diffs.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorVCSInterface) PopupError(msg String, )  {
  classNameV := StringNameFromStr("EditorVCSInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(msg.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
