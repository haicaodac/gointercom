package gointercom

// IntercomUser ..
type IntercomUser struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Name   string `json:"name"`
}

// IntercomUsers ..
type IntercomUsers struct {
	Type       string            `json:"type"`
	TotalCount int               `json:"total_count"`
	Users      []IntercomUser    `json:"users"`
	Pages      IntercomPaginator `json:"pages"`
}

// IntercomConversation ...
type IntercomConversation struct {
	Assignee            map[string]interface{} `json:"assignee"`
	User                map[string]interface{} `json:"user"`
	ConversationMessage map[string]interface{} `json:"conversation_message"`
	ConversationParts   map[string]interface{} `json:"conversation_parts"`
	ConversationRating  map[string]interface{} `json:"conversation_rating"`
	CreatedAt           string                 `json:"created_at"`
	UpdatedAt           string                 `json:"updated_at"`
	Type                string                 `json:"type"`
}

// IntercomPaginator ...
type IntercomPaginator struct {
	Next       string `json:"next"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	TotalPages int    `json:"total_pages"`
}

// IntercomAuthor ...
type IntercomAuthor struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// IntercomNote ...
type IntercomNote struct {
	Type      string         `json:"type"`
	ID        int            `json:"id"`
	CreatedAt string         `json:"created_at"`
	Body      string         `json:"body"`
	Author    IntercomAuthor `json:"author"`
	User      IntercomUser   `json:"user"`
}

// IntercomNotes ...
type IntercomNotes struct {
	Type  string            `json:"type"`
	Pages IntercomPaginator `json:"pages"`
	Notes []IntercomNote    `json:"notes"`
}

// Response ...
type Response struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// Admin ...
type Admin struct {
	Type  string `json:"type"`
	Email string `json:"email"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

// AdminList ...
type AdminList struct {
	Type   string  `json:"type"`
	Admins []Admin `json:"admins"`
}
