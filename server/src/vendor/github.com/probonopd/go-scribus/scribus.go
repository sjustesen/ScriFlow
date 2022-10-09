package scribus

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// https://wiki.scribus.net/canvas/File_Format_Specification_for_Scribus_1.5
// Struct generated using an example Scribus document following https://github.com/miku/zek/issues/14
// Then manually decided which elements can occur multiple times, e.g., changed PAGEOBJECT to []PAGEOBJECT
// in the DOCUMENT struct because one DOCUMENT can have multiple PAGEOBJECTs
// Same for COLOR
// FIXME: Probably many others as well
// FIXME: There must be a better, complete way to generate those structs from e.g., DTDs?

// TODO: Improve completeness
type Document struct {
	XMLName  xml.Name `xml:"SCRIBUSUTF8NEW"`
	Text     string   `xml:",chardata"`
	Version  string   `xml:"Version,attr"`
	DOCUMENT DOCUMENT `xml:"DOCUMENT"`
}

type DOCUMENT struct {
	XMLName                       xml.Name       `xml:"DOCUMENT"`
	Text                          string         `xml:",chardata"`
	ANZPAGES                      string         `xml:"ANZPAGES,attr"`
	PAGEWIDTH                     string         `xml:"PAGEWIDTH,attr"`
	PAGEHEIGHT                    string         `xml:"PAGEHEIGHT,attr"`
	BORDERLEFT                    string         `xml:"BORDERLEFT,attr"`
	BORDERRIGHT                   string         `xml:"BORDERRIGHT,attr"`
	BORDERTOP                     string         `xml:"BORDERTOP,attr"`
	BORDERBOTTOM                  string         `xml:"BORDERBOTTOM,attr"`
	PRESET                        string         `xml:"PRESET,attr"`
	BleedTop                      string         `xml:"BleedTop,attr"`
	BleedLeft                     string         `xml:"BleedLeft,attr"`
	BleedRight                    string         `xml:"BleedRight,attr"`
	BleedBottom                   string         `xml:"BleedBottom,attr"`
	ORIENTATION                   string         `xml:"ORIENTATION,attr"`
	PAGESIZE                      string         `xml:"PAGESIZE,attr"`
	FIRSTNUM                      string         `xml:"FIRSTNUM,attr"`
	BOOK                          string         `xml:"BOOK,attr"`
	AUTOSPALTEN                   string         `xml:"AUTOSPALTEN,attr"`
	ABSTSPALTEN                   string         `xml:"ABSTSPALTEN,attr"`
	UNITS                         string         `xml:"UNITS,attr"`
	DFONT                         string         `xml:"DFONT,attr"`
	DSIZE                         string         `xml:"DSIZE,attr"`
	DCOL                          string         `xml:"DCOL,attr"`
	DGAP                          string         `xml:"DGAP,attr"`
	TabFill                       string         `xml:"TabFill,attr"`
	TabWidth                      string         `xml:"TabWidth,attr"`
	TextDistLeft                  string         `xml:"TextDistLeft,attr"`
	TextDistRight                 string         `xml:"TextDistRight,attr"`
	TextDistBottom                string         `xml:"TextDistBottom,attr"`
	TextDistTop                   string         `xml:"TextDistTop,attr"`
	AUTHOR                        string         `xml:"AUTHOR,attr"`
	COMMENTS                      string         `xml:"COMMENTS,attr"`
	KEYWORDS                      string         `xml:"KEYWORDS,attr"`
	PUBLISHER                     string         `xml:"PUBLISHER,attr"`
	DOCDATE                       string         `xml:"DOCDATE,attr"`
	DOCTYPE                       string         `xml:"DOCTYPE,attr"`
	DOCFORMAT                     string         `xml:"DOCFORMAT,attr"`
	DOCIDENT                      string         `xml:"DOCIDENT,attr"`
	DOCSOURCE                     string         `xml:"DOCSOURCE,attr"`
	DOCLANGINFO                   string         `xml:"DOCLANGINFO,attr"`
	DOCRELATION                   string         `xml:"DOCRELATION,attr"`
	DOCCOVER                      string         `xml:"DOCCOVER,attr"`
	DOCRIGHTS                     string         `xml:"DOCRIGHTS,attr"`
	DOCCONTRIB                    string         `xml:"DOCCONTRIB,attr"`
	TITLE                         string         `xml:"TITLE,attr"`
	SUBJECT                       string         `xml:"SUBJECT,attr"`
	VHOCH                         string         `xml:"VHOCH,attr"`
	VHOCHSC                       string         `xml:"VHOCHSC,attr"`
	VTIEF                         string         `xml:"VTIEF,attr"`
	VTIEFSC                       string         `xml:"VTIEFSC,attr"`
	VKAPIT                        string         `xml:"VKAPIT,attr"`
	BASEGRID                      string         `xml:"BASEGRID,attr"`
	BASEO                         string         `xml:"BASEO,attr"`
	AUTOL                         string         `xml:"AUTOL,attr"`
	UnderlinePos                  string         `xml:"UnderlinePos,attr"`
	UnderlineWidth                string         `xml:"UnderlineWidth,attr"`
	StrikeThruPos                 string         `xml:"StrikeThruPos,attr"`
	StrikeThruWidth               string         `xml:"StrikeThruWidth,attr"`
	GROUPC                        string         `xml:"GROUPC,attr"`
	HCMS                          string         `xml:"HCMS,attr"`
	DPSo                          string         `xml:"DPSo,attr"`
	DPSFo                         string         `xml:"DPSFo,attr"`
	DPuse                         string         `xml:"DPuse,attr"`
	DPgam                         string         `xml:"DPgam,attr"`
	DPbla                         string         `xml:"DPbla,attr"`
	DPPr                          string         `xml:"DPPr,attr"`
	DPIn                          string         `xml:"DPIn,attr"`
	DPInCMYK                      string         `xml:"DPInCMYK,attr"`
	DPIn2                         string         `xml:"DPIn2,attr"`
	DPIn3                         string         `xml:"DPIn3,attr"`
	DISc                          string         `xml:"DISc,attr"`
	DIIm                          string         `xml:"DIIm,attr"`
	ALAYER                        string         `xml:"ALAYER,attr"`
	LANGUAGE                      string         `xml:"LANGUAGE,attr"`
	MINWORDLEN                    string         `xml:"MINWORDLEN,attr"`
	HYCOUNT                       string         `xml:"HYCOUNT,attr"`
	AUTOMATIC                     string         `xml:"AUTOMATIC,attr"`
	AUTOCHECK                     string         `xml:"AUTOCHECK,attr"`
	GUIDELOCK                     string         `xml:"GUIDELOCK,attr"`
	SnapToGuides                  string         `xml:"SnapToGuides,attr"`
	SnapToGrid                    string         `xml:"SnapToGrid,attr"`
	SnapToElement                 string         `xml:"SnapToElement,attr"`
	MINGRID                       string         `xml:"MINGRID,attr"`
	MAJGRID                       string         `xml:"MAJGRID,attr"`
	SHOWGRID                      string         `xml:"SHOWGRID,attr"`
	SHOWGUIDES                    string         `xml:"SHOWGUIDES,attr"`
	Showcolborders                string         `xml:"showcolborders,attr"`
	PreviewMode                   string         `xml:"previewMode,attr"`
	SHOWFRAME                     string         `xml:"SHOWFRAME,attr"`
	SHOWControl                   string         `xml:"SHOWControl,attr"`
	SHOWLAYERM                    string         `xml:"SHOWLAYERM,attr"`
	SHOWMARGIN                    string         `xml:"SHOWMARGIN,attr"`
	SHOWBASE                      string         `xml:"SHOWBASE,attr"`
	SHOWPICT                      string         `xml:"SHOWPICT,attr"`
	SHOWLINK                      string         `xml:"SHOWLINK,attr"`
	RulerMode                     string         `xml:"rulerMode,attr"`
	Showrulers                    string         `xml:"showrulers,attr"`
	ShowBleed                     string         `xml:"showBleed,attr"`
	RulerXoffset                  string         `xml:"rulerXoffset,attr"`
	RulerYoffset                  string         `xml:"rulerYoffset,attr"`
	GuideRad                      string         `xml:"GuideRad,attr"`
	GRAB                          string         `xml:"GRAB,attr"`
	POLYC                         string         `xml:"POLYC,attr"`
	POLYF                         string         `xml:"POLYF,attr"`
	POLYR                         string         `xml:"POLYR,attr"`
	POLYIR                        string         `xml:"POLYIR,attr"`
	POLYCUR                       string         `xml:"POLYCUR,attr"`
	POLYOCUR                      string         `xml:"POLYOCUR,attr"`
	POLYS                         string         `xml:"POLYS,attr"`
	ArcStartAngle                 string         `xml:"arcStartAngle,attr"`
	ArcSweepAngle                 string         `xml:"arcSweepAngle,attr"`
	SpiralStartAngle              string         `xml:"spiralStartAngle,attr"`
	SpiralEndAngle                string         `xml:"spiralEndAngle,attr"`
	SpiralFactor                  string         `xml:"spiralFactor,attr"`
	AutoSave                      string         `xml:"AutoSave,attr"`
	AutoSaveTime                  string         `xml:"AutoSaveTime,attr"`
	AutoSaveCount                 string         `xml:"AutoSaveCount,attr"`
	AutoSaveKeep                  string         `xml:"AutoSaveKeep,attr"`
	AUtoSaveInDocDir              string         `xml:"AUtoSaveInDocDir,attr"`
	AutoSaveDir                   string         `xml:"AutoSaveDir,attr"`
	ScratchBottom                 string         `xml:"ScratchBottom,attr"`
	ScratchLeft                   string         `xml:"ScratchLeft,attr"`
	ScratchRight                  string         `xml:"ScratchRight,attr"`
	ScratchTop                    string         `xml:"ScratchTop,attr"`
	GapHorizontal                 string         `xml:"GapHorizontal,attr"`
	GapVertical                   string         `xml:"GapVertical,attr"`
	StartArrow                    string         `xml:"StartArrow,attr"`
	EndArrow                      string         `xml:"EndArrow,attr"`
	PEN                           string         `xml:"PEN,attr"`
	BRUSH                         string         `xml:"BRUSH,attr"`
	PENLINE                       string         `xml:"PENLINE,attr"`
	PENTEXT                       string         `xml:"PENTEXT,attr"`
	StrokeText                    string         `xml:"StrokeText,attr"`
	TextBackGround                string         `xml:"TextBackGround,attr"`
	TextLineColor                 string         `xml:"TextLineColor,attr"`
	TextBackGroundShade           string         `xml:"TextBackGroundShade,attr"`
	TextLineShade                 string         `xml:"TextLineShade,attr"`
	TextPenShade                  string         `xml:"TextPenShade,attr"`
	TextStrokeShade               string         `xml:"TextStrokeShade,attr"`
	STIL                          string         `xml:"STIL,attr"`
	STILLINE                      string         `xml:"STILLINE,attr"`
	WIDTH                         string         `xml:"WIDTH,attr"`
	WIDTHLINE                     string         `xml:"WIDTHLINE,attr"`
	PENSHADE                      string         `xml:"PENSHADE,attr"`
	LINESHADE                     string         `xml:"LINESHADE,attr"`
	BRUSHSHADE                    string         `xml:"BRUSHSHADE,attr"`
	CPICT                         string         `xml:"CPICT,attr"`
	PICTSHADE                     string         `xml:"PICTSHADE,attr"`
	CSPICT                        string         `xml:"CSPICT,attr"`
	PICTSSHADE                    string         `xml:"PICTSSHADE,attr"`
	PICTSCX                       string         `xml:"PICTSCX,attr"`
	PICTSCY                       string         `xml:"PICTSCY,attr"`
	PSCALE                        string         `xml:"PSCALE,attr"`
	PASPECT                       string         `xml:"PASPECT,attr"`
	EmbeddedPath                  string         `xml:"EmbeddedPath,attr"`
	HalfRes                       string         `xml:"HalfRes,attr"`
	DispX                         string         `xml:"dispX,attr"`
	DispY                         string         `xml:"dispY,attr"`
	Constrain                     string         `xml:"constrain,attr"`
	MINORC                        string         `xml:"MINORC,attr"`
	MAJORC                        string         `xml:"MAJORC,attr"`
	GuideC                        string         `xml:"GuideC,attr"`
	BaseC                         string         `xml:"BaseC,attr"`
	RenderStack                   string         `xml:"renderStack,attr"`
	GridType                      string         `xml:"GridType,attr"`
	PAGEC                         string         `xml:"PAGEC,attr"`
	MARGC                         string         `xml:"MARGC,attr"`
	RANDF                         string         `xml:"RANDF,attr"`
	CurrentProfile                string         `xml:"currentProfile,attr"`
	CalligraphicPenFillColor      string         `xml:"calligraphicPenFillColor,attr"`
	CalligraphicPenLineColor      string         `xml:"calligraphicPenLineColor,attr"`
	CalligraphicPenFillColorShade string         `xml:"calligraphicPenFillColorShade,attr"`
	CalligraphicPenLineColorShade string         `xml:"calligraphicPenLineColorShade,attr"`
	CalligraphicPenLineWidth      string         `xml:"calligraphicPenLineWidth,attr"`
	CalligraphicPenAngle          string         `xml:"calligraphicPenAngle,attr"`
	CalligraphicPenWidth          string         `xml:"calligraphicPenWidth,attr"`
	CalligraphicPenStyle          string         `xml:"calligraphicPenStyle,attr"`
	CheckProfile                  []CheckProfile `xml:"CheckProfile"`
	COLOR                         []COLOR        `xml:"COLOR"`
	HYPHEN                        string         `xml:"HYPHEN"`
	STYLE                         STYLE          `xml:"STYLE"`
	CHARSTYLE                     CHARSTYLE      `xml:"CHARSTYLE"`
	TableStyle                    TableStyle     `xml:"TableStyle"`
	CellStyle                     CellStyle      `xml:"CellStyle"`
	LAYERS                        LAYERS         `xml:"LAYERS"`
	Printer                       Printer        `xml:"Printer"`
	PDF                           PDF            `xml:"PDF"`
	DocItemAttributes             string         `xml:"DocItemAttributes"`
	TablesOfContents              string         `xml:"TablesOfContents"`
	NotesStyles                   NotesStyles    `xml:"NotesStyles"`
	NotesFrames                   string         `xml:"NotesFrames"`
	PageSets                      PageSets       `xml:"PageSets"`
	Sections                      Sections       `xml:"Sections"`
	MASTERPAGE                    MASTERPAGE     `xml:"MASTERPAGE"`
	PAGE                          []PAGE         `xml:"PAGE"`
	PAGEOBJECT                    []PAGEOBJECT   `xml:"PAGEOBJECT"`
}

