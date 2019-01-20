FROM node:11 As builder
COPY . /app
WORKDIR /app
RUN npm install
RUN npm rebuild node-sass
RUN npm install
RUN npm run build

FROM httpd:2.4
COPY --from=builder ./app/public/ /usr/local/apache2/htdocs/
