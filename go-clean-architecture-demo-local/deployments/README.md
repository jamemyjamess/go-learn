# `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).

docker-compose -f deployments/docker-compose.yml up --build -d

docker-compose -f deployments/docker-compose.yml --env-file <.env.uat> up --build -d

uat
docker-compose -f deployments/docker-compose.yml --env-file .env.uat up --build -d