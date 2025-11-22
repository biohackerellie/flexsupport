# AI PROJECT SCAFFOLDING PARAMETERS

## Project Overview
You are helping scaffold a new web application called **RepairDesk** — a lightweight helpdesk/ticketing system tailored initially for a **repair shop** workflow.

There are two primary user roles:
1. **Employee / Front Desk Staff** — creates, edits, and manages tickets and customer info.
2. **Repair Technician** — manages active repair jobs, tracks status, parts, costs, and deadlines.

The goal is to provide a **clear, functional frontend scaffold** and a **well-organized backend plan**, not necessarily a complete production implementation.

---

## Tech Stack & Frameworks
- **Language:** Go (Golang)
- **Templating:** `html/template`
- **Frontend Interactivity:** `htmx` for dynamic updates, `Alpine.js` for interactive components when needed
- **UI Library Support:** allowed to use [`rotmh/shadcn-templ`](https://github.com/rotmh/shadcn-templ) for UI components
- **Storage/DB:** not yet implemented — just produce TODO items and planning guidance
- **Routing:** Go HTTP or `chi` router (you may suggest)
- **Build Goal:** Produce well-organized, human-readable code structure, focusing on layout and extensibility

---

## Primary Deliverables

### 1. Frontend Scaffold
Generate well-structured templates and page scaffolding that includes:
- `base.tmpl` or equivalent for shared layout (nav, title, etc.)
- Templates for:
  - Ticket listing / dashboard (employee view)
  - Ticket entry form / parts input form
  - Technician's detail view with job progress and actions
- htmx components for dynamic sections like:
  - “Mark as completed”
  - “Update status” dropdowns
  - “Add part” inline updates
- Example use of `shadcn-templ` components (buttons, modals, tables)
- Optionally, Alpine snippets for simple UI logic (e.g., modal toggles)

### 2. Backend TODO / Planning Document
Provide a clear high-level TODO list describing:
- Proposed directory structure (`/cmd`, `/internal`, `/templates`, `/static`, etc.)
- Core data models/entities:
  - Customer
  - RepairTicket (status, cost, dueDate, assignedTo)
  - Parts
- Suggested Go route/handler design
- Integration points for frontend templates
- Example endpoint definitions (stub form handlers, update routes)
- Suggestions for extending later (authentication, email, reporting, etc.)

---

## Design & UX Focus
The design should be:
- Clean, minimal, functional
- Favor tables and cards for organizing ticket info
- Emphasize legibility and usability for repair techs
- Responsive layout (works well on desktop/tablet)

---

## Agent Behavior & Output Style
- Start by summarizing the structure you’ll generate.
- Then output Go template and example Go code inline in fenced blocks.
- Use helpful comments to explain file purposes.
- Be practical rather than exhaustive; focus on a functional starting point.
- Reference but do not duplicate external component code unless vital.

---

## Example Interaction Flow
1. The user (Ellie) requests the scaffold.
2. The agent outputs:
   - A project overview summary
   - Suggested directory layout
   - Example template files (`base.tmpl`, `dashboard.tmpl`, etc.)
   - A TODO checklist for the backend
3. Ellie reviews and requests deeper code generation or iteration.

---

## Constraints
- Do not generate sensitive or copyrighted material.
- Focus on functionality and clarity, not lorem-ipsum-heavy mock data.
- Keep Go code idiomatic and templates minimal but extendable.
- Use plain Markdown fenced code blocks for all code output (no HTML formatting).

---

## Goals
Deliver a **ready-to-expand** scaffold that gives Ellie:
- A visual, usable starting frontend
- A solid understanding of what needs to be built next
- A foundation for iterative development of a helpdesk/ticket app

---

*End of parameter specification.*
