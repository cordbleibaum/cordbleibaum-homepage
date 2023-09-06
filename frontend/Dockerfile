FROM node:20 as dist

RUN npm install -g http-server

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install

COPY . .
ENV NODE_ENV=production
RUN npm run build


FROM nginx:1.25

COPY --from=dist /app/dist /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 8080

ENTRYPOINT ["nginx", "-g", "daemon off;"]
