// automatically generated by stateify.

package fs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *StableAttr) StateTypeName() string {
	return "pkg/sentry/fs.StableAttr"
}

func (s *StableAttr) StateFields() []string {
	return []string{
		"Type",
		"DeviceID",
		"InodeID",
		"BlockSize",
		"DeviceFileMajor",
		"DeviceFileMinor",
	}
}

func (s *StableAttr) beforeSave() {}

// +checklocksignore
func (s *StableAttr) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Type)
	stateSinkObject.Save(1, &s.DeviceID)
	stateSinkObject.Save(2, &s.InodeID)
	stateSinkObject.Save(3, &s.BlockSize)
	stateSinkObject.Save(4, &s.DeviceFileMajor)
	stateSinkObject.Save(5, &s.DeviceFileMinor)
}

func (s *StableAttr) afterLoad() {}

// +checklocksignore
func (s *StableAttr) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Type)
	stateSourceObject.Load(1, &s.DeviceID)
	stateSourceObject.Load(2, &s.InodeID)
	stateSourceObject.Load(3, &s.BlockSize)
	stateSourceObject.Load(4, &s.DeviceFileMajor)
	stateSourceObject.Load(5, &s.DeviceFileMinor)
}

func (ua *UnstableAttr) StateTypeName() string {
	return "pkg/sentry/fs.UnstableAttr"
}

func (ua *UnstableAttr) StateFields() []string {
	return []string{
		"Size",
		"Usage",
		"Perms",
		"Owner",
		"AccessTime",
		"ModificationTime",
		"StatusChangeTime",
		"Links",
	}
}

func (ua *UnstableAttr) beforeSave() {}

// +checklocksignore
func (ua *UnstableAttr) StateSave(stateSinkObject state.Sink) {
	ua.beforeSave()
	stateSinkObject.Save(0, &ua.Size)
	stateSinkObject.Save(1, &ua.Usage)
	stateSinkObject.Save(2, &ua.Perms)
	stateSinkObject.Save(3, &ua.Owner)
	stateSinkObject.Save(4, &ua.AccessTime)
	stateSinkObject.Save(5, &ua.ModificationTime)
	stateSinkObject.Save(6, &ua.StatusChangeTime)
	stateSinkObject.Save(7, &ua.Links)
}

func (ua *UnstableAttr) afterLoad() {}

// +checklocksignore
func (ua *UnstableAttr) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ua.Size)
	stateSourceObject.Load(1, &ua.Usage)
	stateSourceObject.Load(2, &ua.Perms)
	stateSourceObject.Load(3, &ua.Owner)
	stateSourceObject.Load(4, &ua.AccessTime)
	stateSourceObject.Load(5, &ua.ModificationTime)
	stateSourceObject.Load(6, &ua.StatusChangeTime)
	stateSourceObject.Load(7, &ua.Links)
}

func (a *AttrMask) StateTypeName() string {
	return "pkg/sentry/fs.AttrMask"
}

func (a *AttrMask) StateFields() []string {
	return []string{
		"Type",
		"DeviceID",
		"InodeID",
		"BlockSize",
		"Size",
		"Usage",
		"Perms",
		"UID",
		"GID",
		"AccessTime",
		"ModificationTime",
		"StatusChangeTime",
		"Links",
	}
}

func (a *AttrMask) beforeSave() {}

// +checklocksignore
func (a *AttrMask) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.Type)
	stateSinkObject.Save(1, &a.DeviceID)
	stateSinkObject.Save(2, &a.InodeID)
	stateSinkObject.Save(3, &a.BlockSize)
	stateSinkObject.Save(4, &a.Size)
	stateSinkObject.Save(5, &a.Usage)
	stateSinkObject.Save(6, &a.Perms)
	stateSinkObject.Save(7, &a.UID)
	stateSinkObject.Save(8, &a.GID)
	stateSinkObject.Save(9, &a.AccessTime)
	stateSinkObject.Save(10, &a.ModificationTime)
	stateSinkObject.Save(11, &a.StatusChangeTime)
	stateSinkObject.Save(12, &a.Links)
}

