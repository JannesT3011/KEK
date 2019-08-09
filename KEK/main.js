const Discord = require("discord.js")
const bot = new Discord.Client()
const {TOKEN, GOOGLE_API_KEY, PREFIX} = require("./config")
bot.music = require("discord.js-musicbot-addon")

bot.on("ready", function() {
    console.log("Ready!")
})

bot.on("message", function(message) {
    if (message.content.startsWith(PREFIX + "emoji")) {
        emote_name = message.content.split(" ")[1]
        img_url = message.content.split(" ")[2]
        try {
            message.guild.createEmoji(img_url, emote_name)
            return message.channel.send("Emoji `" + emote_name + "` added!")
        } catch (error) {
            return message.channel.send("An error occured: " + error)
        }
    }
})

bot.music.start(bot, {
    youtubeKey: GOOGLE_API_KEY,
    play: {
        usage: "{{PREFIX}}play",
    }
})

bot.login(TOKEN)
