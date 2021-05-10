FROM azul/zulu-openjdk-alpine:11
COPY build/libs/* app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/app.jar"]