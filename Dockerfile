FROM node:22-alpine3.19

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install 

COPY src ./src
COPY public ./public
COPY .eslintrc.json next.config.mjs postcss.config.mjs tailwind.config.ts tsconfig.json ./

CMD ["npm", "run", "dev"]
