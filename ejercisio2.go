package main

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Producto struct {
	Nombre   string
	Cantidad int
	Precio   float64
}

func main() {
	var productos []Producto

	fmt.Print("Ingrese el nombre del cliente: ")
	var nombreCliente string
	fmt.Scanln(&nombreCliente)

	fechaFactura := time.Now().Format("02/01/2006")

	for {
		fmt.Print("Ingrese el nombre del producto (o 'fin' para terminar): ")
		var nombreProducto string
		fmt.Scanln(&nombreProducto)
		if nombreProducto == "fin" {
			break
		}
		cantidad, precio := pedirCantidadYPrecio()

		productos = append(productos, Producto{Nombre: nombreProducto, Cantidad: cantidad, Precio: precio})
	}

	// Generar el PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Factura")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Cliente: %s", nombreCliente))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Fecha: %s", fechaFactura))
	pdf.Ln(12)

	// Encabezados
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(70, 10, "Producto", "1", 0, "", false, 0, "")
	pdf.CellFormat(30, 10, "Cantidad", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Precio Unitario", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Total", "1", 0, "C", false, 0, "")
	pdf.Ln(10)

	// Contenido de la tabla
	pdf.SetFont("Arial", "", 12)
	var totalGeneral float64
	for _, p := range productos {
		totalProducto := float64(p.Cantidad) * p.Precio
		totalGeneral += totalProducto
		pdf.CellFormat(70, 10, p.Nombre, "1", 0, "", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf("%d", p.Cantidad), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", p.Precio), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", totalProducto), "1", 0, "C", false, 0, "")
		pdf.Ln(10)
	}

	// Total general
	pdf.SetFont("Arial", "B", 12)
	pdf.Ln(10)
	pdf.CellFormat(70, 10, "", "", 0, "", false, 0, "")
	pdf.CellFormat(30, 10, "", "", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Total General:", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", totalGeneral), "1", 0, "C", false, 0, "")

	// Guardar el archivo PDF
	err := pdf.OutputFileAndClose("factura.pdf")
	if err != nil {
		fmt.Println("Error al generar el PDF:", err)
	} else {
		fmt.Println("Factura generada exitosamente: factura.pdf")
	}
}

// pedirCantidadYPrecio solicita la cantidad y el precio unitario de un producto y valida su entrada
func pedirCantidadYPrecio() (int, float64) {
	var cantidad int
	for {
		fmt.Print("Ingrese la cantidad: ")
		_, err := fmt.Scanln(&cantidad) // Usa Scanln para leer la cantidad
		if err == nil && cantidad > 0 {
			break
		}
		fmt.Println("Cantidad no válida, intente nuevamente.")
	}

	var precio float64
	for {
		fmt.Print("Ingrese el precio unitario: ")
		_, err := fmt.Scanln(&precio) // Usa Scanln para leer el precio
		if err == nil && precio > 0 {
			break
		}
		fmt.Println("Precio no válido, intente nuevamente.")
	}

	return cantidad, precio
}
