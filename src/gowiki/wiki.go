package main

import "fmt"
import "io/ioutil"

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    fileName := p.Title + ".txt"
    return ioutil.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    fileName := title + ".txt"
    body, err := ioutil.ReadFile(fileName)
    if err != nil {
        return nil, err
    }
    return &Page{title, body}, nil
}

func main() {
    p1 := &Page{"Hello", []byte("Hello World !")}
    err := p1.save()
    if err != nil {
        fmt.Println("Error while saving file with name ", p1.Title, ".txt")
    } else {
        fmt.Println("No Error while saving file with name ", p1.Title, ".txt")
    }
    p2, err := loadPage("Hello")
    if err != nil {
        fmt.Println("Error while loading file with name ", p1.Title, ".txt")
    } else {
        fmt.Println("Body of file with name ", p2.Title, ".txt -> \n", string(p2.Body))
    }
}
