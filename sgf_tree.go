package sgf

type SGFTree struct {
	root *SGFElement
}

func SGFTreeParse(str string) (*SGFTree, error) {
	const (
		stateKey = iota
		stateValue
	)
	state := stateKey
	key, value := "", ""
	var elm *SGFElement = nil
	var root *SGFElement = nil
	for _, x := range str {
		switch string(x) {
		case "(":
		case ")":
		case "\n":
		case ";":
			elm = newSGFElement(elm)
			if root == nil {
				root = elm
			}
		case "[":
			state = stateValue
			value = ""
		case "]":
			elm.data[key] = value
			state = stateKey
			key = ""
		default:
			if state == stateKey {
				key += string(x)
			} else {
				value += string(x)
			}
		}
	}
	// root.PrintTraverse()
	return &SGFTree{root: root}, nil
}

func (t *SGFTree) Size() int {
	return t.root.Size()
}

func (t *SGFTree) Title() string {
	return t.root.Title()
}

func (t *SGFTree) Komi() float64 {
	return t.root.Komi()
}

func (t *SGFTree) Result() string {
	return t.root.Result()
}

func (t *SGFTree) PlayerBlack() string {
	return t.root.PlayerBlack()
}

func (t *SGFTree) PlayerWhite() string {
	return t.root.PlayerWhite()
}

func (t *SGFTree) DateTime() string {
	return t.root.DateTime()
}

func (t *SGFTree) Place() string {
	return t.root.Place()
}
