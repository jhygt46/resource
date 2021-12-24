package lang

import (
	"os"
	"fmt"
	"unicode"
    "io/ioutil"
	"encoding/json"
)

var Langs = []string{"Adlam", "Ahom", "Anatolian_Hieroglyphs", "Arabic", "Armenia", "ASCII_Hex_Digit", "Avestan", "Balinese", "Bamum", "Bassa_Vah", "Batak", "Bengali", "Bhaiksuki", "Bidi_Control", "Bopomofo", "Brahmi", "Braille", "Buginese", "Buhid", "C", "Canadian_Aboriginal", "Carian", "Caucasian_Albanian", "Cc", "Cf", "Chakma", "Cham", "Cherokee", "Chorasmian", "Co", "Common", "Coptic", "Cs", "Cuneiform", "Cypriot", "Cyrillic", "Dash", "Deprecated", "Deseret", "Devanagari", "Diacritic", "Digit", "Dives_Akuru", "Dogra", "Duployan", "Egyptian_Hieroglyphs", "Elbasan", "Elymaic", "Ethiopic", "Extender", "Georgian", "Glagolitic", "Gothic", "Grantha", "Greek", "Gujarati", "Gunjala_Gondi", "Gurmukhi", "Han", "Hangul", "Hanifi_Rohingya", "Hanunoo", "Hatran", "Hebrew", "Hex_Digit", "Hiragana", "Hyphen", "Ideographic", "IDS_Binary_Operator", "IDS_Trinary_Operator", "Imperial_Aramaic", "Inherited", "Inscriptional_Pahlavi", "Inscriptional_Parthian", "Javanese", "Join_Control", "Kaithi", "Kannada", "Katakana", "Kayah_Li", "Kharoshthi", "Khitan_Small_Script", "Khmer", "Khojki", "Khudawadi", "L", "Lao", "Latin", "Lepcha", "Letter", "Limbu", "Linear_A", "Linear_B", "Lisu", "Ll", "Lm", "Lo", "Logical_Order_Exception", "Lower", "Lt", "Lu", "Lycian", "Lydian", "M", "Mahajani", "Makasar", "Malayalam", "Mandaic", "Manichaean", "Marchen", "Mark", "Masaram_Gondi", "Mc", "Me", "Medefaidrin", "Meetei_Mayek", "Mende_Kikakui", "Meroitic_Cursive", "Meroitic_Hieroglyphs", "Miao", "Mn", "Modi", "Mongolian", "Mro", "Multani", "Myanmar", "N", "Nabataean", "Nandinagari", "Nd", "New_Tai_Lue", "Newa", "Nko", "Nl", "No", "Noncharacter_Code_Point", "Number", "Nushu", "Nyiakeng_Puachue_Hmong", "Ogham", "Ol_Chiki", "Old_Hungarian", "Old_Italic", "Old_North_Arabian", "Old_Permic", "Old_Persian", "Old_Sogdian", "Old_South_Arabian", "Old_Turkic", "Oriya", "Osage", "Osmanya", "Other", "Other_Alphabetic", "Other_Default_Ignorable_Code_Point", "Other_Grapheme_Extend", "Other_ID_Continue", "Other_ID_Start", "Other_Lowercase", "Other_Math", "Other_Uppercase", "P", "Pahawh_Hmong", "Palmyrene", "Pattern_Syntax", "Pattern_White_Space", "Pau_Cin_Hau", "Pc", "Pd", "Pe", "Pf", "Phags_Pa", "Phoenician", "Pi", "Po", "Prepended_Concatenation_Mark", "Ps", "Psalter_Pahlavi", "Punct", "Quotation_Mark", "Radical", "Regional_Indicator", "Rejang", "Runic", "S", "Samaritan", "Saurashtra", "Sc", "Sentence_Terminal", "Sharada", "Shavian", "Siddham", "SignWriting", "Sinhala", "Sk", "Sm", "So", "Soft_Dotted", "Sogdian", "Sora_Sompeng", "Soyombo", "Space", "STerm", "Sundanese", "Syloti_Nagri", "Symbol", "Syriac", "Tagalog", "Tagbanwa", "Tai_Le", "Tai_Tham", "Tai_Viet", "Takri", "Tamil", "Tangut", "Telugu", "Terminal_Punctuation", "Thaana", "Thai", "Tibetan", "Tifinagh", "Tirhuta", "Title", "Ugaritic", "Unified_Ideograph", "Upper", "Vai", "Variation_Selector", "Wancho", "Warang_Citi", "White_Space", "Yezidi", "Yi", "Z", "Zl", "Zp", "Zs"}

