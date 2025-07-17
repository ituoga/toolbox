package toolbox

func Djb33(seed uint32, k string) uint32 {
	var (
		l = uint32(len(k))
		d = 5381 + seed + l
	)
	if l > 0 {
		d = djb33_loop(d, k)
	}
	return d ^ (d >> 16)
}

func djb33_loop(d uint32, k string) uint32 {
	var (
		l = uint32(len(k))
		i = uint32(0)
	)
loop:
	if i >= l-1 {
		goto exit
	}
	d = (d * 33) ^ uint32(k[i])
	i++
	goto loop
exit:
	return d
}
