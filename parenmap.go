package main

// closerOf maps a bracket/parenthesis/quotation/... opener to a closer.
//
// It is generated automatically using makeparens.py, with a few additions
// listed to the end of the map.
//
// Low-9 quotation marks are left out, since their matchers depends on the
// language. For instance, '‚' (U+201A) matches with '‘' (U+2018) in German, but
// with '’' (U+2019) in Polish.
var closerOf = map[rune]rune{
	'(': ')', // U+28 U+29
	'[': ']', // U+5B U+5D
	'{': '}', // U+7B U+7D
	'«': '»', // U+AB U+BB
	'༺': '༻', // U+F3A U+F3B
	'༼': '༽', // U+F3C U+F3D
	'᚛': '᚜', // U+169B U+169C
	'‘': '’', // U+2018 U+2019
	'“': '”', // U+201C U+201D
	'‹': '›', // U+2039 U+203A
	'⁅': '⁆', // U+2045 U+2046
	'⁽': '⁾', // U+207D U+207E
	'₍': '₎', // U+208D U+208E
	'⌈': '⌉', // U+2308 U+2309
	'⌊': '⌋', // U+230A U+230B
	'〈': '〉', // U+2329 U+232A
	'❨': '❩', // U+2768 U+2769
	'❪': '❫', // U+276A U+276B
	'❬': '❭', // U+276C U+276D
	'❮': '❯', // U+276E U+276F
	'❰': '❱', // U+2770 U+2771
	'❲': '❳', // U+2772 U+2773
	'❴': '❵', // U+2774 U+2775
	'⟅': '⟆', // U+27C5 U+27C6
	'⟦': '⟧', // U+27E6 U+27E7
	'⟨': '⟩', // U+27E8 U+27E9
	'⟪': '⟫', // U+27EA U+27EB
	'⟬': '⟭', // U+27EC U+27ED
	'⟮': '⟯', // U+27EE U+27EF
	'⦃': '⦄', // U+2983 U+2984
	'⦅': '⦆', // U+2985 U+2986
	'⦇': '⦈', // U+2987 U+2988
	'⦉': '⦊', // U+2989 U+298A
	'⦋': '⦌', // U+298B U+298C
	'⦍': '⦎', // U+298D U+298E
	'⦏': '⦐', // U+298F U+2990
	'⦑': '⦒', // U+2991 U+2992
	'⦓': '⦔', // U+2993 U+2994
	'⦕': '⦖', // U+2995 U+2996
	'⦗': '⦘', // U+2997 U+2998
	'⧘': '⧙', // U+29D8 U+29D9
	'⧚': '⧛', // U+29DA U+29DB
	'⧼': '⧽', // U+29FC U+29FD
	'⸂': '⸃', // U+2E02 U+2E03
	'⸄': '⸅', // U+2E04 U+2E05
	'⸉': '⸊', // U+2E09 U+2E0A
	'⸌': '⸍', // U+2E0C U+2E0D
	'⸜': '⸝', // U+2E1C U+2E1D
	'⸠': '⸡', // U+2E20 U+2E21
	'⸢': '⸣', // U+2E22 U+2E23
	'⸤': '⸥', // U+2E24 U+2E25
	'⸦': '⸧', // U+2E26 U+2E27
	'⸨': '⸩', // U+2E28 U+2E29
	'〈': '〉', // U+3008 U+3009
	'《': '》', // U+300A U+300B
	'「': '」', // U+300C U+300D
	'『': '』', // U+300E U+300F
	'【': '】', // U+3010 U+3011
	'〔': '〕', // U+3014 U+3015
	'〖': '〗', // U+3016 U+3017
	'〘': '〙', // U+3018 U+3019
	'〚': '〛', // U+301A U+301B
	'〝': '〞', // U+301D U+301E
	'︗': '︘', // U+FE17 U+FE18
	'︵': '︶', // U+FE35 U+FE36
	'︷': '︸', // U+FE37 U+FE38
	'︹': '︺', // U+FE39 U+FE3A
	'︻': '︼', // U+FE3B U+FE3C
	'︽': '︾', // U+FE3D U+FE3E
	'︿': '﹀', // U+FE3F U+FE40
	'﹁': '﹂', // U+FE41 U+FE42
	'﹃': '﹄', // U+FE43 U+FE44
	'﹇': '﹈', // U+FE47 U+FE48
	'﹙': '﹚', // U+FE59 U+FE5A
	'﹛': '﹜', // U+FE5B U+FE5C
	'﹝': '﹞', // U+FE5D U+FE5E
	'（': '）', // U+FF08 U+FF09
	'［': '］', // U+FF3B U+FF3D
	'｛': '｝', // U+FF5B U+FF5D
	'｟': '｠', // U+FF5F U+FF60
	'｢': '｣', // U+FF62 U+FF63

	// Additions:
	'⎴': '⎵', // U+23B4 U+23B5 (has category Sm instead of Pi/Pf)
	'⏜': '⏝', // U+23DC U+23DD (ditto)
	'⏞': '⏟', // U+23DE U+23DF (ditto)
	'⏠': '⏡', // U+23E0 U+23E1 (ditto)
	'❛': '❜', // U+275B U+275C (has category So instead of Pi/Pf)
	'❝': '❞', // U+275D U+275E (ditto)
	'﴿': '﴾', // U+FD3F U+FD3E (closer appears before opener)
}

var openerOf = map[rune]rune{}

func init() {
	for k, v := range closerOf {
		openerOf[v] = k
	}
}
