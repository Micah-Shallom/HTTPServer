package main

import (
	"errors"
	f "fmt"
	"log"
	"net/http"
	"context"
	"net"
)

const keyServerAddr = "ServerAddr"

func formHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	f.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
	if err := r.ParseForm(); err != nil {
		f.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	f.Fprintf(w, "Post request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	f.Fprintf(w, "Name= %s\n", name)
	f.Fprintf(w, "Address= %s\n", address)
	}  
	
func mainHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	f.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))

	if r.URL.Path != "/main" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	f.Fprintf(w, "Hello!!")
}

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/", fileServer)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/main", mainHandler)

	// if err := http.ListenAndServe(":8000", nil); err != nil {
		// 	log.Fatal(err)
		// }
		
		ctx, cancelCtx := context.WithCancel(context.Background())
		
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	
	serverTwo := &http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	
	go func(){
		f.Printf("Starting server on port 3333\n")
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed){
			f.Printf("Server One has been shutdown")
			}else if err != nil {
			f.Printf("Error Starting Server One")
			log.Fatal(err)
		}
		cancelCtx()
		}()
		
		go func(){
			f.Printf("Starting server on port 4444\n")
			err := serverTwo.ListenAndServe()
			if errors.Is(err, http.ErrServerClosed){
				f.Printf("Server Two has been shutdown")
				}else if err != nil {
					f.Printf("Error Starting Server Two")
			log.Fatal(err)
		}
		cancelCtx()
	}()

	<- ctx.Done()

	// err := http.ListenAndServe(":8000", mux)
	// if errors.Is(err, http.ErrServerClosed){ //checks if the server was shutdown
	// 	f.Printf("Server has been shutdown")
	// }else if err != nil {
	// 	f.Printf("Error starting server")
	// 	log.Fatal(err)
	// }
}
