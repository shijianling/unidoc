/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */
/*
 * The embedded character metrics specified in this file are distributed under the terms listed in
 * ./afms/MustRead.html.
 */

package fonts

import (
	"github.com/unidoc/unidoc/pdf/core"
	"github.com/unidoc/unidoc/pdf/model/textencoding"
)

// fontHelveticaBoldOblique represents the Helvetica-BoldOblique font.
// This is a built-in font and it is assumed that every reader has access to it.
type fontHelveticaBoldOblique struct {
	encoder textencoding.TextEncoder
}

// NewFontHelveticaBoldOblique returns a new instance of the font with a default encoder set (WinAnsiEncoding).
func NewFontHelveticaBoldOblique() fontHelveticaBoldOblique {
	font := fontHelveticaBoldOblique{}
	font.encoder = textencoding.NewWinAnsiTextEncoder() // Default
	return font
}

// Encoder returns the font's text encoder.
func (font fontHelveticaBoldOblique) Encoder() textencoding.TextEncoder {
	return font.encoder
}

// SetEncoder sets the font's text encoder.
func (font fontHelveticaBoldOblique) SetEncoder(encoder textencoding.TextEncoder) {
	font.encoder = encoder
}

// GetGlyphCharMetrics returns character metrics for a given glyph.
func (font fontHelveticaBoldOblique) GetGlyphCharMetrics(glyph string) (CharMetrics, bool) {
	metrics, has := helveticaBoldObliqueCharMetrics[glyph]
	if !has {
		return metrics, false
	}

	return metrics, true
}

// ToPdfObject returns a primitive PDF object representation of the font.
func (font fontHelveticaBoldOblique) ToPdfObject() core.PdfObject {
	obj := &core.PdfIndirectObject{}

	fontDict := core.MakeDict()
	fontDict.Set("Type", core.MakeName("Font"))
	fontDict.Set("Subtype", core.MakeName("Type1"))
	fontDict.Set("BaseFont", core.MakeName("Helvetica-BoldOblique"))
	fontDict.Set("Encoding", font.encoder.ToPdfObject())

	obj.PdfObject = fontDict
	return obj
}

