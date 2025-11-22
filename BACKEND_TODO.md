# RepairDesk Backend Implementation TODO

This document outlines the backend implementation tasks needed to transform the scaffold into a fully functional repair shop helpdesk system.

## Directory Structure

```
flexsupport/
├── cmd/
│   └── repairdesk/
│       └── main.go                 # Application entry point
├── internal/
│   ├── handlers/
│   │   └── handlers.go            # HTTP request handlers
│   ├── models/
│   │   └── ticket.go              # Data models
│   ├── middleware/                 # Custom middleware (auth, logging, etc.)
│   ├── database/                   # Database connection and queries
│   │   ├── db.go                  # Connection management
│   │   ├── migrations/            # SQL migration files
│   │   └── queries/               # SQL query implementations
│   ├── services/                   # Business logic layer
│   │   ├── ticket_service.go
│   │   ├── customer_service.go
│   │   └── notification_service.go
│   └── config/                     # Configuration management
│       └── config.go
├── templates/
│   ├── layouts/
│   │   └── base.tmpl              ✅ Created
│   ├── pages/
│   │   ├── dashboard.tmpl         ✅ Created
│   │   ├── ticket-form.tmpl       ✅ Created
│   │   └── technician-view.tmpl   ✅ Created
│   └── partials/                   # Reusable template components
│       ├── status-badge.tmpl
│       └── notification.tmpl
├── static/
│   ├── css/
│   │   └── styles.css             ✅ Created
│   ├── js/
│   │   └── app.js                 # Optional JavaScript enhancements
│   └── images/
├── migrations/                     # Database migration files
├── go.mod                         ✅ Created
├── go.sum
├── .env.example                    # Environment variable template
└── README.md                       # Project documentation
```

---

## Core Tasks

### 1. Database Setup

#### 1.1 Choose Database Solution
- [ ] **PostgreSQL** (recommended for production)
  - Robust, feature-rich, excellent for structured data
  - Go driver: `github.com/lib/pq` or use with `database/sql`
- [ ] **SQLite** (good for getting started)
  - Zero-config, file-based, perfect for prototyping
  - Go driver: `github.com/mattn/go-sqlite3`
- [ ] **MySQL/MariaDB** (alternative)
  - Widely used, good tooling
  - Go driver: `github.com/go-sql-driver/mysql`

#### 1.2 Create Database Schema
Create migration files in `internal/database/migrations/`:

**001_initial_schema.sql:**
```sql
-- Customers table
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(phone)
);

-- Technicians table
CREATE TABLE technicians (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    is_available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tickets table
CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    status VARCHAR(50) DEFAULT 'new',
    priority VARCHAR(50) DEFAULT 'normal',

    customer_id INT REFERENCES customers(id),
    customer_name VARCHAR(255) NOT NULL,
    customer_phone VARCHAR(50) NOT NULL,
    customer_email VARCHAR(255),

    device_type VARCHAR(100) NOT NULL,
    device_brand VARCHAR(100),
    device_model VARCHAR(100),
    serial_number VARCHAR(255),

    issue_description TEXT NOT NULL,
    internal_notes TEXT,
    estimated_cost DECIMAL(10,2) DEFAULT 0,

    assigned_to INT REFERENCES technicians(id),
    due_date TIMESTAMP,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),

    CHECK (status IN ('new', 'in_progress', 'waiting_parts', 'ready', 'completed', 'cancelled')),
    CHECK (priority IN ('low', 'normal', 'high', 'urgent'))
);

-- Parts table
CREATE TABLE parts (
    id SERIAL PRIMARY KEY,
    ticket_id INT REFERENCES tickets(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    cost DECIMAL(10,2) NOT NULL,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    added_by VARCHAR(255)
);

-- Work notes table
CREATE TABLE work_notes (
    id SERIAL PRIMARY KEY,
    ticket_id INT REFERENCES tickets(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for performance
CREATE INDEX idx_tickets_status ON tickets(status);
CREATE INDEX idx_tickets_assigned ON tickets(assigned_to);
CREATE INDEX idx_tickets_created ON tickets(created_at);
CREATE INDEX idx_tickets_customer ON tickets(customer_phone);
CREATE INDEX idx_parts_ticket ON parts(ticket_id);
CREATE INDEX idx_notes_ticket ON work_notes(ticket_id);
```