type FRAMEOBJECT struct {
	XMLName         xml.Name  `xml:"FRAMEOBJECT"`
	Text            string    `xml:",chardata"`
	InID            string    `xml:"InID,attr"`
	XPOS            string    `xml:"XPOS,attr"`
	YPOS            string    `xml:"YPOS,attr"`
	OwnPage         string    `xml:"OwnPage,attr"`
	ItemID          string    `xml:"ItemID,attr"`
	PTYPE           string    `xml:"PTYPE,attr"`
	WIDTH           string    `xml:"WIDTH,attr"`
	HEIGHT          string    `xml:"HEIGHT,attr"`
	FRTYPE          string    `xml:"FRTYPE,attr"`
	CLIPEDIT        string    `xml:"CLIPEDIT,attr"`
	PWIDTH          string    `xml:"PWIDTH,attr"`
	PLINEART        string    `xml:"PLINEART,attr"`
	LOCALSCX        string    `xml:"LOCALSCX,attr"`
	LOCALSCY        string    `xml:"LOCALSCY,attr"`
	LOCALX          string    `xml:"LOCALX,attr"`
	LOCALY          string    `xml:"LOCALY,attr"`
	LOCALROT        string    `xml:"LOCALROT,attr"`
	PICART          string    `xml:"PICART,attr"`
	SCALETYPE       string    `xml:"SCALETYPE,attr"`
	RATIO           string    `xml:"RATIO,attr"`
	COLUMNS         string    `xml:"COLUMNS,attr"`
	COLGAP          string    `xml:"COLGAP,attr"`
	AUTOTEXT        string    `xml:"AUTOTEXT,attr"`
	EXTRA           string    `xml:"EXTRA,attr"`
	TEXTRA          string    `xml:"TEXTRA,attr"`
	BEXTRA          string    `xml:"BEXTRA,attr"`
	REXTRA          string    `xml:"REXTRA,attr"`
	VAlign          string    `xml:"VAlign,attr"`
	FLOP            string    `xml:"FLOP,attr"`
	PLTSHOW         string    `xml:"PLTSHOW,attr"`
	BASEOF          string    `xml:"BASEOF,attr"`
	TextPathType    string    `xml:"textPathType,attr"`
	TextPathFlipped string    `xml:"textPathFlipped,attr"`
	Path            string    `xml:"path,attr"`
	Copath          string    `xml:"copath,attr"`
	IsInline        string    `xml:"isInline,attr"`
	GXpos           string    `xml:"gXpos,attr"`
	GYpos           string    `xml:"gYpos,attr"`
	GWidth          string    `xml:"gWidth,attr"`
	GHeight         string    `xml:"gHeight,attr"`
	LAYER           string    `xml:"LAYER,attr"`
	NEXTITEM        string    `xml:"NEXTITEM,attr"`
	BACKITEM        string    `xml:"BACKITEM,attr"`
	StoryText       StoryText `xml:"StoryText"`
}

