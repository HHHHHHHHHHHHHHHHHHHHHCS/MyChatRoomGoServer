package G

import (
	"container/list"
	"fmt"
)

type charMap struct {
	theCharMap *list.List
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
	l:=list.New()
	for i:='a';i<='z';i++ {
		l.PushBack(i)
		fmt.Print(l)
	}
	for i:='A';i<='Z';i++ {
		l.PushBack(i)
		fmt.Print(l)
	}

	for i:='0';i<='9';i++ {
		l.PushBack(i)
		fmt.Print(l)
	}

	return _charMap
}

func (p *charMap) checkChar(ch byte) bool {
	return false
}

func (p *charMap) checkString(str string) bool {
	return false
}
