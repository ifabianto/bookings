package models

import "github.com/ifabianto/bookings/internal/forms"

// TemplateData holds data sent from handlers to templates (available to all the pages in website)
type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int
}
