package G

type charMap struct {
	theCharMap []int
}

var _charMap *charMap

func CharMap() *charMap {
	if _charMap == nil {
		_charMap = newCharMap()

	}
	return _charMap
}

func newCharMap() *charMap {
	_charMap = &charMap{}
	return _charMap
}

func (p *charMap) checkChar(ch byte) bool {
	return false
}

func (p *charMap) checkString(str string) bool {
	return false
}
