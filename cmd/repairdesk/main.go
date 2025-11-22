package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"flexsupport/internal/handlers"
)

func main() {
	// Initialize router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	// Parse templates
	templates, err := loadTemplates()
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
	}

	log.Printf("Loaded %d page templates", len(templates))

	// Initialize handlers
	h := handlers.NewHandler(templates)

	// Static files
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Routes
	setupRoutes(r, h)

	// Start server
	port := "8080"
	log.Printf("Starting RepairDesk server on :%s", port)
	log.Printf("Visit http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// loadTemplates parses all template files, creating isolated template sets for each page
func loadTemplates() (map[string]*template.Template, error) {
	// Get list of all template files
	layoutFiles, err := filepath.Glob(filepath.Join("templates", "layouts", "*.tmpl"))
	if err != nil {
		return nil, fmt.Errorf("globbing layouts: %w", err)
	}
	if len(layoutFiles) == 0 {
		return nil, fmt.Errorf("no layout templates found in templates/layouts/")
	}

	pageFiles, err := filepath.Glob(filepath.Join("templates", "pages", "*.tmpl"))
	if err != nil {
		return nil, fmt.Errorf("globbing pages: %w", err)
	}

	partialFiles, _ := filepath.Glob(filepath.Join("templates", "partials", "*.tmpl"))

	// Create a map to hold separate template sets for each page
	templates := make(map[string]*template.Template)

	// Parse each page with its own template set to avoid block name conflicts
	for _, pageFile := range pageFiles {
		// Get the base name without extension to use as the template name
		baseName := filepath.Base(pageFile)
		templateName := baseName[:len(baseName)-len(filepath.Ext(baseName))]

		// Combine layout files with this specific page file
		files := append([]string{}, layoutFiles...)
		files = append(files, pageFile)
		files = append(files, partialFiles...)

		// Parse this combination into its own template set
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", pageFile, err)
		}

		templates[templateName] = tmpl
	}

	return templates, nil
}

// setupRoutes configures all application routes
func setupRoutes(r chi.Router, h *handlers.Handler) {
	// Dashboard
	r.Get("/", h.Dashboard)

	// Tickets
	r.Route("/tickets", func(r chi.Router) {
		r.Get("/", h.ListTickets)
		r.Get("/new", h.NewTicketForm)
		r.Post("/", h.CreateTicket)
		r.Get("/search", h.SearchTickets)
		r.Get("/{id}", h.ViewTicket)
		r.Get("/{id}/edit", h.EditTicketForm)
		r.Post("/{id}", h.UpdateTicket)

		// Ticket actions
		r.Post("/{id}/status", h.UpdateTicketStatus)
		r.Post("/{id}/parts", h.AddPart)
		r.Delete("/{id}/parts/{partId}", h.DeletePart)
		r.Post("/{id}/notes", h.AddNote)
	})

	// Technician view
	r.Get("/technician", h.TechnicianQueue)
	r.Get("/technician/{id}", h.TechnicianTicketView)

	// API endpoints (for htmx)
	r.Route("/api", func(r chi.Router) {
		r.Get("/stats/open", h.GetOpenTicketsCount)
	})
}
