# URL Shortener

Simple URL Shortener

## Setup

1. run `./dev install_docker`, to install  `docker-toolbox`
2. run `./dev build_machine`, to build docker machine `default`
3. load `eval "$(docker-machine env default)"`
4. run `sudo ./dev load_resolver`
5. run `cp .env.example .env`, and change values accordingly
6. run `./dev start`, to start the application
7. access the app on the `DOMAIN` you've set on `.env`

run `./dev` to know the other shortcut commands

## API

### Shorten url

```
curl -XPOST \
    http://$(DOMAIN)/api/shorten \
    -d '{
        "url":"https://example.com",
        "expiration":-1
    }'
```

- _url_ only valid urls are allowed
- _expiration_ optional seconds before shortened url expires
  - if value is `-1` it never expires
  - if no value is supplied a default 24 hour expiration is set

### Expand url

```
curl -XPOST \
    http://pucha.kc/api/expand \
    -d '{
        "short":"QdK"
    }'
```

- returns status `404` when long url is not found
