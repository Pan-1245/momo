import { SlashCommandBuilder, PermissionFlagsBits } from "discord.js";

export default new SlashCommandBuilder()
  .setName("addrole")
  .setDescription("Adds a specified role to a user.")
  .addUserOption((option) =>
    option
      .setName("target")
      .setDescription("The user to give the role to")
      .setRequired(true)
  )
  .addRoleOption((option) =>
    option.setName("role").setDescription("The role to add").setRequired(true)
  )
  .setDefaultMemberPermissions(PermissionFlagsBits.ManageRoles);
