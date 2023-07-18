package main

import "fmt"

func contoh1(){ //eksperimen mengibah isi slice di variabel baru
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]
	
	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]
	
	fmt.Println("\n")
	// Buah "grape" diubah menjadi "pinnaple"
	aFruits[0] = "pinnaple"
	fruits[1] = "pete"
	
	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]
	return
}

func contoh2(){ //eksperimen len() dan cap()
	var fruits = []string{"apple", "grape", "banana", "pineapple"}
	var bFruits = fruits[0:2]

	fmt.Println(cap(bFruits)) // 4
	fmt.Println(len(bFruits)) // 2
	
	var cFruits = append(bFruits, "pear", "peach")
	var dFruits = append(bFruits, "snakefruit", "avocado", "dragonfruit")
	
	fmt.Println(fruits) // [apple grape pear peach]
	fmt.Println(cFruits) // [apple grape pear peach]
	fmt.Println(dFruits) // [apple grape snakefruit avocado dragonfruit]
	return
}

func contoh3(){ //eksperimen copy(target, source)
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	new := make([]string, 5)
	u := copy(new, dst)
	fmt.Println(new) // [watermelon pinnaple apple  ] (slot kosong 2)
	_ = append(new[:3], "4", "5")

	o := copy(dst, new)
	
	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(new) // watermelon pinnaple apple 4 5
	fmt.Println(n)   // 3 (hanya mengambil sebanyak len(dst))
	fmt.Println(u)	// 3 (mengambil semua pada dst -> 3 (sisa 2 slot))
	fmt.Println(o) // 3 (len(dst) tidak bertambah)

	dst = []string{"potato", "potato", "potato"}
	src = []string{"melon", "pear"}
	n = copy(dst, src)
	
	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2
}

func contoh4(){ //slice 3 indeks (a[start:until before:capacity])
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]
	
	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3
	
	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3
	
	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
}

func main(){
	//contoh1()
	//contoh2()
	//contoh3()
	contoh4()
}