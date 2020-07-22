import logging
from telegram.ext import Updater, CommandHandler, MessageHandler, Filters, run_async
from telegram.error import BadRequest
token = ""
group_id = ""

logging.basicConfig(format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
                    level=logging.INFO)
logger = logging.getLogger(__name__)

@run_async
def start(update, context):
    bot = context.bot
    user = update.effective_user
    if not update.message.chat.type == "private":
        return
    chat = bot.getChat(group_id)
    chat_user = True
    try:
        chat_member = chat.get_member(user.id)
    except BadRequest:
        chat_user = False
    if chat_member.status == 'left' or chat_member.status == 'kicked':
        chat_user = False
    if not chat_user:
        return
    update.message.reply_text('YOUR-STRING-HERE')

@run_async
def echo(update, context):
    bot = context.bot
    user = update.effective_user
    if not update.message.chat.type == "private":
        return
    chat = bot.getChat(group_id)
    chat_user = True
    try:
        chat_member = chat.get_member(user.id)
    except BadRequest:
        chat_user = False
    if chat_member.status == 'left' or chat_member.status == 'kicked':
        chat_user = False
    if not chat_user:
        return
    chat = bot.getChat(group_id)
    chat.send_message(update.message.text)

def main():
    updater = Updater(token=token, use_context=True)
    dp = updater.dispatcher
    logger.info("Bot started.")
    dp.add_handler(CommandHandler("start", start))
    dp.add_handler(MessageHandler(Filters.text & ~Filters.command, echo))
    updater.start_polling()
    updater.idle()

if __name__ == '__main__':
    main()