func (a *AttrMask) afterLoad() {}

// +checklocksignore
func (a *AttrMask) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.Type)
	stateSourceObject.Load(1, &a.DeviceID)
	stateSourceObject.Load(2, &a.InodeID)
	stateSourceObject.Load(3, &a.BlockSize)
	stateSourceObject.Load(4, &a.Size)
	stateSourceObject.Load(5, &a.Usage)
	stateSourceObject.Load(6, &a.Perms)
	stateSourceObject.Load(7, &a.UID)
	stateSourceObject.Load(8, &a.GID)
	stateSourceObject.Load(9, &a.AccessTime)
	stateSourceObject.Load(10, &a.ModificationTime)
	stateSourceObject.Load(11, &a.StatusChangeTime)
	stateSourceObject.Load(12, &a.Links)
}

func (p *PermMask) StateTypeName() string {
	return "pkg/sentry/fs.PermMask"
}

func (p *PermMask) StateFields() []string {
	return []string{
		"Read",
		"Write",
		"Execute",
	}
}

func (p *PermMask) beforeSave() {}

// +checklocksignore
func (p *PermMask) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.Read)
	stateSinkObject.Save(1, &p.Write)
	stateSinkObject.Save(2, &p.Execute)
}

func (p *PermMask) afterLoad() {}

// +checklocksignore
func (p *PermMask) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.Read)
	stateSourceObject.Load(1, &p.Write)
	stateSourceObject.Load(2, &p.Execute)
}

func (f *FilePermissions) StateTypeName() string {
	return "pkg/sentry/fs.FilePermissions"
}

func (f *FilePermissions) StateFields() []string {
	return []string{
		"User",
		"Group",
		"Other",
		"Sticky",
		"SetUID",
		"SetGID",
	}
}

func (f *FilePermissions) beforeSave() {}

// +checklocksignore
func (f *FilePermissions) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.User)
	stateSinkObject.Save(1, &f.Group)
	stateSinkObject.Save(2, &f.Other)
	stateSinkObject.Save(3, &f.Sticky)
	stateSinkObject.Save(4, &f.SetUID)
	stateSinkObject.Save(5, &f.SetGID)
}

func (f *FilePermissions) afterLoad() {}

// +checklocksignore
func (f *FilePermissions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.User)
	stateSourceObject.Load(1, &f.Group)
	stateSourceObject.Load(2, &f.Other)
	stateSourceObject.Load(3, &f.Sticky)
	stateSourceObject.Load(4, &f.SetUID)
	stateSourceObject.Load(5, &f.SetGID)
}

func (f *FileOwner) StateTypeName() string {
	return "pkg/sentry/fs.FileOwner"
}

func (f *FileOwner) StateFields() []string {
	return []string{
		"UID",
		"GID",
	}
}

func (f *FileOwner) beforeSave() {}

// +checklocksignore
func (f *FileOwner) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.UID)
	stateSinkObject.Save(1, &f.GID)
}

func (f *FileOwner) afterLoad() {}

// +checklocksignore
func (f *FileOwner) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.UID)
	stateSourceObject.Load(1, &f.GID)
}

func (d *DentAttr) StateTypeName() string {
	return "pkg/sentry/fs.DentAttr"
}

func (d *DentAttr) StateFields() []string {
	return []string{
		"Type",
		"InodeID",
	}
}

func (d *DentAttr) beforeSave() {}

// +checklocksignore
func (d *DentAttr) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.Type)
	stateSinkObject.Save(1, &d.InodeID)
}

func (d *DentAttr) afterLoad() {}

// +checklocksignore
func (d *DentAttr) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.Type)
	stateSourceObject.Load(1, &d.InodeID)
}

func (s *SortedDentryMap) StateTypeName() string {
	return "pkg/sentry/fs.SortedDentryMap"
}

func (s *SortedDentryMap) StateFields() []string {
	return []string{
		"names",
		"entries",
	}
}