#### 1.3 Database Connection Management
Create `internal/database/db.go`:
```go
package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // or your chosen driver
)

type DB struct {
    *sql.DB
}

func New(dataSourceName string) (*DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, fmt.Errorf("opening database: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("pinging database: %w", err)
    }

    // Set connection pool settings
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)

    log.Println("Database connection established")
    return &DB{db}, nil
}
```

#### 1.4 Migration Runner
- [ ] Implement migration runner or use tool like `golang-migrate/migrate`
- [ ] Create up/down migrations for schema changes
- [ ] Add migration command to application

---

### 2. Data Access Layer

#### 2.1 Ticket Repository
Create `internal/database/ticket_repo.go`:
- [ ] `GetTicketByID(id int) (*models.Ticket, error)`
- [ ] `ListTickets(filters TicketFilters) ([]models.Ticket, error)`
- [ ] `CreateTicket(ticket *models.Ticket) error`
- [ ] `UpdateTicket(ticket *models.Ticket) error`
- [ ] `DeleteTicket(id int) error`
- [ ] `UpdateTicketStatus(id int, status string) error`
- [ ] `SearchTickets(query string) ([]models.Ticket, error)`
- [ ] `GetTicketStats() (*models.TicketStats, error)`
- [ ] `GetTicketsByTechnician(techID int) ([]models.Ticket, error)`

#### 2.2 Parts Repository
Create `internal/database/part_repo.go`:
- [ ] `AddPart(part *models.Part) error`
- [ ] `GetPartsByTicket(ticketID int) ([]models.Part, error)`
- [ ] `DeletePart(id int) error`
- [ ] `GetTotalPartsCost(ticketID int) (float64, error)`

#### 2.3 Work Notes Repository
Create `internal/database/note_repo.go`:
- [ ] `AddNote(note *models.WorkNote) error`
- [ ] `GetNotesByTicket(ticketID int) ([]models.WorkNote, error)`
- [ ] `DeleteNote(id int) error`

#### 2.4 Customer Repository
Create `internal/database/customer_repo.go`:
- [ ] `GetOrCreateCustomer(phone string, name string, email string) (*models.Customer, error)`
- [ ] `GetCustomerByPhone(phone string) (*models.Customer, error)`
- [ ] `GetCustomerHistory(customerID int) ([]models.Ticket, error)`

#### 2.5 Technician Repository
Create `internal/database/technician_repo.go`:
- [ ] `GetTechnicians() ([]models.Technician, error)`
- [ ] `GetTechnicianByID(id int) (*models.Technician, error)`
- [ ] `UpdateTechnicianAvailability(id int, available bool) error`
- [ ] `GetTechnicianWorkload(id int) (int, error)`

---

### 3. Business Logic Layer

Create service layer in `internal/services/`:

#### 3.1 Ticket Service
`internal/services/ticket_service.go`:
- [ ] Wrap database operations with business logic
- [ ] Validate ticket data before saving
- [ ] Calculate total costs (labor + parts)
- [ ] Auto-assign tickets to available technicians
- [ ] Handle status transitions with validation
- [ ] Generate ticket numbers/IDs

#### 3.2 Customer Service
`internal/services/customer_service.go`:
- [ ] Deduplicate customers by phone
- [ ] Track customer history
- [ ] Calculate customer lifetime value

