package meander

// Place represents a single place.
type Place struct {
	*googleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photos          []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

// Public gets a public view of a Place.
func (p *Place) Public() interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photos,
		"vicinity": p.Vicinity,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}

type googleResponse struct {
	Results []Place `json:"results"`
}
type googleGeometry struct {
	googleLocation `json:"location"`
}
type googleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type googlePhoto struct {
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}
