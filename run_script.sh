sudo apt install make
make run-images
curl -X POST http://35.199.37.126/api/account/register -d '{"username": "john", "password": "john123456"}'
