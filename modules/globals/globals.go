package globals

type Globals struct {
	If     []IfOrUnless `xml:"if"`
	Unless []IfOrUnless `xml:"unless"`
}

type IfOrUnless struct {
	Variant            *string `xml:"variant,attr,omitempty"`
	HasVariant         *string `xml:"has-variant,attr,omitempty"`
	Constant           *string `xml:"constant,attr,omitempty"`
	ConstantValue      *string `xml:"constant-value,attr,omitempty"`
	ConstantComparison *string `xml:"constant-comparison,attr,omitempty"`
	Items              []byte  `xml:",innerxml"`
}
