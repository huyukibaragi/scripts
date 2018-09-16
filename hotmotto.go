// main.go
package main

import (
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("https://nettomotto.jp/?sitetype=d"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	identity := page.FindByName("id") // ID, Passの要素を取得し、値を設定
	password := page.FindByName("password")
	identity.Fill("mail")
	password.Fill("pass")
	// formをサブミット
	if err := page.FindByClass("login_btn").Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}
	page.Find("body > main > section.service_block > div.service_box01 > div.service_takeout > p > a > img").Click()
	page.Find("#form5 > input").Click()
	page.Find("#row_1520 > td:nth-child(15) > input[type=\"image\"]").Click()

}
