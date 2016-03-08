package sgf

import (
	"fmt"
	"strconv"
)

type SGFElement struct {
	parent *SGFElement
	child  *SGFElement
	data   map[string]string
}

func (e *SGFElement) String() string {
	return fmt.Sprintf("{parent: %p, data: %v}", e.parent, e.data)
}

func (e *SGFElement) Size() int {
	size, _ := strconv.Atoi(e.data["SZ"])
	return size
}

func (e *SGFElement) Title() string {
	return e.data["EV"]
}

func (e *SGFElement) Komi() float64 {
	n, _ := strconv.ParseFloat(e.data["KM"], 64)
	return n
}

func (e *SGFElement) Result() string {
	return e.data["RE"]
}

func (e *SGFElement) PlayerBlack() string {
	return e.data["PB"]
}

func (e *SGFElement) PlayerWhite() string {
	return e.data["PW"]
}

func (e *SGFElement) DateTime() string {
	return e.data["DT"]
}

func (e *SGFElement) Place() string {
	return e.data["PC"]
}

func (e *SGFElement) PrintTraverse() {
	elm := e
	for {
		fmt.Println(elm)
		if elm.child == nil {
			break
		}
		elm = elm.child
	}
}

func newSGFElement(e *SGFElement) *SGFElement {
	elm := &SGFElement{parent: e, data: make(map[string]string), child: nil}
	if e != nil {
		e.child = elm
	}
	return elm
}
