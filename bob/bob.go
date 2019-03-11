package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	remark = strings.TrimSpace(remark)

	silence := remark == ""
	isYelling := (remark == strings.ToUpper(remark)) && (strings.ContainsAny(remark, alpha))
	isQuestion := false
	//this is because if the string is empty then you cant access that index and theres an error
	if len(remark) > 0 {
		isQuestion = remark[len(remark)-1:] == "?"
	}

	if silence {
		return "Fine. Be that way!"
	}

	if isQuestion && isYelling {
		return "Calm down, I know what I'm doing!"

	} else if isQuestion {
		return "Sure."

	} else if isYelling {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