#### 3.3 Notification Service
`internal/services/notification_service.go`:
- [ ] Send SMS notifications (integrate Twilio or similar)
- [ ] Send email notifications (integrate SendGrid or SMTP)
- [ ] Notify customer when ticket status changes
- [ ] Notify customer when repair is ready
- [ ] Send reminder notifications

---

### 4. Complete Handler Implementation

Update `internal/handlers/handlers.go`:

#### 4.1 Replace Mock Data
- [ ] Remove all `getMock*()` functions
- [ ] Inject database repositories into Handler struct
- [ ] Update all handlers to use real database queries

#### 4.2 Form Validation
- [ ] Add input validation for all forms
- [ ] Sanitize user input to prevent XSS
- [ ] Return helpful error messages

#### 4.3 Error Handling
- [ ] Implement consistent error responses
- [ ] Log errors properly
- [ ] Show user-friendly error pages

#### 4.4 HTMX Response Handling
- [ ] Return proper HTML fragments for htmx requests
- [ ] Use appropriate htmx headers (HX-Trigger, HX-Redirect, etc.)
- [ ] Implement loading states

---

### 5. Authentication & Authorization

Create `internal/middleware/auth.go`:

#### 5.1 User Authentication
- [ ] Implement session management (use `gorilla/sessions`)
- [ ] Create login/logout handlers
- [ ] Password hashing with bcrypt
- [ ] Login form template

#### 5.2 Role-Based Access Control
- [ ] Define roles: Admin, Front Desk, Technician
- [ ] Middleware to check user permissions
- [ ] Restrict routes based on roles
- [ ] Show/hide UI elements based on permissions

#### 5.3 Session Management
- [ ] Store user info in session
- [ ] Add CSRF protection
- [ ] Implement "remember me" functionality

---

### 6. Configuration Management

Create `internal/config/config.go`:

#### 6.1 Environment Variables
Create `.env.example`:
```env
# Database
DATABASE_URL=postgres://user:pass@localhost:5432/repairdesk?sslmode=disable

# Server
PORT=8080
ENV=development

# Session
SESSION_SECRET=your-secret-key-change-in-production

# Email (optional)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@example.com
SMTP_PASS=your-password

# SMS (optional)
TWILIO_ACCOUNT_SID=your-account-sid
TWILIO_AUTH_TOKEN=your-auth-token
TWILIO_PHONE_NUMBER=+1234567890
```

#### 6.2 Config Loader
- [ ] Use `godotenv` to load .env files
- [ ] Parse and validate configuration
- [ ] Provide sensible defaults
- [ ] Make config available to handlers/services

---

### 7. Additional Features

#### 7.1 Search Functionality
- [ ] Full-text search across tickets
- [ ] Search by customer name, phone, device
- [ ] Search by ticket ID
- [ ] Filter by date range

#### 7.2 Reporting
- [ ] Daily/weekly/monthly ticket reports
- [ ] Technician performance metrics
- [ ] Revenue reports
- [ ] Export to CSV/PDF

#### 7.3 File Uploads
- [ ] Allow uploading photos of damaged devices
- [ ] Store files in `static/uploads/`
- [ ] Display photos in ticket view
- [ ] Security: validate file types, limit sizes

#### 7.4 Print Functionality
- [ ] Printable ticket receipts
- [ ] Work order printouts for technicians
- [ ] Customer invoices

#### 7.5 Notifications
- [ ] Email notifications on status changes
- [ ] SMS reminders for pickups
- [ ] In-app notifications

#### 7.6 Dashboard Enhancements
- [ ] Real-time updates with WebSocket or SSE
- [ ] Charts and graphs for metrics
- [ ] Quick filters and saved views

---

### 8. Testing

#### 8.1 Unit Tests
- [ ] Test database repositories
- [ ] Test business logic in services
- [ ] Test utility functions
- [ ] Aim for >70% coverage

#### 8.2 Integration Tests
- [ ] Test HTTP handlers end-to-end
- [ ] Use test database
- [ ] Test htmx interactions

