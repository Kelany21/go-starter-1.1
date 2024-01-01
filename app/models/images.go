package models

import "starter-golang-new/helpers"

/**
* request images
 */
type ImageRequest struct {
	Images []string `json:"images"`
}

/*
* return with module name
 */
func ImageModule() string {
	return helpers.ModuleName("images")
}

/*
* return with route name
 */
func ImageRoute() string {
	return helpers.ModuleRoute(ImageModule())
}


/*
* return with model name
 */
func ImageModel() string {
	return helpers.ModelModel(PageModule())
}
