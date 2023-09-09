package backend

import "AKO/ch/teko/ako/controller"

func main() {
	handler := controller.MakeApi()
	handler.Run()
}