type CheckProfile struct {
	XMLName                          xml.Name `xml:"CheckProfile"`
	Text                             string   `xml:",chardata"`
	Name                             string   `xml:"Name,attr"`
	IgnoreErrors                     string   `xml:"ignoreErrors,attr"`
	AutoCheck                        string   `xml:"autoCheck,attr"`
	CheckGlyphs                      string   `xml:"checkGlyphs,attr"`
	CheckOrphans                     string   `xml:"checkOrphans,attr"`
	CheckOverflow                    string   `xml:"checkOverflow,attr"`
	CheckPictures                    string   `xml:"checkPictures,attr"`
	CheckPartFilledImageFrames       string   `xml:"checkPartFilledImageFrames,attr"`
	CheckResolution                  string   `xml:"checkResolution,attr"`
	CheckTransparency                string   `xml:"checkTransparency,attr"`
	MinResolution                    string   `xml:"minResolution,attr"`
	MaxResolution                    string   `xml:"maxResolution,attr"`
	CheckAnnotations                 string   `xml:"checkAnnotations,attr"`
	CheckRasterPDF                   string   `xml:"checkRasterPDF,attr"`
	CheckForGIF                      string   `xml:"checkForGIF,attr"`
	IgnoreOffLayers                  string   `xml:"ignoreOffLayers,attr"`
	CheckNotCMYKOrSpot               string   `xml:"checkNotCMYKOrSpot,attr"`
	CheckDeviceColorsAndOutputIntent string   `xml:"checkDeviceColorsAndOutputIntent,attr"`
	CheckFontNotEmbedded             string   `xml:"checkFontNotEmbedded,attr"`
	CheckFontIsOpenType              string   `xml:"checkFontIsOpenType,attr"`
	CheckAppliedMasterDifferentSide  string   `xml:"checkAppliedMasterDifferentSide,attr"`
	CheckEmptyTextFrames             string   `xml:"checkEmptyTextFrames,attr"`
}

type COLOR struct {
	XMLName  xml.Name `xml:"COLOR"`
	Text     string   `xml:",chardata"`
	NAME     string   `xml:"NAME,attr"`
	CMYK     string   `xml:"CMYK,attr"`
	SPACE    string   `xml:"SPACE,attr"`
	C        string   `xml:"C,attr"`
	M        string   `xml:"M,attr"`
	Y        string   `xml:"Y,attr"`
	K        string   `xml:"K,attr"`
	R        string   `xml:"R,attr"`
	G        string   `xml:"G,attr"`
	B        string   `xml:"B,attr"`
	Register string   `xml:"Register,attr"`
}

