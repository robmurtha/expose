package exiv2

// Fields contains field data organized by group EXIF, XMP, IPTC.
type Fields struct {
	EXIF []*Field `json:"exif,omitempty"`
	XMP  []*Field `json:"xmp,omitempty"`
	IPTC []*Field `json:"iptc,omitempty"`
}

// Field contains metadata for one field as parsed from exiv2.
type Field struct {
	// Group is the top level group name such as video, audio, Image, Photo
	Group string `json:"group"`

	// Name is the tag name of the field and is generally camel case
	Name string `json:"name"`
	// Type depends on the container, for XMP XmpText is common and may also contain other data types.

	Type string `json:"type"`
	// Size of field, or number of elements if a collection.

	Count int `json:"count"`

	// Value holds the data value.
	Value string `json:"value"`
}
