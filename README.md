# upload_file_service
###About
- This is upload file service 
- This service is also used as resource service
- Auth middleware is used for check access token(jwt token) if it is valid
### Start project
- Change config in config/config.yaml
  + Change dsn to connect to mysql
  + Change storage to choose where to write image file when user upload it
- go mod tidy
- go run ./main.go
### Test Feature
- Use Authorization Server to get access token
- Use postman to call API
- Paste token to Authorization Header(type: Bearer Token)
- Choose form-data option -> key: file, value: image file
- Call http://localhost:5000/upload to upload image
- Go to DB and storage path(in config) to check if file is successfully updated