// Helvetica-BoldOblique font metics loaded from afms/Helvetica-BoldOblique.afm.  See afms/MustRead.html for license information.
var helveticaBoldObliqueCharMetrics = map[string]CharMetrics{
	"A":              {GlyphName: "A", Wx: 722.000000, Wy: 0.000000},
	"AE":             {GlyphName: "AE", Wx: 1000.000000, Wy: 0.000000},
	"Aacute":         {GlyphName: "Aacute", Wx: 722.000000, Wy: 0.000000},
	"Abreve":         {GlyphName: "Abreve", Wx: 722.000000, Wy: 0.000000},
	"Acircumflex":    {GlyphName: "Acircumflex", Wx: 722.000000, Wy: 0.000000},
	"Adieresis":      {GlyphName: "Adieresis", Wx: 722.000000, Wy: 0.000000},
	"Agrave":         {GlyphName: "Agrave", Wx: 722.000000, Wy: 0.000000},
	"Amacron":        {GlyphName: "Amacron", Wx: 722.000000, Wy: 0.000000},
	"Aogonek":        {GlyphName: "Aogonek", Wx: 722.000000, Wy: 0.000000},
	"Aring":          {GlyphName: "Aring", Wx: 722.000000, Wy: 0.000000},
	"Atilde":         {GlyphName: "Atilde", Wx: 722.000000, Wy: 0.000000},
	"B":              {GlyphName: "B", Wx: 722.000000, Wy: 0.000000},
	"C":              {GlyphName: "C", Wx: 722.000000, Wy: 0.000000},
	"Cacute":         {GlyphName: "Cacute", Wx: 722.000000, Wy: 0.000000},
	"Ccaron":         {GlyphName: "Ccaron", Wx: 722.000000, Wy: 0.000000},
	"Ccedilla":       {GlyphName: "Ccedilla", Wx: 722.000000, Wy: 0.000000},
	"D":              {GlyphName: "D", Wx: 722.000000, Wy: 0.000000},
	"Dcaron":         {GlyphName: "Dcaron", Wx: 722.000000, Wy: 0.000000},
	"Dcroat":         {GlyphName: "Dcroat", Wx: 722.000000, Wy: 0.000000},
	"Delta":          {GlyphName: "Delta", Wx: 612.000000, Wy: 0.000000},
	"E":              {GlyphName: "E", Wx: 667.000000, Wy: 0.000000},
	"Eacute":         {GlyphName: "Eacute", Wx: 667.000000, Wy: 0.000000},
	"Ecaron":         {GlyphName: "Ecaron", Wx: 667.000000, Wy: 0.000000},
	"Ecircumflex":    {GlyphName: "Ecircumflex", Wx: 667.000000, Wy: 0.000000},
	"Edieresis":      {GlyphName: "Edieresis", Wx: 667.000000, Wy: 0.000000},
	"Edotaccent":     {GlyphName: "Edotaccent", Wx: 667.000000, Wy: 0.000000},
	"Egrave":         {GlyphName: "Egrave", Wx: 667.000000, Wy: 0.000000},
	"Emacron":        {GlyphName: "Emacron", Wx: 667.000000, Wy: 0.000000},
	"Eogonek":        {GlyphName: "Eogonek", Wx: 667.000000, Wy: 0.000000},
	"Eth":            {GlyphName: "Eth", Wx: 722.000000, Wy: 0.000000},
	"Euro":           {GlyphName: "Euro", Wx: 556.000000, Wy: 0.000000},
	"F":              {GlyphName: "F", Wx: 611.000000, Wy: 0.000000},
	"G":              {GlyphName: "G", Wx: 778.000000, Wy: 0.000000},
	"Gbreve":         {GlyphName: "Gbreve", Wx: 778.000000, Wy: 0.000000},
	"Gcommaaccent":   {GlyphName: "Gcommaaccent", Wx: 778.000000, Wy: 0.000000},
	"H":              {GlyphName: "H", Wx: 722.000000, Wy: 0.000000},
	"I":              {GlyphName: "I", Wx: 278.000000, Wy: 0.000000},
	"Iacute":         {GlyphName: "Iacute", Wx: 278.000000, Wy: 0.000000},
	"Icircumflex":    {GlyphName: "Icircumflex", Wx: 278.000000, Wy: 0.000000},
	"Idieresis":      {GlyphName: "Idieresis", Wx: 278.000000, Wy: 0.000000},
	"Idotaccent":     {GlyphName: "Idotaccent", Wx: 278.000000, Wy: 0.000000},
	"Igrave":         {GlyphName: "Igrave", Wx: 278.000000, Wy: 0.000000},
	"Imacron":        {GlyphName: "Imacron", Wx: 278.000000, Wy: 0.000000},
	"Iogonek":        {GlyphName: "Iogonek", Wx: 278.000000, Wy: 0.000000},
	"J":              {GlyphName: "J", Wx: 556.000000, Wy: 0.000000},
	"K":              {GlyphName: "K", Wx: 722.000000, Wy: 0.000000},
	"Kcommaaccent":   {GlyphName: "Kcommaaccent", Wx: 722.000000, Wy: 0.000000},
	"L":              {GlyphName: "L", Wx: 611.000000, Wy: 0.000000},
	"Lacute":         {GlyphName: "Lacute", Wx: 611.000000, Wy: 0.000000},
	"Lcaron":         {GlyphName: "Lcaron", Wx: 611.000000, Wy: 0.000000},
	"Lcommaaccent":   {GlyphName: "Lcommaaccent", Wx: 611.000000, Wy: 0.000000},
	"Lslash":         {GlyphName: "Lslash", Wx: 611.000000, Wy: 0.000000},
	"M":              {GlyphName: "M", Wx: 833.000000, Wy: 0.000000},
	"N":              {GlyphName: "N", Wx: 722.000000, Wy: 0.000000},
	"Nacute":         {GlyphName: "Nacute", Wx: 722.000000, Wy: 0.000000},
	"Ncaron":         {GlyphName: "Ncaron", Wx: 722.000000, Wy: 0.000000},
	"Ncommaaccent":   {GlyphName: "Ncommaaccent", Wx: 722.000000, Wy: 0.000000},
	"Ntilde":         {GlyphName: "Ntilde", Wx: 722.000000, Wy: 0.000000},
	"O":              {GlyphName: "O", Wx: 778.000000, Wy: 0.000000},
	"OE":             {GlyphName: "OE", Wx: 1000.000000, Wy: 0.000000},
	"Oacute":         {GlyphName: "Oacute", Wx: 778.000000, Wy: 0.000000},
	"Ocircumflex":    {GlyphName: "Ocircumflex", Wx: 778.000000, Wy: 0.000000},
	"Odieresis":      {GlyphName: "Odieresis", Wx: 778.000000, Wy: 0.000000},
	"Ograve":         {GlyphName: "Ograve", Wx: 778.000000, Wy: 0.000000},
	"Ohungarumlaut":  {GlyphName: "Ohungarumlaut", Wx: 778.000000, Wy: 0.000000},
	"Omacron":        {GlyphName: "Omacron", Wx: 778.000000, Wy: 0.000000},
	"Oslash":         {GlyphName: "Oslash", Wx: 778.000000, Wy: 0.000000},
	"Otilde":         {GlyphName: "Otilde", Wx: 778.000000, Wy: 0.000000},
	"P":              {GlyphName: "P", Wx: 667.000000, Wy: 0.000000},
	"Q":              {GlyphName: "Q", Wx: 778.000000, Wy: 0.000000},
	"R":              {GlyphName: "R", Wx: 722.000000, Wy: 0.000000},
	"Racute":         {GlyphName: "Racute", Wx: 722.000000, Wy: 0.000000},
	"Rcaron":         {GlyphName: "Rcaron", Wx: 722.000000, Wy: 0.000000},
	"Rcommaaccent":   {GlyphName: "Rcommaaccent", Wx: 722.000000, Wy: 0.000000},
	"S":              {GlyphName: "S", Wx: 667.000000, Wy: 0.000000},
	"Sacute":         {GlyphName: "Sacute", Wx: 667.000000, Wy: 0.000000},
	"Scaron":         {GlyphName: "Scaron", Wx: 667.000000, Wy: 0.000000},
	"Scedilla":       {GlyphName: "Scedilla", Wx: 667.000000, Wy: 0.000000},
	"Scommaaccent":   {GlyphName: "Scommaaccent", Wx: 667.000000, Wy: 0.000000},
	"T":              {GlyphName: "T", Wx: 611.000000, Wy: 0.000000},
	"Tcaron":         {GlyphName: "Tcaron", Wx: 611.000000, Wy: 0.000000},
	"Tcommaaccent":   {GlyphName: "Tcommaaccent", Wx: 611.000000, Wy: 0.000000},
	"Thorn":          {GlyphName: "Thorn", Wx: 667.000000, Wy: 0.000000},
	"U":              {GlyphName: "U", Wx: 722.000000, Wy: 0.000000},
	"Uacute":         {GlyphName: "Uacute", Wx: 722.000000, Wy: 0.000000},
	"Ucircumflex":    {GlyphName: "Ucircumflex", Wx: 722.000000, Wy: 0.000000},
	"Udieresis":      {GlyphName: "Udieresis", Wx: 722.000000, Wy: 0.000000},
	"Ugrave":         {GlyphName: "Ugrave", Wx: 722.000000, Wy: 0.000000},
	"Uhungarumlaut":  {GlyphName: "Uhungarumlaut", Wx: 722.000000, Wy: 0.000000},
	"Umacron":        {GlyphName: "Umacron", Wx: 722.000000, Wy: 0.000000},
	"Uogonek":        {GlyphName: "Uogonek", Wx: 722.000000, Wy: 0.000000},
	"Uring":          {GlyphName: "Uring", Wx: 722.000000, Wy: 0.000000},
	"V":              {GlyphName: "V", Wx: 667.000000, Wy: 0.000000},
	"W":              {GlyphName: "W", Wx: 944.000000, Wy: 0.000000},
	"X":              {GlyphName: "X", Wx: 667.000000, Wy: 0.000000},
	"Y":              {GlyphName: "Y", Wx: 667.000000, Wy: 0.000000},
	"Yacute":         {GlyphName: "Yacute", Wx: 667.000000, Wy: 0.000000},
	"Ydieresis":      {GlyphName: "Ydieresis", Wx: 667.000000, Wy: 0.000000},
	"Z":              {GlyphName: "Z", Wx: 611.000000, Wy: 0.000000},
	"Zacute":         {GlyphName: "Zacute", Wx: 611.000000, Wy: 0.000000},
	"Zcaron":         {GlyphName: "Zcaron", Wx: 611.000000, Wy: 0.000000},
	"Zdotaccent":     {GlyphName: "Zdotaccent", Wx: 611.000000, Wy: 0.000000},
	"a":              {GlyphName: "a", Wx: 556.000000, Wy: 0.000000},
	"aacute":         {GlyphName: "aacute", Wx: 556.000000, Wy: 0.000000},
	"abreve":         {GlyphName: "abreve", Wx: 556.000000, Wy: 0.000000},
	"acircumflex":    {GlyphName: "acircumflex", Wx: 556.000000, Wy: 0.000000},
	"acute":          {GlyphName: "acute", Wx: 333.000000, Wy: 0.000000},
	"adieresis":      {GlyphName: "adieresis", Wx: 556.000000, Wy: 0.000000},
	"ae":             {GlyphName: "ae", Wx: 889.000000, Wy: 0.000000},
	"agrave":         {GlyphName: "agrave", Wx: 556.000000, Wy: 0.000000},
	"amacron":        {GlyphName: "amacron", Wx: 556.000000, Wy: 0.000000},
	"ampersand":      {GlyphName: "ampersand", Wx: 722.000000, Wy: 0.000000},
	"aogonek":        {GlyphName: "aogonek", Wx: 556.000000, Wy: 0.000000},
	"aring":          {GlyphName: "aring", Wx: 556.000000, Wy: 0.000000},
	"asciicircum":    {GlyphName: "asciicircum", Wx: 584.000000, Wy: 0.000000},
	"asciitilde":     {GlyphName: "asciitilde", Wx: 584.000000, Wy: 0.000000},
	"asterisk":       {GlyphName: "asterisk", Wx: 389.000000, Wy: 0.000000},
	"at":             {GlyphName: "at", Wx: 975.000000, Wy: 0.000000},
	"atilde":         {GlyphName: "atilde", Wx: 556.000000, Wy: 0.000000},
	"b":              {GlyphName: "b", Wx: 611.000000, Wy: 0.000000},
	"backslash":      {GlyphName: "backslash", Wx: 278.000000, Wy: 0.000000},
	"bar":            {GlyphName: "bar", Wx: 280.000000, Wy: 0.000000},
	"braceleft":      {GlyphName: "braceleft", Wx: 389.000000, Wy: 0.000000},
	"braceright":     {GlyphName: "braceright", Wx: 389.000000, Wy: 0.000000},
	"bracketleft":    {GlyphName: "bracketleft", Wx: 333.000000, Wy: 0.000000},
	"bracketright":   {GlyphName: "bracketright", Wx: 333.000000, Wy: 0.000000},
	"breve":          {GlyphName: "breve", Wx: 333.000000, Wy: 0.000000},
	"brokenbar":      {GlyphName: "brokenbar", Wx: 280.000000, Wy: 0.000000},
	"bullet":         {GlyphName: "bullet", Wx: 350.000000, Wy: 0.000000},
	"c":              {GlyphName: "c", Wx: 556.000000, Wy: 0.000000},
	"cacute":         {GlyphName: "cacute", Wx: 556.000000, Wy: 0.000000},
	"caron":          {GlyphName: "caron", Wx: 333.000000, Wy: 0.000000},
	"ccaron":         {GlyphName: "ccaron", Wx: 556.000000, Wy: 0.000000},
	"ccedilla":       {GlyphName: "ccedilla", Wx: 556.000000, Wy: 0.000000},
	"cedilla":        {GlyphName: "cedilla", Wx: 333.000000, Wy: 0.000000},
	"cent":           {GlyphName: "cent", Wx: 556.000000, Wy: 0.000000},
	"circumflex":     {GlyphName: "circumflex", Wx: 333.000000, Wy: 0.000000},
	"colon":          {GlyphName: "colon", Wx: 333.000000, Wy: 0.000000},
	"comma":          {GlyphName: "comma", Wx: 278.000000, Wy: 0.000000},
	"commaaccent":    {GlyphName: "commaaccent", Wx: 250.000000, Wy: 0.000000},
	"copyright":      {GlyphName: "copyright", Wx: 737.000000, Wy: 0.000000},
	"currency":       {GlyphName: "currency", Wx: 556.000000, Wy: 0.000000},
	"d":              {GlyphName: "d", Wx: 611.000000, Wy: 0.000000},
	"dagger":         {GlyphName: "dagger", Wx: 556.000000, Wy: 0.000000},
	"daggerdbl":      {GlyphName: "daggerdbl", Wx: 556.000000, Wy: 0.000000},
	"dcaron":         {GlyphName: "dcaron", Wx: 743.000000, Wy: 0.000000},
	"dcroat":         {GlyphName: "dcroat", Wx: 611.000000, Wy: 0.000000},
	"degree":         {GlyphName: "degree", Wx: 400.000000, Wy: 0.000000},
	"dieresis":       {GlyphName: "dieresis", Wx: 333.000000, Wy: 0.000000},
	"divide":         {GlyphName: "divide", Wx: 584.000000, Wy: 0.000000},
	"dollar":         {GlyphName: "dollar", Wx: 556.000000, Wy: 0.000000},
	"dotaccent":      {GlyphName: "dotaccent", Wx: 333.000000, Wy: 0.000000},
	"dotlessi":       {GlyphName: "dotlessi", Wx: 278.000000, Wy: 0.000000},
	"e":              {GlyphName: "e", Wx: 556.000000, Wy: 0.000000},
	"eacute":         {GlyphName: "eacute", Wx: 556.000000, Wy: 0.000000},
	"ecaron":         {GlyphName: "ecaron", Wx: 556.000000, Wy: 0.000000},
	"ecircumflex":    {GlyphName: "ecircumflex", Wx: 556.000000, Wy: 0.000000},
	"edieresis":      {GlyphName: "edieresis", Wx: 556.000000, Wy: 0.000000},
	"edotaccent":     {GlyphName: "edotaccent", Wx: 556.000000, Wy: 0.000000},
	"egrave":         {GlyphName: "egrave", Wx: 556.000000, Wy: 0.000000},
	"eight":          {GlyphName: "eight", Wx: 556.000000, Wy: 0.000000},
	"ellipsis":       {GlyphName: "ellipsis", Wx: 1000.000000, Wy: 0.000000},
	"emacron":        {GlyphName: "emacron", Wx: 556.000000, Wy: 0.000000},
	"emdash":         {GlyphName: "emdash", Wx: 1000.000000, Wy: 0.000000},
	"endash":         {GlyphName: "endash", Wx: 556.000000, Wy: 0.000000},
	"eogonek":        {GlyphName: "eogonek", Wx: 556.000000, Wy: 0.000000},
	"equal":          {GlyphName: "equal", Wx: 584.000000, Wy: 0.000000},
	"eth":            {GlyphName: "eth", Wx: 611.000000, Wy: 0.000000},
	"exclam":         {GlyphName: "exclam", Wx: 333.000000, Wy: 0.000000},
	"exclamdown":     {GlyphName: "exclamdown", Wx: 333.000000, Wy: 0.000000},
	"f":              {GlyphName: "f", Wx: 333.000000, Wy: 0.000000},
	"fi":             {GlyphName: "fi", Wx: 611.000000, Wy: 0.000000},
	"five":           {GlyphName: "five", Wx: 556.000000, Wy: 0.000000},
	"fl":             {GlyphName: "fl", Wx: 611.000000, Wy: 0.000000},
	"florin":         {GlyphName: "florin", Wx: 556.000000, Wy: 0.000000},
	"four":           {GlyphName: "four", Wx: 556.000000, Wy: 0.000000},
	"fraction":       {GlyphName: "fraction", Wx: 167.000000, Wy: 0.000000},
	"g":              {GlyphName: "g", Wx: 611.000000, Wy: 0.000000},
	"gbreve":         {GlyphName: "gbreve", Wx: 611.000000, Wy: 0.000000},
	"gcommaaccent":   {GlyphName: "gcommaaccent", Wx: 611.000000, Wy: 0.000000},
	"germandbls":     {GlyphName: "germandbls", Wx: 611.000000, Wy: 0.000000},
	"grave":          {GlyphName: "grave", Wx: 333.000000, Wy: 0.000000},
	"greater":        {GlyphName: "greater", Wx: 584.000000, Wy: 0.000000},
	"greaterequal":   {GlyphName: "greaterequal", Wx: 549.000000, Wy: 0.000000},
	"guillemotleft":  {GlyphName: "guillemotleft", Wx: 556.000000, Wy: 0.000000},
	"guillemotright": {GlyphName: "guillemotright", Wx: 556.000000, Wy: 0.000000},
	"guilsinglleft":  {GlyphName: "guilsinglleft", Wx: 333.000000, Wy: 0.000000},
	"guilsinglright": {GlyphName: "guilsinglright", Wx: 333.000000, Wy: 0.000000},
	"h":              {GlyphName: "h", Wx: 611.000000, Wy: 0.000000},
	"hungarumlaut":   {GlyphName: "hungarumlaut", Wx: 333.000000, Wy: 0.000000},
	"hyphen":         {GlyphName: "hyphen", Wx: 333.000000, Wy: 0.000000},
	"i":              {GlyphName: "i", Wx: 278.000000, Wy: 0.000000},
	"iacute":         {GlyphName: "iacute", Wx: 278.000000, Wy: 0.000000},
	"icircumflex":    {GlyphName: "icircumflex", Wx: 278.000000, Wy: 0.000000},
	"idieresis":      {GlyphName: "idieresis", Wx: 278.000000, Wy: 0.000000},
	"igrave":         {GlyphName: "igrave", Wx: 278.000000, Wy: 0.000000},
	"imacron":        {GlyphName: "imacron", Wx: 278.000000, Wy: 0.000000},
	"iogonek":        {GlyphName: "iogonek", Wx: 278.000000, Wy: 0.000000},
	"j":              {GlyphName: "j", Wx: 278.000000, Wy: 0.000000},
	"k":              {GlyphName: "k", Wx: 556.000000, Wy: 0.000000},
	"kcommaaccent":   {GlyphName: "kcommaaccent", Wx: 556.000000, Wy: 0.000000},
	"l":              {GlyphName: "l", Wx: 278.000000, Wy: 0.000000},
	"lacute":         {GlyphName: "lacute", Wx: 278.000000, Wy: 0.000000},
	"lcaron":         {GlyphName: "lcaron", Wx: 400.000000, Wy: 0.000000},
	"lcommaaccent":   {GlyphName: "lcommaaccent", Wx: 278.000000, Wy: 0.000000},
	"less":           {GlyphName: "less", Wx: 584.000000, Wy: 0.000000},
	"lessequal":      {GlyphName: "lessequal", Wx: 549.000000, Wy: 0.000000},
	"logicalnot":     {GlyphName: "logicalnot", Wx: 584.000000, Wy: 0.000000},
	"lozenge":        {GlyphName: "lozenge", Wx: 494.000000, Wy: 0.000000},
	"lslash":         {GlyphName: "lslash", Wx: 278.000000, Wy: 0.000000},
	"m":              {GlyphName: "m", Wx: 889.000000, Wy: 0.000000},
	"macron":         {GlyphName: "macron", Wx: 333.000000, Wy: 0.000000},
	"minus":          {GlyphName: "minus", Wx: 584.000000, Wy: 0.000000},
	"mu":             {GlyphName: "mu", Wx: 611.000000, Wy: 0.000000},
	"multiply":       {GlyphName: "multiply", Wx: 584.000000, Wy: 0.000000},
	"n":              {GlyphName: "n", Wx: 611.000000, Wy: 0.000000},
	"nacute":         {GlyphName: "nacute", Wx: 611.000000, Wy: 0.000000},
	"ncaron":         {GlyphName: "ncaron", Wx: 611.000000, Wy: 0.000000},
	"ncommaaccent":   {GlyphName: "ncommaaccent", Wx: 611.000000, Wy: 0.000000},
	"nine":           {GlyphName: "nine", Wx: 556.000000, Wy: 0.000000},
	"notequal":       {GlyphName: "notequal", Wx: 549.000000, Wy: 0.000000},
	"ntilde":         {GlyphName: "ntilde", Wx: 611.000000, Wy: 0.000000},
	"numbersign":     {GlyphName: "numbersign", Wx: 556.000000, Wy: 0.000000},
	"o":              {GlyphName: "o", Wx: 611.000000, Wy: 0.000000},
	"oacute":         {GlyphName: "oacute", Wx: 611.000000, Wy: 0.000000},
	"ocircumflex":    {GlyphName: "ocircumflex", Wx: 611.000000, Wy: 0.000000},
	"odieresis":      {GlyphName: "odieresis", Wx: 611.000000, Wy: 0.000000},
	"oe":             {GlyphName: "oe", Wx: 944.000000, Wy: 0.000000},
	"ogonek":         {GlyphName: "ogonek", Wx: 333.000000, Wy: 0.000000},
	"ograve":         {GlyphName: "ograve", Wx: 611.000000, Wy: 0.000000},
	"ohungarumlaut":  {GlyphName: "ohungarumlaut", Wx: 611.000000, Wy: 0.000000},
	"omacron":        {GlyphName: "omacron", Wx: 611.000000, Wy: 0.000000},
	"one":            {GlyphName: "one", Wx: 556.000000, Wy: 0.000000},
	"onehalf":        {GlyphName: "onehalf", Wx: 834.000000, Wy: 0.000000},
	"onequarter":     {GlyphName: "onequarter", Wx: 834.000000, Wy: 0.000000},
	"onesuperior":    {GlyphName: "onesuperior", Wx: 333.000000, Wy: 0.000000},
	"ordfeminine":    {GlyphName: "ordfeminine", Wx: 370.000000, Wy: 0.000000},
	"ordmasculine":   {GlyphName: "ordmasculine", Wx: 365.000000, Wy: 0.000000},
	"oslash":         {GlyphName: "oslash", Wx: 611.000000, Wy: 0.000000},
	"otilde":         {GlyphName: "otilde", Wx: 611.000000, Wy: 0.000000},
	"p":              {GlyphName: "p", Wx: 611.000000, Wy: 0.000000},
	"paragraph":      {GlyphName: "paragraph", Wx: 556.000000, Wy: 0.000000},
	"parenleft":      {GlyphName: "parenleft", Wx: 333.000000, Wy: 0.000000},
	"parenright":     {GlyphName: "parenright", Wx: 333.000000, Wy: 0.000000},
	"partialdiff":    {GlyphName: "partialdiff", Wx: 494.000000, Wy: 0.000000},
	"percent":        {GlyphName: "percent", Wx: 889.000000, Wy: 0.000000},
	"period":         {GlyphName: "period", Wx: 278.000000, Wy: 0.000000},
	"periodcentered": {GlyphName: "periodcentered", Wx: 278.000000, Wy: 0.000000},
	"perthousand":    {GlyphName: "perthousand", Wx: 1000.000000, Wy: 0.000000},
	"plus":           {GlyphName: "plus", Wx: 584.000000, Wy: 0.000000},
	"plusminus":      {GlyphName: "plusminus", Wx: 584.000000, Wy: 0.000000},
	"q":              {GlyphName: "q", Wx: 611.000000, Wy: 0.000000},
	"question":       {GlyphName: "question", Wx: 611.000000, Wy: 0.000000},
	"questiondown":   {GlyphName: "questiondown", Wx: 611.000000, Wy: 0.000000},
	"quotedbl":       {GlyphName: "quotedbl", Wx: 474.000000, Wy: 0.000000},
	"quotedblbase":   {GlyphName: "quotedblbase", Wx: 500.000000, Wy: 0.000000},
	"quotedblleft":   {GlyphName: "quotedblleft", Wx: 500.000000, Wy: 0.000000},
	"quotedblright":  {GlyphName: "quotedblright", Wx: 500.000000, Wy: 0.000000},
	"quoteleft":      {GlyphName: "quoteleft", Wx: 278.000000, Wy: 0.000000},
	"quoteright":     {GlyphName: "quoteright", Wx: 278.000000, Wy: 0.000000},
	"quotesinglbase": {GlyphName: "quotesinglbase", Wx: 278.000000, Wy: 0.000000},
	"quotesingle":    {GlyphName: "quotesingle", Wx: 238.000000, Wy: 0.000000},
	"r":              {GlyphName: "r", Wx: 389.000000, Wy: 0.000000},
	"racute":         {GlyphName: "racute", Wx: 389.000000, Wy: 0.000000},
	"radical":        {GlyphName: "radical", Wx: 549.000000, Wy: 0.000000},
	"rcaron":         {GlyphName: "rcaron", Wx: 389.000000, Wy: 0.000000},
	"rcommaaccent":   {GlyphName: "rcommaaccent", Wx: 389.000000, Wy: 0.000000},
	"registered":     {GlyphName: "registered", Wx: 737.000000, Wy: 0.000000},
	"ring":           {GlyphName: "ring", Wx: 333.000000, Wy: 0.000000},
	"s":              {GlyphName: "s", Wx: 556.000000, Wy: 0.000000},
	"sacute":         {GlyphName: "sacute", Wx: 556.000000, Wy: 0.000000},
	"scaron":         {GlyphName: "scaron", Wx: 556.000000, Wy: 0.000000},
	"scedilla":       {GlyphName: "scedilla", Wx: 556.000000, Wy: 0.000000},
	"scommaaccent":   {GlyphName: "scommaaccent", Wx: 556.000000, Wy: 0.000000},
	"section":        {GlyphName: "section", Wx: 556.000000, Wy: 0.000000},
	"semicolon":      {GlyphName: "semicolon", Wx: 333.000000, Wy: 0.000000},
	"seven":          {GlyphName: "seven", Wx: 556.000000, Wy: 0.000000},
	"six":            {GlyphName: "six", Wx: 556.000000, Wy: 0.000000},
	"slash":          {GlyphName: "slash", Wx: 278.000000, Wy: 0.000000},
	"space":          {GlyphName: "space", Wx: 278.000000, Wy: 0.000000},
	"sterling":       {GlyphName: "sterling", Wx: 556.000000, Wy: 0.000000},
	"summation":      {GlyphName: "summation", Wx: 600.000000, Wy: 0.000000},
	"t":              {GlyphName: "t", Wx: 333.000000, Wy: 0.000000},
	"tcaron":         {GlyphName: "tcaron", Wx: 389.000000, Wy: 0.000000},
	"tcommaaccent":   {GlyphName: "tcommaaccent", Wx: 333.000000, Wy: 0.000000},
	"thorn":          {GlyphName: "thorn", Wx: 611.000000, Wy: 0.000000},
	"three":          {GlyphName: "three", Wx: 556.000000, Wy: 0.000000},
	"threequarters":  {GlyphName: "threequarters", Wx: 834.000000, Wy: 0.000000},
	"threesuperior":  {GlyphName: "threesuperior", Wx: 333.000000, Wy: 0.000000},
	"tilde":          {GlyphName: "tilde", Wx: 333.000000, Wy: 0.000000},
	"trademark":      {GlyphName: "trademark", Wx: 1000.000000, Wy: 0.000000},
	"two":            {GlyphName: "two", Wx: 556.000000, Wy: 0.000000},
	"twosuperior":    {GlyphName: "twosuperior", Wx: 333.000000, Wy: 0.000000},
	"u":              {GlyphName: "u", Wx: 611.000000, Wy: 0.000000},
	"uacute":         {GlyphName: "uacute", Wx: 611.000000, Wy: 0.000000},
	"ucircumflex":    {GlyphName: "ucircumflex", Wx: 611.000000, Wy: 0.000000},
	"udieresis":      {GlyphName: "udieresis", Wx: 611.000000, Wy: 0.000000},
	"ugrave":         {GlyphName: "ugrave", Wx: 611.000000, Wy: 0.000000},
	"uhungarumlaut":  {GlyphName: "uhungarumlaut", Wx: 611.000000, Wy: 0.000000},
	"umacron":        {GlyphName: "umacron", Wx: 611.000000, Wy: 0.000000},
	"underscore":     {GlyphName: "underscore", Wx: 556.000000, Wy: 0.000000},
	"uogonek":        {GlyphName: "uogonek", Wx: 611.000000, Wy: 0.000000},
	"uring":          {GlyphName: "uring", Wx: 611.000000, Wy: 0.000000},
	"v":              {GlyphName: "v", Wx: 556.000000, Wy: 0.000000},
	"w":              {GlyphName: "w", Wx: 778.000000, Wy: 0.000000},
	"x":              {GlyphName: "x", Wx: 556.000000, Wy: 0.000000},
	"y":              {GlyphName: "y", Wx: 556.000000, Wy: 0.000000},
	"yacute":         {GlyphName: "yacute", Wx: 556.000000, Wy: 0.000000},
	"ydieresis":      {GlyphName: "ydieresis", Wx: 556.000000, Wy: 0.000000},
	"yen":            {GlyphName: "yen", Wx: 556.000000, Wy: 0.000000},
	"z":              {GlyphName: "z", Wx: 500.000000, Wy: 0.000000},
	"zacute":         {GlyphName: "zacute", Wx: 500.000000, Wy: 0.000000},
	"zcaron":         {GlyphName: "zcaron", Wx: 500.000000, Wy: 0.000000},
	"zdotaccent":     {GlyphName: "zdotaccent", Wx: 500.000000, Wy: 0.000000},
	"zero":           {GlyphName: "zero", Wx: 556.000000, Wy: 0.000000},
}
