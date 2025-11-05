package main

import "fmt"

type Category struct {
	ID   int
	Name string
}

type Supplier struct {
	ID    int
	Name  string
	Email string
	Phone string
}

type Product struct {
	ID       int
	Name     string
	Price    float64
	Stock    int
	Category Category
	Supplier Supplier
}

type Sale struct {
	ID       int
	Product  Product
	Quantity int
	Total    float64
}

var categories []Category
var suppliers []Supplier
var products []Product
var sales []Sale

func main() {
	for {
		fmt.Println("\n=== SISTEMA DE GESTION DE INVENTARIO ===")
		fmt.Println("1. Crear Categoria")
		fmt.Println("2. Crear Proveedor")
		fmt.Println("3. Crear Producto")
		fmt.Println("4. Listar Productos")
		fmt.Println("5. Modificar Stock de Producto")
		fmt.Println("6. Registrar Venta")
		fmt.Println("7. Listar Ventas")
		fmt.Println("0. Salir")
		fmt.Print("Seleccione una opcion: ")
		opcion := readInt()
		switch opcion {
		case 1:
			createCategory()
		case 2:
			createSupplier()
		case 3:
			createProduct()
		case 4:
			listProducts()
		case 5:
			updateProductStock()
		case 6:
			createSale()
		case 7:
			listSales()
		case 0:
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opcion inválida.")
		}
	}
}

func readString() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func readInt() int {
	var num int
	fmt.Scan(&num)
	return num
}

func readFloat() float64 {
	var num float64
	fmt.Scan(&num)
	return num
}

func createCategory() {
	fmt.Print("ID Categoria: ")
	id := readInt()
	fmt.Print("Nombre Categoria: ")
	name := readString()
	categories = append(categories, Category{ID: id, Name: name})
	fmt.Println("Categoria creada.")
}

func createSupplier() {
	fmt.Print("ID Proveedor: ")
	id := readInt()
	fmt.Print("Nombre: ")
	name := readString()
	fmt.Print("Email: ")
	email := readString()
	fmt.Print("Telefono: ")
	phone := readString()
	suppliers = append(suppliers, Supplier{ID: id, Name: name, Email: email, Phone: phone})
	fmt.Println("Proveedor creado.")
}

func createProduct() {
	fmt.Print("ID Producto: ")
	id := readInt()
	fmt.Print("Nombre Producto: ")
	name := readString()
	fmt.Print("Precio: ")
	price := readFloat()
	fmt.Print("Stock Inicial: ")
	stock := readInt()
	fmt.Println("Seleccione Categoria por ID:")
	listCategories()
	catID := readInt()
	var cat Category
	for _, c := range categories {
		if c.ID == catID {
			cat = c
		}
	}
	fmt.Println("Seleccione Proveedor por ID:")
	listSuppliers()
	supID := readInt()
	var sup Supplier
	for _, s := range suppliers {
		if s.ID == supID {
			sup = s
		}
	}
	products = append(products, Product{ID: id, Name: name, Price: price, Stock: stock, Category: cat, Supplier: sup})
	fmt.Println("Producto creado.")
}

func listProducts() {
	fmt.Println("\n--- LISTA DE PRODUCTOS ---")
	for _, p := range products {
		fmt.Printf("ID:%d | %s | Precio: %.2f | Stock: %d | Categoria:%s | Proveedor:%s\n",
			p.ID, p.Name, p.Price, p.Stock, p.Category.Name, p.Supplier.Name)
	}
}

func listCategories() {
	for _, c := range categories {
		fmt.Printf("ID:%d - %s\n", c.ID, c.Name)
	}
}

func listSuppliers() {
	for _, s := range suppliers {
		fmt.Printf("ID:%d - %s\n", s.ID, s.Name)
	}
}

func updateProductStock() {
	fmt.Print("Ingrese ID del producto: ")
	id := readInt()
	for i := range products {
		if products[i].ID == id {
			fmt.Print("Cantidad a agregar o quitar: ")
			amount := readInt()
			products[i].Stock += amount
			fmt.Println("Stock actualizado.")
			return
		}
	}
	fmt.Println("Producto no encontrado.")
}

func createSale() {
	fmt.Print("ID Venta: ")
	id := readInt()
	fmt.Print("ID Producto: ")
	pid := readInt()
	for _, p := range products {
		if p.ID == pid {
			fmt.Print("Cantidad: ")
			q := readInt()
			sale := Sale{ID: id, Product: p, Quantity: q}
			sale.Total = float64(q) * p.Price
			sales = append(sales, sale)
			fmt.Println("Venta registrada.")
			return
		}
	}
	fmt.Println("Producto no encontrado.")
}

func listSales() {
	fmt.Println("\n--- LISTA DE VENTAS ---")
	for _, s := range sales {
		fmt.Printf("ID Venta:%d | Prod:%s | Cant:%d | Total: %.2f\n",
			s.ID, s.Product.Name, s.Quantity, s.Total)
	}
}
