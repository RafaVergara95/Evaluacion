package main

type Producto struct {
	Id                int64  `json:"id"`
	Nombre            string `json:"nombre"`
	Tipo              string `json:"tipo"`
	Precio            int    `json:"precio"`
	Fecha_actualizada string `json:"fecha_actualizada"`
}
