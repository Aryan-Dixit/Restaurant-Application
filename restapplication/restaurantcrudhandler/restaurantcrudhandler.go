package restaurantcrudhandler

import (
	"MAD_Rest_Assign/dbrepository"
	"MAD_Rest_Assign/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	logger "log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	restauranterrors "MAD_Rest_Assign/restapplication/packages/errors"
	"MAD_Rest_Assign/restapplication/packages/httphandlers"
	mthdroutr "MAD_Rest_Assign/restapplication/packages/mthdrouter"
	"MAD_Rest_Assign/restapplication/packages/resputl"
)

type RestaurantCrudHandler struct {
	httphandlers.BaseHandler
	mongoSession *dbrepository.MongoRepository
}

func NewRestaurantCrudHandler(mongoSession *dbrepository.MongoRepository) *RestaurantCrudHandler {
	return &RestaurantCrudHandler{mongoSession: mongoSession}
}

func (p *RestaurantCrudHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

func (p *RestaurantCrudHandler) Get(r *http.Request) resputl.SrvcRes {

	queryParams, ok := r.URL.Query()["typeOfFood"]
	logger.Print(queryParams)
	if ok {
		if queryParams[0] != "" {
			resp, err := p.mongoSession.FindByTypeOfFood(queryParams[0])
			if err != nil {
				return resputl.ReponseCustomError(err)
			}
			res := transformobjListToResponse(resp)
			return resputl.Response200OK(res)
		} else {
			return resputl.SimpleBadRequest("Invalid Input Data")
		}
	}
	queryParams, ok = r.URL.Query()["name"]
	logger.Print(queryParams)
	if ok {
		if queryParams[0] != "" {
			resp, err := p.mongoSession.FindByName(queryParams[0])
			if err != nil {
				return resputl.ReponseCustomError(err)
			}
			res := transformobjListToResponse(resp)
			return resputl.Response200OK(res)
		} else {
			return resputl.SimpleBadRequest("Invalid Input Data")
		}
	}
	queryParams, ok = r.URL.Query()["postcode"]
	logger.Print(queryParams)
	if ok {
		if queryParams[0] != "" {
			resp, err := p.mongoSession.FindByTypeOfPostCode(queryParams[0])
			if err != nil {
				return resputl.ReponseCustomError(err)
			}
			res := transformobjListToResponse(resp)
			return resputl.Response200OK(res)
		} else {
			return resputl.SimpleBadRequest("Invalid Input Data")
		}
	}

	queryParams, ok = r.URL.Query()["search"]
	logger.Print(queryParams)
	if ok {
		if queryParams[0] != "" {
			resp, err := p.mongoSession.Search(queryParams[0])
			if err != nil {
				return resputl.ReponseCustomError(err)
			}
			res := transformobjListToResponse(resp)
			return resputl.Response200OK(res)
		} else {
			return resputl.SimpleBadRequest("Invalid Input Data")
		}
	}
	pathParam := mux.Vars(r)
	usID := pathParam["id"]
	if usID == "" {
		resp, err := p.mongoSession.GetAll()

		if err != nil {
			return resputl.ReponseCustomError(err)
		}

		responseObj := transformobjListToResponse(resp)

		return resputl.Response200OK(responseObj)
	}

	NewusID := domain.StringToID(usID)
	obj, err := p.mongoSession.Get(NewusID)

	if err != nil {
		return resputl.ProcessError(restauranterrors.NotFoundError("User Object Not found"), "")
	}

	return resputl.Response200OK(obj)

}
func (p *RestaurantCrudHandler) Post(r *http.Request) resputl.SrvcRes {
	fmt.Println("hello")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resputl.ReponseCustomError(err)
	}

	e, err := ValidateRestaurantCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		return resputl.SimpleBadRequest("Invalid Input Data")

	}
	logger.Printf("Received POST request to Create schedule %s ", string(body))
	var requestdata *RestaurantCreateReqDTO
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	f := dbrepository.Factory{}
	userObj := f.NewUser(requestdata.Name, requestdata.Address, requestdata.AddressLine2, requestdata.URL, requestdata.Outcode, requestdata.Postcode, requestdata.Rating, requestdata.Type_Of_Food)
	id, err := p.mongoSession.Store(userObj)
	Newid := bson.ObjectId(id).Hex()
	if err != nil {
		logger.Fatalf("Error while creating in DB: %v", err)
		return resputl.ProcessError(restauranterrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK(&RestaurantCreateRespDTO{ID: Newid})
}

func (p *RestaurantCrudHandler) Put(r *http.Request) resputl.SrvcRes {
	restaurantid := mux.Vars(r)["id"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	e, err := ValidateRestaurantCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		return resputl.SimpleBadRequest("Invalid Input Data")

	}
	logger.Printf("Received PUT request to Create schedule %s ", string(body))
	var requestdata *RestaurantCreateReqDTO
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	f := dbrepository.Factory{}
	userObj := f.NewUser(requestdata.Name, requestdata.Address, requestdata.AddressLine2, requestdata.URL, requestdata.Outcode, requestdata.Postcode, requestdata.Rating, requestdata.Type_Of_Food)
	err = p.mongoSession.Update(userObj, restaurantid)
	if err != nil {
		logger.Fatalf("Error while creating in DB: %v", err)
		return resputl.ProcessError(restauranterrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK("Succesfully Updated")
}

func (p *RestaurantCrudHandler) Delete(r *http.Request) resputl.SrvcRes {
	pathParam := mux.Vars(r)
	usID := pathParam["id"]
	if usID == "" {
		return resputl.Response200OK("No Object ID")
	} else {
		NewusID := domain.StringToID(usID)
		err := p.mongoSession.Delete(NewusID)
		if err != nil {
			return resputl.ProcessError(restauranterrors.NotFoundError("User Object Not found"), "")
		}

		return resputl.Response200OK("Data Succesfully Deleted")

	}

}
