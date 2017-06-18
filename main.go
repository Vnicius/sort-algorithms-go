package main

import (
  "./sort"
  "io/ioutil"
  "strings"
  "strconv"
  "os"
  )

func main(){
  data, err := ioutil.ReadFile(os.Args[1])

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

  sort.SelectionSort(arr)
  sort.InsertionSort(arr)
  sort.MergeSort(arr)
  sort.QuickSort(arr)
  sort.HeapSort(arr)
}
