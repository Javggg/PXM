package variables

import "encoding/xml"

type Variables struct {
	XMLName xml.Name `xml:"variables"`
	VariableContainer
	Scope *string `xml:"scope,attr,omitempty"`
}

type BaseVariable struct {
	ID *string `xml:"id,attr,omitempty"`
}

type VariableContainer struct {
	Variable       []Variable       `xml:"variable"`
	Array          []Array          `xml:"array"`
	WithTeam       []WithTeam       `xml:"with-team"`
	PlayerLocation []PlayerLocation `xml:"player-location"`
	Score          []Score          `xml:"score"`
	TimeLimit      []TimeLimit      `xml:"timelimit"`
	MaxBuildHeight []MaxBuildHeight `xml:"maxbuildheight"`
	WorldTime      []WorldTime      `xml:"worldtime"`
}

type Variable struct {
	XMLName xml.Name `xml:"variable"`
	BaseVariable
	Exclusive *string `xml:"exclusive,attr,omitempty"`
	Scope     *string `xml:"scope,attr,omitempty"`
	Default   *string `xml:"default,attr,omitempty"`
}

type Array struct {
	XMLName xml.Name `xml:"array"`
	BaseVariable
	Size    string  `xml:"size,attr"`
	Scope   *string `xml:"scope,attr,omitempty"`
	Default *string `xml:"default,attr,omitempty"`
}

type WithTeam struct {
	XMLName xml.Name `xml:"with-team"`
	BaseVariable
	Var  string `xml:"var,attr"`
	Team string `xml:"team,attr"`
}

type PlayerLocation struct {
	XMLName xml.Name `xml:"player-location"`
	BaseVariable
	Component string `xml:"component,attr"`
}

type Score struct {
	XMLName xml.Name `xml:"score"`
	BaseVariable
}

type TimeLimit struct {
	XMLName xml.Name `xml:"timelimit"`
	BaseVariable
}

type MaxBuildHeight struct {
	XMLName xml.Name `xml:"maxbuildheight"`
	BaseVariable
}

type WorldTime struct {
	XMLName xml.Name `xml:"worldtime"`
	BaseVariable
}
