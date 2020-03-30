package response

import (
	"fmt"
	"encoding/json"
	"net/http"

	"spookify/model"
)

// HandleSuccess represent process of convert the standard response message
func HandleSuccess(resp http.ResponseWriter, data interface{}) {
	// Wrapper for response when success
	returnData := model.Resp {
		Status: "SUCCESS!",
		Message: "Process ended without error.",
		Data:    data,
	}

	// Convert the response to form json
	jsonData, err := json.Marshal(returnData)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipHandler.handleSuccess] Error when do json marshalling for error handling: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

// HandleError represent process of convert the standard response message
func HandleError(resp http.ResponseWriter, message string) {
	// Wrapper for response when error
	data := model.Resp {
		Status: "Fail",
		Message: message,
	}

	// Convert the response to form json
	jsonData, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipHandler.handleError] Error when do json marshalling for error handling  %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}