package main

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
    "strconv"
    "os"
    "encoding/json"
)

func main() {
    var res *http.Response

    for i := 41575; i <= 99999; i++ {
    url := "https://vmp.telkomsel.com/api/packages/000" + strconv.Itoa(i)
    payload := strings.NewReader("{\"toBeSubscribedTo\":false}")
    req, err := http.NewRequest("PUT", url, payload)
    if err != nil {
        //Specific error handling would depend on scenario
        fmt.Printf("%v\n", err)
        return
    }
    req.Header.Add("authorization", "Bearer .......")
    req.Header.Add("channelid", "VMP")
    req.Header.Add("Content-Type", "application/json")
    
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        //Specific error handling would depend on scenario
        fmt.Printf("%v\n", err)
        return
    }
    // defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        //Specific error handling would depend on scenario
        fmt.Printf("%v\n", err)
        return
    }
    var result map[string]interface{}
    json.Unmarshal([]byte(string(body)), &result)
    str := fmt.Sprint(result["message"])
    if(str != "SYS-UXP-0021"){
        file, err := os.OpenFile("text1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println(err)
        }
        defer file.Close()
        if _, err := file.WriteString("id: 000"+strconv.Itoa(i)+" "+str+"\n"); err != nil {
            fmt.Println(err)
        }
    }
    if(str == "BIZ-UXP-0002"){
        file0, err0 := os.OpenFile("tidak cukup1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err0 != nil {
            fmt.Println(err0)
        }
        defer file0.Close()
        if _, err0 := file0.WriteString("id: 000"+strconv.Itoa(i)+" "+str+"\n"); err0 != nil {
            fmt.Println(err0)
        }
    }
    if(result["reason"] == "Success"){
        file1, err1 := os.OpenFile("berhasil1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err1 != nil {
            fmt.Println(err1)
        }
        defer file1.Close()
        if _, err1 := file1.WriteString("id: 000"+strconv.Itoa(i)+" "+str+"\n"); err1 != nil {
            fmt.Println(err1)
        }
    }
    // file, err := os.Create("text.txt")
    if err != nil {
        return
    }
    // defer file.Close()
    // file.WriteString(str)
    fmt.Println("id: 000"+strconv.Itoa(i)+" " +string(body))
    res.Body.Close()
    }

    defer res.Body.Close()

}