func (s *SortedDentryMap) beforeSave() {}

// +checklocksignore
func (s *SortedDentryMap) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.names)
	stateSinkObject.Save(1, &s.entries)
}

func (s *SortedDentryMap) afterLoad() {}

// +checklocksignore
func (s *SortedDentryMap) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.names)
	stateSourceObject.Load(1, &s.entries)
}

func (d *Dirent) StateTypeName() string {
	return "pkg/sentry/fs.Dirent"
}

func (d *Dirent) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"userVisible",
		"Inode",
		"name",
		"parent",
		"deleted",
		"mounted",
		"children",
	}
}

// +checklocksignore
func (d *Dirent) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	var childrenValue map[string]*Dirent
	childrenValue = d.saveChildren()
	stateSinkObject.SaveValue(7, childrenValue)
	stateSinkObject.Save(0, &d.AtomicRefCount)
	stateSinkObject.Save(1, &d.userVisible)
	stateSinkObject.Save(2, &d.Inode)
	stateSinkObject.Save(3, &d.name)
	stateSinkObject.Save(4, &d.parent)
	stateSinkObject.Save(5, &d.deleted)
	stateSinkObject.Save(6, &d.mounted)
}

// +checklocksignore
func (d *Dirent) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.AtomicRefCount)
	stateSourceObject.Load(1, &d.userVisible)
	stateSourceObject.Load(2, &d.Inode)
	stateSourceObject.Load(3, &d.name)
	stateSourceObject.Load(4, &d.parent)
	stateSourceObject.Load(5, &d.deleted)
	stateSourceObject.Load(6, &d.mounted)
	stateSourceObject.LoadValue(7, new(map[string]*Dirent), func(y interface{}) { d.loadChildren(y.(map[string]*Dirent)) })
	stateSourceObject.AfterLoad(d.afterLoad)
}

func (c *DirentCache) StateTypeName() string {
	return "pkg/sentry/fs.DirentCache"
}

func (c *DirentCache) StateFields() []string {
	return []string{
		"maxSize",
		"limit",
	}
}

func (c *DirentCache) beforeSave() {}

// +checklocksignore
func (c *DirentCache) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	if !state.IsZeroValue(&c.currentSize) {
		state.Failf("currentSize is %#v, expected zero", &c.currentSize)
	}
	if !state.IsZeroValue(&c.list) {
		state.Failf("list is %#v, expected zero", &c.list)
	}
	stateSinkObject.Save(0, &c.maxSize)
	stateSinkObject.Save(1, &c.limit)
}

func (c *DirentCache) afterLoad() {}

// +checklocksignore
func (c *DirentCache) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.maxSize)
	stateSourceObject.Load(1, &c.limit)
}

func (d *DirentCacheLimiter) StateTypeName() string {
	return "pkg/sentry/fs.DirentCacheLimiter"
}

func (d *DirentCacheLimiter) StateFields() []string {
	return []string{
		"max",
	}
}

func (d *DirentCacheLimiter) beforeSave() {}

// +checklocksignore
func (d *DirentCacheLimiter) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	if !state.IsZeroValue(&d.count) {
		state.Failf("count is %#v, expected zero", &d.count)
	}
	stateSinkObject.Save(0, &d.max)
}

func (d *DirentCacheLimiter) afterLoad() {}

// +checklocksignore
func (d *DirentCacheLimiter) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.max)
}

func (l *direntList) StateTypeName() string {
	return "pkg/sentry/fs.direntList"
}

func (l *direntList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *direntList) beforeSave() {}

// +checklocksignore
func (l *direntList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *direntList) afterLoad() {}

// +checklocksignore
func (l *direntList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *direntEntry) StateTypeName() string {
	return "pkg/sentry/fs.direntEntry"
}

func (e *direntEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *direntEntry) beforeSave() {}

// +checklocksignore
func (e *direntEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *direntEntry) afterLoad() {}

