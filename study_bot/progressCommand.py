from discord.ext import commands
from dotenv import load_dotenv
import os
import pymongo
from pymongo import MongoClient

# MongoClient 설정 (MongoDB 연결)
client = MongoClient(os.getenv('MONGO_URI'))
db = client[os.getenv('MONGO_DATABASE')]
progress_collection = db[os.getenv('MONGO_COLLECTION')]

class ProgressCommand(commands.Cog):
    def __init__(self, bot):
        self.bot = bot
