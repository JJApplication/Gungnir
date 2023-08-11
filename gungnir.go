/*
   Create: 2023/8/3
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

func Run(root string) {
	h := handleFs(root)
	serveMux(h)
}
