# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Installation des dépendances système nécessaires pour compiler
RUN apk add --no-cache git

# Copie des fichiers de définition de module
COPY go.mod go.sum ./
COPY proto ./proto 
# On modifie le go.mod pour pointer vers le dossier local proto
RUN go mod edit -replace github.com/jobmate-org/jobmate-proto=./proto
RUN go mod download

# Copie du code source
COPY . .

# Build du binaire (CGO_ENABLED=0 pour un binaire statique léger)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Runtime (Image ultra-légère)
FROM gcr.io/distroless/static-debian11
COPY --from=builder /app/main /
CMD ["/main"]