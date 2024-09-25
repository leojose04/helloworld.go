package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func imprimir() {
	fmt.Println("texto desde la funcion imprimir")
}

func hola_string(s string) string {
	return "hola " + s
}

func sumar(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("a y b son iguales")
	} else {
		fmt.Println("a y b son diferentes")
	}

}

func arreglo() {
	var aprendices [5]string
	aprendices[0] = "sebastian"
	aprendices[1] = "leonardo"
	aprendices[2] = "franck"
	aprendices[3] = "juam jose"
	aprendices[4] = "jaider"
	fmt.Print(aprendices[1])
}

func tipo_datos() {
	var texto string = "leonardo"
	var numero int = 1000
	var decimal float64 = 0.00001
	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))

}

func converstringboolean () {
	var palabra string = "true"
	boolean, err := strconv.ParseBool(palabra)
	fmt.Println(boolean,reflect.TypeOf(boolean))
	fmt.Println("este es el error", err )
}

func convertstring(){
	var palabra_bool bool = true
	strbool := strconv.FormatBool(palabra_bool)
	fmt.Println(strbool,reflect.TypeOf(strbool))


}

func ejercisio15 () {
	var nombre, apellido string = "Leonardo", "Cruz"
	fmt.Println(nombre, apellido)
}
func ejercisios15 (){
	var (
		nombre  string = "Leonardo"
		edad int = 20
		pensionado bool = false
	)
	fmt.Println("nombre: ", nombre)
	fmt.Println("edad: ", edad)
	fmt.Println("pensionado: ", pensionado)
}
func ejercisio16 () {
	var nombre string
	var edad int
	var peso float64
	var estudiante bool
	fmt.Println("nombre: ", nombre)
	fmt.Println("edad: ", edad)
	fmt.Println("peso: ", peso)
	fmt.Println("estudiante: ", estudiante)

}

func ejercisio17 () {
	nombre := "leonardo da vinci"
	edad := 70
	peso := 75
	estudiante := false
	fmt.Println("nombre: ", nombre)
	fmt.Println("edad: ", edad)
	fmt.Println("peso: ", peso)
	fmt.Println("estudiante: ", estudiante)
}

func ejercisios17 () {
	var profesion = "Deportista"
	sueldo := 1000000
	fmt.Println("profesion: ", profesion)
	fmt.Println("sueldo: ", sueldo)
}

func ejercisio18 () {
	var var1 = "este es el nivel1"
	var var2 = "este es el nivel2"
	{
		var var3 = "este es el nivel3"
        fmt.Println(var3)
	}
	fmt.Println(var1)
	fmt.Println(var2)
}

func ejercisio19 () {
	color := "rojo"
	fmt.Println(color, &color)
}
func ejercisios19 () {
	vehiculo1 := "rojo"
	fmt.Println("el vehiculo1 es", vehiculo1)

	vehiculo2 := vehiculo1
	fmt.Println("el vehiculo2 es", vehiculo2)

	vehiculo3 := &vehiculo1
	fmt.Println("el vehiculo es", *vehiculo3)

	vehiculo1 = "gris"

	// fmt.Println("el vehiculo1 ahora es", vehiculo1, &vehiculo1)
	// fmt.Println("el vehiculo2 sigue siendo", vehiculo2, &vehiculo2)
	// fmt.Println("el vehiculo3 sigue apuntando a", *vehiculo3, vehiculo3)
}

//func ejercisio20 (altura float32) float32 {
//	altura = altura * 3.28
//	return altura
//}
//var altura float32 = 1.70

func ejercisio20_2(altura *float32) float32 {
    *altura = *altura - 0.10
    return *altura
}

var altura float32 = 1.70





func main() {
	//fmt.Println("Texto desde la funcion main")
	//imprimir()
	//fmt.Println(hola_string("leonardo"))
	//fmt.Println(sumar(3, 5))
	//comparar_bool()
	//arreglo()
	//tipo_datos()
	//converstringboolean()
	//convertstring()
	//ejercisio15()
	//ejercisios15()
	//ejercisio16()
	//ejercisio17()
	//ejercisios17()
	//ejercisio18()
	//ejercisio19()
	//ejercisios19()
	//fmt.Println("la altura es:", altura, "mts")
    //fmt.Println("la altura es:", altura, "mts")
    //fmt.Println("la nueva altura es:", altura, "mts")
	//fmt.Println("La altura es:", altura, "mts")
	//fmt.Println("La altura es:", ejercisio20(altura), " pies")
	//fmt.Println("La nueva altura es:", altura, "mts")
	fmt.Println("La altura es:", altura, "mts")
	fmt.Println("Al envejecer:", ejercisio20_2(&altura), "mts")
	fmt.Println("Despues de envejecer:", altura, "mts")
	//fmt.Println("el area de un circulo cuyo radio es 3 es:", area(3))

	





}
