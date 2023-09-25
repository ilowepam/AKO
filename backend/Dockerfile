# Verwende das offizielle Go-Image als Basis
FROM golang:latest

# Setze den Arbeitsverzeichnis-Pfad im Container
WORKDIR /go/src/app

# Kopiere den Go-Code in das Arbeitsverzeichnis im Container
COPY ./ako/backend /go/src/app

# Kompiliere den Go-Code
RUN go build -o main .

# Exponiere den Port, auf dem die Anwendung läuft
EXPOSE 8080

# Starte die Anwendung
CMD ["./main"]
