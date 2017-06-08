package main

import "encoding/json"
import "os"
import "log"

type config struct {
  Name string `json:"name"`
  Type string `json:"type"`
  Data *json.RawMessage `json:"data"`
}

type  ftp struct {
  Url string `json:"url"`
}

func main(){
  file, err := os.Open("test.json")
  if err!=nil {
    log.Println("Error ",  err)
    return
  }
  defer file.Close()
  conf := config{}
  decoder := json.NewDecoder(file)
  err =  decoder.Decode(&conf)

  if err !=  nil {
    log.Println("Error decoding:"+err.Error())
  }

  if(conf.Type=="ftp"){
    f:= ftp{}
    json.Unmarshal(*conf.Data,&f)
    log.Println(f.Url)
  }
}
