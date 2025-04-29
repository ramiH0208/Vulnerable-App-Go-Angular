# Étape 1 : build de l'application Angular
FROM node:20 as build

WORKDIR /app

COPY . .

RUN npm install && npm run build --projects xss-lab --configuration production

# Étape 2 : image finale avec nginx
FROM nginx:alpine

# Copie des fichiers buildés dans nginx
COPY --from=build /app/dist/angular-xss-lab /usr/share/nginx/html

# Configuration de base Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80