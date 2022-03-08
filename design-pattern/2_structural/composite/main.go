package main

import "fmt"

// 组合是一种结构型设计模式，你可以使用它将对象组合成树状结构，并且能像使用独立对象一样使用它们。
func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}

// 组件接口
type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) getName() string {
	return f.name
}

// 组合
type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

// 叶子
type component interface {
	search(string)
}
