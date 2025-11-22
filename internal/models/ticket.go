package models

import "time"

// Ticket represents a repair ticket in the system
type Ticket struct {
	ID          int       `json:"id"`
	Status      string    `json:"status"` // new, in_progress, waiting_parts, ready, completed
	Priority    string    `json:"priority"` // low, normal, high, urgent

	// Customer information
	CustomerName  string `json:"customer_name"`
	CustomerPhone string `json:"customer_phone"`
	CustomerEmail string `json:"customer_email"`

	// Device information
	DeviceType   string `json:"device_type"`
	DeviceBrand  string `json:"device_brand"`
	DeviceModel  string `json:"device_model"`
	SerialNumber string `json:"serial_number"`

	// Repair details
	IssueDescription string  `json:"issue_description"`
	InternalNotes    string  `json:"internal_notes"`
	EstimatedCost    float64 `json:"estimated_cost"`

	// Assignment and scheduling
	AssignedTo string     `json:"assigned_to"`
	DueDate    *time.Time `json:"due_date"`

	// Metadata
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`

	// Related data (loaded via joins)
	Parts          []Part     `json:"parts,omitempty"`
	Notes          []WorkNote `json:"notes,omitempty"`
	TotalPartsCost float64    `json:"total_parts_cost"`
}

// Part represents a replacement part or material used in a repair
type Part struct {
	ID       int     `json:"id"`
	TicketID int     `json:"ticket_id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Cost     float64 `json:"cost"`
	AddedAt  time.Time `json:"added_at"`
	AddedBy  string  `json:"added_by"`
}

// WorkNote represents a work log entry or note on a ticket
type WorkNote struct {
	ID        int       `json:"id"`
	TicketID  int       `json:"ticket_id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Timestamp time.Time `json:"timestamp"`
}

// Customer represents customer information (for future use)
type Customer struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	TotalOrders int       `json:"total_orders"`
}

// Technician represents a repair technician user
type Technician struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	ActiveJobs   int    `json:"active_jobs"`
	IsAvailable  bool   `json:"is_available"`
}

// TicketStats represents dashboard statistics
type TicketStats struct {
	OpenTickets     int `json:"open_tickets"`
	InProgress      int `json:"in_progress"`
	Overdue         int `json:"overdue"`
	CompletedToday  int `json:"completed_today"`
}

// StatusClass returns the Tailwind CSS class for the ticket status badge
func (t *Ticket) StatusClass() string {
	switch t.Status {
	case "new":
		return "bg-blue-100 text-blue-800"
	case "in_progress":
		return "bg-yellow-100 text-yellow-800"
	case "waiting_parts":
		return "bg-orange-100 text-orange-800"
	case "ready":
		return "bg-green-100 text-green-800"
	case "completed":
		return "bg-gray-100 text-gray-800"
	default:
		return "bg-gray-100 text-gray-800"
	}
}

// StatusDisplay returns a human-readable status string
func (t *Ticket) StatusDisplay() string {
	switch t.Status {
	case "new":
		return "New"
	case "in_progress":
		return "In Progress"
	case "waiting_parts":
		return "Waiting for Parts"
	case "ready":
		return "Ready for Pickup"
	case "completed":
		return "Completed"
	default:
		return t.Status
	}
}

// TotalCost calculates the total cost including parts and estimated labor
func (t *Ticket) TotalCost() float64 {
	return t.EstimatedCost + t.TotalPartsCost
}

// IsOverdue checks if the ticket is past its due date
func (t *Ticket) IsOverdue() bool {
	if t.DueDate == nil {
		return false
	}
	return time.Now().After(*t.DueDate) && t.Status != "completed"
}
