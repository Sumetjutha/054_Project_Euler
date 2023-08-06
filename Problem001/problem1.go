package main

import "fmt"

// ฟังก์ชัน multiplesOf3and5 หาผลรวมของจำนวนที่เป็นเลขคู่ของ 3 หรือ 5 ที่น้อยกว่า number
func multiplesOf3and5(number int) int {
	totalSum := 0                 // กำหนดตัวแปร totalSum เพื่อเก็บผลรวม
	for i := 1; i < number; i++ { // วนลูปทุกตัวเลขตั้งแต่ 1 จนถึง number-1
		if i%3 == 0 || i%5 == 0 { // ตรวจสอบว่าตัวเลข i เป็นเลขคู่ของ 3 หรือ 5 หรือไม่
			totalSum += i // หากเป็นจริง ให้เพิ่มค่า i เข้าไปในผลรวม totalSum
		}
	}
	return totalSum // คืนค่าผลรวมของจำนวนที่เป็นเลขคู่ของ 3 หรือ 5 ที่น้อยกว่า number
}

func main() {
	limitValue := 10                       // กำหนดค่า limitValue เป็น 10
	result := multiplesOf3and5(limitValue) // เรียกใช้ฟังก์ชัน multiplesOf3and5 และเก็บผลลัพธ์ในตัวแปร result
	fmt.Printf("ผลรวมของจำนวนที่เป็นเลขคู่ของ 3 หรือ 5 ที่น้อยกว่า %d คือ: %d\n", limitValue, result)
}
