sudo apt install make
make run-images
curl -X POST http://localhost/api/account/register -d '{"username": "john", "password": "john123456"}'
