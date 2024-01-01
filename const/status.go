package _const

/// status
const (
	ACTIVE     = "activate"
	DEACTIVATE = "deactivate"
	BLOCK      = "block"
	ALL        = "all"
	TRASH      = "trash"
	DELETE     = "delete"
	/// quick edit
	QUIQK_EDIT_URL = "/quick-edit/:id"
	UPDATE_URL     = "/:id"
	STORE_URL      = ""
	SHOW_URL       = "/show/:id"
	LIST_URL       = ""
	DELETE_URL     = "/:id"
	/// stander url for change status
	ACTIVE_URL   = "/active/:id"
	DEACTIVE_URL = "/de-active/:id"
	TRASH_URL    = "/trash/:id"
	BLOCK_URL    = "/block/:id"
	/// stander bulck status change
	BULK_ACTIVE_URL   = "/bulk-active"
	BULK_DEACTIVE_URL = "/bulk-de-active"
	BULK_TRASH_URL    = "/bulk-trash"
	BULK_BLOCK_URL    = "/bulk-block"
	BULK_DELETE_URL   = "/bulk-delete"
	Status            = "/status"
)

type IDS struct {
	Ids []int `json:"ids"`
}

func AllStatusOnly() map[string]interface{} {
	///// define default actions
	m := make(map[string]interface{})
	m["all"] = ALL
	return m
}

func AllAndActiveStatus() map[string]interface{} {
	///// define default actions
	m := make(map[string]interface{})
	m["all"] = ALL
	m["active"] = ACTIVE

	return m
}

func AllAndActiveAndDeactiveStatus() map[string]interface{} {
	///// define default actions
	m := make(map[string]interface{})
	m["all"] = ALL
	m["active"] = ACTIVE
	m["deactivate"] = DEACTIVATE
	m["trash"] = TRASH

	return m
}
