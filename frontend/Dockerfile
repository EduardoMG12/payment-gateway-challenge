FROM node:20-slim AS development

WORKDIR /app

RUN corepack enable

COPY package.json pnpm-lock.yaml ./
RUN pnpm install

COPY . .

EXPOSE 5173

CMD ["pnpm", "run", "dev", "--host", "0.0.0.0"]

FROM development AS builder
RUN pnpm run build

FROM nginx:stable-alpine AS production
COPY --from=builder /app/dist /usr/share/nginx/html