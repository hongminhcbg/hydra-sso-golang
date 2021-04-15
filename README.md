# hydra-sso-golang
1. Init database
    docker run \
    --network hydraguide \
    --name ory-hydra-example--postgres \
    -e POSTGRES_USER=hydra \
    -e POSTGRES_PASSWORD=secret \
    -e POSTGRES_DB=hydra \
    -v hydra-postgres:/var/lib/postgresql/data \
    -d postgres:13

2. Init system secret
   export SECRETS_SYSTEM=$(export LC_CTYPE=C; cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
3. Init DNS
   export DSN=postgres://hydra:secret@ory-hydra-example--postgres:5432/hydra?sslmode=disable
4. Show env - optional
   docker run -it --rm --entrypoint hydra oryd/hydra:v1.10.1 help serve
5. Migrate database
   docker run -it --rm \
   --network hydraguide \
   oryd/hydra:v1.10.1 \
   migrate sql --yes $DSN
   
6. Run hydra

    docker run -d \
    --name ory-hydra-example--hydra \
    --network hydraguide \
    -p 4444:4444 \
    -p 4445:4445 \
    -e SECRETS_SYSTEM=$SECRETS_SYSTEM \
    -e DSN=$DSN \
    -e URLS_SELF_ISSUER=https://localhost:3000/ \
    -e URLS_CONSENT=http://localhost:3000/consent \
    -e URLS_LOGIN=http://localhost:3000/login \
    oryd/hydra:v1.10.1 serve -c /etc/config/hydra/hydra.yml all --dangerous-force-http
