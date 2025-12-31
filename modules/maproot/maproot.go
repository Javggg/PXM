package maproot

import (
	"encoding/xml"
	"pxm/modules/actions"
	"pxm/modules/authors"
	"pxm/modules/constants"
	"pxm/modules/filters"
	"pxm/modules/regions"
	"pxm/modules/replacements"
	"pxm/modules/variables"
)

type Map struct {
	XMLName          xml.Name                   `xml:"map"`
	Proto            string                     `xml:"proto,attr"`
	MinServerVersion *string                    `xml:"min-server-version,attr"`
	MaxServerVersion *string                    `xml:"max-server-version,attr"`
	Internal         *string                    `xml:"internal,attr"`
	Constants        *constants.Constants       `xml:"constants,omitempty"`
	Authors          *authors.Authors           `xml:"authors,omitempty"`
	Contributors     *authors.Contributors      `xml:"contributors,omitempty"`
	Regions          *regions.Regions           `xml:"regions,omitempty"`
	Filters          *filters.Filters           `xml:"filters,omitempty"`
	Variables        *variables.Variables       `xml:"variables,omitempty"`
	Replacements     *replacements.Replacements `xml:"replacements,omitempty"`
	Actions          *actions.Actions           `xml:"actions,omitempty"`
}
