# GoCCUHours
中正大學自動填寫時數(兼任助理、學習型臨時工)的自動化腳本終端工具
![terminal](https://i.imgur.com/vobpyUl.png)

# 需求(Requirements)
* `Google Chrome`版本請升級至77.0以上
* 下載對應作業系統的`chromedriver`放置到同級目錄
  * [**下載連結**](https://chromedriver.storage.googleapis.com/index.html?path=77.0.3865.40/)

# 安裝(Install) 
## 方法 I
直接下載以編譯好的最新檔案
* [**下載連結**](https://github.com/curtis992250/GoCCUHours/releases)

## 方法 II
有golang環境者可自行編譯
```cmd
git clone https://github.com/curtis992250/GoCCUHours.git

cd CCUHours
```

### Windows
```cmd
go build -o CCUHours.exe .
```

### Mac & Linux
```shell
go build -o CCUHours .
```

# 使用(Usage)
* windows
  * ```cmd
    CCUHours.exe
    ```
  * 或是雙擊`CCUHours.exe`執行

* Mac & Linux
  * ```cmd
    ./CCUHours
    ```

# 詳細使用方法
* AllowAction 裡有明確定義合法的`action`
* 除了AllowAction有數字(`1,2,3...`子選單可進入)外，皆使用`action`作為開頭
* `action`分別為`set`(設置),`add`(添加),`rm`(刪除),`run`(執行功能),`exit`(離開選單)

## set
設置項目的值
```cmd
set UserName 你的帳號
```
or UserName對應的編號為`1`
```cmd
set 1 你的帳號
```

## add 
添加項目(用於添加工作事項、日期、時段等)
```cmd
add mywork 工作項目1
add work10 工作項目2
add something 工作項目3 
```
* 已有的項目不可被添加


## rm
刪除已有的項目
```cmd
rm mywork
```
or `work1`對應編號為`1`
```cmd
rm 1
```



# Known Issue
* 目前手邊無**勞僱型臨時工**帳號可測試，所以無法填寫此身分的時數