package blog

type UserCreateArg struct {
}

type UserCreateRet struct {
}

type UserListArg struct {
	Name string `json:"name" form:"name"`
}

type UserListRet struct {
	List []UserListRow
}

type UserListRow struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
