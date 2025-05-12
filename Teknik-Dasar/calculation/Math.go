package calculation

func Math(number1 int, number2 int) int { //Fungsi ini bisa di panggil ke package yang berbeda, karena huruf depanya kapital (artinya fungsi public)
	return minus(number1, number2) //Fungsi add bisa di panggil di dalam package yang sama, tetapi fungsi add tidak bisa di panggi di package yang berbeda
}

func minus(n1 int, n2 int) int { //Kalau ini hanya bisa di panggil di package yang sama saja, karena huruf epanya tidak kapital (artinya private)
	return n1 - n2
}

func add(n1 int, n2 int) int {
	return n1 + n2
}

func Multiple(number1 int, number2 int) int {
	return number1 * number2
}
