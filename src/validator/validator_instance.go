package validator

import (
	"strings"
	
)

func traverse(t *trav, tags string) (er fielderr) {

	if len(tags) == 0 {
		return
	}

	tagArray := strings.Split(tags, ",")
	
	er.fieldname =  t.fieldname;
		
	
	for _, tag := range tagArray {
		if !bakedValidator[tag](t.val) {
			er.err += errMsg[tag];
			return;
		}
	}
	return
}
