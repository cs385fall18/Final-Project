sudo apt install make
make run-images
curl -X POST localhost/api/account/register -d '{"username": "john", "password": "john123456"}'
# loadbalancer ip address or the vm's external ip address
