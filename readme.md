

# Instructions (kinda)
- `docker build --tag wmrss_init .`
- `docker run -d -p 8080:8080 wmrss_init`

## Alternative Method
- update the Makefile and `make run`


## Route

`/rss/v1/weather-maps?pretty=true`

__note__: passing query param `pretty=true` will pretty print, anything else
will minifiy
