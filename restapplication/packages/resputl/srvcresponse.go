package resputl

import (
	"encoding/json"
	"fmt"
	"net/http"

	resterrors "MAD_Rest_Assign/restapplication/packages/errors"
)

type EmptyStruct struct {
}

type SrvcRes struct {
	Code     int
	Response interface{}
	Message  string
	Headers  map[string]string
}

func marshalResponse(r interface{}) ([]byte, error) {
	return json.MarshalIndent(r, "", "")
}

func (s *SrvcRes) RenderResponse(w http.ResponseWriter) {
	if s.Headers == nil {
		s.Headers = map[string]string{"Content-Type": "application/json",
			"Access-Control-Allow-Headers": "source,x-authorization-token,Content-Type",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*"}
	}
	for h, val := range s.Headers {
		w.Header().Set(h, val)
	}
	var statusBool bool
	switch s.Code {
	case http.StatusOK:
		statusBool = true
	default:
		statusBool = false
	}

	formatted := map[string]interface{}{
		"responseData": s.Response,
		"message":      s.Message,
		"status":       statusBool}

	data, _ := marshalResponse(formatted)
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	w.WriteHeader(s.Code)
	fmt.Fprint(w, string(data))
}

func Simple200OK(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusOK, inf, message, nil}
}

func Simple404Response(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusNotFound, inf, message, nil}
}

func Simple422Response(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusUnprocessableEntity, inf, message, nil}
}

func PreconditionFailed(message string) SrvcRes {
	var inf EmptyStruct
	return SrvcRes{http.StatusPreconditionFailed, inf, message, nil}
}

func OptionsResponseOK(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusOK, inf, message, nil}
}

func SimpleBadRequest(message string) SrvcRes {
	return SrvcRes{http.StatusBadRequest, "{}", message, nil}
}

func Response200OK(response interface{}) SrvcRes {
	return SrvcRes{http.StatusOK, response, "OK", nil}

}

func ResponseNotImplemented(response interface{}) SrvcRes {
	return SrvcRes{http.StatusNotImplemented, "{}", "Method not implementd", nil}
}

func ReponseCustomError(err error) SrvcRes {
	var inf EmptyStruct
	cusErr := err.(*resterrors.CustomError)
	return SrvcRes{cusErr.GetStatusCode(), inf, cusErr.GetMessage(), nil}
}

func ReponseInternalError() SrvcRes {
	var inf EmptyStruct
	return SrvcRes{http.StatusInternalServerError, inf, "Internal Server Error", nil}
}

func ProcessError(err error, inf interface{}) SrvcRes {

	switch err.(type) {
	case *resterrors.CustomError:
		return ReponseCustomError(err)
	default:
		return ReponseInternalError()

	}
}
