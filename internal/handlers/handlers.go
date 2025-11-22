package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"flexsupport/internal/models"
)

// Handler holds dependencies for HTTP handlers
type Handler struct {
	templates map[string]*template.Template
	// TODO: Add database connection, services, etc.
}

// NewHandler creates a new Handler instance
func NewHandler(templates map[string]*template.Template) *Handler {
	return &Handler{
		templates: templates,
	}
}

// Dashboard renders the main dashboard view
func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	// TODO: Fetch real data from database
	data := map[string]interface{}{
		"CurrentUser": "Admin User",
		"Stats": models.TicketStats{
			OpenTickets:    12,
			InProgress:     5,
			Overdue:        2,
			CompletedToday: 3,
		},
		"Tickets": getMockTickets(),
	}

	h.render(w, "dashboard", data)
}

// ListTickets handles the ticket listing page
func (h *Handler) ListTickets(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement filtering based on query parameters
	status := r.URL.Query().Get("status")
	search := r.URL.Query().Get("search")

	log.Printf("Listing tickets with status=%s, search=%s", status, search)

	// For now, return mock data
	h.renderPartial(w, "ticket-list", map[string]interface{}{
		"Tickets": getMockTickets(),
	})
}

// NewTicketForm renders the new ticket form
func (h *Handler) NewTicketForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"CurrentUser": "Admin User",
		"Ticket":      &models.Ticket{},
		"Technicians": getMockTechnicians(),
	}

	h.render(w, "ticket-form", data)
}

// CreateTicket handles ticket creation
func (h *Handler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	// TODO: Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// TODO: Validate and save to database
	log.Printf("Creating ticket: %+v", r.PostForm)

	// For now, redirect to dashboard
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ViewTicket renders the ticket detail view
func (h *Handler) ViewTicket(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// TODO: Fetch from database
	ticket := getMockTicket(id)

	data := map[string]interface{}{
		"CurrentUser": "Admin User",
		"Ticket":      ticket,
	}

	h.render(w, "technician-view", data)
}

// EditTicketForm renders the edit ticket form
func (h *Handler) EditTicketForm(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// TODO: Fetch from database
	ticket := getMockTicket(id)

	data := map[string]interface{}{
		"CurrentUser": "Admin User",
		"Ticket":      ticket,
		"Technicians": getMockTechnicians(),
	}

	h.render(w, "ticket-form", data)
}

// UpdateTicket handles ticket updates
func (h *Handler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// TODO: Update database
	log.Printf("Updating ticket %s: %+v", idStr, r.PostForm)

	http.Redirect(w, r, "/tickets/"+idStr, http.StatusSeeOther)
}

// SearchTickets handles ticket search (htmx endpoint)
func (h *Handler) SearchTickets(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	log.Printf("Searching tickets: %s", search)

	// TODO: Implement search
	h.renderPartial(w, "ticket-list", map[string]interface{}{
		"Tickets": getMockTickets(),
	})
}

// UpdateTicketStatus handles status updates (htmx endpoint)
func (h *Handler) UpdateTicketStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	status := r.FormValue("status")

	log.Printf("Updating ticket %s status to: %s", idStr, status)

	// TODO: Update database

	// Return updated status badge
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<span id="status-badge" class="px-3 py-1 inline-flex text-sm leading-5 font-semibold rounded-full bg-yellow-100 text-yellow-800">%s</span>`, status)
}

// AddPart adds a part to a ticket (htmx endpoint)
func (h *Handler) AddPart(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	log.Printf("Adding part to ticket %s: %+v", idStr, r.PostForm)

	// TODO: Save to database

	// Return new part HTML
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<div class="flex items-center justify-between p-3 bg-gray-50 rounded-md">
		<div class="flex-1">
			<span class="text-sm font-medium text-gray-900">%s</span>
			<span class="text-sm text-gray-500 ml-2">× %s</span>
		</div>
		<div class="flex items-center gap-3">
			<span class="text-sm font-medium text-gray-900">$%s</span>
		</div>
	</div>`, r.FormValue("part_name"), r.FormValue("quantity"), r.FormValue("cost"))
}

// DeletePart removes a part from a ticket (htmx endpoint)
func (h *Handler) DeletePart(w http.ResponseWriter, r *http.Request) {
	ticketID := chi.URLParam(r, "id")
	partID := chi.URLParam(r, "partId")

	log.Printf("Deleting part %s from ticket %s", partID, ticketID)

	// TODO: Delete from database

	w.WriteHeader(http.StatusOK)
}

// AddNote adds a work note to a ticket (htmx endpoint)
func (h *Handler) AddNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	note := r.FormValue("note")

	log.Printf("Adding note to ticket %s: %s", idStr, note)

	// TODO: Save to database

	// Return new note HTML
	w.Header().Set("Content-Type", "text/html")
	now := time.Now().Format("Jan 2, 2006 3:04 PM")
	fmt.Fprintf(w, `<div class="p-3 bg-gray-50 rounded-md">
		<div class="flex justify-between items-start mb-1">
			<span class="text-sm font-medium text-gray-900">Current User</span>
			<span class="text-xs text-gray-500">%s</span>
		</div>
		<p class="text-sm text-gray-700 whitespace-pre-line">%s</p>
	</div>`, now, note)
}

