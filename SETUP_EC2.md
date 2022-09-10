sudo yum update -y
sudo yum install -y docker
sudo yum install -y git
sudo service docker start
sudo usermod -a -G docker ec2-user
reboot 

sudo service docker start

docker info

git@github.com:chayuto/go-iot-api.git
git clone git@github.com:chayuto/go-iot-api.git

 cd /mnt/c
chmod 600 ricco-aws.pem
 ssh ec2-user@3.26.78.153 -i ricco-aws.pem -v
 cp /mnt/c/ricco-aws.pem ricco-aws.pem 