package main

import (
	"apiRest/models"
	_ "apiRest/routers"
	"apiRest/utils"
	"log"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "github.com/mattn/go-sqlite3"
)

func CORSMiddleware(ctx *context.Context) {
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ctx.Output.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Output.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if ctx.Input.Method() == "OPTIONS" {
		ctx.Output.SetStatus(200)
		ctx.Output.Body([]byte(""))
		return
	}
}
func init() {
	// Registra el driver de SQLite3
	if err := orm.RegisterDriver("sqlite3", orm.DRSqlite); err != nil {
		log.Fatalf("Error al registrar el driver de SQLite3: %v", err)
	} else {
		log.Println("Driver de SQLite3 registrado exitosamente.")
	}

	// Registra el modelo de datos
	orm.RegisterModel(new(models.Consumidor), new(models.Producto), new(models.Orden))

	// Configura y registra la conexión a la base de datos
	dbDriver, _ := beego.AppConfig.String("dbDriver")
	dbConnStr, _ := beego.AppConfig.String("dbConnStr")
	if err := orm.RegisterDataBase("default", dbDriver, dbConnStr); err != nil {
		log.Fatalf("Error al registrar la base de datos: %v", err)
	} else {
		log.Println("Base de datos registrada exitosamente.")
	}
}

func main() {
	o := orm.NewOrm()
	// Verificar si ya hay registros en la tabla de productos
	existingRecords, err := o.QueryTable(new(models.Producto)).Count()
	if err != nil {
		log.Fatalf("Error al consultar la base de datos: %v", err)
	}
	// Si no hay registros, sincronizar la base de datos e insertar datos de prueba
	if existingRecords == 0 {
		if err := orm.RunSyncdb("default", false, true); err != nil {
			log.Fatalf("Error al sincronizar la base de datos: %v", err)
		} else {
			log.Println("Base de datos sincronizada exitosamente.")
			utils.InsertarDatosPrueba()
		}
	} else {
		log.Println("La base de datos ya tiene registros. No se sincronizará ni se insertarán datos de prueba.")
	}
	beego.InsertFilter("*", beego.BeforeRouter, CORSMiddleware)
	beego.Run()
}
