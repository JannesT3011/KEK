FROM node:10

WORKDIR /usr/KEK

COPY package*.json ./

RUN npm install

COPY . .

CMD [ "node", "main.js" ]
