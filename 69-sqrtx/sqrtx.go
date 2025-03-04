func mySqrt(x int) int {
    var l, r, res = 0, x, 0

    for (l <= r) {
        m := l + ((r - l) / 2)

        if (math.Pow(float64(m), float64(2)) > float64(x)) {
            r = m - 1
        } else if (math.Pow(float64(m), float64(2)) < float64(x)) {
            l = m + 1
            res = m
        } else {
            return m
        }
    }
    return res


}