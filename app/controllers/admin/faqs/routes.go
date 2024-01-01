package faqs

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
)

/**
* all admin modules route will store here
 */
func Routes(r *gin.RouterGroup) *gin.RouterGroup {
	name := models.FaqRoute()
	/// all routes
	r.GET(name+_const.LIST_URL, Index)
	r.GET(name+_const.Status, Status)
	r.POST(name+_const.STORE_URL, Store)
	r.POST(name+"/store", StoreFaqWithAnswers)
	r.POST(name+"/update/:id", UpdateFaqWithAnswers)
	r.PUT(name+_const.UPDATE_URL, Update)
	r.GET(name+_const.SHOW_URL, Show)
	r.DELETE(name+_const.DELETE_URL, Delete)
	r.PATCH(name+_const.QUIQK_EDIT_URL, QuickEdit)
	//// status route
	r.PATCH(name+_const.ACTIVE_URL, Active)
	r.PATCH(name+_const.DEACTIVE_URL, DeActive)
	r.PATCH(name+_const.TRASH_URL, Trash)
	/// bulk status
	r.PATCH(name+_const.BULK_ACTIVE_URL, BulkActive)
	r.PATCH(name+_const.BULK_DEACTIVE_URL, BulkDeActive)
	r.PATCH(name+_const.BULK_TRASH_URL, BulkTrash)
	r.PATCH(name+_const.BULK_DELETE_URL, BulkDelete)

	answer := models.AnswerRoute()
	r.POST(answer+_const.STORE_URL, StoreAnswer)
	r.PUT(answer+_const.UPDATE_URL, UpdateAnswer)

	//customs Url
	r.GET(name+"/list", GetAllFaqs)
	r.GET(name+"/show-faq/:id", ShowFaq)
	return r
}
