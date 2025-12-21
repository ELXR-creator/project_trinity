from telegram import Update
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes
import os
from datetime import datetime

BOT_TOKEN = os.getenv("BOT_TOKEN")



if not BOT_TOKEN:
    raise RuntimeError("token is missing")

def date_and_time_logger():
    now = datetime.now()
    date_str = now.strftime("%Y-%m-%d")      # 2025-01-07
    time_str = now.strftime("%H:%M:%S")      # 21:34:10
    day_name = now.strftime("%A")            # Monday

    
async def handle_start(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    await update.message.reply_text("I am alive!")
    await update.message.reply_animation(
        animation="https://media.giphy.com/media/3o7aD2saalBwwftBIY/giphy.gif"
    )


async def handle_task(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    await update.message.reply_text("What is your task?")


def main():
    app = ApplicationBuilder().token(BOT_TOKEN).build()

    app.add_handler(CommandHandler("start", handle_start))
    app.add_handler(CommandHandler("task", handle_task))

    app.run_polling()


if __name__ == "__main__":
    main()

