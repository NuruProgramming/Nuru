/* 
############ Sorting Algorithm ##############
             
             By @VictorKariuki

        https://github.com/VictorKariuki

#############################################
*/

slice = unda(arr,start, end) {
    result = []
    wakati (start < end) {
        result = result + [arr[start]]
        start = start + 1
    }
    rudisha result
}

merge = unda(left, right) {
    result = []
    lLen = left.idadi()
    rLen = right.idadi()
     l = 0
     r = 0
    wakati (l < lLen && r < rLen) {
        kama (left[l] < right[r]) {
            result = result + [left[l]]
            l = l + 1
        } sivyo {
            result = result + [right[r]]
            r = r + 1
        }
    }
    andika(result)
}


mergeSort = unda(arr){
    len = arr.idadi()
    andika("arr is ", arr," of length ", len)
    kama (len < 2) {
        rudisha arr
    }
    andika("len is greater than or == to 2", len > 1)

    mid = (len / 2)
    andika("arr has a mid point of ", mid)
   
    left = slice(arr, 0, mid)
    right = slice(arr, mid, len)
    andika("left slice is ", left)
    andika("right slice is ", right)
    sortedLeft = mergeSort(left)
    sortedRight = mergeSort(right)
    andika("sortedLeft is ", sortedLeft)
    andika("sortedRight is ", sortedRight)
    rudisha merge(sortedLeft, sortedRight)
}

arr = [6, 5, 3, 1, 8, 7, 2, 4]
sortedArray = mergeSort(arr)
andika(sortedArray)
