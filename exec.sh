GOARCH=wasm GOOS=js go build -o lib.wasm 
goexec 'http.ListenAndServe(":8000", http.FileServer(http.Dir(".")))'