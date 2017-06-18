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
  return mergesort(arr)
}

func mergesort(arr []int) []int{
  if len(arr) > 1{
    mid := int(len(arr)/2)
    left,right := mergesort(arr[:mid]), mergesort(arr[mid:])
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

func QuickSort(arr []int) []int{
  return quicksort(arr,0,len(arr)-1)
}

func partition(arr []int, begin,end int) ([]int, int){
  pivot := arr[begin]
  top := begin

  for i := begin; i <= end; i++{
    if arr[i] < pivot{
      arr[top] = arr[i]
      arr[i] = arr[top+1]
    }
  }

  arr[top] = pivot
  return arr,top
}

func quicksort(arr []int, begin,end int) []int{
  if begin < end{
      arr,mid := partition(arr,begin,end)

      arr = quicksort(arr,begin,mid)
      arr = quicksort(arr,mid+1,end)
  }

  return arr
}

//##########################

func HeapSort(arr []int) []int{
  arr = buildHeap(arr)
  n := len(arr) - 1

  for i := n; i >= 0; i--{
      arr = swap(arr,0,i)
      n--
      arr = maxHeap(arr,0,n)
  }

  return arr
}

func swap(arr []int, i,j int) []int{
  aux := arr[i]
  arr[i], arr[j] = arr[j],aux
  return arr
}

func maxHeap(arr []int, i, length int) []int{
  left := 2 * i
  right := 2 * i + 1
  bigger := i

  if left <= length && arr[left] > arr[i]{
    bigger = left
  }

  if right <= length && arr[right] > arr[bigger]{
    bigger = right
  }

  if bigger != i{
    arr = swap(arr,i,bigger)
    arr = maxHeap(arr,bigger,length)
  }

  return arr
}

func buildHeap(arr []int) []int{
  n := len(arr) - 1

  for i:= int(n/2); i >= 0; i--{
    maxHeap(arr,i,n)
  }

  return arr
}
