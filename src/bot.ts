import { Client, GatewayIntentBits } from "discord.js";

require("dotenv").config();

const TOKEN = process.env.TOKEN as string;

const client = new Client({ intents: [GatewayIntentBits.Guilds] });

client.on("ready", () => {
  if (client.user !== null) {
    console.log(`Logged in as ${client.user.tag}!`);
  }
});

client.on("interactionCreate", async (interaction) => {
  if (!interaction.isChatInputCommand()) return;

  if (interaction.commandName === "ping") {
    await interaction.reply("Pong!");
  }
});

client.login(TOKEN);
