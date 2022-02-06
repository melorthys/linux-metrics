package main

import (
	"fmt"
	"os/exec"
)

func main() {
	ram := "free | grep Mem | awk '{ printf(\"%.1f\", $3/$2 * 100.0) }'"
	cpu := "top -bn2 | grep '%Cpu' | tail -1 | grep -P '(....|...) id,'|awk '{print 100-$8}'"
	disk := "df -h --output=pcent / | awk 'NR==2{print $1}' | rev | cut -c 2- | rev"

	ram_output, err := exec.Command("/bin/sh", "-c", ram).Output()
	if err != nil {
		fmt.Println("error")
	}

	var cpu_output []byte
	cpu_output, err = exec.Command("/bin/sh", "-c", cpu).Output()
	if err != nil {
		fmt.Println("error")
	}

	var disk_output []byte
	disk_output, err = exec.Command("/bin/sh", "-c", disk).Output()
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println("ram  = % " + string(ram_output))
	fmt.Print("cpu  = % " + string(cpu_output))
	fmt.Print("disk = % " + string(disk_output))

}
