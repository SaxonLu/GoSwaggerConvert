Go-Swagger Demo
=============

目前進度

- [x] Header Token 金鑰
- [x] 01dev 連結
- [x] Try It Out Button On Work
- [x] Request Example Value | Model
- [x] Expected Response Model
- [x] API 分類 Tags (目前顯示default)
- [ ] Try It Out Button Disabled

安裝
-----

請依系統下載對應的安裝檔案 [GitHub路徑](https://github.com/go-swagger/go-swagger/releases/tag/v0.21.0)
* * *
※Windows → swagger_windows_amd64.exe
* * *
下載後設定環境變數即可
* * *

撰寫
-----
- 要點
 - 整體都是以註解方式進行
 - // 後需空格

- Header部分

  

|Swagger |語意 |
|-------------------------------|-----------------------------|
|`Package main Go-Swagger MenuAPI`|package _ Title名稱 |
|`簡單的 go-swagger` |Title解釋 |
|`Host: 01cx.txuntest.com`|此API連結目標|
|`Version: 0.0.1`|此次Swagger文件版本|
|`Security:`|啟用認證|
|`SecurityDefinitions:`|認證設定|
|` swagger:meta`|Header結尾符|

  

- Struct部分

|Swagger |語意 |
|-------------------------------|-----------------------------|
|`swagger:model userData`|型別(model parameters response) _ 用途|
|`in:body` |未填寫則以param方式呈現|
|`json:"user_name"`|轉型成json讓swagger讀取，如需採用param型態則可以不用|

  

- Function部分

|Swagger |語意 |
|-------------------------------|-----------------------------|
|`swagger:route post /api Tags fdReq`|RestfulAPI型態 _ Request網址 _ 自定義分類 _ Request Model|
|`Responses: 200: fdResp`|Response代碼 _ Response Model|

 - 產生檔案
   - 於根目錄 執行 `swagger_windows_amd64.exe generate spec -o ./swagger.json`即可

 - 使用Docker部屬
   - 於終端機輸入 `docker pull swaggerapi/swagger-ui` 下載檔案
   - `docker run -p 80:8080 -e SWAGGER_JSON=/mapi/swagger.json -v c:/work/go/src/meeline/mapi swaggerapi/swagger-ui` 部屬
   - 其中 80:8080 → 啟動埠號:8080 EX:9010:8080 → localhost:9010
   - -v 後的路徑則為映像檔取得Swagger.json的路徑
    