package main

import (
	"encoding/json"
	"net/http"
)

const (
	UsersApi    = "https://jsonplaceholder.typicode.com/users"
	PhotosApi   = "https://jsonplaceholder.typicode.com/photos"
	CommentsApi = "https://jsonplaceholder.typicode.com/comments"
)

type (
	ApiResponse struct {
		StatCode    uint32      `json:"stat_code"`
		StatMessage string      `json:"stat_message"`
		Data        interface{} `json:"data,omitempty"`
	}

	UsersList struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Address  struct {
			Street  string `json:"street"`
			Suite   string `json:"suite"`
			City    string `json:"city"`
			ZipCode string `json:"zipcode"`
			Geo     struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			} `json:"geo"`
		} `json:"address"`
		Phone   string `json:"phone"`
		Website string `json:"website"`
		Company struct {
			Name        string `json:"name"`
			CatchPhrase string `json:"catchPhrase"`
			Bs          string `json:"bs"`
		}
	}

	PhotosList struct {
		AlbumID      int    `json:"albumId"`
		ID           int    `json:"id"`
		Title        string `json:"title"`
		Url          string `json:"url"`
		ThumbnailUrl string `json:"thumbnailUrl"`
	}

	CommentsList struct {
		PostID int    `json:"postId"`
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Body   string `json:"body"`
	}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		JSONResponse(w, ApiResponse{StatCode: http.StatusOK, StatMessage: "Ping"})
	})
	http.HandleFunc("/users", UsersHandler)
	http.HandleFunc("/photos", PhotosHandler)
	http.HandleFunc("/comments", CommentsHandler)

	http.ListenAndServe(":3000", nil)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []UsersList{}
	res, err := http.Get(UsersApi)

	if err != nil {
		JSONResponse(w, ApiResponse{StatCode: http.StatusUnprocessableEntity, StatMessage: err.Error()})
		return
	}

	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		JSONResponse(w, ApiResponse{StatCode: http.StatusUnprocessableEntity, StatMessage: err.Error()})
		return
	}

	JSONResponse(w, ApiResponse{StatCode: http.StatusOK, StatMessage: "Successfully", Data: users})
}

func PhotosHandler(w http.ResponseWriter, r *http.Request) {
	photos := []PhotosList{}
	res, err := http.Get(PhotosApi)

	if err != nil {
		JSONResponse(w, ApiResponse{StatCode: http.StatusUnprocessableEntity, StatMessage: err.Error()})
		return
	}

	if err := json.NewDecoder(res.Body).Decode(&photos); err != nil {
		JSONResponse(w, ApiResponse{StatCode: http.StatusUnprocessableEntity, StatMessage: err.Error()})
		return
	}

	JSONResponse(w, ApiResponse{StatCode: http.StatusOK, StatMessage: "Successfully", Data: photos})
}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	comments := []CommentsList{}
	res, err := http.Get(CommentsApi)

	if err != nil {
		JSONResponse(w, ApiResponse{StatCode: http.StatusUnprocessableEntity, StatMessage: err.Error()})
		return
	}

	if err := json.NewDecoder(res.Body).Decode(&comments); err != nil {
		JSONResponse(w, ApiResponse{StatCode: http.StatusUnprocessableEntity, StatMessage: err.Error()})
		return
	}

	JSONResponse(w, ApiResponse{StatCode: http.StatusOK, StatMessage: "Successfully", Data: comments})
}

func JSONResponse(w http.ResponseWriter, api ApiResponse) {
	resBytes, err := json.Marshal(&api)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte("Parse json response failed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(api.StatCode))
	w.Write(resBytes)
}
