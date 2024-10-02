## 說明
此為抓取抓取房價 的爬蟲server 
需要申請
1. telegram的bot token 將其填寫到config/config.go中的 Telegram_token 並填入自己的chatid
當
    591(https://sale.591.com.tw/)
或
    實價登陸比價王(https://www.houseprice.tw/)
有新的物件產生時 即可透過Telegram即時通知

2. 修改要查的區域
    ### 2.1. repository/buyscraper/getNewItem (實價登陸比價王)
        修改要查詢的區域 及 產品類型
    ### 2.2. repository/scraper/getNewItem (591)
        修改要查詢的區域 及 產品類型
##

## 使用方法
### 抓取必要套件
go get .

### 執行程式
go run main.go
