package config

//var bind map[string]string = map[string]string{
//	"app.httpAddress":             "APP_HTTP_ADDRESS",
//	"app.grpcAddress":             "APP_GRPC_ADDRESS",
//	"app.env":                     "APP_ENV",
//	"app.secretKey":               "APP_SECRET_KEY",
//	"oauth2.google.client_id":     "OAUTH2_GOOGLE_CLIENT_ID",
//	"oauth2.google.client_secret": "OAUTH2_GOOGLE_CLIENT_SECRET",
//	"client.url":                  "CLIENT_URL",
//	"database.mysql":              "DATABASE_MYSQL",
//}
//

var bind map[string]string = map[string]string{
	"APP_HTTP_ADDRESS":            "app.httpAddress",
	"APP_GRPC_ADDRESS":            "app.grpcAddress",
	"APP_ENV":                     "app.env",
	"APP_SECRET_KEY":              "app.secretKey",
	"OAUTH2_GOOGLE_CLIENT_ID":     "oauth2.google.client_id",
	"OAUTH2_GOOGLE_CLIENT_SECRET": "oauth2.google.client_secret",
	"CLIENT_URL":                  "client.url",
	"DATABASE_MYSQL":              "database.mysql",
	"DATABASE_POSTGRES":           "database.postgres",
}
