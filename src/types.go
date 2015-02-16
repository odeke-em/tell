package tell

type Association struct {
	Label string             `"json:label,string"`
	Meta  *map[string]string `"json"`
}

type Entity struct {
	Id string `"json:id,string"`
	// Associations maps labels to Associations
	Associations *map[string]Association
	Meta         *map[string]string `"json"`
}

type Area struct {
	// Id is the unique identifier, could be a url, name etc
	Id       string `"json:id,string"`
	Segments *map[string]Segment
}

type Segment struct {
	Id          string             `"json:id,string"`
	Annotations *map[string]string `"json"`
}
