package two_pointer


func BackspaceStrCompare(S string, _ string) bool {
	fast := 1
	slow := 0


	for fast < len(S) {
		if S[fast] == bs {
			S[slow] = ""
		} else {
			if S[fast] != "#" && S[slow] == "" {
				S[slow] = S[fast]
			}
			slow += 1
		}
		fast += 1
	}
	return false
}
