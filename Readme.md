### **第一週：Goroutine 基礎與 Channels 初體驗 (Week 1: Goroutine Fundamentals & Introduction to Channels)**

| 天數 | 主題 | 學習重點 | 心得 |
| :--- | :--- | :--- | :--- |
| **Day 1** | **並行 (Concurrency) vs. 平行 (Parallelism)** | - 釐清這兩個在非同步編程中常被混淆的核心概念。 <br> - 了解 Golang 如何透過 Goroutine 實現並行，並在多核心處理器上達到平行處理。 | **【Golang併發之始】Concurrency還是Parallelism？別再傻傻分不清楚！** |
| **Day 2** | **初探 Goroutine** | - 學習 `go` 關鍵字，如何啟動一個 Goroutine。 <br> - 了解 Goroutine 是由 Go runtime 管理的輕量級執行緒。 <br> - 觀察主 Goroutine (main goroutine) 結束後，所有子 Goroutine 也會跟著結束的現象。 | **【Hello, Goroutine!】用一行`go`關鍵字，開啟並行的奇幻旅程** |
| **Day 3** | **使用 `sync.WaitGroup` 等待 Goroutine** | - 學習如何使用 `sync.WaitGroup` 來等待一組 Goroutine 完成任務。 <br> - 掌握 `Add`, `Done`, `Wait` 三個主要方法的使用時機。 | **【Goroutine的同步與等待】`sync.WaitGroup`：你的團隊合作好幫手** |
| **Day 4** | **Channels 入門：Goroutine 間的溝通** | - 學習 Golang 的核心理念："Do not communicate by sharing memory; instead, share memory by communicating"。 <br> - 學習使用 `make(chan Type)` 建立 Channel。 <br> - 掌握使用 `<-` 運算子發送與接收資料。 | **【Goroutine的橋樑】Channels：優雅地在並行世界中傳遞訊息** |
| **Day 5** | **無緩衝 (Unbuffered) vs. 有緩衝 (Buffered) Channels** | - 理解無緩衝 Channel 的同步阻塞特性。 <br> - 學習有緩衝 Channel 的非同步特性以及其容量限制。 <br> - 探討兩者不同的使用情境。 | **【Channels的雙面刃】無緩衝的同步之舞 vs. 有緩衝的非同步效率** |
| **Day 6** | **使用 `range` 遍歷 Channel** | - 學習如何使用 `for...range` 優雅地從 Channel 中接收資料，直到 Channel 被關閉。 <br> - 了解 `close()` 函數如何關閉一個 Channel，以及其對接收方的影響。 | **【優雅地接收】用`for...range`遍歷Channel，直到世界盡頭** |
| **Day 7** | **`select` 陳述式：多路 Channel 監聽** | - 學習 `select` 如何像 `switch` 一樣，但用於處理 Channel 操作。 <br> - 實作監聽多個 Channel，並處理最先準備好的那一個。 <br> - 了解 `default` case 如何實現非阻塞的 Channel 操作。 | **【並行的十字路口】`select`：在多個Channel之間做出選擇** |

### **第二週：深入 Channels 與同步機制 (Week 2: Advanced Channels & Synchronization)**

