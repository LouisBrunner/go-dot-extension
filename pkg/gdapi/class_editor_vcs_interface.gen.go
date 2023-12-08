// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorVCSInterface struct {
  obj gdc.ObjectPtr
}

func (me *EditorVCSInterface) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorVCSInterface) BaseClass() string {
  return "EditorVCSInterface"
}

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

func  (me *EditorVCSInterface) XInitialize(project_path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XSetCredentials(username String, password String, ssh_public_key_path String, ssh_private_key_path String, ssh_passphrase String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XGetModifiedFilesData() { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XStageFile(file_path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XUnstageFile(file_path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XDiscardFile(file_path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XCommit(msg String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XGetDiff(identifier String, area int, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XShutDown() { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XGetVcsName() { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XGetPreviousCommits(max_commits int, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XGetBranchList() { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XGetRemotes() { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XCreateBranch(branch_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XRemoveBranch(branch_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XCreateRemote(remote_name String, remote_url String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XRemoveRemote(remote_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XGetCurrentBranchName() { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XCheckoutBranch(branch_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XPull(remote String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XPush(remote String, force bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XFetch(remote String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) XGetLineDiff(file_path String, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) CreateDiffLine(new_line_no int, old_line_no int, content String, status String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) CreateDiffHunk(old_start int, new_start int, old_lines int, new_lines int, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) CreateDiffFile(new_file String, old_file String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) CreateCommit(msg String, author String, id String, unix_timestamp int, offset_minutes int, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) CreateStatusFile(file_path String, change_type EditorVCSInterfaceChangeType, area EditorVCSInterfaceTreeArea, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) AddDiffHunksIntoDiffFile(diff_file Dictionary, diff_hunks Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) AddLineDiffsIntoDiffHunk(diff_hunk Dictionary, line_diffs Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorVCSInterface) PopupError(msg String, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