// Default paragraph styles
// FIXME: Probably not complete
type STYLE struct {
	XMLName                xml.Name `xml:"STYLE"`
	Text                   string   `xml:",chardata"`
	NAME                   string   `xml:"NAME,attr"`
	DefaultStyle           string   `xml:"DefaultStyle,attr"`
	ALIGN                  string   `xml:"ALIGN,attr"`
	DIRECTION              string   `xml:"DIRECTION,attr"`
	LINESPMode             string   `xml:"LINESPMode,attr"`
	LINESP                 string   `xml:"LINESP,attr"`
	INDENT                 string   `xml:"INDENT,attr"`
	RMARGIN                string   `xml:"RMARGIN,attr"`
	FIRST                  string   `xml:"FIRST,attr"`
	VOR                    string   `xml:"VOR,attr"`
	NACH                   string   `xml:"NACH,attr"`
	ParagraphEffectOffset  string   `xml:"ParagraphEffectOffset,attr"`
	DROP                   string   `xml:"DROP,attr"`
	DROPLIN                string   `xml:"DROPLIN,attr"`
	Bullet                 string   `xml:"Bullet,attr"`
	Numeration             string   `xml:"Numeration,attr"`
	HyphenConsecutiveLines string   `xml:"HyphenConsecutiveLines,attr"`
	BCOLOR                 string   `xml:"BCOLOR,attr"`
	BSHADE                 string   `xml:"BSHADE,attr"`
	PARENT                 string   `xml:"PARENT,attr"`
	FONT                   string   `xml:"FONT,attr"`
	FCOLOR                 string   `xml:"FCOLOR,attr"`
}

type CHARSTYLE struct {
	XMLName       xml.Name `xml:"CHARSTYLE"`
	Text          string   `xml:",chardata"`
	CNAME         string   `xml:"CNAME,attr"`
	DefaultStyle  string   `xml:"DefaultStyle,attr"`
	FONT          string   `xml:"FONT,attr"`
	FONTSIZE      string   `xml:"FONTSIZE,attr"`
	FONTFEATURES  string   `xml:"FONTFEATURES,attr"`
	FEATURES      string   `xml:"FEATURES,attr"`
	FCOLOR        string   `xml:"FCOLOR,attr"`
	FSHADE        string   `xml:"FSHADE,attr"`
	HyphenWordMin string   `xml:"HyphenWordMin,attr"`
	SCOLOR        string   `xml:"SCOLOR,attr"`
	BGCOLOR       string   `xml:"BGCOLOR,attr"`
	BGSHADE       string   `xml:"BGSHADE,attr"`
	SSHADE        string   `xml:"SSHADE,attr"`
	TXTSHX        string   `xml:"TXTSHX,attr"`
	TXTSHY        string   `xml:"TXTSHY,attr"`
	TXTOUT        string   `xml:"TXTOUT,attr"`
	TXTULP        string   `xml:"TXTULP,attr"`
	TXTULW        string   `xml:"TXTULW,attr"`
	TXTSTP        string   `xml:"TXTSTP,attr"`
	TXTSTW        string   `xml:"TXTSTW,attr"`
	SCALEH        string   `xml:"SCALEH,attr"`
	SCALEV        string   `xml:"SCALEV,attr"`
	BASEO         string   `xml:"BASEO,attr"`
	KERN          string   `xml:"KERN,attr"`
	LANGUAGE      string   `xml:"LANGUAGE,attr"`
}

type TableStyle struct {
	XMLName           xml.Name          `xml:"TableStyle"`
	Text              string            `xml:",chardata"`
	NAME              string            `xml:"NAME,attr"`
	DefaultStyle      string            `xml:"DefaultStyle,attr"`
	FillColor         string            `xml:"FillColor,attr"`
	FillShade         string            `xml:"FillShade,attr"`
	TableBorderLeft   TableBorderLeft   `xml:"TableBorderLeft"`
	TableBorderRight  TableBorderRight  `xml:"TableBorderRight"`
	TableBorderTop    TableBorderTop    `xml:"TableBorderTop"`
	TableBorderBottom TableBorderBottom `xml:"TableBorderBottom"`
}

type TableBorderLeft struct {
	XMLName         xml.Name        `xml:"TableBorderLeft"`
	Text            string          `xml:",chardata"`
	TableBorderLine TableBorderLine `xml:"TableBorderLine"`
}

type TableBorderLine struct {
	XMLName  xml.Name `xml:"TableBorderLine"`
	Text     string   `xml:",chardata"`
	Width    string   `xml:"Width,attr"`
	PenStyle string   `xml:"PenStyle,attr"`
	Color    string   `xml:"Color,attr"`
	Shade    string   `xml:"Shade,attr"`
}

type TableBorderRight struct {
	XMLName         xml.Name        `xml:"TableBorderRight"`
	Text            string          `xml:",chardata"`
	TableBorderLine TableBorderLine `xml:"TableBorderLine"`
}

type TableBorderTop struct {
	XMLName         xml.Name        `xml:"TableBorderTop"`
	Text            string          `xml:",chardata"`
	TableBorderLine TableBorderLine `xml:"TableBorderLine"`
}

type TableBorderBottom struct {
	XMLName         xml.Name        `xml:"TableBorderBottom"`
	Text            string          `xml:",chardata"`
	TableBorderLine TableBorderLine `xml:"TableBorderLine"`
}

type CellStyle struct {
	XMLName           xml.Name          `xml:"CellStyle"`
	Text              string            `xml:",chardata"`
	NAME              string            `xml:"NAME,attr"`
	DefaultStyle      string            `xml:"DefaultStyle,attr"`
	FillColor         string            `xml:"FillColor,attr"`
	FillShade         string            `xml:"FillShade,attr"`
	LeftPadding       string            `xml:"LeftPadding,attr"`
	RightPadding      string            `xml:"RightPadding,attr"`
	TopPadding        string            `xml:"TopPadding,attr"`
	BottomPadding     string            `xml:"BottomPadding,attr"`
	TableBorderLeft   TableBorderLeft   `xml:"TableBorderLeft"`
	TableBorderRight  TableBorderRight  `xml:"TableBorderRight"`
	TableBorderTop    TableBorderTop    `xml:"TableBorderTop"`
	TableBorderBottom TableBorderBottom `xml:"TableBorderBottom"`
}

type LAYERS []struct {
	XMLName  xml.Name `xml:"LAYERS"`
	Text     string   `xml:",chardata"`
	NUMMER   string   `xml:"NUMMER,attr"`
	LEVEL    string   `xml:"LEVEL,attr"`
	NAME     string   `xml:"NAME,attr"`
	SICHTBAR string   `xml:"SICHTBAR,attr"`
	DRUCKEN  string   `xml:"DRUCKEN,attr"`
	EDIT     string   `xml:"EDIT,attr"`
	SELECT   string   `xml:"SELECT,attr"`
	FLOW     string   `xml:"FLOW,attr"`
	TRANS    string   `xml:"TRANS,attr"`
	BLEND    string   `xml:"BLEND,attr"`
	OUTL     string   `xml:"OUTL,attr"`
	LAYERC   string   `xml:"LAYERC,attr"`
}

