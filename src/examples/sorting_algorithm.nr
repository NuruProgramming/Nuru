/* 
############ Sorting Algorithm ##############
             
             By @VictorKariuki

        https://github.com/VictorKariuki

#############################################
*/

fanya slice = unda(arr,start, end) {
    fanya result = []
    wakati (start < end) {
        result = result + [arr[start]]
        start = start + 1
    }
    rudisha result
}

fanya merge = unda(left, right) {
    fanya result = []
    fanya lLen = idadi(left)
    fanya rLen = idadi(right)
    fanya  l = 0
    fanya  r = 0
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


fanya mergeSort = unda(arr){
    fanya len = idadi(arr)
    andika("arr is ", arr," of length ", len)
    kama (len < 2) {
        rudisha arr
    }
    andika("len is greater than or == to 2", len > 1)

    fanya mid = (len / 2)
    andika("arr has a mid point of ", mid)
   
    fanya left = slice(arr, 0, mid)
    fanya right = slice(arr, mid, len)
    andika("left slice is ", left)
    andika("right slice is ", right)
    fanya sortedLeft = mergeSort(left)
    fanya sortedRight = mergeSort(right)
    andika("sortedLeft is ", sortedLeft)
    andika("sortedRight is ", sortedRight)
    rudisha merge(sortedLeft, sortedRight)
}

fanya arr = [6, 5, 3, 1, 8, 7, 2, 4]
fanya sortedArray = mergeSort(arr)
andika(sortedArray)