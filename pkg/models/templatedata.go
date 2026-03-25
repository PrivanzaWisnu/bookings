package models

// Holds Data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]string
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string // Cross Site Request Forgery Token
	Flash     string
	Warning   string
	Error     string
}
