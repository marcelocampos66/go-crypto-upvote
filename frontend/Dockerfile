FROM node:14.18.1 AS build
WORKDIR /usr/app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx AS prod
COPY ./nginx.conf ./etc/nginx/nginx.conf
COPY --from=build /usr/app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]