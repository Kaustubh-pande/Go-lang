func compareTriplets(a []int32, b []int32) []int32 {
    var (
        alise int32
        bob int32
        resultarray []int32
        
    )
    n := len(a)
    for i:=0;i<n;i++{
        j :=i
        if a[i] > b[j]{
            alise++
        }else if a[i] < b[j]{
                 bob++
              }
    }
    resultarray = append(resultarray,alise,bob)
    return  resultarray

}