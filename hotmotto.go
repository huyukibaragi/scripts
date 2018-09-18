package main

import (
	"log"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",             // headlessモードの指定
			"--window-size=1280,800", // ウィンドウサイズの指定
		}),
		agouti.Debug,
	)
	if err := driver.Start(); err != nil {
		log.Fatal(err)
	}
	defer driver.Stop()
	page, err := driver.NewPage()
	if err != nil {
		log.Fatal(err)
	}
	page.Navigate("https://nettomotto.jp/?sitetype=d")
	page.FindByName("id").Fill("")
	page.FindByName("password").Fill("")
	// formをサブミット
	if err := page.FindByClass("login_btn").Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}
	page.Find("body > main > section.service_block > div.service_box01 > div.service_takeout > p > a > img").Click()
	page.Find("#form5 > input").Click()
	time.Sleep(1 * time.Second) // 1秒待つ
	page.Find("#row_1520 > td:nth-child(15) > input[type=\"image\"]").Click()
	time.Sleep(1 * time.Second) // 1秒待つ
	page.Find("#productsForm_2372 > div.push > input").Fill("1")
	time.Sleep(1 * time.Second) // 1秒待つ
	page.Find("#productsForm_2372 > div.btn_cartin").Click()
	time.Sleep(1 * time.Second) // 1秒待つ
	page.Find("#btnCart > img").Click()
	time.Sleep(1 * time.Second) // 1秒待つ
	page.Find("#cartForm > div.center > p > input[type=\"image\"]").Click()
	time.Sleep(1 * time.Second) // 1秒待つ
	page.FindByName("card_expire_month").Fill("")
	page.FindByName("card_expire_year").Fill("")
	page.FindByName("card_securitycode").Fill("")
	page.FindByName("btn_next").Click()
	//time.Sleep(4 * time.Second)   // 4秒待つ
	//page.Screenshot("Screen.png") // スクリーンショット
}
