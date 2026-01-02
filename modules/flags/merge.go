package flags

import (
	"encoding/xml"
	"strings"
)

func (f *Flags) Merge(other Flags) {
	f.Flags = append(f.Flags, other.Flags...)
	f.Nets = append(f.Nets, other.Nets...)
	f.Posts = append(f.Posts, other.Posts...)
}

func (p *Post) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*p = Post{}

	for _, attr := range start.Attr {
		val := attr.Value
		switch attr.Name.Local {
		case "id":
			p.ID = &val
		case "name":
			p.Name = &val
		case "owner":
			p.Owner = &val
		case "permanent":
			p.Permanent = &val
		case "sequential":
			p.Sequential = &val
		case "points-rate":
			p.PointsRate = &val
		case "recover-time":
			p.RecoverTime = &val
		case "respawn-time":
			p.RespawnTime = &val
		case "respawn-speed":
			p.RespawnSpeed = &val
		case "yaw":
			p.Yaw = &val
		case "fallback":
			p.Fallback = &val
		}
	}

	var text strings.Builder

	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}

		switch t := tok.(type) {

		case xml.CharData:
			if strings.TrimSpace(string(t)) != "" {
				text.Write(t)
			}

		case xml.StartElement:
			if t.Name.Local == "post" {
				var child Post
				if err := d.DecodeElement(&child, &t); err != nil {
					return err
				}
				p.Posts = append(p.Posts, child)
			}

		case xml.EndElement:
			if t.Name.Local == start.Name.Local {
				p.Value = strings.TrimSpace(text.String())
				return nil
			}
		}
	}
}

func (p Post) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "post"}

	addAttr := func(name string, val *string) {
		if val != nil {
			start.Attr = append(start.Attr, xml.Attr{
				Name:  xml.Name{Local: name},
				Value: *val,
			})
		}
	}

	addAttr("id", p.ID)
	addAttr("name", p.Name)
	addAttr("owner", p.Owner)
	addAttr("permanent", p.Permanent)
	addAttr("sequential", p.Sequential)
	addAttr("points-rate", p.PointsRate)
	addAttr("recover-time", p.RecoverTime)
	addAttr("respawn-time", p.RespawnTime)
	addAttr("respawn-speed", p.RespawnSpeed)
	addAttr("yaw", p.Yaw)
	addAttr("fallback", p.Fallback)

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if len(p.Posts) > 0 {
		for _, child := range p.Posts {
			if err := e.Encode(child); err != nil {
				return err
			}
		}
	} else if p.Value != "" {
		if err := e.EncodeToken(xml.CharData(p.Value)); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}