// TechnicianQueue shows the technician's work queue
func (h *Handler) TechnicianQueue(w http.ResponseWriter, r *http.Request) {
	// TODO: Filter tickets by assigned technician
	data := map[string]interface{}{
		"CurrentUser": "Tech User",
		"Tickets":     getMockTickets(),
	}

	h.render(w, "dashboard", data)
}

// TechnicianTicketView shows the detailed technician view of a ticket
func (h *Handler) TechnicianTicketView(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	ticket := getMockTicket(id)

	data := map[string]interface{}{
		"CurrentUser": "Tech User",
		"Ticket":      ticket,
	}

	h.render(w, "technician-view", data)
}

// GetOpenTicketsCount returns the count of open tickets (htmx endpoint)
func (h *Handler) GetOpenTicketsCount(w http.ResponseWriter, r *http.Request) {
	// TODO: Query database
	fmt.Fprintf(w, "12")
}

// render executes a template with base layout
func (h *Handler) render(w http.ResponseWriter, templateName string, data map[string]interface{}) {
	// Get the template set for this page
	tmpl, ok := h.templates[templateName]
	if !ok {
		log.Printf("Template %s not found", templateName)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Execute the specific page template, which will use the base layout
	if err := tmpl.ExecuteTemplate(w, templateName, data); err != nil {
		log.Printf("Error rendering template %s: %v", templateName, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// renderPartial executes a template without the base layout
func (h *Handler) renderPartial(w http.ResponseWriter, templateName string, data map[string]interface{}) {
	// For partials, try to find them in any template set (use the first one)
	for _, tmpl := range h.templates {
		if err := tmpl.ExecuteTemplate(w, templateName, data); err == nil {
			return
		}
	}
	log.Printf("Error rendering partial %s: template not found", templateName)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

// Mock data functions - TODO: Replace with real database queries

func getMockTickets() []models.Ticket {
	return []models.Ticket{
		{
			ID:               1001,
			Status:           "new",
			Priority:         "high",
			CustomerName:     "John Doe",
			CustomerPhone:    "(555) 123-4567",
			DeviceType:       "Smartphone",
			DeviceModel:      "iPhone 13 Pro",
			IssueDescription: "Cracked screen, needs replacement",
			AssignedTo:       "Mike Tech",
			DueDate:          timePtr(time.Now().Add(48 * time.Hour)),
		},
		{
			ID:               1002,
			Status:           "in_progress",
			Priority:         "normal",
			CustomerName:     "Jane Smith",
			CustomerPhone:    "(555) 987-6543",
			DeviceType:       "Laptop",
			DeviceModel:      "MacBook Pro 2020",
			IssueDescription: "Battery not charging",
			AssignedTo:       "Sarah Tech",
		},
	}
}

func getMockTicket(id int) models.Ticket {
	dueDate := time.Now().Add(48 * time.Hour)
	return models.Ticket{
		ID:               id,
		Status:           "in_progress",
		Priority:         "high",
		CustomerName:     "John Doe",
		CustomerPhone:    "(555) 123-4567",
		CustomerEmail:    "john@example.com",
		DeviceType:       "Smartphone",
		DeviceBrand:      "Apple",
		DeviceModel:      "iPhone 13 Pro",
		SerialNumber:     "ABC123456789",
		IssueDescription: "Screen is completely shattered after being dropped. Touch functionality still works but glass is unsafe.",
		EstimatedCost:    150.00,
		AssignedTo:       "Mike Tech",
		DueDate:          &dueDate,
		CreatedAt:        time.Now().Add(-24 * time.Hour),
		UpdatedAt:        time.Now(),
		CreatedBy:        "Front Desk",
		Parts: []models.Part{
			{ID: 1, Name: "iPhone 13 Pro Screen Assembly", Quantity: 1, Cost: 89.99},
			{ID: 2, Name: "Screen Adhesive", Quantity: 1, Cost: 5.99},
		},
		Notes: []models.WorkNote{
			{
				ID:        1,
				Author:    "Mike Tech",
				Content:   "Customer confirmed backup was done. Safe to proceed.",
				Timestamp: time.Now().Add(-2 * time.Hour),
			},
		},
		TotalPartsCost: 95.98,
	}
}

func getMockTechnicians() []models.Technician {
	return []models.Technician{
		{ID: 1, Name: "Mike Tech", ActiveJobs: 3, IsAvailable: true},
		{ID: 2, Name: "Sarah Tech", ActiveJobs: 2, IsAvailable: true},
		{ID: 3, Name: "Bob Repair", ActiveJobs: 5, IsAvailable: false},
	}
}

func timePtr(t time.Time) *time.Time {
	return &t
}
