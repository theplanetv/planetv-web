FROM node:22-alpine3.19

WORKDIR /web-next

COPY web-next/package.json web-next/package-lock.json ./

RUN npm install

COPY web-next/src ./src
COPY web-next/.eslintrc.json web-next/next.config.mjs web-next/postcss.config.mjs web-next/tailwind.config.ts web-next/tsconfig.json ./

RUN npm run build

CMD ["npm", "run", "dev"]
