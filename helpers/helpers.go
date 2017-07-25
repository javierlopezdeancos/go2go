package helpers

type Headers struct {
	Title 		string
	TitleSide	string
	Break		string
	Split 		string
	Given 		string
	When  		string
	Then  		string
}

func GetHeaders () Headers {
	return Headers{
		Title: 		"***************************************************************",
		TitleSide: 	"*                 							                   *",
		Break: 		"\n",
		Split: 		"---------------------------------------------------------------",
		Given: 		"\n-------------------------GIVEN---------------------------------",
		When:  		"\n------------------------WHEN-----------------------------------",
		Then:  		"\n------------------------THEN-----------------------------------",
	}
}