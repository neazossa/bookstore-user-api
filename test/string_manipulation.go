package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "pq: duplicate key value violates unique constraint \"m_user_index_2\" "

	if strings.Contains(str, "m_user_index_2") {
		fmt.Println("error duplicate email")
	}
}