| 天數 | 主題 | 學習重點 | 心得 |
| :--- | :--- | :--- | :--- |
| **Day 8** | **單向 Channel (Directional Channels)** | - 學習 `chan<- T` (只能發送) 和 `<-chan T` (只能接收) 的宣告與使用。 <br> - 探討單向 Channel 如何提升函式簽章的可讀性與型別安全性。 | **【Channel的安全守則】用單向Channel，打造更穩健的API** |
| **Day 9** | **計時器 (`time.Ticker`) 與延遲 (`time.Timer`)** | - 學習如何使用 `time.NewTicker` 建立定時觸發的 Channel。 <br> - 學習如何使用 `time.NewTimer` 或 `time.After` 建立一次性延遲的 Channel。 | **【時間的魔法師】在Goroutine中優雅地處理定時與延遲** |
| **Day 10**| **Goroutine 洩漏 (Leak) 的成因與預防 (一)** | - 了解什麼是 Goroutine 洩漏：Goroutine 永遠阻塞且無法被回收。 <br> - 分析因為 Channel 接收方或發送方永久等待而造成的洩漏。 <br> - 實作一個會發生洩漏的範例並修復它。 | **【記憶體的小偷】你的Goroutine正在悄悄洩漏嗎？（上）** |
| **Day 11**| **`context` 套件：控制 Goroutine 的生命週期** | - 學習 `context.Context` 的基本概念與四大功能：`WithCancel`, `WithDeadline`, `WithTimeout`, `WithValue`。 <br> - 實作如何透過 `context` 來優雅地取消一個或多個 Goroutine。 | **【Goroutine的生命控制器】`context`：優雅地發出取消訊號** |
| **Day 12**| **Goroutine 洩漏的成因與預防 (二)** | - 結合 `context` 與 `select`，展示如何避免因為等待 Channel 而造成的洩漏。 <br> - 探討在複雜場景下，如何確保每個 Goroutine 都有明確的退出機制。 | **【記憶體的小偷】你的Goroutine正在悄悄洩漏嗎？（下）** |
| **Day 13**| **互斥鎖 (`sync.Mutex`) 與競爭條件 (Race Condition)** | - 了解什麼是競爭條件 (Race Condition)。 <br> - 學習使用 `sync.Mutex` 來保護共享資源，避免多個 Goroutine 同時存取。 <br> - 介紹 `go run -race` 指令來檢測競爭條件。 | **【共享資源的守護者】`sync.Mutex`：避免並行世界中的數據混亂** |
| **Day 14**| **讀寫鎖 (`sync.RWMutex`)** | - 學習 `sync.RWMutex` 如何在「讀多寫少」的場景下提升效能。 <br> - 比較 `Mutex` 與 `RWMutex` 的效能差異與使用情境。 | **【讀寫效能優化】`sync.RWMutex`：讓你的讀取操作飛起來** |

### **第三週：常見併發模式 (Week 3: Common Concurrency Patterns)**

| 天數 | 主題 | 學習重點 | 心得 |
| :--- | :--- | :--- | :--- |
| **Day 15**| **Worker Pool 模式 (一)** | - 了解 Worker Pool 模式的用途：限制併發數量，重複利用 Goroutine 資源。 <br> - 實作一個基本的 Worker Pool，包含任務 Channel 與結果 Channel。 | **【併發任務管理】Worker Pool模式：打造你的Goroutine大軍 (上)** |
| **Day 16**| **Worker Pool 模式 (二)** | - 優化 Worker Pool，加入 `context` 以便能優雅地關閉所有 Worker。 <br> - 探討如何動態調整 Worker 的數量。 | **【併發任務管理】Worker Pool模式：打造你的Goroutine大軍 (下)** |
| **Day 17**| **Fan-in, Fan-out 模式** | - **Fan-out**：一個生產者將任務分發給多個消費者 Goroutine。 <br> - **Fan-in**：多個生產者 Goroutine 將結果匯總到一個 Channel。 <br> - 實作一個結合 Fan-out 與 Fan-in 的 Pipeline。 | **【數據處理的流水線】Fan-in, Fan-out：分工與合作的藝術** |
| **Day 18**| **Pipeline 模式** | - 學習如何將多個處理階段串連起來，每個階段都是一個 Goroutine，透過 Channel 連接。 <br> - 實作一個多階段的數據處理 Pipeline，例如：生成數字 -> 計算平方 -> 列印結果。 | **【打造數據處理工廠】Pipeline模式：讓數據在Goroutine間流動** |
| **Day 19**| **Rate Limiting (速率限制)** | - 學習如何控制事件發生的速率，常用於 API 請求或資源密集型任務。 <br> - 運用 `time.Ticker` 結合 Channel 實作一個簡單的速率限制器。 | **【溫柔地請求】Goroutine速率限制：別把你的API打爆了** |
| **Day 20**| **錯誤處理與 `sync.ErrGroup`** | - 探討在多個 Goroutine 中如何有效地處理和傳遞錯誤。 <br> - 學習使用 `golang.org/x/sync/errgroup` 套件，簡化多 Goroutine 的錯誤處理與同步。 | **【一個都不能錯】在並行世界中，如何優雅地處理錯誤？** |
| **Day 21**| **`sync.Once` 的使用** | - 學習如何使用 `sync.Once` 來確保某個初始化操作只會被執行一次，即使在多個 Goroutine 呼叫下。 <br> - 探討其在單例模式 (Singleton) 中的應用。 | **【只做一次的承諾】`sync.Once`：確保你的初始化萬無一失** |

