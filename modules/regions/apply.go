package regions

import (
	"pxm/modules/filters"
	"pxm/modules/kits"
)

type Apply struct {
	InlineRegion            *string                  `xml:"region,attr,omitempty"`
	InlineEnter             *string                  `xml:"enter,attr,omitempty"`
	InlineLeave             *string                  `xml:"leave,attr,omitempty"`
	InlineBlock             *string                  `xml:"block,attr,omitempty"`
	InlineBlockPlace        *string                  `xml:"block-place,attr,omitempty"`
	InlineBlockPlaceAgainst *string                  `xml:"block-place-against,attr,omitempty"`
	InlineBlockBreak        *string                  `xml:"block-break,attr,omitempty"`
	InlineBlockPhysics      *string                  `xml:"block-physics,attr,omitempty"`
	InlineUse               *string                  `xml:"use,attr,omitempty"`
	InlineKit               *string                  `xml:"kit,attr,omitempty"`
	InlineLendKit           *string                  `xml:"lend-kit,attr,omitempty"`
	InlineFilter            *string                  `xml:"filter,attr,omitempty"`
	Velocity                *string                  `xml:"velocity,attr,omitempty"`
	Message                 *string                  `xml:"message,attr,omitempty"`
	EarlyWarning            *string                  `xml:"early-warning,attr,omitempty"`
	Region                  *RegionContainer         `xml:"region,omitempty"`
	Enter                   *filters.FilterContainer `xml:"enter,omitempty"`
	Leave                   *filters.FilterContainer `xml:"leave,omitempty"`
	Block                   *filters.FilterContainer `xml:"block,omitempty"`
	BlockPlace              *filters.FilterContainer `xml:"block-place,omitempty"`
	BlockPlaceAgainst       *filters.FilterContainer `xml:"block-place-against,omitempty"`
	BlockBreak              *filters.FilterContainer `xml:"block-break,omitempty"`
	BlockPhysics            *filters.FilterContainer `xml:"block-physics,omitempty"`
	Use                     *filters.FilterContainer `xml:"use,omitempty"`
	Kit                     *kits.Kit                `xml:"kit,omitempty"`
	LendKit                 *kits.Kit                `xml:"lend-kit,omitempty"`
	Filter                  *filters.FilterContainer `xml:"filter,omitempty"`
}
