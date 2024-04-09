<div align="center">
  <a href="readme_zh.md">中文版 README</a>
  </br>
  <p>My communication channel :<a href="https://t.me/cliptalk">https://t.me/cliptalk</a>
</div>
### Project Introduction:
This is a 1:1 mirror version of the official Poe website, completely replicated.

It is not afraid of official resource updates and automatically synchronizes with official resources.

Account cookies can be set to achieve shared login-free access.

Due to some reasons, the WebSocket communication is not fully implemented, so free-flow conversation without the need for translation is temporarily not supported.

It is a semi-finished product, and those who are interested can improve it themselves.

Demo site: [https://poe.atvai.com](https://poe.atvai.com)

The demo site does not provide any technical support and is only for learning and exchange purposes. It is maintained casually.

Telegram exchange channel: [https://t.me/cliptalk](https://t.me/cliptalk)

### Project Deployment
#### 1.1 Preparation
1. Install 1panel or BT panel. If you are a server novice, it is recommended to install one of these. Here, I am using 1panel. (If you are an expert, you can directly use Nginx for reverse proxy.)
2. Install Docker and Docker Compose. (It is recommended to use Docker Compose. For specific installation instructions, please Google or consult GPT.)
3. Clone the project to your local machine.
```shell
git clone https://github.com/petiky/fkpoe.git
```
#### 1.2 Deployment
```shell
cd fkpoe
vim docker-compose.yml
# Modify LOCAL_BASE_URL to your domain name
docker-compose up -d
```
or
```shell
mv .env.example .env
vim .env
# Modify LOCAL_BASE_URL to your domain name

./fkpoe
```
#### 1.3 Configuration
`I am using 1panel's reverse proxy here. If you are using BT panel, you can also use its reverse proxy. I won't go into details here.`
1. Create a new website, select reverse proxy, fill in your domain name, set the frontend request path to `/`, name it `root` (or any other name you like), and set the proxy address to your server address. Then save it.
2. Configure an SSL certificate for your domain name. (I won't elaborate on this. If you don't know how to do it, please Google it.)
   Now, reopen your domain name, and you should see an official website that is exactly the same as Poe. Registration, login, and conversation are all identical. (Due to IP issues, login may be recognized as a bot.)

### Precautions
1. Please do not modify any files in the project, as it may cause the project to malfunction.
2. Please do not modify any files in the project, as it may cause the project to malfunction.
3. If you encounter errors such as 403, please check if your reverse proxy configuration is correct or if your IP is clean, among other factors.
4. Please do not directly use the login or registration interfaces to log in or register, as there is a high probability of account banning. If you use them, please use the sessionKey to log in yourself.
5. This project is for learning and exchange purposes only and must not be used for commercial purposes. Otherwise, you will bear the consequences.
6. This project is for learning and exchange purposes only and must not be used for commercial purposes. Otherwise, you will bear the consequences.
7. This project is for learning and exchange purposes only and must not be used for commercial purposes. Otherwise, you will bear the consequences.
8. This project is for learning and exchange purposes only and must not be used for commercial purposes. Otherwise, you will bear the consequences.

### Rewarding the Author
If you find this project helpful, you can reward me. Thank you!

USDT TRC address: `TBZdWC2y1b2DPLK6awnEfShUu7x9XRY4xp`