package valueobject

type PokemonItemFotoResponse struct {
	Id              int                   `json:"id"`
	Name            string                `json:"name"`
	FotoURL         string                `json:"url_image"`
	Height          int                   `json:"height"`
	Weight          int                   `json:"weight"`
	BaseExperiencia int                   `json:"base_experience"`
	Types           []PokemonTypeResponse `json:"types"`
	Habilidades     []AbilityResponse     `json:"abilities"`
}

type PokemonResponse struct {
	ID             int                   `json:"id"`
	Name           string                `json:"name"`
	BaseExperience int                   `json:"base_experience"`
	Height         int                   `json:"height"`
	Weight         int                   `json:"weight"`
	Sprites        SpritesResponse       `json:"sprites"`
	Types          []PokemonTypeResponse `json:"types"`
	Stats          []StatResponse        `json:"stats"`
	Abilities      []AbilityResponse     `json:"abilities"`
}

type SpritesResponse struct {
	FrontDefault string `json:"front_default"`
	BackDefault  string `json:"back_default"`
	FrontShiny   string `json:"front_shiny"`
	BackShiny    string `json:"back_shiny"`
}

type PokemonTypeResponse struct {
	Slot int                  `json:"slot"`
	Type TypeResourceResponse `json:"type"`
}

type TypeResourceResponse struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type StatResponse struct {
	BaseStat int                  `json:"base_stat"`
	Effort   int                  `json:"effort"`
	Stat     StatResourceResponse `json:"stat"`
}

type StatResourceResponse struct {
	Name string `json:"name"`
}

type AbilityResponse struct {
	Ability  AbilityResourceResponse `json:"ability"`
	IsHidden bool                    `json:"is_hidden"`
	Slot     int                     `json:"slot"`
}

type AbilityResourceResponse struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonListResponse struct {
	Count    int                   `json:"count"`
	Next     string                `json:"next"`
	Previous string                `json:"previous"`
	Results  []PokemonItemResponse `json:"results"`
}
type PokemonItemResponse struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
