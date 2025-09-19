package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/BR7T/BookShelf/models"
)

func GetBookByISBN(ISBN string , w http.ResponseWriter , r *http.Request)(error){
	urlOpenLibrary := fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:%v&format=json&jscmd=details" , ISBN)
	
	var book models.BookResponse

	resp , err := http.Get(urlOpenLibrary)
	if err != nil{
		return fmt.Errorf("erro ao buscar isbn")
	}
	defer resp.Body.Close()

	body , err := io.ReadAll(resp.Body)
	if err != nil{
		return fmt.Errorf("erro ao ler resposta: %v" , err)
	}


	err = json.Unmarshal(body , &book)
	if err != nil{
		return fmt.Errorf("erro ao decodificar JSON: %v" , err)
	}
	
	for _  , bookData := range book{
		fmt.Printf("Title: %s \n" , bookData.Details.Title)
	}
	fmt.Print(urlOpenLibrary , "\n")
	fmt.Print("teste")



	
	return nil
}