type Printer struct {
	XMLName            xml.Name `xml:"Printer"`
	Text               string   `xml:",chardata"`
	FirstUse           string   `xml:"firstUse,attr"`
	ToFile             string   `xml:"toFile,attr"`
	UseAltPrintCommand string   `xml:"useAltPrintCommand,attr"`
	OutputSeparations  string   `xml:"outputSeparations,attr"`
	UseSpotColors      string   `xml:"useSpotColors,attr"`
	UseColor           string   `xml:"useColor,attr"`
	MirrorH            string   `xml:"mirrorH,attr"`
	MirrorV            string   `xml:"mirrorV,attr"`
	UseICC             string   `xml:"useICC,attr"`
	DoGCR              string   `xml:"doGCR,attr"`
	DoClip             string   `xml:"doClip,attr"`
	SetDevParam        string   `xml:"setDevParam,attr"`
	UseDocBleeds       string   `xml:"useDocBleeds,attr"`
	CropMarks          string   `xml:"cropMarks,attr"`
	BleedMarks         string   `xml:"bleedMarks,attr"`
	RegistrationMarks  string   `xml:"registrationMarks,attr"`
	ColorMarks         string   `xml:"colorMarks,attr"`
	IncludePDFMarks    string   `xml:"includePDFMarks,attr"`
	PSLevel            string   `xml:"PSLevel,attr"`
	PDLanguage         string   `xml:"PDLanguage,attr"`
	MarkLength         string   `xml:"markLength,attr"`
	MarkOffset         string   `xml:"markOffset,attr"`
	BleedTop           string   `xml:"BleedTop,attr"`
	BleedLeft          string   `xml:"BleedLeft,attr"`
	BleedRight         string   `xml:"BleedRight,attr"`
	BleedBottom        string   `xml:"BleedBottom,attr"`
	Printer            string   `xml:"printer,attr"`
	Filename           string   `xml:"filename,attr"`
	SeparationName     string   `xml:"separationName,attr"`
	PrinterCommand     string   `xml:"printerCommand,attr"`
}

type PDF struct {
	XMLName           xml.Name `xml:"PDF"`
	Text              string   `xml:",chardata"`
	FirstUse          string   `xml:"firstUse,attr"`
	Thumbnails        string   `xml:"Thumbnails,attr"`
	Articles          string   `xml:"Articles,attr"`
	Bookmarks         string   `xml:"Bookmarks,attr"`
	Compress          string   `xml:"Compress,attr"`
	CMethod           string   `xml:"CMethod,attr"`
	Quality           string   `xml:"Quality,attr"`
	EmbedPDF          string   `xml:"EmbedPDF,attr"`
	MirrorH           string   `xml:"MirrorH,attr"`
	MirrorV           string   `xml:"MirrorV,attr"`
	Clip              string   `xml:"Clip,attr"`
	RangeSel          string   `xml:"rangeSel,attr"`
	RangeTxt          string   `xml:"rangeTxt,attr"`
	RotateDeg         string   `xml:"RotateDeg,attr"`
	PresentMode       string   `xml:"PresentMode,attr"`
	RecalcPic         string   `xml:"RecalcPic,attr"`
	FontEmbedding     string   `xml:"FontEmbedding,attr"`
	Grayscale         string   `xml:"Grayscale,attr"`
	RGBMode           string   `xml:"RGBMode,attr"`
	UseProfiles       string   `xml:"UseProfiles,attr"`
	UseProfiles2      string   `xml:"UseProfiles2,attr"`
	Binding           string   `xml:"Binding,attr"`
	PicRes            string   `xml:"PicRes,attr"`
	Resolution        string   `xml:"Resolution,attr"`
	Version           string   `xml:"Version,attr"`
	Intent            string   `xml:"Intent,attr"`
	Intent2           string   `xml:"Intent2,attr"`
	SolidP            string   `xml:"SolidP,attr"`
	ImageP            string   `xml:"ImageP,attr"`
	PrintP            string   `xml:"PrintP,attr"`
	InfoString        string   `xml:"InfoString,attr"`
	BTop              string   `xml:"BTop,attr"`
	BLeft             string   `xml:"BLeft,attr"`
	BRight            string   `xml:"BRight,attr"`
	BBottom           string   `xml:"BBottom,attr"`
	UseDocBleeds      string   `xml:"useDocBleeds,attr"`
	CropMarks         string   `xml:"cropMarks,attr"`
	BleedMarks        string   `xml:"bleedMarks,attr"`
	RegistrationMarks string   `xml:"registrationMarks,attr"`
	ColorMarks        string   `xml:"colorMarks,attr"`
	DocInfoMarks      string   `xml:"docInfoMarks,attr"`
	MarkLength        string   `xml:"markLength,attr"`
	MarkOffset        string   `xml:"markOffset,attr"`
	ImagePr           string   `xml:"ImagePr,attr"`
	PassOwner         string   `xml:"PassOwner,attr"`
	PassUser          string   `xml:"PassUser,attr"`
	Permissions       string   `xml:"Permissions,attr"`
	Encrypt           string   `xml:"Encrypt,attr"`
	UseLayers         string   `xml:"UseLayers,attr"`
	UseLpi            string   `xml:"UseLpi,attr"`
	UseSpotColors     string   `xml:"UseSpotColors,attr"`
	DoMultiFile       string   `xml:"doMultiFile,attr"`
	DisplayBookmarks  string   `xml:"displayBookmarks,attr"`
	DisplayFullscreen string   `xml:"displayFullscreen,attr"`
	DisplayLayers     string   `xml:"displayLayers,attr"`
	DisplayThumbs     string   `xml:"displayThumbs,attr"`
	HideMenuBar       string   `xml:"hideMenuBar,attr"`
	HideToolBar       string   `xml:"hideToolBar,attr"`
	FitWindow         string   `xml:"fitWindow,attr"`
	OpenAfterExport   string   `xml:"openAfterExport,attr"`
	PageLayout        string   `xml:"PageLayout,attr"`
	OpenAction        string   `xml:"openAction,attr"`
	Fonts             Fonts    `xml:"Fonts"`
	Subset            []Subset `xml:"Subset"`
	LPI               []LPI    `xml:"LPI"`
}

type Fonts struct {
	XMLName xml.Name `xml:"Fonts"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"Name,attr"`
}

type Subset struct {
	XMLName xml.Name `xml:"Subset"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"Name,attr"`
}

type LPI struct {
	XMLName      xml.Name `xml:"LPI"`
	Text         string   `xml:",chardata"`
	Color        string   `xml:"Color,attr"`
	Frequency    string   `xml:"Frequency,attr"`
	Angle        string   `xml:"Angle,attr"`
	SpotFunction string   `xml:"SpotFunction,attr"`
}

type NotesStyle struct {
	XMLName     xml.Name `xml:"notesStyle"`
	Text        string   `xml:",chardata"`
	Name        string   `xml:"Name,attr"`
	Start       string   `xml:"Start,attr"`
	Endnotes    string   `xml:"Endnotes,attr"`
	Type        string   `xml:"Type,attr"`
	Range       string   `xml:"Range,attr"`
	Prefix      string   `xml:"Prefix,attr"`
	Suffix      string   `xml:"Suffix,attr"`
	AutoHeight  string   `xml:"AutoHeight,attr"`
	AutoWidth   string   `xml:"AutoWidth,attr"`
	AutoRemove  string   `xml:"AutoRemove,attr"`
	AutoWeld    string   `xml:"AutoWeld,attr"`
	SuperNote   string   `xml:"SuperNote,attr"`
	SuperMaster string   `xml:"SuperMaster,attr"`
	MarksStyle  string   `xml:"MarksStyle,attr"`
	NotesStyle  string   `xml:"NotesStyle,attr"`
}

