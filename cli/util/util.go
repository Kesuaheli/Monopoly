package util

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/Kesuaheli/monopoly/lang"
	"golang.org/x/text/language"
)

var (
	SelectedLanguage language.Tag
)

// SelectableInput propts the user to select one element of the given slice.
//
// skipOne determines if the propt should be skipped if all contains only one element.
// afterSelection, if not nil, is executed immediately after the user entered their selection and
// before the confirmation message. If afterSelection returns true the confirmation message is
// skipped.
func SelectableInput[T any](head string, all []T, skipOne bool, afterSelection func(selected T, i int) bool) (T, int) {
	if len(all) == 0 {
		var t T
		return t, -1
	} else if skipOne && len(all) == 1 {
		return all[0], 0
	}

	selection := strings.Builder{}
	selection.WriteString(fmt.Sprintf(lang.MustLocalize("cli.input.choose", SelectedLanguage), head))
	selection.WriteByte('\n')
	digitsInAll := int(math.Floor(math.Log10(float64(len(all))))) + 1
	for n, item := range all {
		selection.WriteString(fmt.Sprintf("- [%*d] %s\n", digitsInAll, n+1, lang.LocalizeInterface(&item, SelectedLanguage)))
	}
	fmt.Printf(selection.String())

	selectedNum := NumberInput(1, len(all)) - 1
	selected := all[selectedNum]
	if afterSelection != nil && afterSelection(selected, selectedNum) {
		return selected, selectedNum
	}
	fmt.Printf(lang.MustLocalize("cli.input.selected", SelectedLanguage)+"\n", fmt.Sprintf("%s", lang.LocalizeInterface(selected, SelectedLanguage)))
	return selected, selectedNum
}

// NumberInput propts the user to enter a number in the given range (both min and max are inclusive)
func NumberInput(min, max int) (selected int) {
	for {
		fmt.Printf(lang.MustLocalize("cli.input.select", SelectedLanguage), min, max)
		var bScan []byte
		fmt.Scan(&bScan)
		n, err := strconv.ParseInt(string(bScan), 10, 0)

		if err == nil && n >= int64(min) && n <= int64(max) {
			return int(n)
		}
	}
}

// SliceDeleteElements removes every element from base that is in delete.
//
// When SliceDeleteElements removes m elements, it might not modify the elements
// base[len(base)-m:len(base)]. If those elements contain pointers you might consider zeroing those
// elements so that objects they reference can be garbage collected.
func SliceDeleteElements[S ~[]E, E comparable](base, delete S) S {
	return slices.DeleteFunc(base, func(v E) bool {
		return slices.Contains(delete, v)
	})
}

// ToUpperFirst returns s with only the first character mapped to their upper case.
func ToUpperFirst(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}
