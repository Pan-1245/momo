import {
  REST,
  Client,
  Routes,
  GatewayIntentBits,
  PermissionFlagsBits,
} from "discord.js";

import { config } from "./config.js";
import roles from "./commands/roles.js";

const client = new Client({
  intents: [
    GatewayIntentBits.Guilds,
    GatewayIntentBits.GuildMessages,
    GatewayIntentBits.MessageContent,
    GatewayIntentBits.GuildModeration,
  ],
});

const rest = new REST({ version: "10" }).setToken(config.DISCORD_TOKEN);

client.on("ready", () => {
  console.log(`Logged in as ${client.user.tag}!`);
});

client.on("interactionCreate", async (interaction) => {
  if (!interaction.isChatInputCommand()) return;

  if (interaction.commandName === "addrole") {
    const targetMember = interaction.options.getMember("target");
    const role = interaction.options.getRole("role");

    // Permission check (make sure the bot can manage the role)
    if (
      !targetMember.guild.members.me.permissions.has(
        PermissionFlagsBits.ManageRoles
      ) ||
      role.position >= targetMember.guild.members.me.roles.highest.position
    ) {
      return interaction.reply({
        content: "I don't have permission to manage that role.",
        ephemeral: true,
      });
    }

    // Check if the target member already has the role
    if (targetMember.roles.cache.has(role.id)) {
      return interaction.reply({
        content: "That user already has that role.",
        ephemeral: true,
      });
    }

    try {
      await targetMember.roles.add(role);
      await interaction.reply(
        `Successfully added the ${role.name} role to ${targetMember.user.tag}.`
      );
    } catch (err) {
      console.error(err);
      await interaction.reply({
        content: "There was an error adding the role.",
        ephemeral: true,
      });
    }
  }
});

client.on("messageDeleteBulk", async (messages) => {
  console.log(`${messages.size} deleted.`);
});

async function main() {
  const commands = [roles];

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

    client.login(config.DISCORD_TOKEN);
  } catch (err) {
    console.log(err);
  }
}

main();
