#Command  for  building  the docker  container
docker  build   -t app-002-ai-service  .
#Command for  running  the  docker  container
docker run -p 3081:3081 app-002-ai-service