// +checklocksignore
func (e *direntEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (l *eventList) StateTypeName() string {
	return "pkg/sentry/fs.eventList"
}

func (l *eventList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *eventList) beforeSave() {}

// +checklocksignore
func (l *eventList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *eventList) afterLoad() {}

// +checklocksignore
func (l *eventList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *eventEntry) StateTypeName() string {
	return "pkg/sentry/fs.eventEntry"
}

func (e *eventEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *eventEntry) beforeSave() {}

// +checklocksignore
func (e *eventEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *eventEntry) afterLoad() {}

// +checklocksignore
func (e *eventEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (f *File) StateTypeName() string {
	return "pkg/sentry/fs.File"
}

func (f *File) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"UniqueID",
		"Dirent",
		"flags",
		"async",
		"FileOperations",
		"offset",
	}
}

// +checklocksignore
func (f *File) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.AtomicRefCount)
	stateSinkObject.Save(1, &f.UniqueID)
	stateSinkObject.Save(2, &f.Dirent)
	stateSinkObject.Save(3, &f.flags)
	stateSinkObject.Save(4, &f.async)
	stateSinkObject.Save(5, &f.FileOperations)
	stateSinkObject.Save(6, &f.offset)
}

// +checklocksignore
func (f *File) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.AtomicRefCount)
	stateSourceObject.Load(1, &f.UniqueID)
	stateSourceObject.Load(2, &f.Dirent)
	stateSourceObject.Load(3, &f.flags)
	stateSourceObject.Load(4, &f.async)
	stateSourceObject.LoadWait(5, &f.FileOperations)
	stateSourceObject.Load(6, &f.offset)
	stateSourceObject.AfterLoad(f.afterLoad)
}

func (f *overlayFileOperations) StateTypeName() string {
	return "pkg/sentry/fs.overlayFileOperations"
}

func (f *overlayFileOperations) StateFields() []string {
	return []string{
		"upper",
		"lower",
		"dirCursor",
	}
}

func (f *overlayFileOperations) beforeSave() {}

// +checklocksignore
func (f *overlayFileOperations) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.upper)
	stateSinkObject.Save(1, &f.lower)
	stateSinkObject.Save(2, &f.dirCursor)
}

func (f *overlayFileOperations) afterLoad() {}

// +checklocksignore
func (f *overlayFileOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.upper)
	stateSourceObject.Load(1, &f.lower)
	stateSourceObject.Load(2, &f.dirCursor)
}

func (omi *overlayMappingIdentity) StateTypeName() string {
	return "pkg/sentry/fs.overlayMappingIdentity"
}

func (omi *overlayMappingIdentity) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"id",
		"overlayFile",
	}
}

func (omi *overlayMappingIdentity) beforeSave() {}

// +checklocksignore
func (omi *overlayMappingIdentity) StateSave(stateSinkObject state.Sink) {
	omi.beforeSave()
	stateSinkObject.Save(0, &omi.AtomicRefCount)
	stateSinkObject.Save(1, &omi.id)
	stateSinkObject.Save(2, &omi.overlayFile)
}

func (omi *overlayMappingIdentity) afterLoad() {}

// +checklocksignore
func (omi *overlayMappingIdentity) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &omi.AtomicRefCount)
	stateSourceObject.Load(1, &omi.id)
	stateSourceObject.Load(2, &omi.overlayFile)
}

func (m *MountSourceFlags) StateTypeName() string {
	return "pkg/sentry/fs.MountSourceFlags"
}

func (m *MountSourceFlags) StateFields() []string {
	return []string{
		"ReadOnly",
		"NoAtime",
		"ForcePageCache",
		"NoExec",
	}
}

func (m *MountSourceFlags) beforeSave() {}

// +checklocksignore
func (m *MountSourceFlags) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.ReadOnly)
	stateSinkObject.Save(1, &m.NoAtime)
	stateSinkObject.Save(2, &m.ForcePageCache)
	stateSinkObject.Save(3, &m.NoExec)
}

func (m *MountSourceFlags) afterLoad() {}

