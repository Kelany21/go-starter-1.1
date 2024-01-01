package helpers

import (
	"fmt"
	"github.com/bykovme/gotrans"
)

func DoneUpdate(moduleName string) string {
	return fmt.Sprintf(gotrans.Tr(GetCurrentLang(), "done_update_item"), gotrans.Tr(GetCurrentLang(), moduleName))
}

func DoneDelete(moduleName string) string {
	return fmt.Sprintf(gotrans.Tr(GetCurrentLang(), "done_delete_item"), gotrans.Tr(GetCurrentLang(), moduleName))
}

func DoneActivationPack(activateCount int, deactivateCount int) string {
	if activateCount != 0 && deactivateCount != 0 {
		return gotrans.Tr(GetCurrentLang(), "activated_pack") + gotrans.Tr(GetCurrentLang(), "and") + gotrans.Tr(GetCurrentLang(), "deactivated_pack")
	} else if activateCount == 0 {
		return gotrans.Tr(GetCurrentLang(), "deactivated_pack")
	} else {
		return gotrans.Tr(GetCurrentLang(), "activated_pack")
	}
}

func DoneActivate() string {
	return gotrans.Tr(GetCurrentLang(), "activated")
}

func DoneTrash() string {
	return gotrans.Tr(GetCurrentLang(), "trashed")
}

func DoneDeactivate() string {
	return gotrans.Tr(GetCurrentLang(), "deactivated")
}

func DoneGetItem() string {
	return gotrans.Tr(GetCurrentLang(), "done_get_item")
}

func DoneCreateItem(moduleName string) string {
	return fmt.Sprintf(gotrans.Tr(GetCurrentLang(), "done_created_item"), gotrans.Tr(GetCurrentLang(), moduleName))
}

func DoneGetAllItems() string {
	return gotrans.Tr(GetCurrentLang(), "get_all_items")
}

func ItemNotFound() string {
	return gotrans.Tr(GetCurrentLang(), "item_not_found")
}

func Wrong() string {
	return gotrans.Tr(GetCurrentLang(), "wrong")
}
