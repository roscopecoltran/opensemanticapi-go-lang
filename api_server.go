package main

/**
 * @TODO do the json encoding the proper way
 * extend the api to more possibilities
 */

import (
    // "github.com/monbro/opensemanticapi/api" // @TODO put the logic into this file
    "log"
    "github.com/codegangsta/martini"
     "github.com/monbro/opensemanticapi/database"
     "strings"
)

func main() {
    log.Println("Starting API server now.")

    // configurate the database connection
    Db := new(database.Database)
    Db.Init("", 10)

    // set up a nice glas of martini
    m := martini.Classic()

    // http://localhost:3000
    m.Get("/", func() string {
        return `{"result":"Hello world!"}`
    })

    // http://localhost:3000/relation/database
    m.Get("/relation/:word", func(params martini.Params) string {
        relations := Db.GetPopularWordRelations(params["word"])
        return `{"result": ["`+strings.Join(relations, `", "`)+`"]}`
    })

    // actual launch the server
    m.Run()
}