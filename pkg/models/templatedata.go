package models

// TemplateData holds data sent from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	DataMap map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}