func BuildLang(){

	WriteUnicode("Cc", unicode.Cc)
	WriteUnicode("Cf", unicode.Cf)
	WriteUnicode("Co", unicode.Co)
	WriteUnicode("Cs", unicode.Cs)
	WriteUnicode("Digit", unicode.Digit)
	WriteUnicode("Nd", unicode.Nd)
	WriteUnicode("Letter", unicode.Letter)
	WriteUnicode("L", unicode.L)
	WriteUnicode("Lm", unicode.Lm)
	WriteUnicode("Lo", unicode.Lo)
	WriteUnicode("Lower", unicode.Lower)
	WriteUnicode("Ll", unicode.Ll)
	WriteUnicode("Mark", unicode.Mark)
	WriteUnicode("M", unicode.M)
	WriteUnicode("Mc", unicode.Mc)
	WriteUnicode("Me", unicode.Me)
	WriteUnicode("Mn", unicode.Mn)
	WriteUnicode("Nl", unicode.Nl)
	WriteUnicode("No", unicode.No)
	WriteUnicode("Number", unicode.Number)
	WriteUnicode("N", unicode.N)
	WriteUnicode("Other", unicode.Other)
	WriteUnicode("C", unicode.C)
	WriteUnicode("Pc", unicode.Pc)
	WriteUnicode("Pd", unicode.Pd)
	WriteUnicode("Pe", unicode.Pe)
	WriteUnicode("Pf", unicode.Pf)
	WriteUnicode("Pi", unicode.Pi)
	WriteUnicode("Po", unicode.Po)
	WriteUnicode("Ps", unicode.Ps)
	WriteUnicode("Punct", unicode.Punct)
	WriteUnicode("P", unicode.P)
	WriteUnicode("Sc", unicode.Sc)
	WriteUnicode("Sk", unicode.Sk)
	WriteUnicode("Sm", unicode.Sm)
	WriteUnicode("So", unicode.So)
	WriteUnicode("Space", unicode.Space)
	WriteUnicode("Z", unicode.Z)
	WriteUnicode("Symbol", unicode.Symbol)
	WriteUnicode("S", unicode.S)
	WriteUnicode("Title", unicode.Title)
	WriteUnicode("Lt", unicode.Lt)
	WriteUnicode("Upper", unicode.Upper)
	WriteUnicode("Lu", unicode.Lu)
	WriteUnicode("Zl", unicode.Zl)
	WriteUnicode("Zp", unicode.Zp)
	WriteUnicode("Zs", unicode.Zs)



	WriteUnicode("ASCII_Hex_Digit", unicode.ASCII_Hex_Digit)
	WriteUnicode("Bidi_Control", unicode.Bidi_Control)
	WriteUnicode("Dash", unicode.Dash)
	WriteUnicode("Deprecated", unicode.Deprecated)
	WriteUnicode("Diacritic", unicode.Diacritic)
	WriteUnicode("Extender", unicode.Extender)
	WriteUnicode("Hex_Digit", unicode.Hex_Digit)
	WriteUnicode("Hyphen", unicode.Hyphen)
	WriteUnicode("IDS_Binary_Operator", unicode.IDS_Binary_Operator)
	WriteUnicode("IDS_Trinary_Operator", unicode.IDS_Trinary_Operator)
	WriteUnicode("Ideographic", unicode.Ideographic)
	WriteUnicode("Join_Control", unicode.Join_Control)
	WriteUnicode("Logical_Order_Exception", unicode.Logical_Order_Exception)
	WriteUnicode("Noncharacter_Code_Point", unicode.Noncharacter_Code_Point)
	WriteUnicode("Other_Alphabetic", unicode.Other_Alphabetic)
	WriteUnicode("Other_Default_Ignorable_Code_Point", unicode.Other_Default_Ignorable_Code_Point)
	WriteUnicode("Other_Grapheme_Extend", unicode.Other_Grapheme_Extend)
	WriteUnicode("Other_ID_Continue", unicode.Other_ID_Continue)
	WriteUnicode("Other_ID_Start", unicode.Other_ID_Start)
	WriteUnicode("Other_Lowercase", unicode.Other_Lowercase)
	WriteUnicode("Other_Math", unicode.Other_Math)
	WriteUnicode("Other_Uppercase", unicode.Other_Uppercase)
	WriteUnicode("Pattern_Syntax", unicode.Pattern_Syntax)
	WriteUnicode("Pattern_White_Space", unicode.Pattern_White_Space)
	WriteUnicode("Prepended_Concatenation_Mark", unicode.Prepended_Concatenation_Mark)
	WriteUnicode("Quotation_Mark", unicode.Quotation_Mark)
	WriteUnicode("Radical", unicode.Radical)
	WriteUnicode("Regional_Indicator", unicode.Regional_Indicator)
	WriteUnicode("STerm", unicode.STerm)
	WriteUnicode("Sentence_Terminal", unicode.Sentence_Terminal)
	WriteUnicode("Soft_Dotted", unicode.Soft_Dotted)
	WriteUnicode("Terminal_Punctuation", unicode.Terminal_Punctuation)
	WriteUnicode("Unified_Ideograph", unicode.Unified_Ideograph)
	WriteUnicode("Variation_Selector", unicode.Variation_Selector)
	WriteUnicode("White_Space", unicode.White_Space)



	WriteUnicode("Adlam", unicode.Adlam)
	WriteUnicode("Ahom", unicode.Ahom)
	WriteUnicode("Anatolian_Hieroglyphs", unicode.Anatolian_Hieroglyphs)
	WriteUnicode("Arabic", unicode.Arabic)
	WriteUnicode("Armenian", unicode.Armenian)
	WriteUnicode("Avestan", unicode.Avestan)
	WriteUnicode("Balinese", unicode.Balinese)
	WriteUnicode("Bamum", unicode.Bamum)
	WriteUnicode("Bassa_Vah", unicode.Bassa_Vah)
	WriteUnicode("Batak", unicode.Batak)
	WriteUnicode("Bengali", unicode.Bengali)
	WriteUnicode("Bhaiksuki", unicode.Bhaiksuki)
	WriteUnicode("Bopomofo", unicode.Bopomofo)
	WriteUnicode("Brahmi", unicode.Brahmi)
	WriteUnicode("Braille", unicode.Braille)
	WriteUnicode("Buginese", unicode.Buginese)
	WriteUnicode("Buhid", unicode.Buhid)
	WriteUnicode("Canadian_Aboriginal", unicode.Canadian_Aboriginal)
	WriteUnicode("Carian", unicode.Carian)
	WriteUnicode("Caucasian_Albanian", unicode.Caucasian_Albanian)
	WriteUnicode("Chakma", unicode.Chakma)
	WriteUnicode("Cham", unicode.Cham)
	WriteUnicode("Cherokee", unicode.Cherokee)
	WriteUnicode("Chorasmian", unicode.Chorasmian)
	WriteUnicode("Common", unicode.Common)
	WriteUnicode("Coptic", unicode.Coptic)
	WriteUnicode("Cuneiform", unicode.Cuneiform)
	WriteUnicode("Cypriot", unicode.Cypriot)
	WriteUnicode("Cyrillic", unicode.Cyrillic)
	WriteUnicode("Deseret", unicode.Deseret)
	WriteUnicode("Devanagari", unicode.Devanagari)
	WriteUnicode("Dives_Akuru", unicode.Dives_Akuru)
	WriteUnicode("Dogra", unicode.Dogra)
	WriteUnicode("Duployan", unicode.Duployan)
	WriteUnicode("Egyptian_Hieroglyphs", unicode.Egyptian_Hieroglyphs)
	WriteUnicode("Elbasan", unicode.Elbasan)
	WriteUnicode("Elymaic", unicode.Elymaic)
	WriteUnicode("Ethiopic", unicode.Ethiopic)
	WriteUnicode("Georgian", unicode.Georgian)
	WriteUnicode("Glagolitic", unicode.Glagolitic)
	WriteUnicode("Gothic", unicode.Gothic)
	WriteUnicode("Grantha", unicode.Grantha)
	WriteUnicode("Greek", unicode.Greek)
	WriteUnicode("Gujarati", unicode.Gujarati)
	WriteUnicode("Gunjala_Gondi", unicode.Gunjala_Gondi)
	WriteUnicode("Gurmukhi", unicode.Gurmukhi)
	WriteUnicode("Han", unicode.Han)
	WriteUnicode("Hangul", unicode.Hangul)
	WriteUnicode("Hanifi_Rohingya", unicode.Hanifi_Rohingya)
	WriteUnicode("Hanunoo", unicode.Hanunoo)
	WriteUnicode("Hatran", unicode.Hatran)
	WriteUnicode("Hebrew", unicode.Hebrew)
	WriteUnicode("Hiragana", unicode.Hiragana)
	WriteUnicode("Imperial_Aramaic", unicode.Imperial_Aramaic)
	WriteUnicode("Inherited", unicode.Inherited)
	WriteUnicode("Inscriptional_Pahlavi", unicode.Inscriptional_Pahlavi)
	WriteUnicode("Inscriptional_Parthian", unicode.Inscriptional_Parthian)
	WriteUnicode("Javanese", unicode.Javanese)
	WriteUnicode("Kaithi", unicode.Kaithi)
	WriteUnicode("Kannada", unicode.Kannada)
	WriteUnicode("Katakana", unicode.Katakana)
	WriteUnicode("Kayah_Li", unicode.Kayah_Li)
	WriteUnicode("Kharoshthi", unicode.Kharoshthi)
	WriteUnicode("Khitan_Small_Script", unicode.Khitan_Small_Script)
	WriteUnicode("Khmer", unicode.Khmer)
	WriteUnicode("Khojki", unicode.Khojki)
	WriteUnicode("Khudawadi", unicode.Khudawadi)
	WriteUnicode("Lao", unicode.Lao)
	WriteUnicode("Latin", unicode.Latin)
	WriteUnicode("Lepcha", unicode.Lepcha)
	WriteUnicode("Limbu", unicode.Limbu)
	WriteUnicode("Linear_A", unicode.Linear_A)
	WriteUnicode("Linear_B", unicode.Linear_B)
	WriteUnicode("Lisu", unicode.Lisu)
	WriteUnicode("Lycian", unicode.Lycian)
	WriteUnicode("Lydian", unicode.Lydian)
	WriteUnicode("Mahajani", unicode.Mahajani)
	WriteUnicode("Makasar", unicode.Makasar)
	WriteUnicode("Malayalam", unicode.Malayalam)
	WriteUnicode("Mandaic", unicode.Mandaic)
	WriteUnicode("Manichaean", unicode.Manichaean)
	WriteUnicode("Marchen", unicode.Marchen)
	WriteUnicode("Masaram_Gondi", unicode.Masaram_Gondi)
	WriteUnicode("Medefaidrin", unicode.Medefaidrin)
	WriteUnicode("Meetei_Mayek", unicode.Meetei_Mayek)
	WriteUnicode("Mende_Kikakui", unicode.Mende_Kikakui)
	WriteUnicode("Meroitic_Cursive", unicode.Meroitic_Cursive)
	WriteUnicode("Meroitic_Hieroglyphs", unicode.Meroitic_Hieroglyphs)
	WriteUnicode("Miao", unicode.Miao)
	WriteUnicode("Modi", unicode.Modi)
	WriteUnicode("Mongolian", unicode.Mongolian)
	WriteUnicode("Mro", unicode.Mro)
	WriteUnicode("Multani", unicode.Multani)
	WriteUnicode("Myanmar", unicode.Myanmar)
	WriteUnicode("Nabataean", unicode.Nabataean)
	WriteUnicode("Nandinagari", unicode.Nandinagari)
	WriteUnicode("New_Tai_Lue", unicode.New_Tai_Lue)
	WriteUnicode("Newa", unicode.Newa)
	WriteUnicode("Nko", unicode.Nko)
	WriteUnicode("Nushu", unicode.Nushu)
	WriteUnicode("Nyiakeng_Puachue_Hmong", unicode.Nyiakeng_Puachue_Hmong)
	WriteUnicode("Ogham", unicode.Ogham)
	WriteUnicode("Ol_Chiki", unicode.Ol_Chiki)
	WriteUnicode("Old_Hungarian", unicode.Old_Hungarian)
	WriteUnicode("Old_Italic", unicode.Old_Italic)
	WriteUnicode("Old_North_Arabian", unicode.Old_North_Arabian)
	WriteUnicode("Old_Permic", unicode.Old_Permic)
	WriteUnicode("Old_Persian", unicode.Old_Persian)
	WriteUnicode("Old_Sogdian", unicode.Old_Sogdian)
	WriteUnicode("Old_South_Arabian", unicode.Old_South_Arabian)
	WriteUnicode("Old_Turkic", unicode.Old_Turkic)
	WriteUnicode("Oriya", unicode.Oriya)
	WriteUnicode("Osage", unicode.Osage)
	WriteUnicode("Osmanya", unicode.Osmanya)
	WriteUnicode("Pahawh_Hmong", unicode.Pahawh_Hmong)
	WriteUnicode("Palmyrene", unicode.Palmyrene)
	WriteUnicode("Pau_Cin_Hau", unicode.Pau_Cin_Hau)
	WriteUnicode("Phags_Pa", unicode.Phags_Pa)
	WriteUnicode("Phoenician", unicode.Phoenician)
	WriteUnicode("Psalter_Pahlavi", unicode.Psalter_Pahlavi)
	WriteUnicode("Rejang", unicode.Rejang)
	WriteUnicode("Runic", unicode.Runic)
	WriteUnicode("Samaritan", unicode.Samaritan)
	WriteUnicode("Saurashtra", unicode.Saurashtra)
	WriteUnicode("Sharada", unicode.Sharada)
	WriteUnicode("Shavian", unicode.Shavian)
	WriteUnicode("Siddham", unicode.Siddham)
	WriteUnicode("SignWriting", unicode.SignWriting)
	WriteUnicode("Sinhala", unicode.Sinhala)
	WriteUnicode("Sogdian", unicode.Sogdian)
	WriteUnicode("Sora_Sompeng", unicode.Sora_Sompeng)
	WriteUnicode("Soyombo", unicode.Soyombo)
	WriteUnicode("Sundanese", unicode.Sundanese)
	WriteUnicode("Syloti_Nagri", unicode.Syloti_Nagri)
	WriteUnicode("Syriac", unicode.Syriac)
	WriteUnicode("Tagalog", unicode.Tagalog)
	WriteUnicode("Tagbanwa", unicode.Tagbanwa)
	WriteUnicode("Tai_Le", unicode.Tai_Le)
	WriteUnicode("Tai_Tham", unicode.Tai_Tham)
	WriteUnicode("Tai_Viet", unicode.Tai_Viet)
	WriteUnicode("Takri", unicode.Takri)
	WriteUnicode("Tamil", unicode.Tamil)
	WriteUnicode("Tangut", unicode.Tangut)
	WriteUnicode("Telugu", unicode.Telugu)
	WriteUnicode("Thaana", unicode.Thaana)
	WriteUnicode("Thai", unicode.Thai)
	WriteUnicode("Tibetan", unicode.Tibetan)
	WriteUnicode("Tifinagh", unicode.Tifinagh)
	WriteUnicode("Tirhuta", unicode.Tirhuta)
	WriteUnicode("Ugaritic", unicode.Ugaritic)
	WriteUnicode("Vai", unicode.Vai)
	WriteUnicode("Wancho", unicode.Wancho)
	WriteUnicode("Warang_Citi", unicode.Warang_Citi)
	WriteUnicode("Yezidi", unicode.Yezidi)
	WriteUnicode("Yi", unicode.Yi)

}
func WriteUnicode(name string, uni *unicode.RangeTable){

	total := 1114111
	var r rune
	var runes []rune
	for i:=0; i < total; i++ {

		r = int32(i)
		if unicode.Is(uni, r) { runes = append(runes, r) }
	
	}
	
	file, _ := json.MarshalIndent(runes, "", " ")
	err := ioutil.WriteFile("lang/"+name, file, 0777)
	if err != nil { fmt.Println(err) }

}
func GetUnicodes(langs []string) ([]int32, bool) {
	
	res := []int32{}
	for _, val := range langs {
		arr, is := GetUnicode(val)
		if is {
			res = append(res, arr...)
		}else{
			return nil, false
		}
	}
	return res, true

}
func GetUnicode(uni string) ([]int32, bool) {
	file := "lang/"+uni
	if FileExists("lang/"+file) {
		jsonFiltro, err := os.Open(file)
		if err == nil {
			byteValueFiltro, _ := ioutil.ReadAll(jsonFiltro)
			defer jsonFiltro.Close()
			data := []int32{}
			if err := json.Unmarshal(byteValueFiltro, &data); err == nil {
				return data, true
			}
		}
		return nil, false
	}
	return nil, false
}
func FileExists(name string) bool {
    if fi, err := os.Stat(name); err == nil {
        if fi.Mode().IsRegular() {
            return true
        }
    }
    return false
}
