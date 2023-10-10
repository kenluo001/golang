/*
 * @Author: Ken Luo ken.luo.5227@gmail.com
 * @Date: 2023-10-10 11:23:52
 * @FilePath: /golang/base/main.go
 */

package main

	import (
		"net/http"
		"fmt"
	)
	func main()  {
		http.Handle("/", http.FileServer(http.Dir(".")))
		http.ListenAndServe(":8080", nil)
		fmt.Println("hello")
	}