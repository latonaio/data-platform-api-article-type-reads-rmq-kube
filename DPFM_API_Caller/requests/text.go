package requests

type Text struct {
	ArticleType     	string  `json:"ArticleType"`
	Language          	string  `json:"Language"`
	ArticleTypeName		string  `json:"ArticleTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
