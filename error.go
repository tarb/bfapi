package bfapi

import (
	json "encoding/json"

	"github.com/tarb/util/www"
)

//
type APIError struct {
	Faultcode   string `json:"faultcode"`
	Faultstring string `json:"faultstring"`
	Detail      struct {
		Exceptionname         string    `json:"exceptionname"`
		AccountAPINGException Exception `json:"AccountAPINGException"`
		APINGException        Exception `json:"APINGException"`
	} `json:"detail"`
}

//
type Exception struct {
	RequestUUID  string `json:"requestUUID"`
	ErrorCode    string `json:"errorCode"`
	ErrorDetails string `json:"errorDetails"`
}

func (e APIError) Error() string {
	if e.Detail.AccountAPINGException.ErrorCode != "" {
		return e.Detail.AccountAPINGException.ErrorCode
	} else if e.Detail.APINGException.ErrorCode != "" {
		return e.Detail.APINGException.ErrorCode
	} else {
		return e.Faultstring
	}
}

func statusToAPINGException(err error) error {
	if err == nil {
		return nil
	}

	serr, ok := err.(www.StatusError)
	if !ok {
		return err
	}

	if len(serr.Body) > 0 {
		var apiErr APIError

		jsonErr := json.Unmarshal(serr.Body, &apiErr)
		if jsonErr != nil {
			// return the original err, not the json err
			return err
		}

		return apiErr
	}

	return err
}