// +checklocksignore
func (m *MountSourceFlags) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.ReadOnly)
	stateSourceObject.Load(1, &m.NoAtime)
	stateSourceObject.Load(2, &m.ForcePageCache)
	stateSourceObject.Load(3, &m.NoExec)
}

func (f *FileFlags) StateTypeName() string {
	return "pkg/sentry/fs.FileFlags"
}

func (f *FileFlags) StateFields() []string {
	return []string{
		"Direct",
		"NonBlocking",
		"DSync",
		"Sync",
		"Append",
		"Read",
		"Write",
		"Pread",
		"Pwrite",
		"Directory",
		"Async",
		"LargeFile",
		"NonSeekable",
		"Truncate",
	}
}

func (f *FileFlags) beforeSave() {}

// +checklocksignore
func (f *FileFlags) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.Direct)
	stateSinkObject.Save(1, &f.NonBlocking)
	stateSinkObject.Save(2, &f.DSync)
	stateSinkObject.Save(3, &f.Sync)
	stateSinkObject.Save(4, &f.Append)
	stateSinkObject.Save(5, &f.Read)
	stateSinkObject.Save(6, &f.Write)
	stateSinkObject.Save(7, &f.Pread)
	stateSinkObject.Save(8, &f.Pwrite)
	stateSinkObject.Save(9, &f.Directory)
	stateSinkObject.Save(10, &f.Async)
	stateSinkObject.Save(11, &f.LargeFile)
	stateSinkObject.Save(12, &f.NonSeekable)
	stateSinkObject.Save(13, &f.Truncate)
}

func (f *FileFlags) afterLoad() {}

// +checklocksignore
func (f *FileFlags) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.Direct)
	stateSourceObject.Load(1, &f.NonBlocking)
	stateSourceObject.Load(2, &f.DSync)
	stateSourceObject.Load(3, &f.Sync)
	stateSourceObject.Load(4, &f.Append)
	stateSourceObject.Load(5, &f.Read)
	stateSourceObject.Load(6, &f.Write)
	stateSourceObject.Load(7, &f.Pread)
	stateSourceObject.Load(8, &f.Pwrite)
	stateSourceObject.Load(9, &f.Directory)
	stateSourceObject.Load(10, &f.Async)
	stateSourceObject.Load(11, &f.LargeFile)
	stateSourceObject.Load(12, &f.NonSeekable)
	stateSourceObject.Load(13, &f.Truncate)
}

func (i *Inode) StateTypeName() string {
	return "pkg/sentry/fs.Inode"
}

func (i *Inode) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"InodeOperations",
		"StableAttr",
		"LockCtx",
		"Watches",
		"MountSource",
		"overlay",
	}
}

func (i *Inode) beforeSave() {}

// +checklocksignore
func (i *Inode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.AtomicRefCount)
	stateSinkObject.Save(1, &i.InodeOperations)
	stateSinkObject.Save(2, &i.StableAttr)
	stateSinkObject.Save(3, &i.LockCtx)
	stateSinkObject.Save(4, &i.Watches)
	stateSinkObject.Save(5, &i.MountSource)
	stateSinkObject.Save(6, &i.overlay)
}

func (i *Inode) afterLoad() {}

// +checklocksignore
func (i *Inode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.AtomicRefCount)
	stateSourceObject.Load(1, &i.InodeOperations)
	stateSourceObject.Load(2, &i.StableAttr)
	stateSourceObject.Load(3, &i.LockCtx)
	stateSourceObject.Load(4, &i.Watches)
	stateSourceObject.Load(5, &i.MountSource)
	stateSourceObject.Load(6, &i.overlay)
}

func (l *LockCtx) StateTypeName() string {
	return "pkg/sentry/fs.LockCtx"
}

func (l *LockCtx) StateFields() []string {
	return []string{
		"Posix",
		"BSD",
	}
}

func (l *LockCtx) beforeSave() {}

// +checklocksignore
func (l *LockCtx) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.Posix)
	stateSinkObject.Save(1, &l.BSD)
}

