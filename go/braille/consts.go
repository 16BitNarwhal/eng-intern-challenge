package braille

var BrailleToEnglish = map[string]string{
	"O.....": "a",
	"O.O...": "b",
	"OO....": "c",
	"OO.O..": "d",
	"O..O..": "e",
	"OOO...": "f",
	"OOOO..": "g",
	"O.OO..": "h",
	".O.O..": "i",
	".OOO..": "j",
	"O...O.": "k",
	"O.O.O.": "l",
	"OO..O.": "m",
	"OO.OO.": "n",
	"O..OO.": "o",
	"OOO.O.": "p",
	"OOOOO.": "q",
	"O.OOO.": "r",
	".OO.O.": "s",
	".OOOO.": "t",
	"O...OO": "u",
	"O.O.OO": "v",
	".OOO.O": "w",
	"OO..OO": "x",
	"OO.OOO": "y",
	"O..OOO": "z",
	"......": " ",
	".....O": "capital",
	".O.OOO": "number",
}

var BrailleToNumber = map[string]string{
	"O.....": "1",
	"O.O...": "2",
	"OO....": "3",
	"OO.O..": "4",
	"O..O..": "5",
	"OOO...": "6",
	"OOOO..": "7",
	"O.OO..": "8",
	".OO...": "9",
	".OOO..": "0",
}

var EnglishToBraille = map[string]string{}
var NumberToBraille = map[string]string{}

func init() {
	for braille, english := range BrailleToEnglish {
		EnglishToBraille[english] = braille
	}
	for braille, number := range BrailleToNumber {
		NumberToBraille[number] = braille
	}
}