type NotesStyles struct {
	XMLName    xml.Name   `xml:"NotesStyles"`
	Text       string     `xml:",chardata"`
	NotesStyle NotesStyle `xml:"notesStyle"`
}

type PageSets struct {
	XMLName xml.Name `xml:"PageSets"`
	Text    string   `xml:",chardata"`
	Set     []Set    `xml:"Set"`
}

type Set struct {
	XMLName   xml.Name    `xml:"Set"`
	Text      string      `xml:",chardata"`
	Name      string      `xml:"Name,attr"`
	FirstPage string      `xml:"FirstPage,attr"`
	Rows      string      `xml:"Rows,attr"`
	Columns   string      `xml:"Columns,attr"`
	PageNames []PageNames `xml:"PageNames"`
}

type PageNames struct {
	XMLName xml.Name `xml:"PageNames"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"Name,attr"`
}

type Sections struct {
	XMLName xml.Name `xml:"Sections"`
	Text    string   `xml:",chardata"`
	Section Section  `xml:"Section"`
}

type Section struct {
	XMLName    xml.Name `xml:"Section"`
	Text       string   `xml:",chardata"`
	Number     string   `xml:"Number,attr"`
	Name       string   `xml:"Name,attr"`
	From       string   `xml:"From,attr"`
	To         string   `xml:"To,attr"`
	Type       string   `xml:"Type,attr"`
	Start      string   `xml:"Start,attr"`
	Reversed   string   `xml:"Reversed,attr"`
	Active     string   `xml:"Active,attr"`
	FillChar   string   `xml:"FillChar,attr"`
	FieldWidth string   `xml:"FieldWidth,attr"`
}

type MASTERPAGE struct {
	XMLName               xml.Name `xml:"MASTERPAGE"`
	Text                  string   `xml:",chardata"`
	PAGEXPOS              string   `xml:"PAGEXPOS,attr"`
	PAGEYPOS              string   `xml:"PAGEYPOS,attr"`
	PAGEWIDTH             string   `xml:"PAGEWIDTH,attr"`
	PAGEHEIGHT            string   `xml:"PAGEHEIGHT,attr"`
	BORDERLEFT            string   `xml:"BORDERLEFT,attr"`
	BORDERRIGHT           string   `xml:"BORDERRIGHT,attr"`
	BORDERTOP             string   `xml:"BORDERTOP,attr"`
	BORDERBOTTOM          string   `xml:"BORDERBOTTOM,attr"`
	NUM                   string   `xml:"NUM,attr"`
	NAM                   string   `xml:"NAM,attr"`
	MNAM                  string   `xml:"MNAM,attr"`
	Size                  string   `xml:"Size,attr"`
	Orientation           string   `xml:"Orientation,attr"`
	LEFT                  string   `xml:"LEFT,attr"`
	PRESET                string   `xml:"PRESET,attr"`
	VerticalGuides        string   `xml:"VerticalGuides,attr"`
	HorizontalGuides      string   `xml:"HorizontalGuides,attr"`
	AGhorizontalAutoGap   string   `xml:"AGhorizontalAutoGap,attr"`
	AGverticalAutoGap     string   `xml:"AGverticalAutoGap,attr"`
	AGhorizontalAutoCount string   `xml:"AGhorizontalAutoCount,attr"`
	AGverticalAutoCount   string   `xml:"AGverticalAutoCount,attr"`
	AGhorizontalAutoRefer string   `xml:"AGhorizontalAutoRefer,attr"`
	AGverticalAutoRefer   string   `xml:"AGverticalAutoRefer,attr"`
	AGSelection           string   `xml:"AGSelection,attr"`
	PageEffectDuration    string   `xml:"pageEffectDuration,attr"`
	PageViewDuration      string   `xml:"pageViewDuration,attr"`
	EffectType            string   `xml:"effectType,attr"`
	Dm                    string   `xml:"Dm,attr"`
	M                     string   `xml:"M,attr"`
	Di                    string   `xml:"Di,attr"`
}

type PAGE struct {
	XMLName               xml.Name `xml:"PAGE"`
	Text                  string   `xml:",chardata"`
	PAGEXPOS              string   `xml:"PAGEXPOS,attr"`
	PAGEYPOS              string   `xml:"PAGEYPOS,attr"`
	PAGEWIDTH             string   `xml:"PAGEWIDTH,attr"`
	PAGEHEIGHT            string   `xml:"PAGEHEIGHT,attr"`
	BORDERLEFT            string   `xml:"BORDERLEFT,attr"`
	BORDERRIGHT           string   `xml:"BORDERRIGHT,attr"`
	BORDERTOP             string   `xml:"BORDERTOP,attr"`
	BORDERBOTTOM          string   `xml:"BORDERBOTTOM,attr"`
	NUM                   string   `xml:"NUM,attr"`
	NAM                   string   `xml:"NAM,attr"`
	MNAM                  string   `xml:"MNAM,attr"`
	Size                  string   `xml:"Size,attr"`
	Orientation           string   `xml:"Orientation,attr"`
	LEFT                  string   `xml:"LEFT,attr"`
	PRESET                string   `xml:"PRESET,attr"`
	VerticalGuides        string   `xml:"VerticalGuides,attr"`
	HorizontalGuides      string   `xml:"HorizontalGuides,attr"`
	AGhorizontalAutoGap   string   `xml:"AGhorizontalAutoGap,attr"`
	AGverticalAutoGap     string   `xml:"AGverticalAutoGap,attr"`
	AGhorizontalAutoCount string   `xml:"AGhorizontalAutoCount,attr"`
	AGverticalAutoCount   string   `xml:"AGverticalAutoCount,attr"`
	AGhorizontalAutoRefer string   `xml:"AGhorizontalAutoRefer,attr"`
	AGverticalAutoRefer   string   `xml:"AGverticalAutoRefer,attr"`
	AGSelection           string   `xml:"AGSelection,attr"`
	PageEffectDuration    string   `xml:"pageEffectDuration,attr"`
	PageViewDuration      string   `xml:"pageViewDuration,attr"`
	EffectType            string   `xml:"effectType,attr"`
	Dm                    string   `xml:"Dm,attr"`
	M                     string   `xml:"M,attr"`
	Di                    string   `xml:"Di,attr"`
}

