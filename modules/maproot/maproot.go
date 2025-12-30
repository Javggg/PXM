package maproot

import (
	"encoding/xml"
	"pxm/modules/authors"
	"pxm/modules/filters"
	"pxm/modules/regions"
)

type Map struct {
	XMLName          xml.Name              `xml:"map"`
	Proto            string                `xml:"proto,attr"`
	MinServerVersion *string               `xml:"min-server-version,attr"`
	MaxServerVersion *string               `xml:"max-server-version,attr"`
	Internal         *bool                 `xml:"internal,attr"`
	Authors          *authors.Authors      `xml:"authors,omitempty"`
	Contributors     *authors.Contributors `xml:"contributors,omitempty"`
	Regions          *regions.Regions      `xml:"regions,omitempty"`
	Filters          *filters.Filters      `xml:"filters,omitempty"`
}
