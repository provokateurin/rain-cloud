package common

//nolint:gochecknoglobals
var DummyOCSMeta = struct {
	Itemsperpage *string `json:"itemsperpage,omitempty"`
	Message      *string `json:"message,omitempty"`
	Status       string  `json:"status"`
	Statuscode   int     `json:"statuscode"`
	Totalitems   *string `json:"totalitems,omitempty"`
}{
	Itemsperpage: nil,
	Message:      nil,
	Status:       "ok",
	Statuscode:   200,
	Totalitems:   nil,
}
