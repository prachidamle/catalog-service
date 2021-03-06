package parse

import (
	"strconv"
	"strings"
)

func TemplateURLPath(path string) (string, string, string, int, bool) {
	pathSplit := strings.Split(path, ":")
	switch len(pathSplit) {
	case 2:
		catalog := pathSplit[0]
		template := pathSplit[1]
		templateSplit := strings.Split(template, "*")
		templateBase := ""
		switch len(templateSplit) {
		case 1:
			template = templateSplit[0]
		case 2:
			templateBase = templateSplit[0]
			template = templateSplit[1]
		default:
			return "", "", "", 0, false
		}
		return catalog, template, templateBase, -1, true
	case 3:
		catalog := pathSplit[0]
		template := pathSplit[1]
		revision, err := strconv.Atoi(pathSplit[2])
		if err != nil {
			return "", "", "", 0, false
		}
		templateSplit := strings.Split(template, "*")
		templateBase := ""
		switch len(templateSplit) {
		case 1:
			template = templateSplit[0]
		case 2:
			templateBase = templateSplit[0]
			template = templateSplit[1]
		default:
			return "", "", "", 0, false
		}
		return catalog, template, templateBase, revision, true
	default:
		return "", "", "", 0, false
	}
}

func TemplatePath(path string) (string, string, bool) {
	split := strings.Split(path, "/")
	if len(split) < 2 {
		return "", "", false
	}

	base := ""
	dirSplit := strings.SplitN(split[0], "-", 2)
	if len(dirSplit) > 1 {
		base = dirSplit[0]
	}

	return base, split[1], true
}

func VersionPath(path string) (string, string, int, bool) {
	base, template, parsedCorrectly := TemplatePath(path)
	if !parsedCorrectly {
		return "", "", 0, false
	}

	split := strings.Split(path, "/")
	if len(split) < 3 {
		return "", "", 0, false
	}

	revision, err := strconv.Atoi(split[2])
	if err != nil {
		return "", "", 0, false
	}

	return base, template, revision, true
}
