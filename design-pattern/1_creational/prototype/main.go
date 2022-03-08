package main

import "fmt"

// 原型是一种创建型设计模式，使你能够【复制对象，甚至是复杂对象】，而又【无需使代码依赖】它们所属的类。

// 所有的原型类【都必须有一个通用的接口】，使得【即使在对象所属的具体类未知】的情况下【也能复制对象】。
// 原型对象可以生成自身的完整副本，因为相同类的对象可以相互访问对方的私有成员变量。

// Todo ??? 好像没什么用
func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		children: []inode{file1},
		name:     "Folder1",
	}

	folder2 := &folder{
		children: []inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}

// 原型接口
type inode interface {
	print(string)
	clone() inode
}

// 具体原型
type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}

// 具体原型
type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
