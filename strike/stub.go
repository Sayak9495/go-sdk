package strike

import (
	"fmt"
	"encoding/json"
)

func Create(handler string, next string) *Body_structure{
	response := Body_structure{
		ActionHandler:handler,
		NextActionHandler:next,
	}
	
	return &response
}

func (create *Body_structure) Question(context string) *Transaction_structure {
    
	t := Transaction_structure{
		Question: Question_structure{
			QContext: context,
		},
	}
	
	return &t
}

func (t *Transaction_structure) Answer(create *Body_structure, multiple_select bool) *Transaction_structure {

	q := t.Question
	t = &Transaction_structure{
		Question: q,
		Answer1: Answer_structure{
			MultipleSelect: multiple_select,
		},
	}
	
	Update_Question_Array(create,t)
	
	return t
}

//Helper function

func Update_Question_Array(create *Body_structure,t *Transaction_structure){
	newArray := []Transaction_structure{}

	if len(create.QuestionArray)-1 < 0 {
		newArray = create.QuestionArray[:0]	
	} else{
		newArray = create.QuestionArray[:len(create.QuestionArray)-1]
	}
	
	newArray = append(newArray,*t)
}

func Update_QCard_Array(qcard [][]Card_Row_Object,t []Card_Row_Object) [][]Card_Row_Object{
	newArray := [][]Card_Row_Object{}

	if len(qcard)-1 < 0 {
		newArray = qcard[:0]	
	} else{
		newArray = qcard[:len(qcard)-1]
	}
	
	newArray = append(newArray,t)
	return newArray
}

func (create Response_wrapper_structure) ToJson() ([]byte) {
	b, err := json.Marshal(create)
    if err != nil {
        fmt.Println(err)
    }
	return b
}

func (create *Body_structure) Wrapper () Response_wrapper_structure {
	wrapper := Response_wrapper_structure{
		Status: 201,
		Body: create,
	}
	return wrapper
}
