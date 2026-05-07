# API

OpenAPI: `api/openapi.yaml`

Base URL:

```text
http://localhost:8080
```

Create a job:

```sh
curl -F "audio=@song.wav" http://localhost:8080/api/v1/jobs
```

Read job status:

```sh
curl http://localhost:8080/api/v1/jobs/JOB_ID
```
