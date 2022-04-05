# NYCU.ONE

因為 Dcard 2022 Backend Intern 的關係，把之前用 Python 寫的[短網址服務](https://ntust.me)重寫。

除了短網址外，之後也打算加入 Text Snippet、File、限時、密碼（甚至串 WebAuthn？） 等等功能（因為自己常常會用到XD） 我覺得短網址其實是瑞士刀的概念，各種「用一個網址」能解決的事情都很適合放上來，甚至像是留言板、Link Tree 等等，不過需要好好規劃過。

## 懺悔

因為 4/4 才注意到實習的資訊，加上對 Golang 不熟悉，所以目前架構很單純（aka 技術債一堆）。 

## 架構說明

Backend: Golang (Gin)
Testing: go-sqlmock + testify
DB: MySQL

Shorten Algorithm: concat(hex(id) + randString(len=3)) 

不知道思路有什麼好講的，因為這次我做的架構很單純，就只是把參數收進來然後存入資料庫而已，基本的 CRUD。

選擇 testify 單純只是想把我的程式碼變簡短，以這個專案來說，其實它沒有幫到我太多，也可以直接手刻 assert.Equals 。 選擇 go-sqlmock 則是因為需要靠他來幫我 mock db & sql query，讓我的測試可以專注在 Controller 邏輯本身。 因為程式碼本身沒有很複雜，就沒有另外再拆出 Model。

縮短網址的演算法，有考慮過用 Hash 的方式做，但因為時間有限，後來想說用 Auto Increment ID 偷懶一下，但因為只有 ID (int) 又太醜，所以就用 ID 當作 Namespace 搭配隨機產生的三個字元去組合出縮短後（localhost/1ABC）的部份。 有把部分不好閱讀的字元避開。另外因為十進位 ID 會長很快，所以用 16 進位。 另外，選擇這個方式也是因為，方便之後想到更好的演算法，可以去相容。

## 待解問題

- 對資料庫頻繁存取
    - Redis
    - 縮短網址的演算法有沒有可能改成某種 hash function？
- 事前的網址資安掃描
- 網址的移除規則（如果有惡意程式等等）

## To Do

參考 TODO.md