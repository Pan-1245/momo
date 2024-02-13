import { REST, Client, GatewayIntentBits } from "discord.js";
import { config } from "./config";

//#region Client intents
const client = new Client({
  intents: [
    GatewayIntentBits.Guilds,
    GatewayIntentBits.GuildMessages,
    GatewayIntentBits.MessageContent,
    GatewayIntentBits.GuildModeration,
  ],
});
//#endregion

const rest = new REST({ version: "10" }).setToken(config.DISCORD_TOKEN);

//#region Function
async function main() {
  const commands = [];

  try {
    console.log("Started refreshing application (/) commands.");
    await rest.put(
      Routes.applicationGuildCommands(
        config.DISCORD_CLIENT_ID,
        config.DISCORD_GUILD_ID
      ),
      {
        body: commands,
      }
    );

    client.login(DISCORD_TOKEN);
  } catch (err) {
    console.log(err);
  }
}
//#endregion

client.on("ready", () => {
  console.log(`Logged in as ${client.user.tag}!`);
});

//#region Interaction
client.on("interactionCreate", async (interaction) => {
  if (!interaction.isChatInputCommand()) return;
});
//#endregion

main();
