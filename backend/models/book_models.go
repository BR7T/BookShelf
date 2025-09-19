package models

type BookResponse map[string]struct {
	ISBN string `json:"ISBN"`
    Details struct {
        Title      string   `json:"title"`
        Authors    []struct {
            Name string `json:"name"`
        } `json:"authors"`
        Publishers []string `json:"publishers"`
    } `json:"details"`
}

