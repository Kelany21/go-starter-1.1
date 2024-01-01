package config

/**
* any function will add here will not follow the
* permissions please make sure that you need this
* it will be security issue
 */
func GlobalPermissions() map[string]string {
	m := make(map[string]string)
	//m["GetUserByToken"] = "GetUserByToken"
	return m
}
