package designpattern

import "fmt"

// Sample 1: Computer
type Computer interface {
	Use()
}

type HaseeComputer struct{}

func (h HaseeComputer) Use() {
	fmt.Println("Use HaseeComputer")
}

type LenovoComputer struct{}

func (l LenovoComputer) Use() {
	fmt.Println("Use LenovoComputer")
}

// 使用NewComputer就可以实现简单工厂的功能，没有单独创建一个简单工厂类
// 用户通过传入computerBrand参数，从而选择对应的创建对象
func NewComputer(computerBrand string) Computer {
	if computerBrand == "" {
		return nil
	}
	if computerBrand == "HaseeComputer" {
		return HaseeComputer{}
	} else if computerBrand == "LenovoComputer" {
		return LenovoComputer{}
	}
	return nil
}

// Sample2 : Translator
type Translator interface {
	Translate(string) string
}

type GermanTranslator struct{}

func (g GermanTranslator) Translate(words string) string {
	return "words in German"
}

type EnglishTranslator struct{}

func (e EnglishTranslator) Translate(words string) string {
	return "words in English"
}

type JanpaneseTranslator struct{}

func (j JanpaneseTranslator) Translate(words string) string {
	return "words in Janpanese"
}

func NewTranslator(language string) Translator {
	if language == "" {
		return nil
	}
	var translator Translator
	switch language {
	case "German":
		translator = GermanTranslator{}
	case "English":
		translator = EnglishTranslator{}
	case "Janpanese":
		translator = JanpaneseTranslator{}
	}
	return translator
}
