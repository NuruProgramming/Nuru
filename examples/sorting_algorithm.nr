* 
############ Sorting Algorithm ##############
             
             By @VictorKariuki

        https://github.com/VictorKariuki

#############################################
*/

acha slice = fn(arr,start, end) {
    acha result = []
    wakati (start < end) {
        result = result + [arr[start]]
        start = start + 1
    }
    rudisha result
}

acha merge = fn(left, right) {
    acha result = []
    acha lLen = idadi(left)
    acha rLen = idadi(right)
    acha  l = 0
    acha  r = 0
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


acha mergeSort = fn(arr){
    acha len = idadi(arr)
    andika("arr is ", arr," of length ", len)
    kama (len < 2) {
        rudisha arr
    }
    andika("len is greater than or == to 2", len > 1)

    acha mid = (len / 2)
    andika("arr has a mid point of ", mid)
   
    acha left = slice(arr, 0, mid)
    acha right = slice(arr, mid, len)
    andika("left slice is ", left)
    andika("right slice is ", right)
    acha sortedLeft = mergeSort(left)
    acha sortedRight = mergeSort(right)
    andika("sortedLeft is ", sortedLeft)
    andika("sortedRight is ", sortedRight)
    rudisha merge(sortedLeft, sortedRight)
}

acha arr = [6, 5, 3, 1, 8, 7, 2, 4]
acha sortedArray = mergeSort(arr)
andika(sortedArray)