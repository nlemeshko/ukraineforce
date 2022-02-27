#Ukraine force DDOS

[![docker](https://github.com/nlemeshko/ukraineforce/actions/workflows/docker.yaml/badge.svg)](https://github.com/nlemeshko/ukraineforce/actions/workflows/docker.yaml)

### Variables:

- **headers.json**
- **body.json**

### Build:
```docker build -t ukraineforce .```


### Usage:
``` udda [URL] [METHOD] [THREADS] [BODY_JSON_FILE_PATH] [HEADERS_JSON_FILE_PATH] [flags] ```

Flags:
-h, --help   help for udda


Example:

```docker run --rm -d ukraineforce ./udda udda http://127.0.0.1:5894/api/v1/oauth/token POST 20 body.json headers.json```

### Alternative usage

- ```git clone https://github.com/nlemeshko/ukraineforce.git```
- ```Change headers and body```
- ```docker run --rm -d -v ./body.json:/home/udda/body.json -v ./headers.json:/home/udda/headers.json mdsn/ukraineforce ./udda udda [URL] POST 20 body.json headers.json```

### Fork usage:

- **Fork repository and clone**
- **Add to repository three secrets**
    - **DOCKERHUB_USERNAME** - Your Dockerhub username
    - **DOCKERHUB_TOKEN** - Your Dockerhub password
    - **DOCKERHUB_PATH** - Your path to Dockerhub
- **Change headers, body and push**
- **run command:**
    - ``` docker run --rm -d [DOCKERHUB_PATH] ./udda udda [URL] POST 20 body.json headers.json```