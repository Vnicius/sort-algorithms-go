package main

import (
  "fmt"
  "./sort"
  "io/ioutil"
  "strings"
  "strconv"
  )

func main(){
  data, err := ioutil.ReadFile("in.txt")

  if err != nil{
    panic(err)
  }

  values := strings.Fields(string(data))
  var arr []int

  for _,v := range values{
    val, err := strconv.Atoi(v)
    if err == nil{
      arr = append(arr,val)
    }
  }

  fmt.Println(sort.SelectionSort(arr))
  fmt.Println(sort.InsertionSort(arr))
  fmt.Println(sort.MergeSort(arr))
}
