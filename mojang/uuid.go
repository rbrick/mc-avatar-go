package mojang

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	UuidPattern = regexp.MustCompile("^([a-fA-F0-9]{8})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{12})$")
)

type Uuid string

func (u Uuid) WithDashes() string {
	groups := UuidPattern.FindAllStringSubmatch(string(u), -1)
	if groups == nil {
		return string(u)
	}

	return strings.ToLower(fmt.Sprintf("%s-%s-%s-%s-%s", groups[0][1], groups[0][2], groups[0][3], groups[0][4], groups[0][5]))
}

func (u Uuid) WithoutDashes() string {
	groups := UuidPattern.FindAllStringSubmatch(string(u), -1)
	if groups == nil {
		return string(u)
	}

	return strings.ToLower(fmt.Sprintf("%s%s%s%s%s", groups[0][1], groups[0][2], groups[0][3], groups[0][4], groups[0][5]))
}