#### 8.3 Manual Testing
- [ ] Create test tickets
- [ ] Test all workflows
- [ ] Test on different browsers
- [ ] Test responsive design on mobile

---

### 9. Deployment Preparation

#### 9.1 Production Readiness
- [ ] Add rate limiting middleware
- [ ] Implement proper logging (structured logs)
- [ ] Add health check endpoint (`/health`)
- [ ] Configure CORS if needed
- [ ] Enable HTTPS/TLS
- [ ] Add request ID middleware

#### 9.2 Database Backups
- [ ] Set up automated backups
- [ ] Test restore process
- [ ] Document backup procedures

#### 9.3 Deployment
- [ ] Create Dockerfile
- [ ] Set up CI/CD pipeline
- [ ] Deploy to chosen platform (Fly.io, Railway, DigitalOcean, etc.)
- [ ] Set up monitoring (uptime, errors)
- [ ] Configure domain and SSL

#### 9.4 Documentation
- [ ] Write README with setup instructions
- [ ] Document API endpoints
- [ ] Create user guide
- [ ] Document deployment process

---

## Suggested Implementation Order

### Phase 1: Core Functionality (MVP)
1. Set up database (SQLite for quick start)
2. Implement basic ticket CRUD operations
3. Replace mock data with real database queries
4. Get create/read/update working end-to-end

### Phase 2: Enhanced Features
5. Add parts and work notes functionality
6. Implement search and filtering
7. Add basic authentication (single admin user)
8. Polish the UI and fix bugs

### Phase 3: Production Features
9. Add role-based access control
10. Implement notifications
11. Add reporting and analytics
12. Testing and bug fixes

### Phase 4: Polish & Deploy
13. Performance optimization
14. Security hardening
15. Documentation
16. Deploy to production

---

## Useful Go Libraries

### Core
- **Router**: `github.com/go-chi/chi/v5` ✅ Already included
- **Database**: `github.com/lib/pq` (PostgreSQL) or `github.com/mattn/go-sqlite3` (SQLite)
- **Migrations**: `github.com/golang-migrate/migrate/v4`
- **Environment**: `github.com/joho/godotenv`

### Authentication
- **Sessions**: `github.com/gorilla/sessions`
- **Password**: `golang.org/x/crypto/bcrypt` (standard library)
- **CSRF**: `github.com/gorilla/csrf`

### Utilities
- **Validation**: `github.com/go-playground/validator/v10`
- **UUID**: `github.com/google/uuid`
- **Logging**: `github.com/sirupsen/logrus` or `log/slog` (Go 1.21+)

### External Services
- **Email**: `github.com/sendgrid/sendgrid-go` or standard `net/smtp`
- **SMS**: `github.com/twilio/twilio-go`

### Optional Enhancements
- **Query Builder**: `github.com/Masterminds/squirrel`
- **ORM Alternative**: `github.com/jmoiron/sqlx` (lighter than full ORM)
- **Testing**: `github.com/stretchr/testify`

---

## Notes

- Start simple, iterate based on actual usage
- Focus on the repair shop workflow first
- Keep templates clean and maintainable
- Use htmx for all dynamic interactions to minimize JavaScript
- Prioritize data integrity and reliability
- Make it easy to backup and restore data
- Consider multi-tenancy if planning to serve multiple shops

---

## Questions to Answer Before Building

1. **Single shop or multi-tenant?** (affects database schema)
2. **What payment processing is needed?** (Stripe, Square, etc.)
3. **Inventory management?** (track parts stock levels)
4. **Customer portal?** (let customers check ticket status)
5. **Barcode/QR scanning?** (for device tracking)
6. **Email/SMS budget?** (affects notification strategy)

---

This TODO provides a roadmap from the current scaffold to a production-ready repair shop helpdesk system. Start with Phase 1 to get core functionality working, then iterate based on real-world needs.
