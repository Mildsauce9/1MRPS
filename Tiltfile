allow_k8s_contexts('minikube')

k8s_yaml(['k8s/postgres.yaml', 'k8s/api.yaml'])

docker_build('1mrps-api', '.', dockerfile='Dockerfile',
             only=['go.mod', 'go.sum', 'cmd', 'internal'])

k8s_resource('postgres', port_forwards=5432, labels=['db'])

# Runs goose migrations against the port-forwarded Postgres, then the API can start.
local_resource(
    'migrate',
    cmd='go run github.com/pressly/goose/v3/cmd/goose@latest -dir internal/adapters/postgresql/migrations postgres "postgres://postgres:postgres@localhost:5432/game?sslmode=disable" up',
    resource_deps=['postgres'],
    labels=['db'],
)

k8s_resource('api', port_forwards=8080, resource_deps=['migrate'], labels=['backend'])
