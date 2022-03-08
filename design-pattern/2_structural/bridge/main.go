package bridge

import "fmt"

// 桥接是一种结构型设计模式，可将业务逻辑或一个大类拆分为不同的层次结构，从而能独立地进行开发。

// 层次结构中的第一层 （通常称为抽象部分） 将包含对第二层 （实现部分） 对象的引用。
// 抽象部分将能将一些 （有时是绝大部分） 对自己的调用委派给实现部分的对象。
// 所有的实现部分都有一个通用接口， 因此它们能在抽象部分内部相互替换。
func main() {

	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}

	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()

	winComputer := &windows{}

	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	fmt.Println()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()
}

//  抽象
type computer interface {
	print()
	setPrinter(printer)
}

// 精确抽象
type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

// 精确抽象
type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

// 实施
type printer interface {
	printFile()
}

// 具体实施
type epson struct {
}

func (p *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// 具体实施
type hp struct {
}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}