func (l *LockCtx) afterLoad() {}

// +checklocksignore
func (l *LockCtx) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.Posix)
	stateSourceObject.Load(1, &l.BSD)
}

func (w *Watches) StateTypeName() string {
	return "pkg/sentry/fs.Watches"
}

func (w *Watches) StateFields() []string {
	return []string{
		"ws",
		"unlinked",
	}
}

func (w *Watches) beforeSave() {}

// +checklocksignore
func (w *Watches) StateSave(stateSinkObject state.Sink) {
	w.beforeSave()
	stateSinkObject.Save(0, &w.ws)
	stateSinkObject.Save(1, &w.unlinked)
}

func (w *Watches) afterLoad() {}

// +checklocksignore
func (w *Watches) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &w.ws)
	stateSourceObject.Load(1, &w.unlinked)
}

func (i *Inotify) StateTypeName() string {
	return "pkg/sentry/fs.Inotify"
}

func (i *Inotify) StateFields() []string {
	return []string{
		"id",
		"Queue",
		"events",
		"scratch",
		"nextWatch",
		"watches",
	}
}

func (i *Inotify) beforeSave() {}

// +checklocksignore
func (i *Inotify) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.id)
	stateSinkObject.Save(1, &i.Queue)
	stateSinkObject.Save(2, &i.events)
	stateSinkObject.Save(3, &i.scratch)
	stateSinkObject.Save(4, &i.nextWatch)
	stateSinkObject.Save(5, &i.watches)
}

func (i *Inotify) afterLoad() {}

// +checklocksignore
func (i *Inotify) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.id)
	stateSourceObject.Load(1, &i.Queue)
	stateSourceObject.Load(2, &i.events)
	stateSourceObject.Load(3, &i.scratch)
	stateSourceObject.Load(4, &i.nextWatch)
	stateSourceObject.Load(5, &i.watches)
}

func (e *Event) StateTypeName() string {
	return "pkg/sentry/fs.Event"
}

func (e *Event) StateFields() []string {
	return []string{
		"eventEntry",
		"wd",
		"mask",
		"cookie",
		"len",
		"name",
	}
}

func (e *Event) beforeSave() {}

// +checklocksignore
func (e *Event) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.eventEntry)
	stateSinkObject.Save(1, &e.wd)
	stateSinkObject.Save(2, &e.mask)
	stateSinkObject.Save(3, &e.cookie)
	stateSinkObject.Save(4, &e.len)
	stateSinkObject.Save(5, &e.name)
}

func (e *Event) afterLoad() {}

// +checklocksignore
func (e *Event) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.eventEntry)
	stateSourceObject.Load(1, &e.wd)
	stateSourceObject.Load(2, &e.mask)
	stateSourceObject.Load(3, &e.cookie)
	stateSourceObject.Load(4, &e.len)
	stateSourceObject.Load(5, &e.name)
}

func (w *Watch) StateTypeName() string {
	return "pkg/sentry/fs.Watch"
}

func (w *Watch) StateFields() []string {
	return []string{
		"owner",
		"wd",
		"target",
		"unpinned",
		"mask",
		"pins",
	}
}

func (w *Watch) beforeSave() {}

// +checklocksignore
func (w *Watch) StateSave(stateSinkObject state.Sink) {
	w.beforeSave()
	stateSinkObject.Save(0, &w.owner)
	stateSinkObject.Save(1, &w.wd)
	stateSinkObject.Save(2, &w.target)
	stateSinkObject.Save(3, &w.unpinned)
	stateSinkObject.Save(4, &w.mask)
	stateSinkObject.Save(5, &w.pins)
}

func (w *Watch) afterLoad() {}

// +checklocksignore
func (w *Watch) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &w.owner)
	stateSourceObject.Load(1, &w.wd)
	stateSourceObject.Load(2, &w.target)
	stateSourceObject.Load(3, &w.unpinned)
	stateSourceObject.Load(4, &w.mask)
	stateSourceObject.Load(5, &w.pins)
}

