package routes

var PrivateRoutes = []string{
	"/api/v1/users/profile",
	"/api/v1/classes",
	"/api/v1/tests",
	"/api/v1/tests/content",
	"/api/v1/tests/own",
	"/api/{version}/answersheets/start",
	"/api/v1/answersheets/answer",
	"/api/v1/answersheets/test/{testId}/content",
	"/api/v1/answersheets/test/{test_id}/check-doing",
}
