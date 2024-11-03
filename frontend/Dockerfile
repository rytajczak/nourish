FROM node:23.1.0

WORKDIR /app/

RUN npm install -g pnpm

COPY package.json pnpm-lock.yaml /app/

RUN pnpm install --frozen-lockfile

COPY . /app/

RUN pnpm build

CMD ["node", ".output/server/index.mjs"] 
