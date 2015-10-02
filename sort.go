package quicksort

func Sort(values []int) {
    sort(values, 0, len(values)-1)
}

func sort(values []int, l int, r int) {
    if l >= r {
        return
    }

    pivot := values[l]
    i := l + 1

    for j := l; j <= r; j++ {
        if pivot > values[j] {
            values[i], values[j] = values[j], values[i]
            i++
        }
    }

    values[l], values[i-1] = values[i-1], pivot

    sort(values, l, i-2)
    sort(values, i, r)
}
