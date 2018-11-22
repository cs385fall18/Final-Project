sudo apt install make
make run-images
curl -X POST http://<outer address for container>/api/account/register -d '{"username": "john", "password": "john123456"}'
