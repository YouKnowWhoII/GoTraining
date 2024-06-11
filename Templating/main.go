package main

import (
	"os"
	"text/template"
	"time"
)

// Sales data structure
type Sale struct {
	Date   time.Time
	Amount float64
}

// Employee performance data structure
type Employee struct {
	Name             string
	PerformanceScore int
	Sales            []Sale
}

// Product inventory data structure
type Product struct {
	Name     string
	Quantity int
	Price    float64
}

// Report data structure to hold all types of data
type ReportData struct {
	Sales      []Sale
	Employees  []Employee
	Products   []Product
	ReportDate time.Time
}

// Custom functions
func formatDate(t time.Time) string {
	return t.Format("02-Jan-2006")
}

func totalSales(sales []Sale) float64 {
	total := 0.0
	for _, sale := range sales {
		total += sale.Amount
	}
	return total
}

func totalInventoryValue(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += float64(product.Quantity) * product.Price
	}
	return total
}

func main() {
	// Dummy data
	sales := []Sale{
		{Date: time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC), Amount: 1000.00},
		{Date: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), Amount: 1500.00},
	}

	employees := []Employee{
		{Name: "John Doe", PerformanceScore: 95, Sales: sales},
		{Name: "Jane Smith", PerformanceScore: 88, Sales: sales},
	}

	products := []Product{
		{Name: "Widget A", Quantity: 100, Price: 10.50},
		{Name: "Widget B", Quantity: 50, Price: 20.75},
	}

	reportData := ReportData{
		Sales:      sales,
		Employees:  employees,
		Products:   products,
		ReportDate: time.Now(),
	}

	// Parse templates
	funcMap := template.FuncMap{
		"formatDate":          formatDate,
		"totalSales":          totalSales,
		"totalInventoryValue": totalInventoryValue,
	}

	tmpl, err := template.New("reports").Funcs(funcMap).ParseGlob("Templating/templates/*.template")
	if err != nil {
		panic(err)
	}

	// Generate sales report
	salesFile, err := os.Create("Templating/reports/sales_report.txt")
	if err != nil {
		panic(err)
	}
	defer salesFile.Close()
	tmpl.ExecuteTemplate(salesFile, "sales", reportData)

	// Generate employee performance report
	employeeFile, err := os.Create("Templating/reports/employee_report.txt")
	if err != nil {
		panic(err)
	}
	defer employeeFile.Close()
	tmpl.ExecuteTemplate(employeeFile, "employee", reportData)

	// Generate product inventory report
	inventoryFile, err := os.Create("Templating/reports/inventory_report.txt")
	if err != nil {
		panic(err)
	}
	defer inventoryFile.Close()
	tmpl.ExecuteTemplate(inventoryFile, "inventory", reportData)
}
