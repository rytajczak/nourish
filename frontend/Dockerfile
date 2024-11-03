FROM node:23-alpine as build

WORKDIR /app/

RUN npm install -g pnpm

COPY package.json pnpm-lock.yaml /app/

RUN pnpm install --frozen-lockfile

COPY . /app/

RUN pnpm build

FROM node:23-alpine as prod

WORKDIR /app

COPY --from=build /app/.output /app/.output

CMD ["node", ".output/server/index.mjs"]
