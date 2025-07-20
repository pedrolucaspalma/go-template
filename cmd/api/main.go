package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Just an example handler that may call an usecase. Nothing important here
func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Print no caso de uso, por exemplo")
	w.WriteHeader(201)
	fmt.Fprintf(w, "OK!")
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(mainHandler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		mainHandler = middlewares[i](mainHandler)
	}
	return mainHandler
}

func LoggingMiddleware(anyValue string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// ... all logging logic here
			fmt.Printf("Printou no middleware de log: %s\n", anyValue)
			next(w, r)
		}
	}
}

func AuthMiddleware() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Printou no início do middleware de auth")
			// .. all auth logic here
			jwt := "not valid JWT"
			isJwtValid := jwt == "valid JWT"
			if !isJwtValid {
				w.WriteHeader(401)
				fmt.Println("Printou no meio do middleware de auth. O user não foi autenticado e o fluxo precisa ser interrompido.")
				fmt.Fprint(w, "User could not be authenticated")
				return
			}
			fmt.Println("Printou no final do middleware de auth. O user foi autenticado")
			next(w, r)
		}
	}
}

type ValidationFunc func(w http.ResponseWriter, r *http.Request) []FieldResult
type FieldResult struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	// I cannot use the error type here because this is sent via JSON, and json encoders dont automatically call Error() method
	Errors []string `json:"errors"`
}

func ValidationMiddleware(v ValidationFunc) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			invalidFields := []FieldResult{}
			results := v(w, r)
			for _, r := range results {
				fieldIsValid := len(r.Errors) == 0
				if fieldIsValid {
					continue
				}
				invalidFields = append(invalidFields, r)
			}

			allFieldsAreValid := len(invalidFields) == 0
			if allFieldsAreValid {
				next(w, r)
				return
			}

			w.WriteHeader(422)
			json.NewEncoder(w).Encode(APIResponse{
				Message: "Validation Error",
				Data:    invalidFields,
			})
		}
	}
}

func MockFailedValidation(w http.ResponseWriter, r *http.Request) []FieldResult {
	return []FieldResult{
		{
			Name:     "email",
			Location: "body",
			Errors:   []string{"insufficient length", "invalid format"},
		},
	}
}

type APIResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty,omitzero"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /testing", Chain(MainHandler, LoggingMiddleware("qlqr coisa")))
	mux.HandleFunc("GET /auth-testing", Chain(MainHandler, LoggingMiddleware("qlqr coisa"), AuthMiddleware()))
	mux.HandleFunc("GET /auth-testing-no-log", Chain(MainHandler, AuthMiddleware(), LoggingMiddleware("qlqr coisa")))
	mux.HandleFunc("GET /validation-testing", Chain(MainHandler, LoggingMiddleware("qlqr coisa"), ValidationMiddleware(MockFailedValidation)))
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