### **第四週：進階主題與實戰 (Week 4: Advanced Topics & Real-World Practice)**

| 天數 | 主題 | 學習重點 | 心得 |
| :--- | :--- | :--- | :--- |
| **Day 22**| **原子操作 (`sync/atomic`)** | - 了解原子操作的意義，以及為什麼在某些情況下它比 Mutex 更高效。 <br> - 學習 `atomic` 套件中常用的函式，如 `AddInt64`, `LoadInt64`, `StoreInt64`, `CompareAndSwap`。 | **【極致效能的鎖】原子操作`sync/atomic`：比Mutex更快的選擇？** |
| **Day 23**| **Goroutine 調度器 (Scheduler) 內部機制 (G-P-M 模型)** | - 初步理解 G (Goroutine), P (Processor), M (Machine/Thread) 之間的關係。 <br> - 了解 Go 調度器是如何實現高效的 Goroutine 調度與搶佔的。 | **【深入Go的核心】初探Goroutine調度器：G-P-M模型的奧秘** |
| **Day 24**| **實戰專案 (一)：併發 Web Crawler** | - 設計一個簡單的爬蟲，能夠併發地抓取網頁連結。 <br> - 使用 Goroutine 處理每個頁面的抓取，並使用 Channel 來傳遞新的 URL。 | **【實戰演練】用Goroutine打造一個併發網頁爬蟲 (一)** |
| **Day 25**| **實戰專案 (二)：優化 Web Crawler** | - 加入 Worker Pool 模式來限制同時抓取的連線數。 <br> - 使用 `context` 來控制爬蟲的停止。 <br> - 處理重複抓取的問題。 | **【實戰演練】用Goroutine打造一個併發網頁爬蟲 (二)** |
| **Day 26**| **實戰專案 (三)：併發檔案處理** | - 設計一個程式，可以併發地讀取一個目錄下的所有檔案，並計算行數或關鍵字。 <br> - 使用 Fan-out 分配檔案給不同的 Goroutine，再用 Fan-in 匯總結果。 | **【實戰演練】併發檔案處理：讓你的磁碟飛速運轉** |
| **Day 27**| **併發測試與基準測試 (Benchmarking)** | - 學習如何撰寫測試 Goroutine 的單元測試。 <br> - 學習使用 Go 的測試工具來對併發程式碼進行基準測試，並分析效能。 | **【衡量你的併發效能】如何測試與分析Goroutine的效率？** |
| **Day 28**| **常見陷阱與最佳實踐回顧** | - 總結常見的 Goroutine 陷阱，如：迴圈變數捕獲問題、死鎖等。 <br> - 回顧 Goroutine 與 Channel 使用的最佳實踐。 | **【避坑指南】Golang併發程式設計的十大常見錯誤與最佳實踐** |
| **Day 29**| **展望未來：`go.work` 與結構化併發** | - 探討 Golang 在併發領域的未來發展趨勢。 <br> - 簡介結構化併發 (Structured Concurrency) 的概念及其優點。 | **【Goroutine的未來】從`errgroup`到結構化併發的演進** |
| **Day 30**| **系列總結與心得** | - 回顧 30 天的學習歷程，整理核心知識地圖。| **30天深入Goroutine之旅：我的學習地圖與心得** |
