package category

// Category represents a category for a budget
type Category struct {
	ID              string `json:"id"`
	CategoryGroupID string `json:"category_group_id"`
	Name            string `json:"name"`
	Hidden          bool   `json:"hidden"`
	// Budgeted Budgeted amount in current month in milliunits format
	Budgeted int64 `json:"budgeted"`
	// Activity Activity amount in current month in milliunits format
	Activity int64 `json:"activity"`
	// Balance Balance in current month in milliunits format
	Balance int64 `json:"balance"`
	// Deleted Deleted category groups will only be included in delta requests
	Deleted bool `json:"deleted"`

	Note *string `json:"note"`
	// OriginalCategoryGroupID If category is hidden this is the ID of the category
	// group it originally belonged to before it was hidden
	OriginalCategoryGroupID *string `json:"original_category_group_id"`
}

// Group represents a resumed category group for a budget
type Group struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Hidden  bool   `json:"hidden"`
	Deleted bool   `json:"deleted"`
}

// GroupWithCategories represents a category group for a budget
type GroupWithCategories struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Hidden bool   `json:"hidden"`
	// Deleted Deleted category groups will only be included in delta requests
	Deleted bool `json:"deleted"`

	Categories []*Category `json:"categories"`
}