const Discord = require("discord.js")
const bot = new Discord.Client()
const {TOKEN, GOOGLE_API_KEY, PREFIX} = require("./config")
bot.music = require("discord.js-musicbot-addon")

bot.music.start(bot, {
    youtubeKey: GOOGLE_API_KEY,
    play: {
        usage: "{{PREFIX}}play",
    }
})

bot.login(TOKEN)