package main

import "fmt"

func fiboEvenSum(n int) int {
	sum := 0  // สร้างตัวแปร sum เพื่อเก็บผลรวมของตัวเลขที่เป็นจำนวนคู่ในลำดับฟิโบนักกี้ที่พบ
	prev := 1 // กำหนดตัวแปร prev เป็น 1 เพื่อเริ่มต้นลำดับของตัวเลขฟิโบนักกี้ด้วยตัวเลข 1
	curr := 2 // กำหนดตัวแปร curr เป็น 2 เพื่อเริ่มต้นลำดับของตัวเลขฟิโบนักกี้ด้วยตัวเลข 2

	for curr <= n { // เริ่มการทำงานของลูป for ทำงานเมื่อ curr มีค่าน้อยกว่าหรือเท่ากับ n
		if curr%2 == 0 { // ตรวจสอบว่า curr เป็นจำนวนคู่หรือไม่
			sum += curr // ถ้า curr เป็นจำนวนคู่ ก็เพิ่มค่า curr เข้าไปในผลรวม sum
		}

		next := prev + curr // คำนวณค่าใหม่ของตัวเลขฟิโบนักกี้โดยบวกค่า prev และ curr แล้วเก็บในตัวแปร next
		prev = curr         // อัปเดตค่า prev เป็น curr เพื่อเตรียมคำนวณลำดับถัดไปของฟิโบนักกี้
		curr = next         // อัปเดตค่า curr เป็น next เพื่อเตรียมคำนวณลำดับถัดไปของฟิโบนักกี้
	}

	return sum // คืนค่าผลรวมของตัวเลขที่เป็นจำนวนคู่ในลำดับของตัวเลขฟิโบนักกี้ที่มีค่าไม่เกิน n
}

func main() {
	// เรียกใช้งานฟังก์ชันและแสดงผลลัพธ์
	fmt.Println(fiboEvenSum(10))
	fmt.Println(fiboEvenSum(34))
	fmt.Println(fiboEvenSum(60))
}
