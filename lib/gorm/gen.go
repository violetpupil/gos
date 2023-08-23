// https://gorm.io/gen/
// https://gorm.io/gen/gen_tool.html
// go install gorm.io/gen/tools/gentool@latest
// -dsn data source name
// -modelPkgName generated model code's package name
// -outPath specify a directory for output (default "./dao/query")
// -tables specify tables want to generated from, default all tables. 逗号分隔
// -onlyModel only generate models (without query file)
//
// gentool -onlyModel ^
// -dsn "user:pwd@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local" ^
// -tables "t_account,t_country,t_timezone" ^
// -modelPkgName model ^
// -outPath ./types/model
package gorm
