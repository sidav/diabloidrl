package main

func readKey() string {
	return cw.ReadKey()
}

func readKeyAsync(maxMsSinceKeypress int) string {
	return cw.ReadKeyAsync(maxMsSinceKeypress)
}
