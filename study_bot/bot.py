 import discord
 from discord.ext import commands
 ​
 intents = discord.Intents.default()
 intents.message_content = True
 bot = commands.Bot(command_prefix='!', intents=intents)
 ​
 @bot.event
 async def on_ready():
     await bot.add_cog(HelloCommand(bot))
 ​
 bot.run('STUDY_BOT_TOKEN')
