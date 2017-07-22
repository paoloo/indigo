package main

import (
	"fmt"
	"encoding/json"
  "io/ioutil"
  "os"
)

type DataBase struct {
    Name      string `json:"databasename"`
    Tables map[string]map[string]string `json:"tables"`
}

func (d *DataBase) Init(_name string) {
    d.Name = _name
    d.Tables = map[string]map[string]string{}
}

func (d *DataBase) Set(_table string, _key string, _value string) {
    dt, ok := d.Tables[_table]
    if !ok {
        dt = make(map[string]string)
        d.Tables[_table] = dt
    }
    d.Tables[_table][_key]=_value
}

func (d *DataBase) Get(_table string, _key string) (string){
  if _key == "" {
    bytes, _ := json.Marshal(d.Tables[_table])
    return string(bytes)
  } else {
    return d.Tables[_table][_key]
  }
}

func (d *DataBase) Del(_table string, _key string) {
     delete(d.Tables[_table], _key)
}

func (d *DataBase) Dump() (string){
     bytes, _ := json.MarshalIndent(d,""," ")
     return string(bytes)
}

func (d *DataBase) getFileName() (string){
  return (d.Name + ".JSONdb")
}

func (d *DataBase) Load(){
  _, err := os.Stat(d.getFileName())
  if os.IsNotExist(err) {
    file, err := os.Create(d.Name)
    if err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
    }
    defer file.Close()
  }
  raw, err := ioutil.ReadFile(d.getFileName())
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
  json.Unmarshal(raw, &d)
}

func (d *DataBase) Store(){
  err := ioutil.WriteFile(d.getFileName(), []byte(d.Dump()), 0644)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
}