type PAGEOBJECT struct {
	XMLName           xml.Name     `xml:"PAGEOBJECT"`
	Text              string       `xml:",chardata"`
	XPOS              string       `xml:"XPOS,attr"`
	YPOS              string       `xml:"YPOS,attr"`
	OwnPage           string       `xml:"OwnPage,attr"`
	ItemID            string       `xml:"ItemID,attr"`
	PTYPE             string       `xml:"PTYPE,attr"`
	WIDTH             string       `xml:"WIDTH,attr"`
	HEIGHT            string       `xml:"HEIGHT,attr"`
	FRTYPE            string       `xml:"FRTYPE,attr"`
	CLIPEDIT          string       `xml:"CLIPEDIT,attr"`
	ROT               string       `xml:"ROT,attr"`
	PWIDTH            string       `xml:"PWIDTH,attr"`
	PCOLOR            string       `xml:"PCOLOR,attr"`
	PLINEART          string       `xml:"PLINEART,attr"`
	LOCALSCX          string       `xml:"LOCALSCX,attr"`
	LOCALSCY          string       `xml:"LOCALSCY,attr"`
	LOCALX            string       `xml:"LOCALX,attr"`
	LOCALY            string       `xml:"LOCALY,attr"`
	LOCALROT          string       `xml:"LOCALROT,attr"`
	PICART            string       `xml:"PICART,attr"`
	SCALETYPE         string       `xml:"SCALETYPE,attr"`
	RATIO             string       `xml:"RATIO,attr"`
	TransValue        string       `xml:"TransValue,attr"`
	Path              string       `xml:"path,attr"`
	Copath            string       `xml:"copath,attr"`
	GXpos             string       `xml:"gXpos,attr"`
	GYpos             string       `xml:"gYpos,attr"`
	GWidth            string       `xml:"gWidth,attr"`
	GHeight           string       `xml:"gHeight,attr"`
	LAYER             string       `xml:"LAYER,attr"`
	NEXTITEM          string       `xml:"NEXTITEM,attr"`
	BACKITEM          string       `xml:"BACKITEM,attr"`
	Pagenumber        string       `xml:"Pagenumber,attr"`
	PFILE             string       `xml:"PFILE,attr"`
	IRENDER           string       `xml:"IRENDER,attr"`
	EMBEDDED          string       `xml:"EMBEDDED,attr"`
	COMPRESSIONMETHOD string       `xml:"COMPRESSIONMETHOD,attr"`
	GRExtM            string       `xml:"GRExtM,attr"`
	GRTYPM            string       `xml:"GRTYPM,attr"`
	GRSTARTXM         string       `xml:"GRSTARTXM,attr"`
	GRSTARTYM         string       `xml:"GRSTARTYM,attr"`
	GRENDXM           string       `xml:"GRENDXM,attr"`
	GRENDYM           string       `xml:"GRENDYM,attr"`
	GRFOCALXM         string       `xml:"GRFOCALXM,attr"`
	GRFOCALYM         string       `xml:"GRFOCALYM,attr"`
	GRSCALEM          string       `xml:"GRSCALEM,attr"`
	GRSKEWM           string       `xml:"GRSKEWM,attr"`
	ImageRes          string       `xml:"ImageRes,attr"`
	FillRule          string       `xml:"fillRule,attr"`
	ANNAME            string       `xml:"ANNAME,attr"`
	GroupWidth        string       `xml:"groupWidth,attr"`
	GroupHeight       string       `xml:"groupHeight,attr"`
	GroupClips        string       `xml:"groupClips,attr"`
	PLINEEND          string       `xml:"PLINEEND,attr"`
	PLINEJOIN         string       `xml:"PLINEJOIN,attr"`
	RADRECT           string       `xml:"RADRECT,attr"`
	PCOLOR2           string       `xml:"PCOLOR2,attr"`
	PRFILE            string       `xml:"PRFILE,attr"`
	COLUMNS           string       `xml:"COLUMNS,attr"`
	COLGAP            string       `xml:"COLGAP,attr"`
	AUTOTEXT          string       `xml:"AUTOTEXT,attr"`
	EXTRA             string       `xml:"EXTRA,attr"`
	TEXTRA            string       `xml:"TEXTRA,attr"`
	BEXTRA            string       `xml:"BEXTRA,attr"`
	REXTRA            string       `xml:"REXTRA,attr"`
	VAlign            string       `xml:"VAlign,attr"`
	FLOP              string       `xml:"FLOP,attr"`
	PLTSHOW           string       `xml:"PLTSHOW,attr"`
	BASEOF            string       `xml:"BASEOF,attr"`
	TextPathType      string       `xml:"textPathType,attr"`
	TextPathFlipped   string       `xml:"textPathFlipped,attr"`
	PSTYLE            string       `xml:"PSTYLE,attr"`
	StoryText         StoryText    `xml:"StoryText"`
	PAGEOBJECT        []PAGEOBJECT `xml:"PAGEOBJECT"`
}

type Para struct {
	XMLName                  xml.Name `xml:"para"`
	Text                     string   `xml:",chardata"`
	LINESP                   string   `xml:"LINESP,attr"`
	ParagraphEffectCharStyle string   `xml:"ParagraphEffectCharStyle,attr"`
	ParagraphEffectOffset    string   `xml:"ParagraphEffectOffset,attr"`
	ParagraphEffectIndent    string   `xml:"ParagraphEffectIndent,attr"`
	DROP                     string   `xml:"DROP,attr"`
	Bullet                   string   `xml:"Bullet,attr"`
	BulletStr                string   `xml:"BulletStr,attr"`
	Numeration               string   `xml:"Numeration,attr"`
}

// FIXME: The order of 'ITEXT' and 'Para' in the XML document must not be changed;
// currently we are getting all 'ITEXT's first and then all 'Para's when we write out
// the changed XML, probably due to them being declared as []? How to fix this?
type StoryText struct {
	XMLName       xml.Name        `xml:"StoryText"`
	Text          string          `xml:",chardata"`
	DefaultStyle  DefaultStyle    `xml:"DefaultStyle"`
	ITEXT         []ITEXT         `xml:"ITEXT"`
	Para          []Para          `xml:"para"`
	Trail         Trail           `xml:"trail"`
	StoryTextSpan []StoryTextSpan `xml:"StoryTextSpan"`
}

// FIXME: Working around the above by introducing a phantasy tag 'StoryTextSpan' to group the elements inside
// Luckily Scribus can still open the file
// https://gitlab.com/scribus/scribus/issues/8
// TODO: If we want to parse files that contain this tag ourselves again, we also need to handle this case...
type StoryTextSpan struct {
	Text         string       `xml:",chardata"`
	DefaultStyle DefaultStyle `xml:"DefaultStyle"`
	ITEXT        ITEXT        `xml:"ITEXT"`
	Para         Para         `xml:"para"`
	Trail        Trail        `xml:"trail"`
}

type DefaultStyle struct {
	XMLName        xml.Name `xml:"DefaultStyle"`
	Text           string   `xml:",chardata"`
	LINESP         string   `xml:"LINESP,attr"`
	LINESPMode     string   `xml:"LINESPMode,attr"`
	FCOLOR         string   `xml:"FCOLOR,attr"`
	FONT           string   `xml:"FONT,attr"`
	FONTSIZE       string   `xml:"FONTSIZE,attr"`
	PARENT         string   `xml:"PARENT,attr"`
	CPARENT        string   `xml:"CPARENT,attr"`
	ALIGN          string   `xml:"ALIGN,attr"`
	OpticalMargins string   `xml:"OpticalMargins,attr"`
}

