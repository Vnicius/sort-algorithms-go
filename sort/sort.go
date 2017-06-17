package sort

func SelectionSort(arr []int) []int{
    var min, aux int

    for i,_ := range arr{
      min = i

      for j := i; j < len(arr); j++{
        if arr[j] < arr[min]{
          min = j
        }
      }

      aux = arr[min]
      arr[min] = arr[i]
      arr[i] = aux
    }

    return arr
}

//#######################

func InsertionSort(arr []int) []int{
  var value,key int

  for i := range arr{
    value,key = arr[i],i

    for key > 0 && value < arr[key-1]{
      arr[key] = arr[key-1]
      key -= 1
    }

    arr[key] = value
  }

  return arr
}

//########################

func MergeSort(arr []int) []int{
  if len(arr) > 1{
    mid := int(len(arr)/2)

    left,right := arr[:mid], arr[mid:]

    left = MergeSort(left)
    right = MergeSort(right)

    var i,j,k int

    for i < len(left) && j < len(right){
      if left[i] < right[j]{
        arr[k] = left[i]
        i++
      }else{
        arr[k] = right[j]
        j++
      }
      k++
    }

    for i < len(left){
      arr[k] = left[i]
      i++
      k++
    }

    for j < len(right){
      arr[k] = right[j]
      j++
      k++
    }
  }
  return arr
}

//#################
