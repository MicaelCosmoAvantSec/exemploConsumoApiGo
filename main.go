package main

// Bibliotecas
import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
	"os"
)

// Estrutura base para POST/UPDATE
type Post struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Body  string `json:"body"`
}

// Salva saída em arquivo.json com nome personalizado
func _saveJSONResponseToFile(responseBody []byte, fileName string) error {
    file, err := os.Create(fmt.Sprintf("./response/%sResponse.json", fileName))
    if err != nil {
        return err
    }
    defer file.Close()

    err = ioutil.WriteFile(fmt.Sprintf("./response/%sResponse.json", fileName), responseBody, 0644)
    if err != nil {
        return err
    }

    return nil
}

// Realiza uma requisição GET
func getJson() {
    getResponse, err := http.Get("https://jsonplaceholder.typicode.com/posts")
    if err != nil {
        fmt.Printf("Erro ao realizar a requisição GET: %v\n", err)
    } else {
        defer getResponse.Body.Close()
        responseBody, _ := ioutil.ReadAll(getResponse.Body)
        //fmt.Println(string(responseBody))
		_saveJSONResponseToFile(responseBody, "get")
    }
}

// Realiza uma requisição GET com filtro.
func _doGetRequestWithFilter(filter string) (*http.Response, error) {
    url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?%s", filter)
    return http.Get(url)
}

// Faz chamada da requisição GET, passando parâmetro a ser filtrado
func getJsonWithFilter() {
	getResponse, err := _doGetRequestWithFilter("userId=1")
    if err != nil {
        fmt.Printf("Erro ao realizar a requisição GET: %v\n", err)
    } else {
        defer getResponse.Body.Close()
        responseBody, _ := ioutil.ReadAll(getResponse.Body)
        //fmt.Println(string(responseBody))
		_saveJSONResponseToFile(responseBody, "getWithFilter")
    }
}

// Realiza uma requisição POST
func postJson() {
    postRequestData := Post{Title: "Título do post", Body: "Corpo do post"}
    postRequestBody, _ := json.Marshal(postRequestData)
    postResponse, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postRequestBody))
    if err != nil {
        fmt.Printf("Erro ao realizar a requisição POST: %v\n", err)
    } else {
        defer postResponse.Body.Close()
        postResponseBody, _ := ioutil.ReadAll(postResponse.Body)
        //fmt.Println(string(postResponseBody))
		_saveJSONResponseToFile(postResponseBody, "post")
    }

}

// Realiza uma requisição PUT (UPDATE)
func updateJson() {
    putRequestData := Post{Title: "Título do post atualizado", Body: "Corpo do post atualizado"}
    putRequestBody, _ := json.Marshal(putRequestData)
    putRequest, _ := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(putRequestBody))
    putRequest.Header.Set("Content-Type", "application/json")
    putResponse, err := http.DefaultClient.Do(putRequest)
    if err != nil {
        fmt.Printf("Erro ao realizar a requisição PUT: %v\n", err)
    } else {
        defer putResponse.Body.Close()
        putResponseBody, _ := ioutil.ReadAll(putResponse.Body)
        //fmt.Println(string(putResponseBody))
		_saveJSONResponseToFile(putResponseBody, "update")
    }

}

// Realiza uma requisição DELETE
func deleteJson() {
    deleteRequest, _ := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/posts/1", nil)
    deleteResponse, err := http.DefaultClient.Do(deleteRequest)
    if err != nil {
        fmt.Printf("Erro ao realizar a requisição DELETE: %v\n", err)
    } else {
        defer deleteResponse.Body.Close()
        deleteResponseBody, _ := ioutil.ReadAll(deleteResponse.Body)
        //fmt.Println(string(deleteResponseBody))
		_saveJSONResponseToFile(deleteResponseBody, "delete")
    }
}


// Chamada de todas as funções, 
func main() {
	getJson()
	getJsonWithFilter()
	postJson()
	updateJson()
	deleteJson()
}
