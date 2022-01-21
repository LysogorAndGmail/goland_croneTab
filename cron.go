package main

import (
   "fmt"
   "os"
   "os/signal"
   "syscall"
   "time"

   cron "github.com/robfig/cron/v3"
)

func main() {
   // set scheduler berdasarkan zona waktu sesuai kebutuhan
   berlinTime, _ := time.LoadLocation("Europe/Berlin") 
   scheduler := cron.New(cron.WithLocation(berlinTime))

   // stop scheduler tepat sebelum fungsi berakhir
   defer scheduler.Stop()

   // set task yang akan dijalankan scheduler
   // gunakan crontab string untuk mengatur jadwal
   //scheduler.AddFunc("0 0 1 1 *", func() { SendAutomail("New Year") })
   //scheduler.AddFunc("0 07 10 * *", SendMonthlyBillingAutomail)
   //scheduler.AddFunc("0 09 * * 1-5", NotifyDailyAgenda)
   scheduler.AddFunc("*/1 * * * *", NotifyNewOrder)

   // start scheduler
   go scheduler.Start()

   // trap SIGINT untuk trigger shutdown.
   sig := make(chan os.Signal, 1)
   signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
   <-sig
}
/*
func SendAutomail(automailType string) {
    // ... instruksi untuk mengirim automail berdasarkan automailType
    fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " SendAutomail " + automailType + " telah dijalankan.\n")
}

func SendMonthlyBillingAutomail(){
    // ... instruksi untuk mengirim automail berisi tagihan
    fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " SendMonthlyBillingAutomail telah dijalankan.\n")
}

func NotifyDailyAgenda(){
    // ... instruksi untuk mengirim notifikasi agenda harian
    fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " NotifyDailyAgenda telah dijalankan.\n")
}
*/
func NotifyNewOrder(){
    // ... instruksi untuk mengecek pesanan baru, lalu mengirimkan notifikasi
    //fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " new Text.\n")
    text := time.Now().Format("02.01.2006 15:04:05") + " Alla Moja Lubov \n"
    file, err := os.OpenFile("cronLog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
     
    if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1) 
    }
    defer file.Close() 
    file.WriteString(text)
     
    fmt.Println("Done.")
}