# FlexSupport

A lightweight helpdesk and ticketing system tailored for repair shop workflows. Built with Go, html/template, htmx, and Alpine.js.

## Overview

FlexSupport provides a clean, functional interface for managing repair tickets with two primary user roles:

1. **Employee / Front Desk Staff** — Creates, edits, and manages tickets and customer information
2. **Repair Technician** — Manages active repair jobs, tracks status, parts, costs, and deadlines

## Current Status

**🚧 This is a functional scaffold with mock data.** The frontend UI is complete and interactive, but the backend uses placeholder data. See [BACKEND_TODO.md](BACKEND_TODO.md) for the complete implementation roadmap.

### What's Included

✅ **Frontend Scaffold:**
- Base layout with navigation (`templates/layouts/base.tmpl`)
- Dashboard with ticket listing (`templates/pages/dashboard.tmpl`)
- Ticket entry/edit form (`templates/pages/ticket-form.tmpl`)
- Technician detail view with job progress (`templates/pages/technician-view.tmpl`)
- htmx integration for dynamic updates
- Alpine.js for interactive components
- Tailwind CSS for styling

✅ **Backend Structure:**
- Go HTTP server with chi router (`cmd/repairdesk/main.go`)
- Handler stubs for all routes (`internal/handlers/handlers.go`)
- Data models for tickets, parts, notes, etc. (`internal/models/ticket.go`)
- Organized directory structure ready for expansion

## Quick Start

### Prerequisites

- Go 1.21 or later
- Basic familiarity with Go and HTML templates

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd flexsupport
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/repairdesk/main.go
```

4. Open your browser to `http://localhost:8080`

## Project Structure

```
flexsupport/
├── cmd/repairdesk/         # Application entry point
├── internal/
│   ├── handlers/           # HTTP request handlers
│   ├── models/             # Data models
│   ├── middleware/         # Custom middleware (TODO)
│   ├── database/           # Database layer (TODO)
│   ├── services/           # Business logic (TODO)
│   └── config/             # Configuration (TODO)
├── templates/
│   ├── layouts/            # Base layout templates
│   └── pages/              # Page templates
├── static/
│   ├── css/                # Stylesheets
│   ├── js/                 # JavaScript (optional)
│   └── images/             # Static images
├── go.mod                  # Go module definition
├── BACKEND_TODO.md         # Implementation roadmap
└── README.md               # This file
```

## Available Routes

### Employee/Front Desk Routes
- `GET /` - Dashboard with ticket overview
- `GET /tickets` - List all tickets (with filtering)
- `GET /tickets/new` - New ticket form
- `POST /tickets` - Create ticket
- `GET /tickets/{id}` - View ticket details
- `GET /tickets/{id}/edit` - Edit ticket form
- `POST /tickets/{id}` - Update ticket

### Technician Routes
- `GET /technician` - Technician work queue
- `GET /technician/{id}` - Detailed ticket view with actions
- `POST /tickets/{id}/status` - Update ticket status (htmx)
- `POST /tickets/{id}/parts` - Add parts (htmx)
- `DELETE /tickets/{id}/parts/{partId}` - Remove parts (htmx)
- `POST /tickets/{id}/notes` - Add work notes (htmx)

### API Endpoints (for htmx)
- `GET /api/stats/open` - Get open ticket count
- `GET /tickets/search` - Search tickets

## Features

### Current (with Mock Data)
- ✅ Ticket creation and management
- ✅ Customer information tracking
- ✅ Device details and serial numbers
- ✅ Status tracking (New, In Progress, Waiting for Parts, Ready, Completed)
- ✅ Priority levels (Low, Normal, High, Urgent)
- ✅ Parts and materials tracking
- ✅ Work notes/logs
- ✅ Assignment to technicians
- ✅ Due date tracking
- ✅ Dashboard statistics
- ✅ Real-time updates with htmx
- ✅ Responsive design

### Planned (See BACKEND_TODO.md)
- [ ] Database persistence
- [ ] User authentication
- [ ] Role-based access control
- [ ] Search functionality
- [ ] Email/SMS notifications
- [ ] Reporting and analytics
- [ ] File uploads (device photos)
- [ ] Printable receipts/invoices
- [ ] Customer history tracking

## Next Steps

To transform this scaffold into a production application:

1. **Read [BACKEND_TODO.md](BACKEND_TODO.md)** for the complete implementation roadmap
2. **Set up a database** (SQLite for quick start, PostgreSQL for production)
3. **Implement data persistence** by replacing mock data with real database queries
4. **Add authentication** to secure the application
5. **Test thoroughly** with real repair shop workflows
6. **Deploy** to your chosen hosting platform

## Technology Stack

- **Backend:** Go 1.21+ with `html/template`
- **Router:** chi v5
- **Frontend:** htmx for dynamic updates, Alpine.js for interactivity
- **Styling:** Tailwind CSS (via CDN, can be customized)
- **Database:** Not yet implemented (PostgreSQL/SQLite recommended)

## Development

### Adding New Templates

Templates are organized in three locations:
- `templates/layouts/` - Base layouts (navigation, structure)
- `templates/pages/` - Full page templates
- `templates/partials/` - Reusable components (TODO)

All page templates should extend the base layout:
```html
{{define "title"}}Page Title{{end}}
{{define "content"}}
  <!-- Your content here -->
{{end}}
```

### Adding New Routes

1. Add route definition in `cmd/repairdesk/main.go` in the `setupRoutes` function
2. Implement handler in `internal/handlers/handlers.go`
3. Create or update templates as needed

### Using htmx

htmx attributes are already integrated for dynamic updates:
- `hx-get`, `hx-post`, `hx-delete` - Make requests
- `hx-target` - Specify where to inject response
- `hx-swap` - Control how content is swapped
- `hx-trigger` - Define when requests fire

Example:
```html
<button hx-post="/tickets/123/status"
        hx-vals='{"status": "completed"}'
        hx-target="#status-badge">
    Mark Complete
</button>
```

## Contributing

This is a scaffold project. Contributions welcome for:
- Database implementation
- Authentication system
- Additional features
- Bug fixes
- Documentation improvements

## License

[Your chosen license]

## Support

For questions about implementation, see [BACKEND_TODO.md](BACKEND_TODO.md) or open an issue.

---

Built with Go, htmx, Alpine.js, and Tailwind CSS.
