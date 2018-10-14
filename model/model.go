package model

// Location struct
type Location struct {
	Type   string    `json:"type,omitempty"`
	Coords []float64 `json:"coordinates,omitempty"`
}

//If omitempty is not set, then edges with empty values (0 for int/float, "" for string, false
// for bool) would be created for values not specified explicitly.

// Place struct
type Place struct {
	Place     string   `json:"place"`
	UID       string   `json:"uid,omitempty"`
	Name      string   `json:"name,omitempty"`
	Featured  string   `json:"featured,omitempty"`
	Website   string   `json:"website,omitempty"`
	Tagline   string   `json:"tagline,omitempty"`
	Contact   string   `json:"contact,omitempty"`
	Verified  bool     `json:"verified,omitempty"`
	Location  Location `json:"location,omitempty"`
	Amenity   string   `json:"amenity,omitempty"`
	CreatedAt string   `json:"created_at,omitempty"`
}

// Interest struct
type Interest struct {
	Interest string `json:"interest"`
	UID      string `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
}

// Account struct
type Account struct {
	Account    string      `json:"account"`
	UID        string      `json:"uid,omitempty"`
	Name       string      `json:"name,omitempty"`
	Username   string      `json:"username,omitempty"`
	Password   string      `json:"password,omitempty"`
	Email      string      `json:"email,omitempty"`
	ProfileURL string      `json:"profile_url,omitempty"`
	Verified   bool        `json:"verified,omitempty"`
	CreatedAt  string      `json:"created_at,omitempty"`
	Follows    []*Place    `json:"follows,omitempty"`
	Interested []*Interest `json:"interested,omitempty"`
}