type ITEXT struct {
	XMLName  xml.Name `xml:"ITEXT"`
	Text     string   `xml:",chardata"`
	FONT     string   `xml:"FONT,attr"`
	FONTSIZE string   `xml:"FONTSIZE,attr"`
	FCOLOR   string   `xml:"FCOLOR,attr"`
	CPARENT  string   `xml:"CPARENT,attr"`
	CH       string   `xml:"CH,attr"`
}

type Trail struct {
	Text                  string `xml:",chardata"`
	LINESP                string `xml:"LINESP,attr"`
	LINESPMode            string `xml:"LINESPMode,attr"`
	ParagraphEffectOffset string `xml:"ParagraphEffectOffset,attr"`
	ParagraphEffectIndent string `xml:"ParagraphEffectIndent,attr"`
	DROP                  string `xml:"DROP,attr"`
	Bullet                string `xml:"Bullet,attr"`
	BulletStr             string `xml:"BulletStr,attr"`
	Numeration            string `xml:"Numeration,attr"`
	ALIGN                 string `xml:"ALIGN,attr"`
	OpticalMargins        string `xml:"OpticalMargins,attr"`
}

// readScribusFile reads an existing Scribus file from path and
// returns ScribusDocument, error
func NewScribusDocumentFromFile(path string) (Document, error) {
	var scribusDocument Document
	xmlFile, err := os.Open(path)
	if err != nil {
		return scribusDocument, err
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	err = xml.Unmarshal(byteValue, &scribusDocument)
	if err != nil {
		return scribusDocument, err
	}
	return scribusDocument, nil
}

// writeScribusFile writes out a ScribusDocument to disk at path and returns error
func (scribusDocument Document) WriteScribusFile(path string) error {
	if xmlstring, err := xml.MarshalIndent(scribusDocument, "", "    "); err == nil {
		xmlstring = []byte(xml.Header + strings.Replace(string(xmlstring), "&#xA;", "", -1)) // FIXME: https://forum.golangbridge.org/t/read-xml-change-values-write-back-crippled-file/16253/4
		err = ioutil.WriteFile(path, xmlstring, 0644)
		return err
	} else {
		return err
	}
}

// ChangeText changes the text of an ITEXT
func (itext *ITEXT) ChangeText(text string) {
	itext.CH = text
}

// GetPageObjectsWithText returns pointers to the PAGEOBJECTs with the text in question
func (doc DOCUMENT) GetPageObjectsWithText(text string) []*PAGEOBJECT {
	var pos []*PAGEOBJECT
	for i := range doc.PAGEOBJECT {
		for j := range doc.PAGEOBJECT[i].StoryText.ITEXT {
			if doc.PAGEOBJECT[i].StoryText.ITEXT[j].CH == text {
				pos = append(pos, &doc.PAGEOBJECT[i])
				return pos
			}
		}
	}
	return pos
}

// GetPageObjectByName returns a pointer to the PAGEOBJECT with the given name
// (press F2 in Scribus). Scribus makes sure that these names are unique
func (doc DOCUMENT) GetPageObjectByName(name string) *PAGEOBJECT {
	var po *PAGEOBJECT
	for i := range doc.PAGEOBJECT {
		// TODO: Check whether it has an ANNAME
		if doc.PAGEOBJECT[i].ANNAME == name {
			return &doc.PAGEOBJECT[i]
		}
	}
	return po
}

// GetPageObjectsThatAreType2 returns all images/pictures
func (doc DOCUMENT) GetPageObjectsThatAreType2() []*PAGEOBJECT {
	var pos []*PAGEOBJECT
	for i := range doc.PAGEOBJECT {
		if doc.PAGEOBJECT[i].PTYPE == "2" {
			pos = append(pos, &doc.PAGEOBJECT[i])
		}
	}
	return pos
}

// MovePageObject moves the i'th PAGEOBJECT to the supplied x and y position
func (po *PAGEOBJECT) MovePageObject(i int, xpos int, ypos int) {
	po.XPOS = strconv.Itoa(xpos)
	po.YPOS = strconv.Itoa(ypos)
}

// TODO: ChangeBulletPointsOfPageObject changes the bullet points of of the StoryText
// to the contents of a []string
// We should probably read the first para tag in a StoryText tag that has a BulletStr property, and copy that
func (st *StoryText) ChangeBulletPoints(texts []string) {

	// Get the first ITEXT and the first Para and use them as templates for the ones we are creating

	if len(st.ITEXT) < 1 {
		fmt.Println("ChangeBulletPoints: No ITEXT available")
		return
	}
	templateItext := st.ITEXT[0]
	if len(st.Para) < 1 {
		fmt.Println("ChangeBulletPoints: No Para available")
		return
	}
	templatePara := st.Para[0]

	// Clear the pre-existing ITEXTs and PARAs
	st.ITEXT = nil
	st.Para = nil

	var bulletGroups []StoryTextSpan

	for _, text := range texts {
		templateItext.CH = text
		bulletGroup := StoryTextSpan{
			ITEXT: templateItext,
			Para:  templatePara,
		}
		bulletGroups = append(bulletGroups, bulletGroup)
	}
	st.StoryTextSpan = bulletGroups
}

// TODO: ChangeTextParagraphs changes the paragraphs of the StoryText
// to the contents of a []string
// Crude fix sets the line spacing on the para object (FIXME)
// TODO: Be also able to do this without having
// to supply new text
func (st *StoryText) ChangeTextParagraphs(texts []string) {

	templateDefaultStyle := st.DefaultStyle

	if len(st.ITEXT) < 1 {
		fmt.Println("CrudeFix: No ITEXT available")
		return
	}
	templateItext := st.ITEXT[0]

	templatePara := Para{}
	templatePara.LINESP = templateDefaultStyle.LINESP // Magic; why is this needed?

	templateTrail := st.Trail

	// Clear the pre-existing ITEXTs and PARAs
	st.DefaultStyle = DefaultStyle{}
	st.ITEXT = nil
	st.Para = nil
	st.Trail = Trail{}

	var bulletGroups []StoryTextSpan

	for _, text := range texts {
		templateItext.CH = text
		bulletGroup := StoryTextSpan{
			DefaultStyle: templateDefaultStyle,
			ITEXT:        templateItext,
			Para:         templatePara,
			Trail:        templateTrail,
		}
		bulletGroups = append(bulletGroups, bulletGroup)
	}
	st.StoryTextSpan = bulletGroups
}

// TODO: ChangePictureOfPageObject changes the picture of the i'th PAGEOBJECT
// by changing its PFILE to path, returns error
func (scribusDocument *Document) xxxChangePictureOfPageObject(i int, path string) error {

	return nil
}
