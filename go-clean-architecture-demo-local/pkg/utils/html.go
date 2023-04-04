package utils

import (
	"regexp"
	"strings"
)

type HtmlConvertToText struct {
	RegexpData    regexp.Regexp
	TextToReplace string
}

type HtmlConvertToTextList []HtmlConvertToText

func (htmlConvertToTextList *HtmlConvertToTextList) setRegexp(regexpData regexp.Regexp, textToReplace string) {

	htmlConvertToText := new(HtmlConvertToText)
	htmlConvertToText.RegexpData = regexpData
	htmlConvertToText.TextToReplace = textToReplace

	*htmlConvertToTextList = append(*htmlConvertToTextList, *htmlConvertToText)

}

func (htmlConvertToTextList *HtmlConvertToTextList) setRegexpHtmlToNewLineOnly() {
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`<\s*br\/*>`), "\n")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`<\s*a.*href="(.*?)".*>(.*?)<\/a>`), " $2 (Link->$1) ")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`<\s*\/*.+?>`), "\n")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(` {2,}`), " ")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`\n+\s*`), "\n\n")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`\n`), "<br />")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`(<br \/>)+`), "<br />")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`^(<br \/>)`), "")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`<br \/>$`), "")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`<br \/>`), "\\n")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`&nbsp;`), " ")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`\\n$`), "")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`\\n +$`), "")
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`\\n`), "\\n")
}

func ReplaceRegexpHtmlToNewLineOnly(htmlText string) string {
	htmlConvertToTextList := new(HtmlConvertToTextList)
	htmlConvertToTextList.setRegexpHtmlToNewLineOnly()
	for _, item := range *htmlConvertToTextList {
		htmlText = item.RegexpData.ReplaceAllString(htmlText, item.TextToReplace)
	}

	return htmlText
}

func (htmlConvertToTextList *HtmlConvertToTextList) setRegexpNewLineToText() {
	htmlConvertToTextList.setRegexp(*regexp.MustCompile(`\n`), "\\n")
}

func ReplaceRegexpNewLineToText(htmlText string) string {
	htmlConvertToTextList := new(HtmlConvertToTextList)
	htmlConvertToTextList.setRegexpNewLineToText()
	for _, item := range *htmlConvertToTextList {
		htmlText = item.RegexpData.ReplaceAllString(htmlText, item.TextToReplace)
	}

	return htmlText
}

func ReplaceTextNewLineToNewLine(text string) string {
	text = strings.ReplaceAll(text, "\\n", "\n")
	return text
}
