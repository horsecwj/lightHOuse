#aws ecr get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 831273654673.dkr.ecr.ap-southeast-1.amazonaws.com
docker build --rm -t 831273654673.dkr.ecr.ap-southeast-1.amazonaws.com/light-house/rear-end:latest .
#docker push 831273654673.dkr.ecr.ap-southeast-1.amazonaws.com/light-house/rear-end:latest
docker save 831273654673.dkr.ecr.ap-southeast-1.amazonaws.com/light-house/rear-end:latest -o rear-end.zip
scp rear-end.zip root@47.57.137.142:/root/
ssh root@47.57.137.142 "docker load -i /root/rear-end.zip && docker rm -f lighthouse && bash /root/light-house/run_lighthouse.sh"
