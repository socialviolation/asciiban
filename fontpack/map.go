// Package fontpack Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2023-04-07 10:52:05.941323 +1000 AEST m=+5.362173805
// using data from https://github.com/xero/figlet-fonts
package fontpack

import (
	"strings"
)

func Get(f string) string {
	if val, ok := FontMap[strings.ToLower(f)]; ok {
		return val
	}
	return ANSIShadow
}

var FontMap = map[string]string{
	"amc3line": AMC3Line,
	"amc3liv1": AMC3Liv1,
	"amcaaa01": AMCAAA01,
	"amcneko": AMCNeko,
	"amcrazor": AMCRazor,
	"amcrazor2": AMCRazor2,
	"amcslash": AMCSlash,
	"amcslider": AMCSlider,
	"amcthin": AMCThin,
	"amctubes": AMCTubes,
	"amcuntitled": AMCUntitled,
	"ansiregular": ANSIRegular,
	"ansishadow": ANSIShadow,
	"acrobatic": Acrobatic,
	"alligator": Alligator,
	"alligator2": Alligator2,
	"alligator3": Alligator3,
	"alpha": Alpha,
	"alphabet": Alphabet,
	"amcrazo2": Amcrazo2,
	"amcslder": Amcslder,
	"amcun1": Amcun1,
	"arrows": Arrows,
	"avatar": Avatar,
	"banner": Banner,
	"banner3": Banner3,
	"banner3d": Banner3D,
	"banner4": Banner4,
	"barbwire": Barbwire,
	"basic": Basic,
	"bell": Bell,
	"benjamin": Benjamin,
	"big": Big,
	"bigchief": BigChief,
	"bigmoneyne": BigMoneyne,
	"bigmoneynw": BigMoneynw,
	"bigmoneyse": BigMoneyse,
	"bigmoneysw": BigMoneysw,
	"bigfig": Bigfig,
	"binary": Binary,
	"block": Block,
	"bloody": Bloody,
	"bolger": Bolger,
	"braced": Braced,
	"bright": Bright,
	"broadway": Broadway,
	"bulbhead": Bulbhead,
	"calgphy2": Calgphy2,
	"caligraphy": Caligraphy,
	"caligraphy2": Caligraphy2,
	"calvins": CalvinS,
	"catwalk": Catwalk,
	"chunky": Chunky,
	"coinstak": Coinstak,
	"colossal": Colossal,
	"computer": Computer,
	"contessa": Contessa,
	"contrast": Contrast,
	"cosmic": Cosmic,
	"cosmike": Cosmike,
	"crawford": Crawford,
	"crawford2": Crawford2,
	"cricket": Cricket,
	"cursive": Cursive,
	"cyberlarge": Cyberlarge,
	"cybermedium": Cybermedium,
	"cybersmall": Cybersmall,
	"cygnet": Cygnet,
	"danc4": DANC4,
	"dosrebel": DOSRebel,
	"decimal": Decimal,
	"defleppard": DefLeppard,
	"deltacorpspriest1": DeltaCorpsPriest1,
	"diamond": Diamond,
	"doh": Doh,
	"doom": Doom,
	"dotmatrix": DotMatrix,
	"double": Double,
	"doubleshorts": DoubleShorts,
	"drpepper": DrPepper,
	"eftichess": EftiChess,
	"eftifont": EftiFont,
	"eftiitalic": EftiItalic,
	"eftipiti": EftiPiti,
	"eftirobot": EftiRobot,
	"eftiwall": EftiWall,
	"eftiwater": EftiWater,
	"eftitalic": Eftitalic,
	"electronic": Electronic,
	"elite": Elite,
	"epic": Epic,
	"f1row": F1Row,
	"f3d": F3D,
	"f3dascii": F3DASCII,
	"f3x5": F3x5,
	"f4max": F4Max,
	"f5lineoblique": F5LineOblique,
	"fender": Fender,
	"firefontk": FireFontk,
	"firefonts": FireFonts,
	"fourtops": FourTops,
	"fuzzy": Fuzzy,
	"georgi16": Georgi16,
	"georgia11": Georgia11,
	"goofy": Goofy,
	"gothic": Gothic,
	"graceful": Graceful,
	"gradient": Gradient,
	"graffiti": Graffiti,
	"greek": Greek,
	"henry3d": Henry3D,
	"hex": Hex,
	"hollywood": Hollywood,
	"invita": Invita,
	"isometric1": Isometric1,
	"isometric2": Isometric2,
	"isometric3": Isometric3,
	"isometric4": Isometric4,
	"italic": Italic,
	"ivrit": Ivrit,
	"jsblockletters": JSBlockLetters,
	"jsbracketletters": JSBracketLetters,
	"jscursive": JSCursive,
	"jsstickletters": JSStickLetters,
	"jacky": Jacky,
	"jazmine": Jazmine,
	"jerusalem": Jerusalem,
	"katakana": Katakana,
	"kban": Kban,
	"keyboard": Keyboard,
	"lcd": LCD,
	"larry3d": Larry3D,
	"larry3d2": Larry3D2,
	"lean": Lean,
	"letters": Letters,
	"linux": Linux,
	"lockergnome": Lockergnome,
	"madrid": Madrid,
	"marquee": Marquee,
	"maxfour": Maxfour,
	"mike": Mike,
	"mini": Mini,
	"mirror": Mirror,
	"mnemonic": Mnemonic,
	"morse": Morse,
	"morse2": Morse2,
	"moscow": Moscow,
	"mshebrew210": Mshebrew210,
	"nscript": NScript,
	"ntgreek": NTGreek,
	"nvscript": NVScript,
	"nancyj": Nancyj,
	"nancyjfancy": NancyjFancy,
	"nancyjimproved": NancyjImproved,
	"nancyjunderlined": NancyjUnderlined,
	"nipples": Nipples,
	"o8": O8,
	"os2": OS2,
	"octal": Octal,
	"ogre": Ogre,
	"oldbanner": OldBanner,
	"patorjkhex": PatorjkHeX,
	"patorjkscheese": PatorjksCheese,
	"pawp": Pawp,
	"peaks": Peaks,
	"peaksslant": PeaksSlant,
	"pebbles": Pebbles,
	"pepper": Pepper,
	"poison": Poison,
	"puffy": Puffy,
	"puzzle": Puzzle,
	"rectangles": Rectangles,
	"redphoenix": RedPhoenix,
	"relief": Relief,
	"relief2": Relief2,
	"rev": Rev,
	"reverse": Reverse,
	"roman": Roman,
	"rounded": Rounded,
	"rowancap": RowanCap,
	"rozzo": Rozzo,
	"runic": Runic,
	"runyc": Runyc,
	"sblood": SBlood,
	"slscript": SLScript,
	"santaclara": SantaClara,
	"script": Script,
	"serifcap": Serifcap,
	"shadow": Shadow,
	"shimrod": Shimrod,
	"short": Short,
	"slant": Slant,
	"slantrelief": SlantRelief,
	"slide": Slide,
	"small": Small,
	"smallisometric1": SmallIsometric1,
	"smallkeyboard": SmallKeyboard,
	"smallpoison": SmallPoison,
	"smallscript": SmallScript,
	"smallshadow": SmallShadow,
	"smallslant": SmallSlant,
	"smalltengwar": SmallTengwar,
	"smisome1": Smisome1,
	"smkeyboard": Smkeyboard,
	"smpoison": Smpoison,
	"smscript": Smscript,
	"smshadow": Smshadow,
	"smslant": Smslant,
	"smtengwar": Smtengwar,
	"speed": Speed,
	"srelief": Srelief,
	"stacey": Stacey,
	"stampate": Stampate,
	"stampatello": Stampatello,
	"standard": Standard,
	"starwars": StarWars,
	"stellar": Stellar,
	"stforek": Stforek,
	"stickletters": StickLetters,
	"stop": Stop,
	"straight": Straight,
	"strongerthanall": StrongerThanAll,
	"subzero": SubZero,
	"swan": Swan,
	"this": THIS,
	"tanja": Tanja,
	"tengwar": Tengwar,
	"test1": Test1,
	"theedge": TheEdge,
	"thick": Thick,
	"thin": Thin,
	"threepoint": ThreePoint,
	"ticks": Ticks,
	"ticksslant": TicksSlant,
	"tiles": Tiles,
	"tinkertoy": TinkerToy,
	"tombstone": Tombstone,
	"trek": Trek,
	"tubular": Tubular,
	"twopoint": TwoPoint,
	"usaflag": USAFlag,
	"univers": Univers,
	"wavy": Wavy,
	"weird": Weird,
	"whimsy": Whimsy,
}