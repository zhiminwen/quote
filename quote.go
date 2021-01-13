package quote

import (
	"bytes"
	"log"
	"regexp"
	"strings"
	"text/template"

	"github.com/MakeNowJust/heredoc"
	"github.com/Masterminds/sprig"
)

// Word return the string slice of the words
func Word(line string) []string {
	// return regexp.MustCompile(`\s+`).Split(line, -1)
	return strings.Fields(line)
}

// Line return the lines as a string slice, ignoring empty line, line as comment (started with // or #)
func Line(lines string) []string {
	var result []string
	appendPrevLine := false
	index := 0

	for _, l := range regexp.MustCompile(`\n`).Split(lines, -1) {
		nl := strings.TrimSpace(l)
		if nl == "" || strings.HasPrefix(nl, "//") || strings.HasPrefix(nl, "#") {
			continue
		}
		if appendPrevLine {
			result[index-1] = result[index-1] + "\n " + nl
			appendPrevLine = false
		} else {
			result = append(result, nl)
			index++
		}

		if strings.HasSuffix(nl, "\\") {
			appendPrevLine = true
		}

	}

	return result
}

// A wrapper using hereDoc
func HereDoc(text string) string {
	return heredoc.Doc(text)
}

// A wrapper using hereDocf
func HereDocf(raw string, args ...interface{}) string {
	return heredoc.Docf(raw, args...)
}

// join command line with either & or && or ','
func Cmd(raw string, joinChar string) string {
	return strings.Join(Line(raw), joinChar)
}

// join command line with either & or && or ',' using heredocf
func Cmdf(raw string, joinChar string, args ...interface{}) string {
	return strings.Join(Line(HereDocf(raw, args...)), joinChar)
}

// Create a Map[string]string based template. Call log.Fatalf if any error detected
func Template(text string, data map[string]string) string {
	t, err := template.New("myT").Funcs(sprig.TxtFuncMap()).Parse(text)
	if err != nil {
		log.Fatalf("Failed to create template:%v", err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		log.Fatalf("Failed to execute template:%v", err)
	}

	return buf.String()
}

func TemplateGeneric(text string, data map[string]interface{}) string {
	t, err := template.New("myT").Funcs(sprig.TxtFuncMap()).Parse(text)
	if err != nil {
		log.Fatalf("Failed to create template:%v", err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		log.Fatalf("Failed to execute template:%v", err)
	}

	return buf.String()
}

func CmdTemplate(text string, data map[string]string) string {
	str := Template(text, data)
	return Cmd(str, " && ")
}
