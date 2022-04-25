package main

func createProductos(producto Producto) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO productos (nombre, tipo, precio, fecha_actualizada) VALUES (?, ?, ?, ?)", producto.Nombre, producto.Tipo, producto.Precio, producto.Fecha_actualizada)
	return err
}

func deleteProductos(id int64) error {

	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM productos WHERE id = ?", id)
	return err
}

// It takes the ID to make the update
func updateProductos(producto Producto) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE productos SET  precio = ?, fecha_actualizada = ? WHERE tipo = ?", producto.Precio, producto.Fecha_actualizada, producto.Tipo)
	return err
}
func getProductos() ([]Producto, error) {
	//Declare an array because if there's error, we return it empty
	Productos := []Producto{}
	bd, err := getDB()
	if err != nil {
		return Productos, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT id, nombre, tipo, precio, fecha_actualizada FROM productos")
	if err != nil {
		return Productos, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var Producto Producto
		err = rows.Scan(&Producto.Id, &Producto.Nombre, &Producto.Tipo, &Producto.Precio, &Producto.Fecha_actualizada)
		if err != nil {
			return Productos, err
		}
		// and append it to the array
		Productos = append(Productos, Producto)
	}
	return Productos, nil
}

func getProductosById(id int64) (Producto, error) {
	var supermercado Producto
	bd, err := getDB()
	if err != nil {
		return supermercado, err
	}
	row := bd.QueryRow("SELECT id, nombre, tipo, precio, fecha_actualizada FROM productos WHERE id = ?", id)
	err = row.Scan(&supermercado.Id, &supermercado.Nombre, &supermercado.Tipo, &supermercado.Precio, &supermercado.Fecha_actualizada)
	if err != nil {
		return supermercado, err
	}
	// Success!
	return supermercado, nil
}
