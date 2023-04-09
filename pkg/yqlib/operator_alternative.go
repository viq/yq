package yqlib

func alternativeOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	log.Debugf("-- alternative")
	prefs := crossFunctionPreferences{
		CalcWhenEmpty: true,
		Calculation:   alternativeFunc,
		LhsResultValue: func(lhs *CandidateNode) (*CandidateNode, error) {
			if lhs == nil {
				return nil, nil
			}
			truthy, err := isTruthyNode(lhs)
			if err != nil {
				return nil, err
			}
			if truthy {
				return lhs, nil
			}
			return nil, nil
		},
	}
	return crossFunctionWithPrefs(d, context, expressionNode, prefs)
}

func alternativeFunc(d *dataTreeNavigator, context Context, lhs *CandidateNode, rhs *CandidateNode) (*CandidateNode, error) {
	if lhs == nil {
		return rhs, nil
	}
	if rhs == nil {
		return lhs, nil
	}

	isTrue, err := isTruthyNode(lhs)
	if err != nil {
		return nil, err
	} else if isTrue {
		return lhs, nil
	}
	return rhs, nil
}