func (msrc *MountSource) StateTypeName() string {
	return "pkg/sentry/fs.MountSource"
}

func (msrc *MountSource) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"MountSourceOperations",
		"FilesystemType",
		"Flags",
		"fscache",
		"direntRefs",
	}
}

func (msrc *MountSource) beforeSave() {}

// +checklocksignore
func (msrc *MountSource) StateSave(stateSinkObject state.Sink) {
	msrc.beforeSave()
	stateSinkObject.Save(0, &msrc.AtomicRefCount)
	stateSinkObject.Save(1, &msrc.MountSourceOperations)
	stateSinkObject.Save(2, &msrc.FilesystemType)
	stateSinkObject.Save(3, &msrc.Flags)
	stateSinkObject.Save(4, &msrc.fscache)
	stateSinkObject.Save(5, &msrc.direntRefs)
}

func (msrc *MountSource) afterLoad() {}

// +checklocksignore
func (msrc *MountSource) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &msrc.AtomicRefCount)
	stateSourceObject.Load(1, &msrc.MountSourceOperations)
	stateSourceObject.Load(2, &msrc.FilesystemType)
	stateSourceObject.Load(3, &msrc.Flags)
	stateSourceObject.Load(4, &msrc.fscache)
	stateSourceObject.Load(5, &msrc.direntRefs)
}

func (smo *SimpleMountSourceOperations) StateTypeName() string {
	return "pkg/sentry/fs.SimpleMountSourceOperations"
}

func (smo *SimpleMountSourceOperations) StateFields() []string {
	return []string{
		"keep",
		"revalidate",
		"cacheReaddir",
	}
}

func (smo *SimpleMountSourceOperations) beforeSave() {}

// +checklocksignore
func (smo *SimpleMountSourceOperations) StateSave(stateSinkObject state.Sink) {
	smo.beforeSave()
	stateSinkObject.Save(0, &smo.keep)
	stateSinkObject.Save(1, &smo.revalidate)
	stateSinkObject.Save(2, &smo.cacheReaddir)
}

func (smo *SimpleMountSourceOperations) afterLoad() {}

// +checklocksignore
func (smo *SimpleMountSourceOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &smo.keep)
	stateSourceObject.Load(1, &smo.revalidate)
	stateSourceObject.Load(2, &smo.cacheReaddir)
}

func (o *overlayMountSourceOperations) StateTypeName() string {
	return "pkg/sentry/fs.overlayMountSourceOperations"
}

func (o *overlayMountSourceOperations) StateFields() []string {
	return []string{
		"upper",
		"lower",
	}
}

func (o *overlayMountSourceOperations) beforeSave() {}

// +checklocksignore
func (o *overlayMountSourceOperations) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.upper)
	stateSinkObject.Save(1, &o.lower)
}

func (o *overlayMountSourceOperations) afterLoad() {}

// +checklocksignore
func (o *overlayMountSourceOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.upper)
	stateSourceObject.Load(1, &o.lower)
}

func (ofs *overlayFilesystem) StateTypeName() string {
	return "pkg/sentry/fs.overlayFilesystem"
}

func (ofs *overlayFilesystem) StateFields() []string {
	return []string{}
}

func (ofs *overlayFilesystem) beforeSave() {}

// +checklocksignore
func (ofs *overlayFilesystem) StateSave(stateSinkObject state.Sink) {
	ofs.beforeSave()
}

func (ofs *overlayFilesystem) afterLoad() {}

// +checklocksignore
func (ofs *overlayFilesystem) StateLoad(stateSourceObject state.Source) {
}

func (m *Mount) StateTypeName() string {
	return "pkg/sentry/fs.Mount"
}

func (m *Mount) StateFields() []string {
	return []string{
		"ID",
		"ParentID",
		"root",
		"previous",
	}
}

func (m *Mount) beforeSave() {}

// +checklocksignore
func (m *Mount) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.ID)
	stateSinkObject.Save(1, &m.ParentID)
	stateSinkObject.Save(2, &m.root)
	stateSinkObject.Save(3, &m.previous)
}

