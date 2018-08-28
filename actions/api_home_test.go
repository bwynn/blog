package actions

func (as *ActionSuite) Test_APIHomeHandler() {
	res := as.HTML("/api").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
