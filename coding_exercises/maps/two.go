package main

import "fmt"

func main() {
	m1 := map[int]bool{}
	m1[5] = true // cannot assign bool entry to nil initiated

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	// fmt.Println(m2 == m3) // only if nil
	// if m2 == nil && m3 == nil {
	// 	fmt.Println(m2 == m3)
	// }
	if len(m2) == len(m3) {
		for k, v := range m2 {
			// for j, g := range m3 {
			fmt.Printf("key: %#v, m2 val: %#v, m3 val: %#v, key/val equal: %#v\n", k, m2[k], m3[k], m2[k] == m3[k])
			_ = v
			// }
		}

	}
}
