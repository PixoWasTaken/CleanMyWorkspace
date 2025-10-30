package Clean

func CleanWorkSpace(workspace *[][]*string) *[][]*string {
	if workspace == nil { // Si y'a nil on donne nil, donc Chaise. On touche a rien
		return nil
	}

	for i := range *workspace { // C'est facinant mais on déréférence les pointeurs. C'est vraiment un terme classe en vrai :O
		for j := range (*workspace)[i] { 
			(*workspace)[i][j] = nil
		}
	}

	return workspace
}