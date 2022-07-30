package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Answer single answer from api.
type Answer struct {
	Err  error
	Data interface{}
}

type AnyValidate interface {
	Validate(int) error
}

// ReadJSONAndValidate json to struct and validate struct.
func ReadJSONAndValidate(request *http.Request, data AnyValidate, userID int) error {
	bData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return fmt.Errorf("readAll err: %w", err)
	}

	if err = json.Unmarshal(bData, &data); err != nil {
		return fmt.Errorf("unmarshal err: %w", err)
	}

	if err := data.Validate(userID); err != nil {
		return fmt.Errorf("validate error: %w", err)
	}

	return nil
}

// WriteJSON struct to json.
func WriteJSON(w http.ResponseWriter, res Answer) {
	if res.Err != nil {
		fmt.Printf("request err: %+v\n", res.Err)
		http.Error(w, res.Err.Error(), 400)

		return
	}

	if bData, err := json.Marshal(res.Data); err == nil {
		if _, err = w.Write(bData); err != nil {
			fmt.Printf("Write err: %+v\n", res.Err)
		}
	} else {
		fmt.Printf("Marshal err: %+v\n", err)
		http.Error(w, res.Err.Error(), 400)
	}
}
