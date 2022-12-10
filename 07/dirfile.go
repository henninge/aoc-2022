package main

import "fmt"

const (
	FileType = 1
	DirType  = 2
)

type File struct {
	Name string
	Size int
}

type Entry interface {
	GetSize() int
	GetType() int
	Print(depth int)
}

type Dir struct {
	Name    string
	Parent  *Dir
	Depth   int
	Entries map[string]Entry
}

func (f File) GetSize() int {
	return f.Size
}

func (f File) GetType() int {
	return FileType
}

func (f File) Print(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("- %s (file, size=%d)\n", f.Name, f.Size)
}

func (d Dir) GetSize() (sizeSum int) {
	for _, entry := range d.Entries {
		sizeSum += entry.GetSize()
	}
	return
}

func (d Dir) GetType() int {
	return DirType
}

func (d Dir) Print(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("- %s (dir)\n", d.Name)
	for _, entry := range d.Entries {
		entry.Print(d.Depth + 1)
	}
}

func (d *Dir) Walk(collector func(dir *Dir)) {
	collector(d)
	for _, entry := range d.Entries {
		if entry.GetType() == DirType {
			entry.(*Dir).Walk(collector)
		}
	}
}

func NewDir(name string, parent *Dir) *Dir {
	newdir := &Dir{
		Name:    name,
		Parent:  parent,
		Entries: make(map[string]Entry),
	}
	if parent != nil {
		newdir.Depth = parent.Depth + 1
	}
	return newdir
}
