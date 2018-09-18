package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	bento := bentoSelect() // 弁当を選んでもらい弁当IDを取得する
	orderTime := orderTime()
	countBento := "#productsForm_" + bento + " > div.push > input" // あとで使う弁当セレクト用のcssセレクタの値をここで設定する
	orderBento := "#productsForm_" + bento + " > div.btn_cartin"
	driver := agouti.ChromeDriver( // ChromeDriverをheadlessで設定。
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
	page.Navigate("https://nettomotto.jp/?sitetype=d") // ネット注文の画面を立ち上げ
	page.FindByName("id").Fill("")                     // loginする用の個人情報
	page.FindByName("password").Fill("")
	// formをサブミット
	if err := page.FindByClass("login_btn").Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}
	page.Find("body > main > section.service_block > div.service_box01 > div.service_takeout > p > a > img").Click()
	page.Find("#form5 > input").Click()
	time.Sleep(1 * time.Second) // 1秒待つ
	page.FindByName("target_hour").Fill(orderTime[0])
	page.FindByName("target_minutes").Fill(orderTime[1])
	time.Sleep(1 * time.Second) // 1秒待つ
	page.Find("#row_1520 > td:nth-child(15) > input[type=\"image\"]").Click()
	time.Sleep(1 * time.Second) // 1秒待つ
	page.Find(countBento).Fill("1")
	time.Sleep(1 * time.Second) // 1秒待つ
	page.Find(orderBento).Click()
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

func bentoSelect() string {
	fmt.Println("お好みの弁当をお選びください")
	fmt.Println("0 : ", "肉野菜炒め弁当")
	fmt.Println("1 : ", "特のりタル弁当")
	fmt.Println("2 : ", "ロースかつとじ弁当")
	fmt.Println("3 : ", "極うま親子丼")

	var bento string
	_, err := fmt.Scanln(&bento)
	if err != nil {
		log.Fatal(err)
	}

	switch bento { // 弁当のID値を返却
	case "0":
		bento = "2372"
	case "1":
		bento = "5588"
	case "2":
		bento = "2368"
	case "3":
		bento = "2644"
	}
	return bento
}

func orderTime() []string {
	fmt.Println("注文したい時間を『10:20』の形式で入力ください。現在時間の30分後からのご予約が可能となります。")
	var orderTime string
	_, err := fmt.Scanln(&orderTime)
	if err != nil {
		log.Fatal(err)
	}
	orderTimeAr := strings.Split(orderTime, ":")
	return orderTimeAr
}
