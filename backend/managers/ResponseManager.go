package managers

func GetBaseResponseObject() map[string]interface{} {
	response := make(map[string]interface{})
	response["success"] = false
	response["message"] = "Failed. Something went wrong!"
	response["errors"] = make(map[string]interface{})
	return response
}
