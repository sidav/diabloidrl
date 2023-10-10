package roomgrowinggenerator

func makeRandomTransofrmationForVault(v []string) []string {
	if rnd.Rand(2) == 0 {
		v = transposeVault(v)
	}
	if rnd.Rand(2) == 0 {
		v = mirrorVaultX(v)
	}
	if rnd.Rand(2) == 0 {
		v = mirrorVaultY(v)
	}
	return v
}

func revertString(s string) string {
	newS := make([]byte, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		newS[len(s)-i-1] = s[i]
	}
	return string(newS)
}

func mirrorVaultX(v []string) []string {
	mirrorX := []string{}
	for x := range v {
		mirrorX = append(mirrorX, revertString(v[x]))
	}
	return mirrorX
}

func mirrorVaultY(v []string) []string {
	mirrorY := []string{}
	for x := range v {
		mirrorY = append(mirrorY, v[len(v)-1-x])
	}
	return mirrorY
}

func transposeVault(v []string) []string {
	transp := make([][]byte, len(v[0]))
	for i := range transp {
		transp[i] = make([]byte, len(v))
	}
	for x := range v {
		for y := range v[x] {
			transp[y][x] = v[x][y]
		}
	}
	str := make([]string, 0)
	for i := range transp {
		str = append(str, string(transp[i]))
	}
	return str
}

func doesVaultContainDoor(v []string) bool {
	for x := range v {
		for y := range v[x] {
			if v[x][y] == '+' {
				return true
			}
		}
	}
	return false
}
