package main

import (
	"fmt"

	"github.com/sp4reeee/go-recaptcha-v3-bypass/bypass"
)

func main() {
	fmt.Println(bypass.Bypass("https://www.google.com/recaptcha/api2/anchor?ar=1&k=6Lf9KKUeAAAAAEBlSoTWibwj_zwjKlXeXJNmIWKZ&co=aHR0cHM6Ly9pbWFnZXJlY29nbml6ZS5jb206NDQz&hl=fr&v=SglpK98hSCn2CroR0bKRSJl5&size=invisible&cb=1yguiv8yq1q9", ""))
}