func (m *Mount) afterLoad() {}

// +checklocksignore
func (m *Mount) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.ID)
	stateSourceObject.Load(1, &m.ParentID)
	stateSourceObject.Load(2, &m.root)
	stateSourceObject.Load(3, &m.previous)
}

func (mns *MountNamespace) StateTypeName() string {
	return "pkg/sentry/fs.MountNamespace"
}

func (mns *MountNamespace) StateFields() []string {
	return []string{
		"AtomicRefCount",
		"userns",
		"root",
		"mounts",
		"mountID",
	}
}

func (mns *MountNamespace) beforeSave() {}

// +checklocksignore
func (mns *MountNamespace) StateSave(stateSinkObject state.Sink) {
	mns.beforeSave()
	stateSinkObject.Save(0, &mns.AtomicRefCount)
	stateSinkObject.Save(1, &mns.userns)
	stateSinkObject.Save(2, &mns.root)
	stateSinkObject.Save(3, &mns.mounts)
	stateSinkObject.Save(4, &mns.mountID)
}

func (mns *MountNamespace) afterLoad() {}

// +checklocksignore
func (mns *MountNamespace) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mns.AtomicRefCount)
	stateSourceObject.Load(1, &mns.userns)
	stateSourceObject.Load(2, &mns.root)
	stateSourceObject.Load(3, &mns.mounts)
	stateSourceObject.Load(4, &mns.mountID)
}

func (o *overlayEntry) StateTypeName() string {
	return "pkg/sentry/fs.overlayEntry"
}

func (o *overlayEntry) StateFields() []string {
	return []string{
		"lowerExists",
		"lower",
		"mappings",
		"upper",
		"dirCache",
	}
}

func (o *overlayEntry) beforeSave() {}

// +checklocksignore
func (o *overlayEntry) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.lowerExists)
	stateSinkObject.Save(1, &o.lower)
	stateSinkObject.Save(2, &o.mappings)
	stateSinkObject.Save(3, &o.upper)
	stateSinkObject.Save(4, &o.dirCache)
}

func (o *overlayEntry) afterLoad() {}

// +checklocksignore
func (o *overlayEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.lowerExists)
	stateSourceObject.Load(1, &o.lower)
	stateSourceObject.Load(2, &o.mappings)
	stateSourceObject.Load(3, &o.upper)
	stateSourceObject.Load(4, &o.dirCache)
}

func init() {
	state.Register((*StableAttr)(nil))
	state.Register((*UnstableAttr)(nil))
	state.Register((*AttrMask)(nil))
	state.Register((*PermMask)(nil))
	state.Register((*FilePermissions)(nil))
	state.Register((*FileOwner)(nil))
	state.Register((*DentAttr)(nil))
	state.Register((*SortedDentryMap)(nil))
	state.Register((*Dirent)(nil))
	state.Register((*DirentCache)(nil))
	state.Register((*DirentCacheLimiter)(nil))
	state.Register((*direntList)(nil))
	state.Register((*direntEntry)(nil))
	state.Register((*eventList)(nil))
	state.Register((*eventEntry)(nil))
	state.Register((*File)(nil))
	state.Register((*overlayFileOperations)(nil))
	state.Register((*overlayMappingIdentity)(nil))
	state.Register((*MountSourceFlags)(nil))
	state.Register((*FileFlags)(nil))
	state.Register((*Inode)(nil))
	state.Register((*LockCtx)(nil))
	state.Register((*Watches)(nil))
	state.Register((*Inotify)(nil))
	state.Register((*Event)(nil))
	state.Register((*Watch)(nil))
	state.Register((*MountSource)(nil))
	state.Register((*SimpleMountSourceOperations)(nil))
	state.Register((*overlayMountSourceOperations)(nil))
	state.Register((*overlayFilesystem)(nil))
	state.Register((*Mount)(nil))
	state.Register((*MountNamespace)(nil))
	state.Register((*overlayEntry)(nil))
}
