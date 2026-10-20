package pokeapi

type Resource struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// TODO: flesh out full response structure
type LocationArea struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// GameIndex            int                   `json:"game_index"`
	// EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	// Location             Location              `json:"location"`
	// Names                []Name                `json:"names"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	Stat     Stat `json:"stat"`
	Effort   int  `json:"effort"`
	BaseStat int  `json:"base_stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}

// type EncounterMethodRate struct {
// 	EncounterMethod EncounterMethod           `json:"encounter_method"`
// 	VersionDetails  []EncounterVersionDetails `json:"version_details"`
// }
