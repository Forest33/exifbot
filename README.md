# exifbot

Telegram bot that allows you to extract EXIF data from images.

### Installation

Create database and user
```
CREATE DATABASE exifbot;
CREATE USER exifbot WITH LOGIN PASSWORD 'change-me';
GRANT ALL ON DATABASE exifbot TO exifbot;
```

```
git clone https://github.com/Forest33/exifbot.git
cd exifbot
```

Edit the config file config/exifbot.json

```
docker-compose up -d
```
